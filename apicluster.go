// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/hdresearch/vers-sdk-go/internal/apijson"
	"github.com/hdresearch/vers-sdk-go/internal/param"
	"github.com/hdresearch/vers-sdk-go/internal/requestconfig"
	"github.com/hdresearch/vers-sdk-go/option"
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

// Create a new cluster.
func (r *APIClusterService) New(ctx context.Context, body APIClusterNewParams, opts ...option.RequestOption) (res *APIClusterNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/cluster"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve information on a particular cluster.
func (r *APIClusterService) Get(ctx context.Context, clusterID string, opts ...option.RequestOption) (res *APIClusterGetResponse, err error) {
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
func (r *APIClusterService) List(ctx context.Context, opts ...option.RequestOption) (res *APIClusterListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/cluster"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a cluster.
func (r *APIClusterService) Delete(ctx context.Context, clusterID string, opts ...option.RequestOption) (res *APIClusterDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("api/cluster/%s", clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get the SSH private key for VM access
func (r *APIClusterService) GetSSHKey(ctx context.Context, clusterID string, opts ...option.RequestOption) (res *APIClusterGetSSHKeyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("api/cluster/%s/ssh_key", clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CreateParam struct {
	KernelName param.Field[string] `json:"kernel_name"`
	MemSizeMib param.Field[int64]  `json:"mem_size_mib"`
	RootfsName param.Field[string] `json:"rootfs_name"`
	VcpuCount  param.Field[int64]  `json:"vcpu_count"`
}

func (r CreateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type APIClusterNewResponse struct {
	Data        APIClusterNewResponseData `json:"data,required"`
	DurationNs  int64                     `json:"duration_ns,required"`
	OperationID string                    `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                     `json:"time_start,required"`
	JSON      apiClusterNewResponseJSON `json:"-"`
}

// apiClusterNewResponseJSON contains the JSON metadata for the struct
// [APIClusterNewResponse]
type apiClusterNewResponseJSON struct {
	Data        apijson.Field
	DurationNs  apijson.Field
	OperationID apijson.Field
	TimeStart   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIClusterNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterNewResponseJSON) RawJSON() string {
	return r.raw
}

type APIClusterNewResponseData struct {
	// The cluster's ID.
	ID string `json:"id,required"`
	// How many VMs are currently running on this cluster.
	VmCount int64 `json:"vm_count,required"`
	// The VMs that are children of the cluster, including the root VM.
	Vms []Vm `json:"vms,required"`
	// The ID of the cluster's root VM.
	RootVmID string                        `json:"root_vm_id,nullable"`
	JSON     apiClusterNewResponseDataJSON `json:"-"`
}

// apiClusterNewResponseDataJSON contains the JSON metadata for the struct
// [APIClusterNewResponseData]
type apiClusterNewResponseDataJSON struct {
	ID          apijson.Field
	VmCount     apijson.Field
	Vms         apijson.Field
	RootVmID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIClusterNewResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterNewResponseDataJSON) RawJSON() string {
	return r.raw
}

type APIClusterGetResponse struct {
	Data        APIClusterGetResponseData `json:"data,required"`
	DurationNs  int64                     `json:"duration_ns,required"`
	OperationID string                    `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                     `json:"time_start,required"`
	JSON      apiClusterGetResponseJSON `json:"-"`
}

// apiClusterGetResponseJSON contains the JSON metadata for the struct
// [APIClusterGetResponse]
type apiClusterGetResponseJSON struct {
	Data        apijson.Field
	DurationNs  apijson.Field
	OperationID apijson.Field
	TimeStart   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIClusterGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterGetResponseJSON) RawJSON() string {
	return r.raw
}

type APIClusterGetResponseData struct {
	// The cluster's ID.
	ID string `json:"id,required"`
	// How many VMs are currently running on this cluster.
	VmCount int64 `json:"vm_count,required"`
	// The VMs that are children of the cluster, including the root VM.
	Vms []Vm `json:"vms,required"`
	// The ID of the cluster's root VM.
	RootVmID string                        `json:"root_vm_id,nullable"`
	JSON     apiClusterGetResponseDataJSON `json:"-"`
}

// apiClusterGetResponseDataJSON contains the JSON metadata for the struct
// [APIClusterGetResponseData]
type apiClusterGetResponseDataJSON struct {
	ID          apijson.Field
	VmCount     apijson.Field
	Vms         apijson.Field
	RootVmID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIClusterGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type APIClusterListResponse struct {
	Data        []APIClusterListResponseData `json:"data,required"`
	DurationNs  int64                        `json:"duration_ns,required"`
	OperationID string                       `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                      `json:"time_start,required"`
	JSON      apiClusterListResponseJSON `json:"-"`
}

// apiClusterListResponseJSON contains the JSON metadata for the struct
// [APIClusterListResponse]
type apiClusterListResponseJSON struct {
	Data        apijson.Field
	DurationNs  apijson.Field
	OperationID apijson.Field
	TimeStart   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIClusterListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterListResponseJSON) RawJSON() string {
	return r.raw
}

type APIClusterListResponseData struct {
	// The cluster's ID.
	ID string `json:"id,required"`
	// How many VMs are currently running on this cluster.
	VmCount int64 `json:"vm_count,required"`
	// The VMs that are children of the cluster, including the root VM.
	Vms []Vm `json:"vms,required"`
	// The ID of the cluster's root VM.
	RootVmID string                         `json:"root_vm_id,nullable"`
	JSON     apiClusterListResponseDataJSON `json:"-"`
}

// apiClusterListResponseDataJSON contains the JSON metadata for the struct
// [APIClusterListResponseData]
type apiClusterListResponseDataJSON struct {
	ID          apijson.Field
	VmCount     apijson.Field
	Vms         apijson.Field
	RootVmID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIClusterListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterListResponseDataJSON) RawJSON() string {
	return r.raw
}

type APIClusterDeleteResponse struct {
	Data        APIClusterDeleteResponseData `json:"data,required"`
	DurationNs  int64                        `json:"duration_ns,required"`
	OperationID string                       `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                        `json:"time_start,required"`
	JSON      apiClusterDeleteResponseJSON `json:"-"`
}

// apiClusterDeleteResponseJSON contains the JSON metadata for the struct
// [APIClusterDeleteResponse]
type apiClusterDeleteResponseJSON struct {
	Data        apijson.Field
	DurationNs  apijson.Field
	OperationID apijson.Field
	TimeStart   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIClusterDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type APIClusterDeleteResponseData struct {
	// The cluster's ID.
	ID string `json:"id,required"`
	// How many VMs are currently running on this cluster.
	VmCount int64 `json:"vm_count,required"`
	// The VMs that are children of the cluster, including the root VM.
	Vms []Vm `json:"vms,required"`
	// The ID of the cluster's root VM.
	RootVmID string                           `json:"root_vm_id,nullable"`
	JSON     apiClusterDeleteResponseDataJSON `json:"-"`
}

// apiClusterDeleteResponseDataJSON contains the JSON metadata for the struct
// [APIClusterDeleteResponseData]
type apiClusterDeleteResponseDataJSON struct {
	ID          apijson.Field
	VmCount     apijson.Field
	Vms         apijson.Field
	RootVmID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIClusterDeleteResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterDeleteResponseDataJSON) RawJSON() string {
	return r.raw
}

type APIClusterGetSSHKeyResponse struct {
	Data        string `json:"data,required"`
	DurationNs  int64  `json:"duration_ns,required"`
	OperationID string `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                           `json:"time_start,required"`
	JSON      apiClusterGetSSHKeyResponseJSON `json:"-"`
}

// apiClusterGetSSHKeyResponseJSON contains the JSON metadata for the struct
// [APIClusterGetSSHKeyResponse]
type apiClusterGetSSHKeyResponseJSON struct {
	Data        apijson.Field
	DurationNs  apijson.Field
	OperationID apijson.Field
	TimeStart   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIClusterGetSSHKeyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterGetSSHKeyResponseJSON) RawJSON() string {
	return r.raw
}

type APIClusterNewParams struct {
	Create CreateParam `json:"create,required"`
}

func (r APIClusterNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Create)
}
