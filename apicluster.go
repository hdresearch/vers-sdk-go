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
func (r *APIClusterService) Get(ctx context.Context, clusterIDOrAlias string, opts ...option.RequestOption) (res *APIClusterGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if clusterIDOrAlias == "" {
		err = errors.New("missing required cluster_id_or_alias parameter")
		return
	}
	path := fmt.Sprintf("api/cluster/%s", clusterIDOrAlias)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a cluster's configuration
func (r *APIClusterService) Update(ctx context.Context, clusterIDOrAlias string, body APIClusterUpdateParams, opts ...option.RequestOption) (res *APIClusterUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if clusterIDOrAlias == "" {
		err = errors.New("missing required cluster_id_or_alias parameter")
		return
	}
	path := fmt.Sprintf("api/cluster/%s", clusterIDOrAlias)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
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
func (r *APIClusterService) Delete(ctx context.Context, clusterIDOrAlias string, opts ...option.RequestOption) (res *APIClusterDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if clusterIDOrAlias == "" {
		err = errors.New("missing required cluster_id_or_alias parameter")
		return
	}
	path := fmt.Sprintf("api/cluster/%s", clusterIDOrAlias)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get the SSH private key for VM access
func (r *APIClusterService) GetSSHKey(ctx context.Context, clusterIDOrAlias string, opts ...option.RequestOption) (res *APIClusterGetSSHKeyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if clusterIDOrAlias == "" {
		err = errors.New("missing required cluster_id_or_alias parameter")
		return
	}
	path := fmt.Sprintf("api/cluster/%s/ssh_key", clusterIDOrAlias)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ClusterCreateParams struct {
	ClusterType param.Field[ClusterCreateParamsClusterType] `json:"cluster_type,required"`
	Params      param.Field[interface{}]                    `json:"params,required"`
}

func (r ClusterCreateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ClusterCreateParams) implementsClusterCreateParamsUnion() {}

// Satisfied by [ClusterCreateParamsNewClusterParams],
// [ClusterCreateParamsClusterFromCommitParams], [ClusterCreateParams].
type ClusterCreateParamsUnion interface {
	implementsClusterCreateParamsUnion()
}

type ClusterCreateParamsNewClusterParams struct {
	ClusterType param.Field[ClusterCreateParamsNewClusterParamsClusterType] `json:"cluster_type,required"`
	Params      param.Field[ClusterCreateParamsNewClusterParamsParams]      `json:"params,required"`
}

func (r ClusterCreateParamsNewClusterParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ClusterCreateParamsNewClusterParams) implementsClusterCreateParamsUnion() {}

type ClusterCreateParamsNewClusterParamsClusterType string

const (
	ClusterCreateParamsNewClusterParamsClusterTypeNew ClusterCreateParamsNewClusterParamsClusterType = "new"
)

func (r ClusterCreateParamsNewClusterParamsClusterType) IsKnown() bool {
	switch r {
	case ClusterCreateParamsNewClusterParamsClusterTypeNew:
		return true
	}
	return false
}

type ClusterCreateParamsNewClusterParamsParams struct {
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

func (r ClusterCreateParamsNewClusterParamsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ClusterCreateParamsClusterFromCommitParams struct {
	ClusterType param.Field[ClusterCreateParamsClusterFromCommitParamsClusterType] `json:"cluster_type,required"`
	Params      param.Field[ClusterCreateParamsClusterFromCommitParamsParams]      `json:"params,required"`
}

func (r ClusterCreateParamsClusterFromCommitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ClusterCreateParamsClusterFromCommitParams) implementsClusterCreateParamsUnion() {}

type ClusterCreateParamsClusterFromCommitParamsClusterType string

const (
	ClusterCreateParamsClusterFromCommitParamsClusterTypeFromCommit ClusterCreateParamsClusterFromCommitParamsClusterType = "from_commit"
)

func (r ClusterCreateParamsClusterFromCommitParamsClusterType) IsKnown() bool {
	switch r {
	case ClusterCreateParamsClusterFromCommitParamsClusterTypeFromCommit:
		return true
	}
	return false
}

type ClusterCreateParamsClusterFromCommitParamsParams struct {
	CommitKey        param.Field[string] `json:"commit_key,required"`
	ClusterAlias     param.Field[string] `json:"cluster_alias"`
	FsSizeClusterMib param.Field[int64]  `json:"fs_size_cluster_mib"`
	VmAlias          param.Field[string] `json:"vm_alias"`
}

func (r ClusterCreateParamsClusterFromCommitParamsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ClusterCreateParamsClusterType string

const (
	ClusterCreateParamsClusterTypeNew        ClusterCreateParamsClusterType = "new"
	ClusterCreateParamsClusterTypeFromCommit ClusterCreateParamsClusterType = "from_commit"
)

func (r ClusterCreateParamsClusterType) IsKnown() bool {
	switch r {
	case ClusterCreateParamsClusterTypeNew, ClusterCreateParamsClusterTypeFromCommit:
		return true
	}
	return false
}

type ClusterPatchParams struct {
	Alias param.Field[string] `json:"alias"`
}

func (r ClusterPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type APIClusterNewResponse struct {
	Data          APIClusterNewResponseData          `json:"data,required"`
	DurationNs    int64                              `json:"duration_ns,required"`
	OperationCode APIClusterNewResponseOperationCode `json:"operation_code,required"`
	OperationID   string                             `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                     `json:"time_start,required"`
	JSON      apiClusterNewResponseJSON `json:"-"`
}

// apiClusterNewResponseJSON contains the JSON metadata for the struct
// [APIClusterNewResponse]
type apiClusterNewResponseJSON struct {
	Data          apijson.Field
	DurationNs    apijson.Field
	OperationCode apijson.Field
	OperationID   apijson.Field
	TimeStart     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
	// The size of the cluster's backing file
	FsSizeMib int64 `json:"fs_size_mib,required"`
	// The ID of the cluster's root VM.
	RootVmID string `json:"root_vm_id,required"`
	// How many VMs are currently running on this cluster.
	VmCount int64 `json:"vm_count,required"`
	// The VMs that are children of the cluster, including the root VM.
	Vms []VmDto `json:"vms,required"`
	// Human-readable name assigned to the cluster.
	Alias string                        `json:"alias,nullable"`
	JSON  apiClusterNewResponseDataJSON `json:"-"`
}

// apiClusterNewResponseDataJSON contains the JSON metadata for the struct
// [APIClusterNewResponseData]
type apiClusterNewResponseDataJSON struct {
	ID          apijson.Field
	FsSizeMib   apijson.Field
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

type APIClusterNewResponseOperationCode string

const (
	APIClusterNewResponseOperationCodeListClusters     APIClusterNewResponseOperationCode = "list_clusters"
	APIClusterNewResponseOperationCodeGetCluster       APIClusterNewResponseOperationCode = "get_cluster"
	APIClusterNewResponseOperationCodeCreateCluster    APIClusterNewResponseOperationCode = "create_cluster"
	APIClusterNewResponseOperationCodeDeleteCluster    APIClusterNewResponseOperationCode = "delete_cluster"
	APIClusterNewResponseOperationCodeUpdateCluster    APIClusterNewResponseOperationCode = "update_cluster"
	APIClusterNewResponseOperationCodeGetClusterSSHKey APIClusterNewResponseOperationCode = "get_cluster_ssh_key"
	APIClusterNewResponseOperationCodeListVms          APIClusterNewResponseOperationCode = "list_vms"
	APIClusterNewResponseOperationCodeGetVm            APIClusterNewResponseOperationCode = "get_vm"
	APIClusterNewResponseOperationCodeUpdateVm         APIClusterNewResponseOperationCode = "update_vm"
	APIClusterNewResponseOperationCodeBranchVm         APIClusterNewResponseOperationCode = "branch_vm"
	APIClusterNewResponseOperationCodeCommitVm         APIClusterNewResponseOperationCode = "commit_vm"
	APIClusterNewResponseOperationCodeDeleteVm         APIClusterNewResponseOperationCode = "delete_vm"
	APIClusterNewResponseOperationCodeGetVmSSHKey      APIClusterNewResponseOperationCode = "get_vm_ssh_key"
	APIClusterNewResponseOperationCodeUploadRootfs     APIClusterNewResponseOperationCode = "upload_rootfs"
	APIClusterNewResponseOperationCodeDeleteRootfs     APIClusterNewResponseOperationCode = "delete_rootfs"
	APIClusterNewResponseOperationCodeListRootfs       APIClusterNewResponseOperationCode = "list_rootfs"
)

func (r APIClusterNewResponseOperationCode) IsKnown() bool {
	switch r {
	case APIClusterNewResponseOperationCodeListClusters, APIClusterNewResponseOperationCodeGetCluster, APIClusterNewResponseOperationCodeCreateCluster, APIClusterNewResponseOperationCodeDeleteCluster, APIClusterNewResponseOperationCodeUpdateCluster, APIClusterNewResponseOperationCodeGetClusterSSHKey, APIClusterNewResponseOperationCodeListVms, APIClusterNewResponseOperationCodeGetVm, APIClusterNewResponseOperationCodeUpdateVm, APIClusterNewResponseOperationCodeBranchVm, APIClusterNewResponseOperationCodeCommitVm, APIClusterNewResponseOperationCodeDeleteVm, APIClusterNewResponseOperationCodeGetVmSSHKey, APIClusterNewResponseOperationCodeUploadRootfs, APIClusterNewResponseOperationCodeDeleteRootfs, APIClusterNewResponseOperationCodeListRootfs:
		return true
	}
	return false
}

type APIClusterGetResponse struct {
	Data          APIClusterGetResponseData          `json:"data,required"`
	DurationNs    int64                              `json:"duration_ns,required"`
	OperationCode APIClusterGetResponseOperationCode `json:"operation_code,required"`
	OperationID   string                             `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                     `json:"time_start,required"`
	JSON      apiClusterGetResponseJSON `json:"-"`
}

// apiClusterGetResponseJSON contains the JSON metadata for the struct
// [APIClusterGetResponse]
type apiClusterGetResponseJSON struct {
	Data          apijson.Field
	DurationNs    apijson.Field
	OperationCode apijson.Field
	OperationID   apijson.Field
	TimeStart     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
	// The size of the cluster's backing file
	FsSizeMib int64 `json:"fs_size_mib,required"`
	// The ID of the cluster's root VM.
	RootVmID string `json:"root_vm_id,required"`
	// How many VMs are currently running on this cluster.
	VmCount int64 `json:"vm_count,required"`
	// The VMs that are children of the cluster, including the root VM.
	Vms []VmDto `json:"vms,required"`
	// Human-readable name assigned to the cluster.
	Alias string                        `json:"alias,nullable"`
	JSON  apiClusterGetResponseDataJSON `json:"-"`
}

// apiClusterGetResponseDataJSON contains the JSON metadata for the struct
// [APIClusterGetResponseData]
type apiClusterGetResponseDataJSON struct {
	ID          apijson.Field
	FsSizeMib   apijson.Field
	RootVmID    apijson.Field
	VmCount     apijson.Field
	Vms         apijson.Field
	Alias       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIClusterGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type APIClusterGetResponseOperationCode string

const (
	APIClusterGetResponseOperationCodeListClusters     APIClusterGetResponseOperationCode = "list_clusters"
	APIClusterGetResponseOperationCodeGetCluster       APIClusterGetResponseOperationCode = "get_cluster"
	APIClusterGetResponseOperationCodeCreateCluster    APIClusterGetResponseOperationCode = "create_cluster"
	APIClusterGetResponseOperationCodeDeleteCluster    APIClusterGetResponseOperationCode = "delete_cluster"
	APIClusterGetResponseOperationCodeUpdateCluster    APIClusterGetResponseOperationCode = "update_cluster"
	APIClusterGetResponseOperationCodeGetClusterSSHKey APIClusterGetResponseOperationCode = "get_cluster_ssh_key"
	APIClusterGetResponseOperationCodeListVms          APIClusterGetResponseOperationCode = "list_vms"
	APIClusterGetResponseOperationCodeGetVm            APIClusterGetResponseOperationCode = "get_vm"
	APIClusterGetResponseOperationCodeUpdateVm         APIClusterGetResponseOperationCode = "update_vm"
	APIClusterGetResponseOperationCodeBranchVm         APIClusterGetResponseOperationCode = "branch_vm"
	APIClusterGetResponseOperationCodeCommitVm         APIClusterGetResponseOperationCode = "commit_vm"
	APIClusterGetResponseOperationCodeDeleteVm         APIClusterGetResponseOperationCode = "delete_vm"
	APIClusterGetResponseOperationCodeGetVmSSHKey      APIClusterGetResponseOperationCode = "get_vm_ssh_key"
	APIClusterGetResponseOperationCodeUploadRootfs     APIClusterGetResponseOperationCode = "upload_rootfs"
	APIClusterGetResponseOperationCodeDeleteRootfs     APIClusterGetResponseOperationCode = "delete_rootfs"
	APIClusterGetResponseOperationCodeListRootfs       APIClusterGetResponseOperationCode = "list_rootfs"
)

func (r APIClusterGetResponseOperationCode) IsKnown() bool {
	switch r {
	case APIClusterGetResponseOperationCodeListClusters, APIClusterGetResponseOperationCodeGetCluster, APIClusterGetResponseOperationCodeCreateCluster, APIClusterGetResponseOperationCodeDeleteCluster, APIClusterGetResponseOperationCodeUpdateCluster, APIClusterGetResponseOperationCodeGetClusterSSHKey, APIClusterGetResponseOperationCodeListVms, APIClusterGetResponseOperationCodeGetVm, APIClusterGetResponseOperationCodeUpdateVm, APIClusterGetResponseOperationCodeBranchVm, APIClusterGetResponseOperationCodeCommitVm, APIClusterGetResponseOperationCodeDeleteVm, APIClusterGetResponseOperationCodeGetVmSSHKey, APIClusterGetResponseOperationCodeUploadRootfs, APIClusterGetResponseOperationCodeDeleteRootfs, APIClusterGetResponseOperationCodeListRootfs:
		return true
	}
	return false
}

type APIClusterUpdateResponse struct {
	Data          APIClusterUpdateResponseData          `json:"data,required"`
	DurationNs    int64                                 `json:"duration_ns,required"`
	OperationCode APIClusterUpdateResponseOperationCode `json:"operation_code,required"`
	OperationID   string                                `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                        `json:"time_start,required"`
	JSON      apiClusterUpdateResponseJSON `json:"-"`
}

// apiClusterUpdateResponseJSON contains the JSON metadata for the struct
// [APIClusterUpdateResponse]
type apiClusterUpdateResponseJSON struct {
	Data          apijson.Field
	DurationNs    apijson.Field
	OperationCode apijson.Field
	OperationID   apijson.Field
	TimeStart     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *APIClusterUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type APIClusterUpdateResponseData struct {
	// The cluster's ID.
	ID string `json:"id,required"`
	// The size of the cluster's backing file
	FsSizeMib int64 `json:"fs_size_mib,required"`
	// The ID of the cluster's root VM.
	RootVmID string `json:"root_vm_id,required"`
	// How many VMs are currently running on this cluster.
	VmCount int64 `json:"vm_count,required"`
	// The VMs that are children of the cluster, including the root VM.
	Vms []VmDto `json:"vms,required"`
	// Human-readable name assigned to the cluster.
	Alias string                           `json:"alias,nullable"`
	JSON  apiClusterUpdateResponseDataJSON `json:"-"`
}

// apiClusterUpdateResponseDataJSON contains the JSON metadata for the struct
// [APIClusterUpdateResponseData]
type apiClusterUpdateResponseDataJSON struct {
	ID          apijson.Field
	FsSizeMib   apijson.Field
	RootVmID    apijson.Field
	VmCount     apijson.Field
	Vms         apijson.Field
	Alias       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIClusterUpdateResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterUpdateResponseDataJSON) RawJSON() string {
	return r.raw
}

type APIClusterUpdateResponseOperationCode string

const (
	APIClusterUpdateResponseOperationCodeListClusters     APIClusterUpdateResponseOperationCode = "list_clusters"
	APIClusterUpdateResponseOperationCodeGetCluster       APIClusterUpdateResponseOperationCode = "get_cluster"
	APIClusterUpdateResponseOperationCodeCreateCluster    APIClusterUpdateResponseOperationCode = "create_cluster"
	APIClusterUpdateResponseOperationCodeDeleteCluster    APIClusterUpdateResponseOperationCode = "delete_cluster"
	APIClusterUpdateResponseOperationCodeUpdateCluster    APIClusterUpdateResponseOperationCode = "update_cluster"
	APIClusterUpdateResponseOperationCodeGetClusterSSHKey APIClusterUpdateResponseOperationCode = "get_cluster_ssh_key"
	APIClusterUpdateResponseOperationCodeListVms          APIClusterUpdateResponseOperationCode = "list_vms"
	APIClusterUpdateResponseOperationCodeGetVm            APIClusterUpdateResponseOperationCode = "get_vm"
	APIClusterUpdateResponseOperationCodeUpdateVm         APIClusterUpdateResponseOperationCode = "update_vm"
	APIClusterUpdateResponseOperationCodeBranchVm         APIClusterUpdateResponseOperationCode = "branch_vm"
	APIClusterUpdateResponseOperationCodeCommitVm         APIClusterUpdateResponseOperationCode = "commit_vm"
	APIClusterUpdateResponseOperationCodeDeleteVm         APIClusterUpdateResponseOperationCode = "delete_vm"
	APIClusterUpdateResponseOperationCodeGetVmSSHKey      APIClusterUpdateResponseOperationCode = "get_vm_ssh_key"
	APIClusterUpdateResponseOperationCodeUploadRootfs     APIClusterUpdateResponseOperationCode = "upload_rootfs"
	APIClusterUpdateResponseOperationCodeDeleteRootfs     APIClusterUpdateResponseOperationCode = "delete_rootfs"
	APIClusterUpdateResponseOperationCodeListRootfs       APIClusterUpdateResponseOperationCode = "list_rootfs"
)

func (r APIClusterUpdateResponseOperationCode) IsKnown() bool {
	switch r {
	case APIClusterUpdateResponseOperationCodeListClusters, APIClusterUpdateResponseOperationCodeGetCluster, APIClusterUpdateResponseOperationCodeCreateCluster, APIClusterUpdateResponseOperationCodeDeleteCluster, APIClusterUpdateResponseOperationCodeUpdateCluster, APIClusterUpdateResponseOperationCodeGetClusterSSHKey, APIClusterUpdateResponseOperationCodeListVms, APIClusterUpdateResponseOperationCodeGetVm, APIClusterUpdateResponseOperationCodeUpdateVm, APIClusterUpdateResponseOperationCodeBranchVm, APIClusterUpdateResponseOperationCodeCommitVm, APIClusterUpdateResponseOperationCodeDeleteVm, APIClusterUpdateResponseOperationCodeGetVmSSHKey, APIClusterUpdateResponseOperationCodeUploadRootfs, APIClusterUpdateResponseOperationCodeDeleteRootfs, APIClusterUpdateResponseOperationCodeListRootfs:
		return true
	}
	return false
}

type APIClusterListResponse struct {
	Data          []APIClusterListResponseData        `json:"data,required"`
	DurationNs    int64                               `json:"duration_ns,required"`
	OperationCode APIClusterListResponseOperationCode `json:"operation_code,required"`
	OperationID   string                              `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                      `json:"time_start,required"`
	JSON      apiClusterListResponseJSON `json:"-"`
}

// apiClusterListResponseJSON contains the JSON metadata for the struct
// [APIClusterListResponse]
type apiClusterListResponseJSON struct {
	Data          apijson.Field
	DurationNs    apijson.Field
	OperationCode apijson.Field
	OperationID   apijson.Field
	TimeStart     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
	// The size of the cluster's backing file
	FsSizeMib int64 `json:"fs_size_mib,required"`
	// The ID of the cluster's root VM.
	RootVmID string `json:"root_vm_id,required"`
	// How many VMs are currently running on this cluster.
	VmCount int64 `json:"vm_count,required"`
	// The VMs that are children of the cluster, including the root VM.
	Vms []VmDto `json:"vms,required"`
	// Human-readable name assigned to the cluster.
	Alias string                         `json:"alias,nullable"`
	JSON  apiClusterListResponseDataJSON `json:"-"`
}

// apiClusterListResponseDataJSON contains the JSON metadata for the struct
// [APIClusterListResponseData]
type apiClusterListResponseDataJSON struct {
	ID          apijson.Field
	FsSizeMib   apijson.Field
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

type APIClusterListResponseOperationCode string

const (
	APIClusterListResponseOperationCodeListClusters     APIClusterListResponseOperationCode = "list_clusters"
	APIClusterListResponseOperationCodeGetCluster       APIClusterListResponseOperationCode = "get_cluster"
	APIClusterListResponseOperationCodeCreateCluster    APIClusterListResponseOperationCode = "create_cluster"
	APIClusterListResponseOperationCodeDeleteCluster    APIClusterListResponseOperationCode = "delete_cluster"
	APIClusterListResponseOperationCodeUpdateCluster    APIClusterListResponseOperationCode = "update_cluster"
	APIClusterListResponseOperationCodeGetClusterSSHKey APIClusterListResponseOperationCode = "get_cluster_ssh_key"
	APIClusterListResponseOperationCodeListVms          APIClusterListResponseOperationCode = "list_vms"
	APIClusterListResponseOperationCodeGetVm            APIClusterListResponseOperationCode = "get_vm"
	APIClusterListResponseOperationCodeUpdateVm         APIClusterListResponseOperationCode = "update_vm"
	APIClusterListResponseOperationCodeBranchVm         APIClusterListResponseOperationCode = "branch_vm"
	APIClusterListResponseOperationCodeCommitVm         APIClusterListResponseOperationCode = "commit_vm"
	APIClusterListResponseOperationCodeDeleteVm         APIClusterListResponseOperationCode = "delete_vm"
	APIClusterListResponseOperationCodeGetVmSSHKey      APIClusterListResponseOperationCode = "get_vm_ssh_key"
	APIClusterListResponseOperationCodeUploadRootfs     APIClusterListResponseOperationCode = "upload_rootfs"
	APIClusterListResponseOperationCodeDeleteRootfs     APIClusterListResponseOperationCode = "delete_rootfs"
	APIClusterListResponseOperationCodeListRootfs       APIClusterListResponseOperationCode = "list_rootfs"
)

func (r APIClusterListResponseOperationCode) IsKnown() bool {
	switch r {
	case APIClusterListResponseOperationCodeListClusters, APIClusterListResponseOperationCodeGetCluster, APIClusterListResponseOperationCodeCreateCluster, APIClusterListResponseOperationCodeDeleteCluster, APIClusterListResponseOperationCodeUpdateCluster, APIClusterListResponseOperationCodeGetClusterSSHKey, APIClusterListResponseOperationCodeListVms, APIClusterListResponseOperationCodeGetVm, APIClusterListResponseOperationCodeUpdateVm, APIClusterListResponseOperationCodeBranchVm, APIClusterListResponseOperationCodeCommitVm, APIClusterListResponseOperationCodeDeleteVm, APIClusterListResponseOperationCodeGetVmSSHKey, APIClusterListResponseOperationCodeUploadRootfs, APIClusterListResponseOperationCodeDeleteRootfs, APIClusterListResponseOperationCodeListRootfs:
		return true
	}
	return false
}

type APIClusterDeleteResponse struct {
	// A struct containing information about an attempted cluster deletion request.
	// Reports information in the event of a partial failure so billing can still be
	// udpated appropriately.
	Data          APIClusterDeleteResponseData          `json:"data,required"`
	DurationNs    int64                                 `json:"duration_ns,required"`
	OperationCode APIClusterDeleteResponseOperationCode `json:"operation_code,required"`
	OperationID   string                                `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                        `json:"time_start,required"`
	JSON      apiClusterDeleteResponseJSON `json:"-"`
}

// apiClusterDeleteResponseJSON contains the JSON metadata for the struct
// [APIClusterDeleteResponse]
type apiClusterDeleteResponseJSON struct {
	Data          apijson.Field
	DurationNs    apijson.Field
	OperationCode apijson.Field
	OperationID   apijson.Field
	TimeStart     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *APIClusterDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// A struct containing information about an attempted cluster deletion request.
// Reports information in the event of a partial failure so billing can still be
// udpated appropriately.
type APIClusterDeleteResponseData struct {
	ClusterID string `json:"cluster_id,required"`
	// A struct containing information about an attempted VM deletion request. Reports
	// information in the event of a partial failure so billing can still be udpated
	// appropriately.
	Vms     VmDeleteResponse                 `json:"vms,required"`
	FsError string                           `json:"fs_error,nullable"`
	JSON    apiClusterDeleteResponseDataJSON `json:"-"`
}

// apiClusterDeleteResponseDataJSON contains the JSON metadata for the struct
// [APIClusterDeleteResponseData]
type apiClusterDeleteResponseDataJSON struct {
	ClusterID   apijson.Field
	Vms         apijson.Field
	FsError     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIClusterDeleteResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterDeleteResponseDataJSON) RawJSON() string {
	return r.raw
}

type APIClusterDeleteResponseOperationCode string

const (
	APIClusterDeleteResponseOperationCodeListClusters     APIClusterDeleteResponseOperationCode = "list_clusters"
	APIClusterDeleteResponseOperationCodeGetCluster       APIClusterDeleteResponseOperationCode = "get_cluster"
	APIClusterDeleteResponseOperationCodeCreateCluster    APIClusterDeleteResponseOperationCode = "create_cluster"
	APIClusterDeleteResponseOperationCodeDeleteCluster    APIClusterDeleteResponseOperationCode = "delete_cluster"
	APIClusterDeleteResponseOperationCodeUpdateCluster    APIClusterDeleteResponseOperationCode = "update_cluster"
	APIClusterDeleteResponseOperationCodeGetClusterSSHKey APIClusterDeleteResponseOperationCode = "get_cluster_ssh_key"
	APIClusterDeleteResponseOperationCodeListVms          APIClusterDeleteResponseOperationCode = "list_vms"
	APIClusterDeleteResponseOperationCodeGetVm            APIClusterDeleteResponseOperationCode = "get_vm"
	APIClusterDeleteResponseOperationCodeUpdateVm         APIClusterDeleteResponseOperationCode = "update_vm"
	APIClusterDeleteResponseOperationCodeBranchVm         APIClusterDeleteResponseOperationCode = "branch_vm"
	APIClusterDeleteResponseOperationCodeCommitVm         APIClusterDeleteResponseOperationCode = "commit_vm"
	APIClusterDeleteResponseOperationCodeDeleteVm         APIClusterDeleteResponseOperationCode = "delete_vm"
	APIClusterDeleteResponseOperationCodeGetVmSSHKey      APIClusterDeleteResponseOperationCode = "get_vm_ssh_key"
	APIClusterDeleteResponseOperationCodeUploadRootfs     APIClusterDeleteResponseOperationCode = "upload_rootfs"
	APIClusterDeleteResponseOperationCodeDeleteRootfs     APIClusterDeleteResponseOperationCode = "delete_rootfs"
	APIClusterDeleteResponseOperationCodeListRootfs       APIClusterDeleteResponseOperationCode = "list_rootfs"
)

func (r APIClusterDeleteResponseOperationCode) IsKnown() bool {
	switch r {
	case APIClusterDeleteResponseOperationCodeListClusters, APIClusterDeleteResponseOperationCodeGetCluster, APIClusterDeleteResponseOperationCodeCreateCluster, APIClusterDeleteResponseOperationCodeDeleteCluster, APIClusterDeleteResponseOperationCodeUpdateCluster, APIClusterDeleteResponseOperationCodeGetClusterSSHKey, APIClusterDeleteResponseOperationCodeListVms, APIClusterDeleteResponseOperationCodeGetVm, APIClusterDeleteResponseOperationCodeUpdateVm, APIClusterDeleteResponseOperationCodeBranchVm, APIClusterDeleteResponseOperationCodeCommitVm, APIClusterDeleteResponseOperationCodeDeleteVm, APIClusterDeleteResponseOperationCodeGetVmSSHKey, APIClusterDeleteResponseOperationCodeUploadRootfs, APIClusterDeleteResponseOperationCodeDeleteRootfs, APIClusterDeleteResponseOperationCodeListRootfs:
		return true
	}
	return false
}

type APIClusterGetSSHKeyResponse struct {
	Data          string                                   `json:"data,required"`
	DurationNs    int64                                    `json:"duration_ns,required"`
	OperationCode APIClusterGetSSHKeyResponseOperationCode `json:"operation_code,required"`
	OperationID   string                                   `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                           `json:"time_start,required"`
	JSON      apiClusterGetSSHKeyResponseJSON `json:"-"`
}

// apiClusterGetSSHKeyResponseJSON contains the JSON metadata for the struct
// [APIClusterGetSSHKeyResponse]
type apiClusterGetSSHKeyResponseJSON struct {
	Data          apijson.Field
	DurationNs    apijson.Field
	OperationCode apijson.Field
	OperationID   apijson.Field
	TimeStart     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *APIClusterGetSSHKeyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterGetSSHKeyResponseJSON) RawJSON() string {
	return r.raw
}

type APIClusterGetSSHKeyResponseOperationCode string

const (
	APIClusterGetSSHKeyResponseOperationCodeListClusters     APIClusterGetSSHKeyResponseOperationCode = "list_clusters"
	APIClusterGetSSHKeyResponseOperationCodeGetCluster       APIClusterGetSSHKeyResponseOperationCode = "get_cluster"
	APIClusterGetSSHKeyResponseOperationCodeCreateCluster    APIClusterGetSSHKeyResponseOperationCode = "create_cluster"
	APIClusterGetSSHKeyResponseOperationCodeDeleteCluster    APIClusterGetSSHKeyResponseOperationCode = "delete_cluster"
	APIClusterGetSSHKeyResponseOperationCodeUpdateCluster    APIClusterGetSSHKeyResponseOperationCode = "update_cluster"
	APIClusterGetSSHKeyResponseOperationCodeGetClusterSSHKey APIClusterGetSSHKeyResponseOperationCode = "get_cluster_ssh_key"
	APIClusterGetSSHKeyResponseOperationCodeListVms          APIClusterGetSSHKeyResponseOperationCode = "list_vms"
	APIClusterGetSSHKeyResponseOperationCodeGetVm            APIClusterGetSSHKeyResponseOperationCode = "get_vm"
	APIClusterGetSSHKeyResponseOperationCodeUpdateVm         APIClusterGetSSHKeyResponseOperationCode = "update_vm"
	APIClusterGetSSHKeyResponseOperationCodeBranchVm         APIClusterGetSSHKeyResponseOperationCode = "branch_vm"
	APIClusterGetSSHKeyResponseOperationCodeCommitVm         APIClusterGetSSHKeyResponseOperationCode = "commit_vm"
	APIClusterGetSSHKeyResponseOperationCodeDeleteVm         APIClusterGetSSHKeyResponseOperationCode = "delete_vm"
	APIClusterGetSSHKeyResponseOperationCodeGetVmSSHKey      APIClusterGetSSHKeyResponseOperationCode = "get_vm_ssh_key"
	APIClusterGetSSHKeyResponseOperationCodeUploadRootfs     APIClusterGetSSHKeyResponseOperationCode = "upload_rootfs"
	APIClusterGetSSHKeyResponseOperationCodeDeleteRootfs     APIClusterGetSSHKeyResponseOperationCode = "delete_rootfs"
	APIClusterGetSSHKeyResponseOperationCodeListRootfs       APIClusterGetSSHKeyResponseOperationCode = "list_rootfs"
)

func (r APIClusterGetSSHKeyResponseOperationCode) IsKnown() bool {
	switch r {
	case APIClusterGetSSHKeyResponseOperationCodeListClusters, APIClusterGetSSHKeyResponseOperationCodeGetCluster, APIClusterGetSSHKeyResponseOperationCodeCreateCluster, APIClusterGetSSHKeyResponseOperationCodeDeleteCluster, APIClusterGetSSHKeyResponseOperationCodeUpdateCluster, APIClusterGetSSHKeyResponseOperationCodeGetClusterSSHKey, APIClusterGetSSHKeyResponseOperationCodeListVms, APIClusterGetSSHKeyResponseOperationCodeGetVm, APIClusterGetSSHKeyResponseOperationCodeUpdateVm, APIClusterGetSSHKeyResponseOperationCodeBranchVm, APIClusterGetSSHKeyResponseOperationCodeCommitVm, APIClusterGetSSHKeyResponseOperationCodeDeleteVm, APIClusterGetSSHKeyResponseOperationCodeGetVmSSHKey, APIClusterGetSSHKeyResponseOperationCodeUploadRootfs, APIClusterGetSSHKeyResponseOperationCodeDeleteRootfs, APIClusterGetSSHKeyResponseOperationCodeListRootfs:
		return true
	}
	return false
}

type APIClusterNewParams struct {
	ClusterCreateParams ClusterCreateParamsUnion `json:"cluster_create_params,required"`
}

func (r APIClusterNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ClusterCreateParams)
}

type APIClusterUpdateParams struct {
	ClusterPatchParams ClusterPatchParams `json:"cluster_patch_params,required"`
}

func (r APIClusterUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ClusterPatchParams)
}
