// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vers

import (
	"context"
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

// List all clusters.
func (r *APIClusterService) List(ctx context.Context, opts ...option.RequestOption) (res *APIClusterListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/cluster"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CreateParam struct {
	ClusterType param.Field[CreateClusterType] `json:"cluster_type,required"`
	Params      param.Field[interface{}]       `json:"params,required"`
}

func (r CreateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CreateParam) implementsCreateUnionParam() {}

// Satisfied by [CreateNewClusterParamsParam],
// [CreateClusterFromCommitParamsParam], [CreateParam].
type CreateUnionParam interface {
	implementsCreateUnionParam()
}

type CreateNewClusterParamsParam struct {
	ClusterType param.Field[CreateNewClusterParamsClusterType] `json:"cluster_type,required"`
	Params      param.Field[CreateNewClusterParamsParamsParam] `json:"params,required"`
}

func (r CreateNewClusterParamsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CreateNewClusterParamsParam) implementsCreateUnionParam() {}

type CreateNewClusterParamsClusterType string

const (
	CreateNewClusterParamsClusterTypeNew CreateNewClusterParamsClusterType = "new"
)

func (r CreateNewClusterParamsClusterType) IsKnown() bool {
	switch r {
	case CreateNewClusterParamsClusterTypeNew:
		return true
	}
	return false
}

type CreateNewClusterParamsParamsParam struct {
	ClusterAlias param.Field[string] `json:"cluster_alias"`
	// The amount of total space to allocate to the cluster
	FsSizeClusterMib param.Field[int64] `json:"fs_size_cluster_mib"`
	// The size of the VM filesystem (if smaller than the base image + overhead, will
	// cause an error)
	FsSizeVmMib param.Field[int64]  `json:"fs_size_vm_mib"`
	KernelName  param.Field[string] `json:"kernel_name"`
	MemSizeMib  param.Field[int64]  `json:"mem_size_mib"`
	RootfsName  param.Field[string] `json:"rootfs_name"`
	VcpuCount   param.Field[int64]  `json:"vcpu_count"`
	VmAlias     param.Field[string] `json:"vm_alias"`
}

func (r CreateNewClusterParamsParamsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CreateClusterFromCommitParamsParam struct {
	ClusterType param.Field[CreateClusterFromCommitParamsClusterType] `json:"cluster_type,required"`
	Params      param.Field[CreateClusterFromCommitParamsParamsParam] `json:"params,required"`
}

func (r CreateClusterFromCommitParamsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CreateClusterFromCommitParamsParam) implementsCreateUnionParam() {}

type CreateClusterFromCommitParamsClusterType string

const (
	CreateClusterFromCommitParamsClusterTypeFromCommit CreateClusterFromCommitParamsClusterType = "from_commit"
)

func (r CreateClusterFromCommitParamsClusterType) IsKnown() bool {
	switch r {
	case CreateClusterFromCommitParamsClusterTypeFromCommit:
		return true
	}
	return false
}

type CreateClusterFromCommitParamsParamsParam struct {
	CommitKey        param.Field[string] `json:"commit_key,required"`
	ClusterAlias     param.Field[string] `json:"cluster_alias"`
	FsSizeClusterMib param.Field[int64]  `json:"fs_size_cluster_mib"`
	VmAlias          param.Field[string] `json:"vm_alias"`
}

func (r CreateClusterFromCommitParamsParamsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CreateClusterType string

const (
	CreateClusterTypeNew        CreateClusterType = "new"
	CreateClusterTypeFromCommit CreateClusterType = "from_commit"
)

func (r CreateClusterType) IsKnown() bool {
	switch r {
	case CreateClusterTypeNew, CreateClusterTypeFromCommit:
		return true
	}
	return false
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
	// The ID of the cluster's root VM.
	RootVmID string `json:"root_vm_id,required"`
	// How many VMs are currently running on this cluster.
	VmCount int64 `json:"vm_count,required"`
	// The VMs that are children of the cluster, including the root VM.
	Vms []Vm `json:"vms,required"`
	// Human-readable name assigned to the cluster.
	Alias string                        `json:"alias,nullable"`
	JSON  apiClusterNewResponseDataJSON `json:"-"`
}

// apiClusterNewResponseDataJSON contains the JSON metadata for the struct
// [APIClusterNewResponseData]
type apiClusterNewResponseDataJSON struct {
	ID          apijson.Field
	RootVmID    apijson.Field
	VmCount     apijson.Field
	Vms         apijson.Field
	Alias       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIClusterNewResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterNewResponseDataJSON) RawJSON() string {
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
	// The ID of the cluster's root VM.
	RootVmID string `json:"root_vm_id,required"`
	// How many VMs are currently running on this cluster.
	VmCount int64 `json:"vm_count,required"`
	// The VMs that are children of the cluster, including the root VM.
	Vms []Vm `json:"vms,required"`
	// Human-readable name assigned to the cluster.
	Alias string                         `json:"alias,nullable"`
	JSON  apiClusterListResponseDataJSON `json:"-"`
}

// apiClusterListResponseDataJSON contains the JSON metadata for the struct
// [APIClusterListResponseData]
type apiClusterListResponseDataJSON struct {
	ID          apijson.Field
	RootVmID    apijson.Field
	VmCount     apijson.Field
	Vms         apijson.Field
	Alias       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIClusterListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterListResponseDataJSON) RawJSON() string {
	return r.raw
}

type APIClusterNewParams struct {
	Create CreateUnionParam `json:"create,required"`
}

func (r APIClusterNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Create)
}
