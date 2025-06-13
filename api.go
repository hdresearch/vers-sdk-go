// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vers

import (
	"github.com/hdresearch/vers-sdk-go/option"
)

// APIService contains methods and other services that help with interacting with
// the vers API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIService] method instead.
type APIService struct {
	Options []option.RequestOption
	Cluster *APIClusterService
	Vm      *APIVmService
	Rootfs  *APIRootfService
	Health  *APIHealthService
	Network *APINetworkService
}

// NewAPIService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAPIService(opts ...option.RequestOption) (r *APIService) {
	r = &APIService{}
	r.Options = opts
	r.Cluster = NewAPIClusterService(opts...)
	r.Vm = NewAPIVmService(opts...)
	r.Rootfs = NewAPIRootfService(opts...)
	r.Health = NewAPIHealthService(opts...)
	r.Network = NewAPINetworkService(opts...)
	return
}
