package types

// Pricefeed module event types
const (
	EventTypeRequestPriceUpdated = "request_price_updated"
	EventTypeOracleUpdatedPrice = "oracle_updated_price"
	EventTypeNoValidPrices      = "no_valid_prices"

	AttributeValueCategory = ModuleName
	AttributeRequestID      = "request_id"
	AttributeRequestPrice   = "request_price"
	AttributeOracle        = "oracle"
	AttributeExpiry        = "expiry"
)
