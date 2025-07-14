// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/hdresearch/vers-sdk-go/internal/apijson"
	"github.com/hdresearch/vers-sdk-go/internal/apiquery"
	"github.com/hdresearch/vers-sdk-go/internal/param"
	"github.com/hdresearch/vers-sdk-go/internal/requestconfig"
	"github.com/hdresearch/vers-sdk-go/option"
)

// APIVmService contains methods and other services that help with interacting with
// the vers API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIVmService] method instead.
type APIVmService struct {
	Options []option.RequestOption
}

// NewAPIVmService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAPIVmService(opts ...option.RequestOption) (r *APIVmService) {
	r = &APIVmService{}
	r.Options = opts
	return
}

// Retrieve information on a particular VM.
func (r *APIVmService) Get(ctx context.Context, vmIDOrAlias string, opts ...option.RequestOption) (res *APIVmGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if vmIDOrAlias == "" {
		err = errors.New("missing required vm_id_or_alias parameter")
		return
	}
	path := fmt.Sprintf("api/vm/%s", vmIDOrAlias)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update VM state.
func (r *APIVmService) Update(ctx context.Context, vmIDOrAlias string, body APIVmUpdateParams, opts ...option.RequestOption) (res *APIVmUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if vmIDOrAlias == "" {
		err = errors.New("missing required vm_id_or_alias parameter")
		return
	}
	path := fmt.Sprintf("api/vm/%s", vmIDOrAlias)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all VMs.
func (r *APIVmService) List(ctx context.Context, opts ...option.RequestOption) (res *APIVmListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/vm"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a VM.
func (r *APIVmService) Delete(ctx context.Context, vmIDOrAlias string, body APIVmDeleteParams, opts ...option.RequestOption) (res *APIVmDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if vmIDOrAlias == "" {
		err = errors.New("missing required vm_id_or_alias parameter")
		return
	}
	path := fmt.Sprintf("api/vm/%s", vmIDOrAlias)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Branch a VM.
func (r *APIVmService) Branch(ctx context.Context, vmIDOrAlias string, body APIVmBranchParams, opts ...option.RequestOption) (res *APIVmBranchResponse, err error) {
	opts = append(r.Options[:], opts...)
	if vmIDOrAlias == "" {
		err = errors.New("missing required vm_id_or_alias parameter")
		return
	}
	path := fmt.Sprintf("api/vm/%s/branch", vmIDOrAlias)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Commit a VM.
func (r *APIVmService) Commit(ctx context.Context, vmIDOrAlias string, body APIVmCommitParams, opts ...option.RequestOption) (res *APIVmCommitResponse, err error) {
	opts = append(r.Options[:], opts...)
	if vmIDOrAlias == "" {
		err = errors.New("missing required vm_id_or_alias parameter")
		return
	}
	path := fmt.Sprintf("api/vm/%s/commit", vmIDOrAlias)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the SSH private key for VM access
func (r *APIVmService) GetSSHKey(ctx context.Context, vmIDOrAlias string, opts ...option.RequestOption) (res *APIVmGetSSHKeyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if vmIDOrAlias == "" {
		err = errors.New("missing required vm_id_or_alias parameter")
		return
	}
	path := fmt.Sprintf("api/vm/%s/ssh_key", vmIDOrAlias)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type VmBranchParams struct {
	Alias param.Field[string] `json:"alias"`
}

func (r VmBranchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VmCommitRequestParam struct {
	Tags param.Field[[]string] `json:"tags"`
}

func (r VmCommitRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A struct containing information about an attempted VM deletion request. Reports
// information in the event of a partial failure so billing can still be udpated
// appropriately.
type VmDeleteResponse struct {
	DeletedIDs []string                `json:"deleted_ids,required"`
	Errors     []VmDeleteResponseError `json:"errors,required"`
	JSON       vmDeleteResponseJSON    `json:"-"`
}

// vmDeleteResponseJSON contains the JSON metadata for the struct
// [VmDeleteResponse]
type vmDeleteResponseJSON struct {
	DeletedIDs  apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VmDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vmDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Contains a VM ID and the reason that it could not be deleted.
type VmDeleteResponseError struct {
	ID    string                    `json:"id,required"`
	Error string                    `json:"error,required"`
	JSON  vmDeleteResponseErrorJSON `json:"-"`
}

// vmDeleteResponseErrorJSON contains the JSON metadata for the struct
// [VmDeleteResponseError]
type vmDeleteResponseErrorJSON struct {
	ID          apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VmDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vmDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type VmDto struct {
	// The ID of the VM.
	ID string `json:"id,required"`
	// The IDs of direct children branched from this VM.
	Children []string `json:"children,required"`
	// The VM's cluster ID
	ClusterID string `json:"cluster_id,required"`
	// What is the size of the "disk" allocated to this VM
	FsSizeMib int64 `json:"fs_size_mib,required"`
	// The VM's local IP address on the VM subnet
	IPAddress string `json:"ip_address,required"`
	// How much RAM is allocated to this VM
	MemSizeMib int64 `json:"mem_size_mib,required"`
	// The VM's network configuration
	NetworkInfo VmDtoNetworkInfo `json:"network_info,required"`
	// Whether the VM is running, paused, or not started.
	State VmDtoState `json:"state,required"`
	// How many vCPUs were allocated to this VM
	VcpuCount int64 `json:"vcpu_count,required"`
	// Human-readable name assigned to the VM.
	Alias string `json:"alias,nullable"`
	// The parent VM's ID, if present. If None, then this VM is a root VM.
	ParentID string    `json:"parent_id,nullable"`
	JSON     vmDtoJSON `json:"-"`
}

// vmDtoJSON contains the JSON metadata for the struct [VmDto]
type vmDtoJSON struct {
	ID          apijson.Field
	Children    apijson.Field
	ClusterID   apijson.Field
	FsSizeMib   apijson.Field
	IPAddress   apijson.Field
	MemSizeMib  apijson.Field
	NetworkInfo apijson.Field
	State       apijson.Field
	VcpuCount   apijson.Field
	Alias       apijson.Field
	ParentID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VmDto) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vmDtoJSON) RawJSON() string {
	return r.raw
}

// The VM's network configuration
type VmDtoNetworkInfo struct {
	GuestIP     string               `json:"guest_ip,required"`
	GuestMac    string               `json:"guest_mac,required"`
	SSHPort     int64                `json:"ssh_port,required"`
	Tap0IP      string               `json:"tap0_ip,required"`
	Tap0Name    string               `json:"tap0_name,required"`
	VmNamespace string               `json:"vm_namespace,required"`
	JSON        vmDtoNetworkInfoJSON `json:"-"`
}

// vmDtoNetworkInfoJSON contains the JSON metadata for the struct
// [VmDtoNetworkInfo]
type vmDtoNetworkInfoJSON struct {
	GuestIP     apijson.Field
	GuestMac    apijson.Field
	SSHPort     apijson.Field
	Tap0IP      apijson.Field
	Tap0Name    apijson.Field
	VmNamespace apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VmDtoNetworkInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vmDtoNetworkInfoJSON) RawJSON() string {
	return r.raw
}

// Whether the VM is running, paused, or not started.
type VmDtoState string

const (
	VmDtoStateNotStarted VmDtoState = "Not started"
	VmDtoStateRunning    VmDtoState = "Running"
	VmDtoStatePaused     VmDtoState = "Paused"
)

func (r VmDtoState) IsKnown() bool {
	switch r {
	case VmDtoStateNotStarted, VmDtoStateRunning, VmDtoStatePaused:
		return true
	}
	return false
}

type VmPatchParams struct {
	Alias param.Field[string]             `json:"alias"`
	State param.Field[VmPatchParamsState] `json:"state"`
}

func (r VmPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VmPatchParamsState string

const (
	VmPatchParamsStateRunning VmPatchParamsState = "Running"
	VmPatchParamsStatePaused  VmPatchParamsState = "Paused"
)

func (r VmPatchParamsState) IsKnown() bool {
	switch r {
	case VmPatchParamsStateRunning, VmPatchParamsStatePaused:
		return true
	}
	return false
}

type APIVmGetResponse struct {
	Data          APIVmGetResponseData          `json:"data,required"`
	DurationNs    int64                         `json:"duration_ns,required"`
	OperationCode APIVmGetResponseOperationCode `json:"operation_code,required"`
	OperationID   string                        `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                `json:"time_start,required"`
	JSON      apiVmGetResponseJSON `json:"-"`
}

// apiVmGetResponseJSON contains the JSON metadata for the struct
// [APIVmGetResponse]
type apiVmGetResponseJSON struct {
	Data          apijson.Field
	DurationNs    apijson.Field
	OperationCode apijson.Field
	OperationID   apijson.Field
	TimeStart     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *APIVmGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmGetResponseJSON) RawJSON() string {
	return r.raw
}

type APIVmGetResponseData struct {
	// The ID of the VM.
	ID string `json:"id,required"`
	// The IDs of direct children branched from this VM.
	Children []string `json:"children,required"`
	// The VM's cluster ID
	ClusterID string `json:"cluster_id,required"`
	// What is the size of the "disk" allocated to this VM
	FsSizeMib int64 `json:"fs_size_mib,required"`
	// The VM's local IP address on the VM subnet
	IPAddress string `json:"ip_address,required"`
	// How much RAM is allocated to this VM
	MemSizeMib int64 `json:"mem_size_mib,required"`
	// The VM's network configuration
	NetworkInfo APIVmGetResponseDataNetworkInfo `json:"network_info,required"`
	// Whether the VM is running, paused, or not started.
	State APIVmGetResponseDataState `json:"state,required"`
	// How many vCPUs were allocated to this VM
	VcpuCount int64 `json:"vcpu_count,required"`
	// Human-readable name assigned to the VM.
	Alias string `json:"alias,nullable"`
	// The parent VM's ID, if present. If None, then this VM is a root VM.
	ParentID string                   `json:"parent_id,nullable"`
	JSON     apiVmGetResponseDataJSON `json:"-"`
}

// apiVmGetResponseDataJSON contains the JSON metadata for the struct
// [APIVmGetResponseData]
type apiVmGetResponseDataJSON struct {
	ID          apijson.Field
	Children    apijson.Field
	ClusterID   apijson.Field
	FsSizeMib   apijson.Field
	IPAddress   apijson.Field
	MemSizeMib  apijson.Field
	NetworkInfo apijson.Field
	State       apijson.Field
	VcpuCount   apijson.Field
	Alias       apijson.Field
	ParentID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmGetResponseDataJSON) RawJSON() string {
	return r.raw
}

// The VM's network configuration
type APIVmGetResponseDataNetworkInfo struct {
	GuestIP     string                              `json:"guest_ip,required"`
	GuestMac    string                              `json:"guest_mac,required"`
	SSHPort     int64                               `json:"ssh_port,required"`
	Tap0IP      string                              `json:"tap0_ip,required"`
	Tap0Name    string                              `json:"tap0_name,required"`
	VmNamespace string                              `json:"vm_namespace,required"`
	JSON        apiVmGetResponseDataNetworkInfoJSON `json:"-"`
}

// apiVmGetResponseDataNetworkInfoJSON contains the JSON metadata for the struct
// [APIVmGetResponseDataNetworkInfo]
type apiVmGetResponseDataNetworkInfoJSON struct {
	GuestIP     apijson.Field
	GuestMac    apijson.Field
	SSHPort     apijson.Field
	Tap0IP      apijson.Field
	Tap0Name    apijson.Field
	VmNamespace apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmGetResponseDataNetworkInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmGetResponseDataNetworkInfoJSON) RawJSON() string {
	return r.raw
}

// Whether the VM is running, paused, or not started.
type APIVmGetResponseDataState string

const (
	APIVmGetResponseDataStateNotStarted APIVmGetResponseDataState = "Not started"
	APIVmGetResponseDataStateRunning    APIVmGetResponseDataState = "Running"
	APIVmGetResponseDataStatePaused     APIVmGetResponseDataState = "Paused"
)

func (r APIVmGetResponseDataState) IsKnown() bool {
	switch r {
	case APIVmGetResponseDataStateNotStarted, APIVmGetResponseDataStateRunning, APIVmGetResponseDataStatePaused:
		return true
	}
	return false
}

type APIVmGetResponseOperationCode string

const (
	APIVmGetResponseOperationCodeListClusters     APIVmGetResponseOperationCode = "list_clusters"
	APIVmGetResponseOperationCodeGetCluster       APIVmGetResponseOperationCode = "get_cluster"
	APIVmGetResponseOperationCodeCreateCluster    APIVmGetResponseOperationCode = "create_cluster"
	APIVmGetResponseOperationCodeDeleteCluster    APIVmGetResponseOperationCode = "delete_cluster"
	APIVmGetResponseOperationCodeUpdateCluster    APIVmGetResponseOperationCode = "update_cluster"
	APIVmGetResponseOperationCodeGetClusterSSHKey APIVmGetResponseOperationCode = "get_cluster_ssh_key"
	APIVmGetResponseOperationCodeListVms          APIVmGetResponseOperationCode = "list_vms"
	APIVmGetResponseOperationCodeGetVm            APIVmGetResponseOperationCode = "get_vm"
	APIVmGetResponseOperationCodeUpdateVm         APIVmGetResponseOperationCode = "update_vm"
	APIVmGetResponseOperationCodeBranchVm         APIVmGetResponseOperationCode = "branch_vm"
	APIVmGetResponseOperationCodeCommitVm         APIVmGetResponseOperationCode = "commit_vm"
	APIVmGetResponseOperationCodeDeleteVm         APIVmGetResponseOperationCode = "delete_vm"
	APIVmGetResponseOperationCodeGetVmSSHKey      APIVmGetResponseOperationCode = "get_vm_ssh_key"
	APIVmGetResponseOperationCodeUploadRootfs     APIVmGetResponseOperationCode = "upload_rootfs"
	APIVmGetResponseOperationCodeDeleteRootfs     APIVmGetResponseOperationCode = "delete_rootfs"
	APIVmGetResponseOperationCodeListRootfs       APIVmGetResponseOperationCode = "list_rootfs"
)

func (r APIVmGetResponseOperationCode) IsKnown() bool {
	switch r {
	case APIVmGetResponseOperationCodeListClusters, APIVmGetResponseOperationCodeGetCluster, APIVmGetResponseOperationCodeCreateCluster, APIVmGetResponseOperationCodeDeleteCluster, APIVmGetResponseOperationCodeUpdateCluster, APIVmGetResponseOperationCodeGetClusterSSHKey, APIVmGetResponseOperationCodeListVms, APIVmGetResponseOperationCodeGetVm, APIVmGetResponseOperationCodeUpdateVm, APIVmGetResponseOperationCodeBranchVm, APIVmGetResponseOperationCodeCommitVm, APIVmGetResponseOperationCodeDeleteVm, APIVmGetResponseOperationCodeGetVmSSHKey, APIVmGetResponseOperationCodeUploadRootfs, APIVmGetResponseOperationCodeDeleteRootfs, APIVmGetResponseOperationCodeListRootfs:
		return true
	}
	return false
}

type APIVmUpdateResponse struct {
	Data          APIVmUpdateResponseData          `json:"data,required"`
	DurationNs    int64                            `json:"duration_ns,required"`
	OperationCode APIVmUpdateResponseOperationCode `json:"operation_code,required"`
	OperationID   string                           `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                   `json:"time_start,required"`
	JSON      apiVmUpdateResponseJSON `json:"-"`
}

// apiVmUpdateResponseJSON contains the JSON metadata for the struct
// [APIVmUpdateResponse]
type apiVmUpdateResponseJSON struct {
	Data          apijson.Field
	DurationNs    apijson.Field
	OperationCode apijson.Field
	OperationID   apijson.Field
	TimeStart     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *APIVmUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type APIVmUpdateResponseData struct {
	// The ID of the VM.
	ID string `json:"id,required"`
	// The IDs of direct children branched from this VM.
	Children []string `json:"children,required"`
	// The VM's cluster ID
	ClusterID string `json:"cluster_id,required"`
	// What is the size of the "disk" allocated to this VM
	FsSizeMib int64 `json:"fs_size_mib,required"`
	// The VM's local IP address on the VM subnet
	IPAddress string `json:"ip_address,required"`
	// How much RAM is allocated to this VM
	MemSizeMib int64 `json:"mem_size_mib,required"`
	// The VM's network configuration
	NetworkInfo APIVmUpdateResponseDataNetworkInfo `json:"network_info,required"`
	// Whether the VM is running, paused, or not started.
	State APIVmUpdateResponseDataState `json:"state,required"`
	// How many vCPUs were allocated to this VM
	VcpuCount int64 `json:"vcpu_count,required"`
	// Human-readable name assigned to the VM.
	Alias string `json:"alias,nullable"`
	// The parent VM's ID, if present. If None, then this VM is a root VM.
	ParentID string                      `json:"parent_id,nullable"`
	JSON     apiVmUpdateResponseDataJSON `json:"-"`
}

// apiVmUpdateResponseDataJSON contains the JSON metadata for the struct
// [APIVmUpdateResponseData]
type apiVmUpdateResponseDataJSON struct {
	ID          apijson.Field
	Children    apijson.Field
	ClusterID   apijson.Field
	FsSizeMib   apijson.Field
	IPAddress   apijson.Field
	MemSizeMib  apijson.Field
	NetworkInfo apijson.Field
	State       apijson.Field
	VcpuCount   apijson.Field
	Alias       apijson.Field
	ParentID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmUpdateResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmUpdateResponseDataJSON) RawJSON() string {
	return r.raw
}

// The VM's network configuration
type APIVmUpdateResponseDataNetworkInfo struct {
	GuestIP     string                                 `json:"guest_ip,required"`
	GuestMac    string                                 `json:"guest_mac,required"`
	SSHPort     int64                                  `json:"ssh_port,required"`
	Tap0IP      string                                 `json:"tap0_ip,required"`
	Tap0Name    string                                 `json:"tap0_name,required"`
	VmNamespace string                                 `json:"vm_namespace,required"`
	JSON        apiVmUpdateResponseDataNetworkInfoJSON `json:"-"`
}

// apiVmUpdateResponseDataNetworkInfoJSON contains the JSON metadata for the struct
// [APIVmUpdateResponseDataNetworkInfo]
type apiVmUpdateResponseDataNetworkInfoJSON struct {
	GuestIP     apijson.Field
	GuestMac    apijson.Field
	SSHPort     apijson.Field
	Tap0IP      apijson.Field
	Tap0Name    apijson.Field
	VmNamespace apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmUpdateResponseDataNetworkInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmUpdateResponseDataNetworkInfoJSON) RawJSON() string {
	return r.raw
}

// Whether the VM is running, paused, or not started.
type APIVmUpdateResponseDataState string

const (
	APIVmUpdateResponseDataStateNotStarted APIVmUpdateResponseDataState = "Not started"
	APIVmUpdateResponseDataStateRunning    APIVmUpdateResponseDataState = "Running"
	APIVmUpdateResponseDataStatePaused     APIVmUpdateResponseDataState = "Paused"
)

func (r APIVmUpdateResponseDataState) IsKnown() bool {
	switch r {
	case APIVmUpdateResponseDataStateNotStarted, APIVmUpdateResponseDataStateRunning, APIVmUpdateResponseDataStatePaused:
		return true
	}
	return false
}

type APIVmUpdateResponseOperationCode string

const (
	APIVmUpdateResponseOperationCodeListClusters     APIVmUpdateResponseOperationCode = "list_clusters"
	APIVmUpdateResponseOperationCodeGetCluster       APIVmUpdateResponseOperationCode = "get_cluster"
	APIVmUpdateResponseOperationCodeCreateCluster    APIVmUpdateResponseOperationCode = "create_cluster"
	APIVmUpdateResponseOperationCodeDeleteCluster    APIVmUpdateResponseOperationCode = "delete_cluster"
	APIVmUpdateResponseOperationCodeUpdateCluster    APIVmUpdateResponseOperationCode = "update_cluster"
	APIVmUpdateResponseOperationCodeGetClusterSSHKey APIVmUpdateResponseOperationCode = "get_cluster_ssh_key"
	APIVmUpdateResponseOperationCodeListVms          APIVmUpdateResponseOperationCode = "list_vms"
	APIVmUpdateResponseOperationCodeGetVm            APIVmUpdateResponseOperationCode = "get_vm"
	APIVmUpdateResponseOperationCodeUpdateVm         APIVmUpdateResponseOperationCode = "update_vm"
	APIVmUpdateResponseOperationCodeBranchVm         APIVmUpdateResponseOperationCode = "branch_vm"
	APIVmUpdateResponseOperationCodeCommitVm         APIVmUpdateResponseOperationCode = "commit_vm"
	APIVmUpdateResponseOperationCodeDeleteVm         APIVmUpdateResponseOperationCode = "delete_vm"
	APIVmUpdateResponseOperationCodeGetVmSSHKey      APIVmUpdateResponseOperationCode = "get_vm_ssh_key"
	APIVmUpdateResponseOperationCodeUploadRootfs     APIVmUpdateResponseOperationCode = "upload_rootfs"
	APIVmUpdateResponseOperationCodeDeleteRootfs     APIVmUpdateResponseOperationCode = "delete_rootfs"
	APIVmUpdateResponseOperationCodeListRootfs       APIVmUpdateResponseOperationCode = "list_rootfs"
)

func (r APIVmUpdateResponseOperationCode) IsKnown() bool {
	switch r {
	case APIVmUpdateResponseOperationCodeListClusters, APIVmUpdateResponseOperationCodeGetCluster, APIVmUpdateResponseOperationCodeCreateCluster, APIVmUpdateResponseOperationCodeDeleteCluster, APIVmUpdateResponseOperationCodeUpdateCluster, APIVmUpdateResponseOperationCodeGetClusterSSHKey, APIVmUpdateResponseOperationCodeListVms, APIVmUpdateResponseOperationCodeGetVm, APIVmUpdateResponseOperationCodeUpdateVm, APIVmUpdateResponseOperationCodeBranchVm, APIVmUpdateResponseOperationCodeCommitVm, APIVmUpdateResponseOperationCodeDeleteVm, APIVmUpdateResponseOperationCodeGetVmSSHKey, APIVmUpdateResponseOperationCodeUploadRootfs, APIVmUpdateResponseOperationCodeDeleteRootfs, APIVmUpdateResponseOperationCodeListRootfs:
		return true
	}
	return false
}

type APIVmListResponse struct {
	Data          []APIVmListResponseData        `json:"data,required"`
	DurationNs    int64                          `json:"duration_ns,required"`
	OperationCode APIVmListResponseOperationCode `json:"operation_code,required"`
	OperationID   string                         `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                 `json:"time_start,required"`
	JSON      apiVmListResponseJSON `json:"-"`
}

// apiVmListResponseJSON contains the JSON metadata for the struct
// [APIVmListResponse]
type apiVmListResponseJSON struct {
	Data          apijson.Field
	DurationNs    apijson.Field
	OperationCode apijson.Field
	OperationID   apijson.Field
	TimeStart     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *APIVmListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmListResponseJSON) RawJSON() string {
	return r.raw
}

type APIVmListResponseData struct {
	// The ID of the VM.
	ID string `json:"id,required"`
	// The IDs of direct children branched from this VM.
	Children []string `json:"children,required"`
	// The VM's cluster ID
	ClusterID string `json:"cluster_id,required"`
	// What is the size of the "disk" allocated to this VM
	FsSizeMib int64 `json:"fs_size_mib,required"`
	// The VM's local IP address on the VM subnet
	IPAddress string `json:"ip_address,required"`
	// How much RAM is allocated to this VM
	MemSizeMib int64 `json:"mem_size_mib,required"`
	// The VM's network configuration
	NetworkInfo APIVmListResponseDataNetworkInfo `json:"network_info,required"`
	// Whether the VM is running, paused, or not started.
	State APIVmListResponseDataState `json:"state,required"`
	// How many vCPUs were allocated to this VM
	VcpuCount int64 `json:"vcpu_count,required"`
	// Human-readable name assigned to the VM.
	Alias string `json:"alias,nullable"`
	// The parent VM's ID, if present. If None, then this VM is a root VM.
	ParentID string                    `json:"parent_id,nullable"`
	JSON     apiVmListResponseDataJSON `json:"-"`
}

// apiVmListResponseDataJSON contains the JSON metadata for the struct
// [APIVmListResponseData]
type apiVmListResponseDataJSON struct {
	ID          apijson.Field
	Children    apijson.Field
	ClusterID   apijson.Field
	FsSizeMib   apijson.Field
	IPAddress   apijson.Field
	MemSizeMib  apijson.Field
	NetworkInfo apijson.Field
	State       apijson.Field
	VcpuCount   apijson.Field
	Alias       apijson.Field
	ParentID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmListResponseDataJSON) RawJSON() string {
	return r.raw
}

// The VM's network configuration
type APIVmListResponseDataNetworkInfo struct {
	GuestIP     string                               `json:"guest_ip,required"`
	GuestMac    string                               `json:"guest_mac,required"`
	SSHPort     int64                                `json:"ssh_port,required"`
	Tap0IP      string                               `json:"tap0_ip,required"`
	Tap0Name    string                               `json:"tap0_name,required"`
	VmNamespace string                               `json:"vm_namespace,required"`
	JSON        apiVmListResponseDataNetworkInfoJSON `json:"-"`
}

// apiVmListResponseDataNetworkInfoJSON contains the JSON metadata for the struct
// [APIVmListResponseDataNetworkInfo]
type apiVmListResponseDataNetworkInfoJSON struct {
	GuestIP     apijson.Field
	GuestMac    apijson.Field
	SSHPort     apijson.Field
	Tap0IP      apijson.Field
	Tap0Name    apijson.Field
	VmNamespace apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmListResponseDataNetworkInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmListResponseDataNetworkInfoJSON) RawJSON() string {
	return r.raw
}

// Whether the VM is running, paused, or not started.
type APIVmListResponseDataState string

const (
	APIVmListResponseDataStateNotStarted APIVmListResponseDataState = "Not started"
	APIVmListResponseDataStateRunning    APIVmListResponseDataState = "Running"
	APIVmListResponseDataStatePaused     APIVmListResponseDataState = "Paused"
)

func (r APIVmListResponseDataState) IsKnown() bool {
	switch r {
	case APIVmListResponseDataStateNotStarted, APIVmListResponseDataStateRunning, APIVmListResponseDataStatePaused:
		return true
	}
	return false
}

type APIVmListResponseOperationCode string

const (
	APIVmListResponseOperationCodeListClusters     APIVmListResponseOperationCode = "list_clusters"
	APIVmListResponseOperationCodeGetCluster       APIVmListResponseOperationCode = "get_cluster"
	APIVmListResponseOperationCodeCreateCluster    APIVmListResponseOperationCode = "create_cluster"
	APIVmListResponseOperationCodeDeleteCluster    APIVmListResponseOperationCode = "delete_cluster"
	APIVmListResponseOperationCodeUpdateCluster    APIVmListResponseOperationCode = "update_cluster"
	APIVmListResponseOperationCodeGetClusterSSHKey APIVmListResponseOperationCode = "get_cluster_ssh_key"
	APIVmListResponseOperationCodeListVms          APIVmListResponseOperationCode = "list_vms"
	APIVmListResponseOperationCodeGetVm            APIVmListResponseOperationCode = "get_vm"
	APIVmListResponseOperationCodeUpdateVm         APIVmListResponseOperationCode = "update_vm"
	APIVmListResponseOperationCodeBranchVm         APIVmListResponseOperationCode = "branch_vm"
	APIVmListResponseOperationCodeCommitVm         APIVmListResponseOperationCode = "commit_vm"
	APIVmListResponseOperationCodeDeleteVm         APIVmListResponseOperationCode = "delete_vm"
	APIVmListResponseOperationCodeGetVmSSHKey      APIVmListResponseOperationCode = "get_vm_ssh_key"
	APIVmListResponseOperationCodeUploadRootfs     APIVmListResponseOperationCode = "upload_rootfs"
	APIVmListResponseOperationCodeDeleteRootfs     APIVmListResponseOperationCode = "delete_rootfs"
	APIVmListResponseOperationCodeListRootfs       APIVmListResponseOperationCode = "list_rootfs"
)

func (r APIVmListResponseOperationCode) IsKnown() bool {
	switch r {
	case APIVmListResponseOperationCodeListClusters, APIVmListResponseOperationCodeGetCluster, APIVmListResponseOperationCodeCreateCluster, APIVmListResponseOperationCodeDeleteCluster, APIVmListResponseOperationCodeUpdateCluster, APIVmListResponseOperationCodeGetClusterSSHKey, APIVmListResponseOperationCodeListVms, APIVmListResponseOperationCodeGetVm, APIVmListResponseOperationCodeUpdateVm, APIVmListResponseOperationCodeBranchVm, APIVmListResponseOperationCodeCommitVm, APIVmListResponseOperationCodeDeleteVm, APIVmListResponseOperationCodeGetVmSSHKey, APIVmListResponseOperationCodeUploadRootfs, APIVmListResponseOperationCodeDeleteRootfs, APIVmListResponseOperationCodeListRootfs:
		return true
	}
	return false
}

type APIVmDeleteResponse struct {
	// A struct containing information about an attempted VM deletion request. Reports
	// information in the event of a partial failure so billing can still be udpated
	// appropriately.
	Data          APIVmDeleteResponseData          `json:"data,required"`
	DurationNs    int64                            `json:"duration_ns,required"`
	OperationCode APIVmDeleteResponseOperationCode `json:"operation_code,required"`
	OperationID   string                           `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                   `json:"time_start,required"`
	JSON      apiVmDeleteResponseJSON `json:"-"`
}

// apiVmDeleteResponseJSON contains the JSON metadata for the struct
// [APIVmDeleteResponse]
type apiVmDeleteResponseJSON struct {
	Data          apijson.Field
	DurationNs    apijson.Field
	OperationCode apijson.Field
	OperationID   apijson.Field
	TimeStart     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *APIVmDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// A struct containing information about an attempted VM deletion request. Reports
// information in the event of a partial failure so billing can still be udpated
// appropriately.
type APIVmDeleteResponseData struct {
	DeletedIDs []string                       `json:"deleted_ids,required"`
	Errors     []APIVmDeleteResponseDataError `json:"errors,required"`
	JSON       apiVmDeleteResponseDataJSON    `json:"-"`
}

// apiVmDeleteResponseDataJSON contains the JSON metadata for the struct
// [APIVmDeleteResponseData]
type apiVmDeleteResponseDataJSON struct {
	DeletedIDs  apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmDeleteResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmDeleteResponseDataJSON) RawJSON() string {
	return r.raw
}

// Contains a VM ID and the reason that it could not be deleted.
type APIVmDeleteResponseDataError struct {
	ID    string                           `json:"id,required"`
	Error string                           `json:"error,required"`
	JSON  apiVmDeleteResponseDataErrorJSON `json:"-"`
}

// apiVmDeleteResponseDataErrorJSON contains the JSON metadata for the struct
// [APIVmDeleteResponseDataError]
type apiVmDeleteResponseDataErrorJSON struct {
	ID          apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmDeleteResponseDataError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmDeleteResponseDataErrorJSON) RawJSON() string {
	return r.raw
}

type APIVmDeleteResponseOperationCode string

const (
	APIVmDeleteResponseOperationCodeListClusters     APIVmDeleteResponseOperationCode = "list_clusters"
	APIVmDeleteResponseOperationCodeGetCluster       APIVmDeleteResponseOperationCode = "get_cluster"
	APIVmDeleteResponseOperationCodeCreateCluster    APIVmDeleteResponseOperationCode = "create_cluster"
	APIVmDeleteResponseOperationCodeDeleteCluster    APIVmDeleteResponseOperationCode = "delete_cluster"
	APIVmDeleteResponseOperationCodeUpdateCluster    APIVmDeleteResponseOperationCode = "update_cluster"
	APIVmDeleteResponseOperationCodeGetClusterSSHKey APIVmDeleteResponseOperationCode = "get_cluster_ssh_key"
	APIVmDeleteResponseOperationCodeListVms          APIVmDeleteResponseOperationCode = "list_vms"
	APIVmDeleteResponseOperationCodeGetVm            APIVmDeleteResponseOperationCode = "get_vm"
	APIVmDeleteResponseOperationCodeUpdateVm         APIVmDeleteResponseOperationCode = "update_vm"
	APIVmDeleteResponseOperationCodeBranchVm         APIVmDeleteResponseOperationCode = "branch_vm"
	APIVmDeleteResponseOperationCodeCommitVm         APIVmDeleteResponseOperationCode = "commit_vm"
	APIVmDeleteResponseOperationCodeDeleteVm         APIVmDeleteResponseOperationCode = "delete_vm"
	APIVmDeleteResponseOperationCodeGetVmSSHKey      APIVmDeleteResponseOperationCode = "get_vm_ssh_key"
	APIVmDeleteResponseOperationCodeUploadRootfs     APIVmDeleteResponseOperationCode = "upload_rootfs"
	APIVmDeleteResponseOperationCodeDeleteRootfs     APIVmDeleteResponseOperationCode = "delete_rootfs"
	APIVmDeleteResponseOperationCodeListRootfs       APIVmDeleteResponseOperationCode = "list_rootfs"
)

func (r APIVmDeleteResponseOperationCode) IsKnown() bool {
	switch r {
	case APIVmDeleteResponseOperationCodeListClusters, APIVmDeleteResponseOperationCodeGetCluster, APIVmDeleteResponseOperationCodeCreateCluster, APIVmDeleteResponseOperationCodeDeleteCluster, APIVmDeleteResponseOperationCodeUpdateCluster, APIVmDeleteResponseOperationCodeGetClusterSSHKey, APIVmDeleteResponseOperationCodeListVms, APIVmDeleteResponseOperationCodeGetVm, APIVmDeleteResponseOperationCodeUpdateVm, APIVmDeleteResponseOperationCodeBranchVm, APIVmDeleteResponseOperationCodeCommitVm, APIVmDeleteResponseOperationCodeDeleteVm, APIVmDeleteResponseOperationCodeGetVmSSHKey, APIVmDeleteResponseOperationCodeUploadRootfs, APIVmDeleteResponseOperationCodeDeleteRootfs, APIVmDeleteResponseOperationCodeListRootfs:
		return true
	}
	return false
}

type APIVmBranchResponse struct {
	Data          APIVmBranchResponseData          `json:"data,required"`
	DurationNs    int64                            `json:"duration_ns,required"`
	OperationCode APIVmBranchResponseOperationCode `json:"operation_code,required"`
	OperationID   string                           `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                   `json:"time_start,required"`
	JSON      apiVmBranchResponseJSON `json:"-"`
}

// apiVmBranchResponseJSON contains the JSON metadata for the struct
// [APIVmBranchResponse]
type apiVmBranchResponseJSON struct {
	Data          apijson.Field
	DurationNs    apijson.Field
	OperationCode apijson.Field
	OperationID   apijson.Field
	TimeStart     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *APIVmBranchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmBranchResponseJSON) RawJSON() string {
	return r.raw
}

type APIVmBranchResponseData struct {
	// The ID of the VM.
	ID string `json:"id,required"`
	// The IDs of direct children branched from this VM.
	Children []string `json:"children,required"`
	// The VM's cluster ID
	ClusterID string `json:"cluster_id,required"`
	// What is the size of the "disk" allocated to this VM
	FsSizeMib int64 `json:"fs_size_mib,required"`
	// The VM's local IP address on the VM subnet
	IPAddress string `json:"ip_address,required"`
	// How much RAM is allocated to this VM
	MemSizeMib int64 `json:"mem_size_mib,required"`
	// The VM's network configuration
	NetworkInfo APIVmBranchResponseDataNetworkInfo `json:"network_info,required"`
	// Whether the VM is running, paused, or not started.
	State APIVmBranchResponseDataState `json:"state,required"`
	// How many vCPUs were allocated to this VM
	VcpuCount int64 `json:"vcpu_count,required"`
	// Human-readable name assigned to the VM.
	Alias string `json:"alias,nullable"`
	// The parent VM's ID, if present. If None, then this VM is a root VM.
	ParentID string                      `json:"parent_id,nullable"`
	JSON     apiVmBranchResponseDataJSON `json:"-"`
}

// apiVmBranchResponseDataJSON contains the JSON metadata for the struct
// [APIVmBranchResponseData]
type apiVmBranchResponseDataJSON struct {
	ID          apijson.Field
	Children    apijson.Field
	ClusterID   apijson.Field
	FsSizeMib   apijson.Field
	IPAddress   apijson.Field
	MemSizeMib  apijson.Field
	NetworkInfo apijson.Field
	State       apijson.Field
	VcpuCount   apijson.Field
	Alias       apijson.Field
	ParentID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmBranchResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmBranchResponseDataJSON) RawJSON() string {
	return r.raw
}

// The VM's network configuration
type APIVmBranchResponseDataNetworkInfo struct {
	GuestIP     string                                 `json:"guest_ip,required"`
	GuestMac    string                                 `json:"guest_mac,required"`
	SSHPort     int64                                  `json:"ssh_port,required"`
	Tap0IP      string                                 `json:"tap0_ip,required"`
	Tap0Name    string                                 `json:"tap0_name,required"`
	VmNamespace string                                 `json:"vm_namespace,required"`
	JSON        apiVmBranchResponseDataNetworkInfoJSON `json:"-"`
}

// apiVmBranchResponseDataNetworkInfoJSON contains the JSON metadata for the struct
// [APIVmBranchResponseDataNetworkInfo]
type apiVmBranchResponseDataNetworkInfoJSON struct {
	GuestIP     apijson.Field
	GuestMac    apijson.Field
	SSHPort     apijson.Field
	Tap0IP      apijson.Field
	Tap0Name    apijson.Field
	VmNamespace apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmBranchResponseDataNetworkInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmBranchResponseDataNetworkInfoJSON) RawJSON() string {
	return r.raw
}

// Whether the VM is running, paused, or not started.
type APIVmBranchResponseDataState string

const (
	APIVmBranchResponseDataStateNotStarted APIVmBranchResponseDataState = "Not started"
	APIVmBranchResponseDataStateRunning    APIVmBranchResponseDataState = "Running"
	APIVmBranchResponseDataStatePaused     APIVmBranchResponseDataState = "Paused"
)

func (r APIVmBranchResponseDataState) IsKnown() bool {
	switch r {
	case APIVmBranchResponseDataStateNotStarted, APIVmBranchResponseDataStateRunning, APIVmBranchResponseDataStatePaused:
		return true
	}
	return false
}

type APIVmBranchResponseOperationCode string

const (
	APIVmBranchResponseOperationCodeListClusters     APIVmBranchResponseOperationCode = "list_clusters"
	APIVmBranchResponseOperationCodeGetCluster       APIVmBranchResponseOperationCode = "get_cluster"
	APIVmBranchResponseOperationCodeCreateCluster    APIVmBranchResponseOperationCode = "create_cluster"
	APIVmBranchResponseOperationCodeDeleteCluster    APIVmBranchResponseOperationCode = "delete_cluster"
	APIVmBranchResponseOperationCodeUpdateCluster    APIVmBranchResponseOperationCode = "update_cluster"
	APIVmBranchResponseOperationCodeGetClusterSSHKey APIVmBranchResponseOperationCode = "get_cluster_ssh_key"
	APIVmBranchResponseOperationCodeListVms          APIVmBranchResponseOperationCode = "list_vms"
	APIVmBranchResponseOperationCodeGetVm            APIVmBranchResponseOperationCode = "get_vm"
	APIVmBranchResponseOperationCodeUpdateVm         APIVmBranchResponseOperationCode = "update_vm"
	APIVmBranchResponseOperationCodeBranchVm         APIVmBranchResponseOperationCode = "branch_vm"
	APIVmBranchResponseOperationCodeCommitVm         APIVmBranchResponseOperationCode = "commit_vm"
	APIVmBranchResponseOperationCodeDeleteVm         APIVmBranchResponseOperationCode = "delete_vm"
	APIVmBranchResponseOperationCodeGetVmSSHKey      APIVmBranchResponseOperationCode = "get_vm_ssh_key"
	APIVmBranchResponseOperationCodeUploadRootfs     APIVmBranchResponseOperationCode = "upload_rootfs"
	APIVmBranchResponseOperationCodeDeleteRootfs     APIVmBranchResponseOperationCode = "delete_rootfs"
	APIVmBranchResponseOperationCodeListRootfs       APIVmBranchResponseOperationCode = "list_rootfs"
)

func (r APIVmBranchResponseOperationCode) IsKnown() bool {
	switch r {
	case APIVmBranchResponseOperationCodeListClusters, APIVmBranchResponseOperationCodeGetCluster, APIVmBranchResponseOperationCodeCreateCluster, APIVmBranchResponseOperationCodeDeleteCluster, APIVmBranchResponseOperationCodeUpdateCluster, APIVmBranchResponseOperationCodeGetClusterSSHKey, APIVmBranchResponseOperationCodeListVms, APIVmBranchResponseOperationCodeGetVm, APIVmBranchResponseOperationCodeUpdateVm, APIVmBranchResponseOperationCodeBranchVm, APIVmBranchResponseOperationCodeCommitVm, APIVmBranchResponseOperationCodeDeleteVm, APIVmBranchResponseOperationCodeGetVmSSHKey, APIVmBranchResponseOperationCodeUploadRootfs, APIVmBranchResponseOperationCodeDeleteRootfs, APIVmBranchResponseOperationCodeListRootfs:
		return true
	}
	return false
}

type APIVmCommitResponse struct {
	Data          APIVmCommitResponseData          `json:"data,required"`
	DurationNs    int64                            `json:"duration_ns,required"`
	OperationCode APIVmCommitResponseOperationCode `json:"operation_code,required"`
	OperationID   string                           `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                   `json:"time_start,required"`
	JSON      apiVmCommitResponseJSON `json:"-"`
}

// apiVmCommitResponseJSON contains the JSON metadata for the struct
// [APIVmCommitResponse]
type apiVmCommitResponseJSON struct {
	Data          apijson.Field
	DurationNs    apijson.Field
	OperationCode apijson.Field
	OperationID   apijson.Field
	TimeStart     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *APIVmCommitResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmCommitResponseJSON) RawJSON() string {
	return r.raw
}

type APIVmCommitResponseData struct {
	ClusterID        string                      `json:"cluster_id,required"`
	CommitID         string                      `json:"commit_id,required"`
	HostArchitecture string                      `json:"host_architecture,required"`
	JSON             apiVmCommitResponseDataJSON `json:"-"`
}

// apiVmCommitResponseDataJSON contains the JSON metadata for the struct
// [APIVmCommitResponseData]
type apiVmCommitResponseDataJSON struct {
	ClusterID        apijson.Field
	CommitID         apijson.Field
	HostArchitecture apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *APIVmCommitResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmCommitResponseDataJSON) RawJSON() string {
	return r.raw
}

type APIVmCommitResponseOperationCode string

const (
	APIVmCommitResponseOperationCodeListClusters     APIVmCommitResponseOperationCode = "list_clusters"
	APIVmCommitResponseOperationCodeGetCluster       APIVmCommitResponseOperationCode = "get_cluster"
	APIVmCommitResponseOperationCodeCreateCluster    APIVmCommitResponseOperationCode = "create_cluster"
	APIVmCommitResponseOperationCodeDeleteCluster    APIVmCommitResponseOperationCode = "delete_cluster"
	APIVmCommitResponseOperationCodeUpdateCluster    APIVmCommitResponseOperationCode = "update_cluster"
	APIVmCommitResponseOperationCodeGetClusterSSHKey APIVmCommitResponseOperationCode = "get_cluster_ssh_key"
	APIVmCommitResponseOperationCodeListVms          APIVmCommitResponseOperationCode = "list_vms"
	APIVmCommitResponseOperationCodeGetVm            APIVmCommitResponseOperationCode = "get_vm"
	APIVmCommitResponseOperationCodeUpdateVm         APIVmCommitResponseOperationCode = "update_vm"
	APIVmCommitResponseOperationCodeBranchVm         APIVmCommitResponseOperationCode = "branch_vm"
	APIVmCommitResponseOperationCodeCommitVm         APIVmCommitResponseOperationCode = "commit_vm"
	APIVmCommitResponseOperationCodeDeleteVm         APIVmCommitResponseOperationCode = "delete_vm"
	APIVmCommitResponseOperationCodeGetVmSSHKey      APIVmCommitResponseOperationCode = "get_vm_ssh_key"
	APIVmCommitResponseOperationCodeUploadRootfs     APIVmCommitResponseOperationCode = "upload_rootfs"
	APIVmCommitResponseOperationCodeDeleteRootfs     APIVmCommitResponseOperationCode = "delete_rootfs"
	APIVmCommitResponseOperationCodeListRootfs       APIVmCommitResponseOperationCode = "list_rootfs"
)

func (r APIVmCommitResponseOperationCode) IsKnown() bool {
	switch r {
	case APIVmCommitResponseOperationCodeListClusters, APIVmCommitResponseOperationCodeGetCluster, APIVmCommitResponseOperationCodeCreateCluster, APIVmCommitResponseOperationCodeDeleteCluster, APIVmCommitResponseOperationCodeUpdateCluster, APIVmCommitResponseOperationCodeGetClusterSSHKey, APIVmCommitResponseOperationCodeListVms, APIVmCommitResponseOperationCodeGetVm, APIVmCommitResponseOperationCodeUpdateVm, APIVmCommitResponseOperationCodeBranchVm, APIVmCommitResponseOperationCodeCommitVm, APIVmCommitResponseOperationCodeDeleteVm, APIVmCommitResponseOperationCodeGetVmSSHKey, APIVmCommitResponseOperationCodeUploadRootfs, APIVmCommitResponseOperationCodeDeleteRootfs, APIVmCommitResponseOperationCodeListRootfs:
		return true
	}
	return false
}

type APIVmGetSSHKeyResponse struct {
	Data          string                              `json:"data,required"`
	DurationNs    int64                               `json:"duration_ns,required"`
	OperationCode APIVmGetSSHKeyResponseOperationCode `json:"operation_code,required"`
	OperationID   string                              `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                      `json:"time_start,required"`
	JSON      apiVmGetSSHKeyResponseJSON `json:"-"`
}

// apiVmGetSSHKeyResponseJSON contains the JSON metadata for the struct
// [APIVmGetSSHKeyResponse]
type apiVmGetSSHKeyResponseJSON struct {
	Data          apijson.Field
	DurationNs    apijson.Field
	OperationCode apijson.Field
	OperationID   apijson.Field
	TimeStart     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *APIVmGetSSHKeyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmGetSSHKeyResponseJSON) RawJSON() string {
	return r.raw
}

type APIVmGetSSHKeyResponseOperationCode string

const (
	APIVmGetSSHKeyResponseOperationCodeListClusters     APIVmGetSSHKeyResponseOperationCode = "list_clusters"
	APIVmGetSSHKeyResponseOperationCodeGetCluster       APIVmGetSSHKeyResponseOperationCode = "get_cluster"
	APIVmGetSSHKeyResponseOperationCodeCreateCluster    APIVmGetSSHKeyResponseOperationCode = "create_cluster"
	APIVmGetSSHKeyResponseOperationCodeDeleteCluster    APIVmGetSSHKeyResponseOperationCode = "delete_cluster"
	APIVmGetSSHKeyResponseOperationCodeUpdateCluster    APIVmGetSSHKeyResponseOperationCode = "update_cluster"
	APIVmGetSSHKeyResponseOperationCodeGetClusterSSHKey APIVmGetSSHKeyResponseOperationCode = "get_cluster_ssh_key"
	APIVmGetSSHKeyResponseOperationCodeListVms          APIVmGetSSHKeyResponseOperationCode = "list_vms"
	APIVmGetSSHKeyResponseOperationCodeGetVm            APIVmGetSSHKeyResponseOperationCode = "get_vm"
	APIVmGetSSHKeyResponseOperationCodeUpdateVm         APIVmGetSSHKeyResponseOperationCode = "update_vm"
	APIVmGetSSHKeyResponseOperationCodeBranchVm         APIVmGetSSHKeyResponseOperationCode = "branch_vm"
	APIVmGetSSHKeyResponseOperationCodeCommitVm         APIVmGetSSHKeyResponseOperationCode = "commit_vm"
	APIVmGetSSHKeyResponseOperationCodeDeleteVm         APIVmGetSSHKeyResponseOperationCode = "delete_vm"
	APIVmGetSSHKeyResponseOperationCodeGetVmSSHKey      APIVmGetSSHKeyResponseOperationCode = "get_vm_ssh_key"
	APIVmGetSSHKeyResponseOperationCodeUploadRootfs     APIVmGetSSHKeyResponseOperationCode = "upload_rootfs"
	APIVmGetSSHKeyResponseOperationCodeDeleteRootfs     APIVmGetSSHKeyResponseOperationCode = "delete_rootfs"
	APIVmGetSSHKeyResponseOperationCodeListRootfs       APIVmGetSSHKeyResponseOperationCode = "list_rootfs"
)

func (r APIVmGetSSHKeyResponseOperationCode) IsKnown() bool {
	switch r {
	case APIVmGetSSHKeyResponseOperationCodeListClusters, APIVmGetSSHKeyResponseOperationCodeGetCluster, APIVmGetSSHKeyResponseOperationCodeCreateCluster, APIVmGetSSHKeyResponseOperationCodeDeleteCluster, APIVmGetSSHKeyResponseOperationCodeUpdateCluster, APIVmGetSSHKeyResponseOperationCodeGetClusterSSHKey, APIVmGetSSHKeyResponseOperationCodeListVms, APIVmGetSSHKeyResponseOperationCodeGetVm, APIVmGetSSHKeyResponseOperationCodeUpdateVm, APIVmGetSSHKeyResponseOperationCodeBranchVm, APIVmGetSSHKeyResponseOperationCodeCommitVm, APIVmGetSSHKeyResponseOperationCodeDeleteVm, APIVmGetSSHKeyResponseOperationCodeGetVmSSHKey, APIVmGetSSHKeyResponseOperationCodeUploadRootfs, APIVmGetSSHKeyResponseOperationCodeDeleteRootfs, APIVmGetSSHKeyResponseOperationCodeListRootfs:
		return true
	}
	return false
}

type APIVmUpdateParams struct {
	VmPatchParams VmPatchParams `json:"vm_patch_params,required"`
}

func (r APIVmUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.VmPatchParams)
}

type APIVmDeleteParams struct {
	// Delete children recursively
	Recursive param.Field[bool] `query:"recursive,required"`
}

// URLQuery serializes [APIVmDeleteParams]'s query parameters as `url.Values`.
func (r APIVmDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIVmBranchParams struct {
	VmBranchParams VmBranchParams `json:"vm_branch_params,required"`
}

func (r APIVmBranchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.VmBranchParams)
}

type APIVmCommitParams struct {
	VmCommitRequest VmCommitRequestParam `json:"vm_commit_request,required"`
}

func (r APIVmCommitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.VmCommitRequest)
}
