package types

const (
	// ModuleName The name that will be used throughout the module
	ModuleName = "apis"

	// StoreKey Top level store key where all module items will be stored
	StoreKey = ModuleName

	// RouterKey Top level router key
	RouterKey = ModuleName

	// QuerierRoute is the querier route for gov
	QuerierRoute = ModuleName

	// DefaultParamspace default namestore
	DefaultParamspace = ModuleName
)

var (
	// CurrentPricePrefix prefix for the current price of an asset
	CurrentPricePrefix = []byte{0x00}

	// RawPriceFeedPrefix prefix for the raw apis of an asset
	RawPriceFeedPrefix = []byte{0x01}
)

// CurrentPriceKey returns the prefix for the current price
func CurrentPriceKey(requestID string) []byte {
	return append(CurrentPricePrefix, []byte(requestID)...)
}

// RawPriceKey returns the prefix for the raw price
func RawPriceKey(requestID string) []byte {
	return append(RawPriceFeedPrefix, []byte(requestID)...)
}
