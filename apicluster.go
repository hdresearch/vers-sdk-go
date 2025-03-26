// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firecrackermanager

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/firecracker-manager-go/internal/apijson"
	"github.com/stainless-sdks/firecracker-manager-go/internal/apiquery"
	"github.com/stainless-sdks/firecracker-manager-go/internal/param"
	"github.com/stainless-sdks/firecracker-manager-go/internal/requestconfig"
	"github.com/stainless-sdks/firecracker-manager-go/option"
)

// APIClusterService contains methods and other services that help with interacting
// with the vers API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIClusterService] method instead.
type APIClusterService struct {
	Options []option.RequestOption
}

// NewAPIClusterService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPIClusterService(opts ...option.RequestOption) (r *APIClusterService) {
	r = &APIClusterService{}
	r.Options = opts
	return
}

// Start a new cluster.
func (r *APIClusterService) New(ctx context.Context, body APIClusterNewParams, opts ...option.RequestOption) (res *Cluster, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/cluster"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve information on a particular cluster.
func (r *APIClusterService) Get(ctx context.Context, clusterID string, opts ...option.RequestOption) (res *Cluster, err error) {
	opts = append(r.Options[:], opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("api/cluster/%s", clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all clusters.
func (r *APIClusterService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Cluster, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/cluster"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a cluster.
func (r *APIClusterService) Delete(ctx context.Context, clusterID string, body APIClusterDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("api/cluster/%s", clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

type Cluster struct {
	// The cluster's ID.
	ID string `json:"id,required"`
	// The VMs that are children of the cluster, including the root VM.
	Children []Vm `json:"children,required"`
	// The ID of the cluster's root VM.
	RootVmID string `json:"root_vm_id,required"`
	// How many VMs are currently running on this cluster.
	VmCount int64       `json:"vm_count,required"`
	JSON    clusterJSON `json:"-"`
}

// clusterJSON contains the JSON metadata for the struct [Cluster]
type clusterJSON struct {
	ID          apijson.Field
	Children    apijson.Field
	RootVmID    apijson.Field
	VmCount     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Cluster) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clusterJSON) RawJSON() string {
	return r.raw
}

type APIClusterNewParams struct {
	Body interface{} `json:"body,required"`
}

func (r APIClusterNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type APIClusterDeleteParams struct {
	// If true, will delete the cluster and all its children. Otherwise, will only
	// delete the cluster if it has no children (besides the root VM)
	Force param.Field[bool] `query:"force"`
}

// URLQuery serializes [APIClusterDeleteParams]'s query parameters as `url.Values`.
func (r APIClusterDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
