<!--
order: 3
-->

# Messages

## Posting Prices

An authorized oraclef for a particular request can post the current price for that request using the `MsgPostResponse` type.

```go
// MsgPostResponse struct representing a posted price message.
// Used by oracles to input prices to the apis
type MsgPostResponse struct {
	From     sdk.AccAddress `json:"from" yaml:"from"`           // client that sent in this address
	RequestID string         `json:"request_id" yaml:"request_id"` // asset code used by exchanges/api
	Price    sdk.Dec        `json:"price" yaml:"price"`         // price in decimal (max precision 18)
	Expiry   time.Time      `json:"expiry" yaml:"expiry"`       // expiry time
}
```

### State Modifications

- Update the raw price for the oracle for this request. This replaces any previous price for that oracle.
