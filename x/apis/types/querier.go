package types

// price Takes an [assetcode] and returns CurrentPrice for that asset
// apis Takes an [assetcode] and returns the raw []PostedPrice for that asset
// assets Returns []Assets in the apis system

const (
	// QueryGetParams command for params query
	QueryGetParams = "parameters"
	// QueryRequests command for assets query
	QueryRequests = "requests"
)

// QueryWithRequestIDParams fields for querying information from a specific request
type QueryWithRequestIDParams struct {
	RequestID string
}

// NewQueryWithRequestIDParams creates a new instance of QueryWithRequestIDParams
func NewQueryWithRequestIDParams(requestID string) QueryWithRequestIDParams {
	return QueryWithRequestIDParams{
		RequestID: requestID,
	}
}
