package types

import (
	"errors"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Request an asset in the apis
type Request struct {
	// TODO: rename to ID
	RequestID   string           `json:"request_id" yaml:"request_id"`
	URL string           `json:"url" yaml:"url"`
	Active     bool             `json:"active" yaml:"active"`
}

const (
	exampleAPICall = "https://www.metaweather.com/api/location/search/?query=boston"
)

// NewRequest returns a new Request
func NewRequest(id, URL string, active bool) Request {
	return Request{
		RequestID:   id,
		URL:  URL,
		Active:     active,
	}
}

// String implement fmt.Stringer
func (m Request) String() string {
	return fmt.Sprintf(`Asset:
	Request ID: %s
	URL: %s
	Active: %t`,
		m.RequestID, m.URL, m.Active)
}


// Validate performs a basic validation of the request params
func (m Request) Validate() error {
	if strings.TrimSpace(m.RequestID) == "" {
		return errors.New("request id cannot be blank")
	}
	if strings.TrimSpace(m.URL) == "" {
		return errors.New("url cannot be blank")
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


// SortDecs provides the interface needed to sort sdk.Dec slices
type SortDecs []sdk.Dec

func (a SortDecs) Len() int           { return len(a) }
func (a SortDecs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortDecs) Less(i, j int) bool { return a[i].LT(a[j]) }
