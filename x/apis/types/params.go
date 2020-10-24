package types

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/x/params"
)

// Parameter keys
var (
	KeyRequests     = []byte("Requests")
	DefaultRequests = Requests{}
)

// Params params for apis. Can be altered via governance
type Params struct {
	Requests Requests `json:"requests" yaml:"requests"` //  Array containing the requests supported by the apis
}

// NewParams creates a new AssetParams object
func NewParams(requests Requests) Params {
	return Params{
		Requests: requests,
	}
}

// DefaultParams default params for apis
func DefaultParams() Params {
	return NewParams(DefaultRequests)
}

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
// pairs of apis module's parameters.
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		params.NewParamSetPair(KeyRequests, &p.Requests, validateRequestParams),
	}
}

// String implements fmt.stringer
func (p Params) String() string {
	out := "Params:\n"
	for _, a := range p.Requests {
		out += fmt.Sprintf("%s\n", a.String())
	}
	return strings.TrimSpace(out)
}

// Validate ensure that params have valid values
func (p Params) Validate() error {
	return validateRequestParams(p.Requests)
}

func validateRequestParams(i interface{}) error {
	requests, ok := i.(Requests)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return requests.Validate()
}
