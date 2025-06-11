// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vers

import (
	"github.com/hdresearch/vers-sdk-go/option"
)

// APINetworkService contains methods and other services that help with interacting
// with the vers API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPINetworkService] method instead.
type APINetworkService struct {
	Options []option.RequestOption
}

// NewAPINetworkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPINetworkService(opts ...option.RequestOption) (r *APINetworkService) {
	r = &APINetworkService{}
	r.Options = opts
	return
}
