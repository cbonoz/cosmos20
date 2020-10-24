package simulation

import (
	"math/rand"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// PriceGenerator allows deterministic price generation in simulations
type PriceGenerator struct {
	requests            []string
	currentPrice       map[string]sdk.Dec
	maxPrice           map[string]sdk.Dec
	minPrice           map[string]sdk.Dec
	increment          map[string]sdk.Dec
	currentBlockHeight int64
}

// NewPriceGenerator returns a new request price generator from starting values
func NewPriceGenerator(startingPrice map[string]sdk.Dec) *PriceGenerator {
	p := &PriceGenerator{
		requests:            []string{},
		currentPrice:       startingPrice,
		maxPrice:           map[string]sdk.Dec{},
		minPrice:           map[string]sdk.Dec{},
		increment:          map[string]sdk.Dec{},
		currentBlockHeight: 0,
	}

	divisor := sdk.MustNewDecFromStr("20")

	for requestID, startPrice := range startingPrice {
		p.requests = append(p.requests, requestID)
		// allow 10x price increase
		p.maxPrice[requestID] = sdk.MustNewDecFromStr("10.0").Mul(startPrice)
		// allow 100x price decrease
		p.minPrice[requestID] = sdk.MustNewDecFromStr("0.01").Mul(startPrice)
		// set increment - should we use a random increment?
		p.increment[requestID] = startPrice.Quo(divisor)
	}

	// request prices must be calculated in a deterministic order
	// this sort order defines the the order we update each request
	// price in the step function
	sort.Strings(p.requests)

	return p
}

// Step walks prices to a current block height from the previously called height
// noop if called more than once for the same height
func (p *PriceGenerator) Step(r *rand.Rand, blockHeight int64) {
	if p.currentBlockHeight == blockHeight {
		// step already called for blockHeight
		return
	}

	if p.currentBlockHeight > blockHeight {
		// step is called with a previous blockHeight
		panic("step out of order")
	}

	for _, requestID := range p.requests {
		lastPrice := p.currentPrice[requestID]
		minPrice := p.minPrice[requestID]
		maxPrice := p.maxPrice[requestID]
		increment := p.increment[requestID]
		lastHeight := p.currentBlockHeight

		for lastHeight < blockHeight {
			upDown := r.Intn(2)

			if upDown == 0 {
				lastPrice = sdk.MinDec(lastPrice.Add(increment), maxPrice)
			} else {
				lastPrice = sdk.MaxDec(lastPrice.Sub(increment), minPrice)
			}

			lastHeight++
		}

		p.currentPrice[requestID] = lastPrice
	}

	p.currentBlockHeight = blockHeight
}

// GetCurrentPrice returns price for last blockHeight set by Step
func (p *PriceGenerator) GetCurrentPrice(requestID string) sdk.Dec {
	price, ok := p.currentPrice[requestID]

	if !ok {
		panic("unknown request")
	}

	return price
}
