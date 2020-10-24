package types

import (
	"errors"
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Request an asset in the apis
type Request struct {
	// TODO: rename to ID
	RequestID   string           `json:"request_id" yaml:"request_id"`
	BaseAsset  string           `json:"base_asset" yaml:"base_asset"`
	QuoteAsset string           `json:"quote_asset" yaml:"quote_asset"`
	Oracles    []sdk.AccAddress `json:"oracles" yaml:"oracles"`
	Active     bool             `json:"active" yaml:"active"`
}

// NewRequest returns a new Request
func NewRequest(id, base, quote string, oracles []sdk.AccAddress, active bool) Request {
	return Request{
		RequestID:   id,
		BaseAsset:  base,
		QuoteAsset: quote,
		Oracles:    oracles,
		Active:     active,
	}
}

// String implement fmt.Stringer
func (m Request) String() string {
	return fmt.Sprintf(`Asset:
	Request ID: %s
	Base Asset: %s
	Quote Asset: %s
	Oracles: %s
	Active: %t`,
		m.RequestID, m.BaseAsset, m.QuoteAsset, m.Oracles, m.Active)
}

// Validate performs a basic validation of the request params
func (m Request) Validate() error {
	if strings.TrimSpace(m.RequestID) == "" {
		return errors.New("request id cannot be blank")
	}
	if err := sdk.ValidateDenom(m.BaseAsset); err != nil {
		return fmt.Errorf("invalid base asset: %w", err)
	}
	if err := sdk.ValidateDenom(m.QuoteAsset); err != nil {
		return fmt.Errorf("invalid quote asset: %w", err)
	}
	seenOracles := make(map[string]bool)
	for i, oracle := range m.Oracles {
		if oracle.Empty() {
			return fmt.Errorf("oracle %d is empty", i)
		}
		if seenOracles[oracle.String()] {
			return fmt.Errorf("duplicated oracle %s", oracle)
		}
		seenOracles[oracle.String()] = true
	}
	return nil
}

// Requests array type for oracle
type Requests []Request

// Validate checks if all the requests are valid and there are no duplicated
// entries.
func (ms Requests) Validate() error {
	seenRequests := make(map[string]bool)
	for _, m := range ms {
		if seenRequests[m.RequestID] {
			return fmt.Errorf("duplicated request %s", m.RequestID)
		}
		if err := m.Validate(); err != nil {
			return err
		}
		seenRequests[m.RequestID] = true
	}
	return nil
}

// String implements fmt.Stringer
func (ms Requests) String() string {
	out := "Requests:\n"
	for _, m := range ms {
		out += fmt.Sprintf("%s\n", m.String())
	}
	return strings.TrimSpace(out)
}

// CurrentPrice struct that contains the metadata of a current price for a particular request in the apis module.
type CurrentPrice struct {
	RequestID string  `json:"request_id" yaml:"request_id"`
	Price    sdk.Dec `json:"price" yaml:"price"`
}

// NewCurrentPrice returns an instance of CurrentPrice
func NewCurrentPrice(requestID string, price sdk.Dec) CurrentPrice {
	return CurrentPrice{RequestID: requestID, Price: price}
}

// CurrentPrices type for an array of CurrentPrice
type CurrentPrices []CurrentPrice

// PostedPrice price for request posted by a specific oracle
type PostedPrice struct {
	RequestID      string         `json:"request_id" yaml:"request_id"`
	OracleAddress sdk.AccAddress `json:"oracle_address" yaml:"oracle_address"`
	Price         sdk.Dec        `json:"price" yaml:"price"`
	Expiry        time.Time      `json:"expiry" yaml:"expiry"`
}

// NewPostedPrice returns a new PostedPrice
func NewPostedPrice(requestID string, oracle sdk.AccAddress, price sdk.Dec, expiry time.Time) PostedPrice {
	return PostedPrice{
		RequestID:      requestID,
		OracleAddress: oracle,
		Price:         price,
		Expiry:        expiry,
	}
}

// Validate performs a basic check of a PostedPrice params.
func (pp PostedPrice) Validate() error {
	if strings.TrimSpace(pp.RequestID) == "" {
		return errors.New("request id cannot be blank")
	}
	if pp.OracleAddress.Empty() {
		return errors.New("oracle address cannot be empty")
	}
	if pp.Price.IsNegative() {
		return fmt.Errorf("posted price cannot be negative %s", pp.Price)
	}
	if pp.Expiry.IsZero() {
		return errors.New("expiry time cannot be zero")
	}
	return nil
}

// PostedPrices type for an array of PostedPrice
type PostedPrices []PostedPrice

// Validate checks if all the posted prices are valid and there are no duplicated
// entries.
func (pps PostedPrices) Validate() error {
	seenPrices := make(map[string]bool)
	for _, pp := range pps {
		if pp.OracleAddress != nil && seenPrices[pp.RequestID+pp.OracleAddress.String()] {
			return fmt.Errorf("duplicated posted price for marked id %s and oracle address %s", pp.RequestID, pp.OracleAddress)
		}

		if err := pp.Validate(); err != nil {
			return err
		}
		seenPrices[pp.RequestID+pp.OracleAddress.String()] = true
	}

	return nil
}

// implement fmt.Stringer
func (cp CurrentPrice) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Request ID: %s
Price: %s`, cp.RequestID, cp.Price))
}

// implement fmt.Stringer
func (pp PostedPrice) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Request ID: %s
Oracle Address: %s
Price: %s
Expiry: %s`, pp.RequestID, pp.OracleAddress, pp.Price, pp.Expiry))
}

// String implements fmt.Stringer
func (ps PostedPrices) String() string {
	out := "Posted Prices:\n"
	for _, p := range ps {
		out += fmt.Sprintf("%s\n", p.String())
	}
	return strings.TrimSpace(out)
}

// SortDecs provides the interface needed to sort sdk.Dec slices
type SortDecs []sdk.Dec

func (a SortDecs) Len() int           { return len(a) }
func (a SortDecs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortDecs) Less(i, j int) bool { return a[i].LT(a[j]) }
