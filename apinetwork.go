// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vers

import (
	"context"
	"net/http"

	"github.com/hdresearch/vers-sdk-go/internal/apijson"
	"github.com/hdresearch/vers-sdk-go/internal/requestconfig"
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

// Get network information
func (r *APINetworkService) GetInfo(ctx context.Context, opts ...option.RequestOption) (res *Info, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/network"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Info struct {
	NumNetworks          int64    `json:"num_networks,required"`
	NumNetworksAvailable int64    `json:"num_networks_available,required"`
	NumNetworksInUse     int64    `json:"num_networks_in_use,required"`
	JSON                 infoJSON `json:"-"`
}

// infoJSON contains the JSON metadata for the struct [Info]
type infoJSON struct {
	NumNetworks          apijson.Field
	NumNetworksAvailable apijson.Field
	NumNetworksInUse     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *Info) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r infoJSON) RawJSON() string {
	return r.raw
}
