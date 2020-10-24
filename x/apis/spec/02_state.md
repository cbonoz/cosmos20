<!--
order: 2
-->

# State

## Parameters and genesis state

`Parameters` determine which requests are tracked by the apis and which oracles are authorized to post prices for a given request. There is only one active parameter set at any given time. Updates to parameters can be made via on-chain parameter update proposals.

```go
// Params params for apis. Can be altered via governance
type Params struct {
	Requests Requests `json:"requests" yaml:"requests"` //  Array containing the requests supported by the apis
}

// Request an asset in the apis
type Request struct {
	RequestID   string           `json:"request_id" yaml:"request_id"`
	BaseAsset  string           `json:"base_asset" yaml:"base_asset"`
	QuoteAsset string           `json:"quote_asset" yaml:"quote_asset"`
	Oracles    []sdk.AccAddress `json:"oracles" yaml:"oracles"`
	Active     bool             `json:"active" yaml:"active"`
}

type Requests []Request
```

`GenesisState` defines the state that must be persisted when the blockchain stops/stars in order for the normal function of the apis to resume.

```go
// GenesisState - apis state that must be provided at genesis
type GenesisState struct {
	Params       Params        `json:"params" yaml:"params"`
	PostedPrices []PostedPrice `json:"posted_prices" yaml:"posted_prices"`
}

// PostedPrice price for request posted by a specific oracle
type PostedPrice struct {
	RequestID      string         `json:"request_id" yaml:"request_id"`
	OracleAddress sdk.AccAddress `json:"oracle_address" yaml:"oracle_address"`
	Price         sdk.Dec        `json:"price" yaml:"price"`
	Expiry        time.Time      `json:"expiry" yaml:"expiry"`
}

type PostedPrices []PostedPrice
```
