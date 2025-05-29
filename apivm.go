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
func (r *APIVmService) Get(ctx context.Context, vmID string, opts ...option.RequestOption) (res *APIVmGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("api/vm/%s", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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
func (r *APIVmService) Delete(ctx context.Context, vmID string, body APIVmDeleteParams, opts ...option.RequestOption) (res *APIVmDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("api/vm/%s", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Branch a VM.
func (r *APIVmService) Branch(ctx context.Context, vmID string, opts ...option.RequestOption) (res *APIVmBranchResponse, err error) {
	opts = append(r.Options[:], opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("api/vm/%s/branch", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Commit a VM.
func (r *APIVmService) Commit(ctx context.Context, vmID string, opts ...option.RequestOption) (res *APIVmCommitResponse, err error) {
	opts = append(r.Options[:], opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("api/vm/%s/commit", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Get the SSH private key for VM access
func (r *APIVmService) GetSSHKey(ctx context.Context, vmID string, opts ...option.RequestOption) (res *APIVmGetSSHKeyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("api/vm/%s/ssh_key", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update VM state.
func (r *APIVmService) UpdateState(ctx context.Context, vmID string, body APIVmUpdateStateParams, opts ...option.RequestOption) (res *APIVmUpdateStateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("api/vm/%s", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type PatchRequestParam struct {
	Action param.Field[PatchRequestAction] `json:"action,required"`
}

func (r PatchRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PatchRequestAction string

const (
	PatchRequestActionPause  PatchRequestAction = "pause"
	PatchRequestActionResume PatchRequestAction = "resume"
)

func (r PatchRequestAction) IsKnown() bool {
	switch r {
	case PatchRequestActionPause, PatchRequestActionResume:
		return true
	}
	return false
}

type Vm struct {
	// The ID of the VM.
	ID string `json:"id,required"`
	// The IDs of direct children branched from this VM.
	Children []string `json:"children,required"`
	// The VM's cluster ID
	ClusterID string `json:"cluster_id,required"`
	// The VM's local IP address on the VM subnet
	IPAddress string `json:"ip_address,required"`
	// The VM's network configuration
	NetworkInfo VmNetworkInfo `json:"network_info,required"`
	// Whether the VM is running, paused, or not started.
	State VmState `json:"state,required"`
	// The parent VM's ID, if present. If None, then this VM is a root VM.
	ParentID string `json:"parent_id,nullable"`
	JSON     vmJSON `json:"-"`
}

// vmJSON contains the JSON metadata for the struct [Vm]
type vmJSON struct {
	ID          apijson.Field
	Children    apijson.Field
	ClusterID   apijson.Field
	IPAddress   apijson.Field
	NetworkInfo apijson.Field
	State       apijson.Field
	ParentID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Vm) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vmJSON) RawJSON() string {
	return r.raw
}

// The VM's network configuration
type VmNetworkInfo struct {
	GuestIP     string            `json:"guest_ip,required"`
	GuestMac    string            `json:"guest_mac,required"`
	SSHPort     int64             `json:"ssh_port,required"`
	Tap0IP      string            `json:"tap0_ip,required"`
	Tap0Name    string            `json:"tap0_name,required"`
	VmNamespace string            `json:"vm_namespace,required"`
	JSON        vmNetworkInfoJSON `json:"-"`
}

// vmNetworkInfoJSON contains the JSON metadata for the struct [VmNetworkInfo]
type vmNetworkInfoJSON struct {
	GuestIP     apijson.Field
	GuestMac    apijson.Field
	SSHPort     apijson.Field
	Tap0IP      apijson.Field
	Tap0Name    apijson.Field
	VmNamespace apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VmNetworkInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vmNetworkInfoJSON) RawJSON() string {
	return r.raw
}

// Whether the VM is running, paused, or not started.
type VmState string

const (
	VmStateNotStarted VmState = "Not started"
	VmStateRunning    VmState = "Running"
	VmStatePaused     VmState = "Paused"
)

func (r VmState) IsKnown() bool {
	switch r {
	case VmStateNotStarted, VmStateRunning, VmStatePaused:
		return true
	}
	return false
}

type APIVmGetResponse struct {
	Data        APIVmGetResponseData `json:"data,required"`
	DurationNs  int64                `json:"duration_ns,required"`
	OperationID string               `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                `json:"time_start,required"`
	JSON      apiVmGetResponseJSON `json:"-"`
}

// apiVmGetResponseJSON contains the JSON metadata for the struct
// [APIVmGetResponse]
type apiVmGetResponseJSON struct {
	Data        apijson.Field
	DurationNs  apijson.Field
	OperationID apijson.Field
	TimeStart   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
	// The VM's local IP address on the VM subnet
	IPAddress string `json:"ip_address,required"`
	// The VM's network configuration
	NetworkInfo APIVmGetResponseDataNetworkInfo `json:"network_info,required"`
	// Whether the VM is running, paused, or not started.
	State APIVmGetResponseDataState `json:"state,required"`
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
	IPAddress   apijson.Field
	NetworkInfo apijson.Field
	State       apijson.Field
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

type APIVmListResponse struct {
	Data        []APIVmListResponseData `json:"data,required"`
	DurationNs  int64                   `json:"duration_ns,required"`
	OperationID string                  `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                 `json:"time_start,required"`
	JSON      apiVmListResponseJSON `json:"-"`
}

// apiVmListResponseJSON contains the JSON metadata for the struct
// [APIVmListResponse]
type apiVmListResponseJSON struct {
	Data        apijson.Field
	DurationNs  apijson.Field
	OperationID apijson.Field
	TimeStart   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
	// The VM's local IP address on the VM subnet
	IPAddress string `json:"ip_address,required"`
	// The VM's network configuration
	NetworkInfo APIVmListResponseDataNetworkInfo `json:"network_info,required"`
	// Whether the VM is running, paused, or not started.
	State APIVmListResponseDataState `json:"state,required"`
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
	IPAddress   apijson.Field
	NetworkInfo apijson.Field
	State       apijson.Field
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

type APIVmDeleteResponse struct {
	Data        APIVmDeleteResponseData `json:"data,required"`
	DurationNs  int64                   `json:"duration_ns,required"`
	OperationID string                  `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                   `json:"time_start,required"`
	JSON      apiVmDeleteResponseJSON `json:"-"`
}

// apiVmDeleteResponseJSON contains the JSON metadata for the struct
// [APIVmDeleteResponse]
type apiVmDeleteResponseJSON struct {
	Data        apijson.Field
	DurationNs  apijson.Field
	OperationID apijson.Field
	TimeStart   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type APIVmDeleteResponseData struct {
	// The ID of the VM.
	ID string `json:"id,required"`
	// The IDs of direct children branched from this VM.
	Children []string `json:"children,required"`
	// The VM's cluster ID
	ClusterID string `json:"cluster_id,required"`
	// The VM's local IP address on the VM subnet
	IPAddress string `json:"ip_address,required"`
	// The VM's network configuration
	NetworkInfo APIVmDeleteResponseDataNetworkInfo `json:"network_info,required"`
	// Whether the VM is running, paused, or not started.
	State APIVmDeleteResponseDataState `json:"state,required"`
	// The parent VM's ID, if present. If None, then this VM is a root VM.
	ParentID string                      `json:"parent_id,nullable"`
	JSON     apiVmDeleteResponseDataJSON `json:"-"`
}

// apiVmDeleteResponseDataJSON contains the JSON metadata for the struct
// [APIVmDeleteResponseData]
type apiVmDeleteResponseDataJSON struct {
	ID          apijson.Field
	Children    apijson.Field
	ClusterID   apijson.Field
	IPAddress   apijson.Field
	NetworkInfo apijson.Field
	State       apijson.Field
	ParentID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmDeleteResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmDeleteResponseDataJSON) RawJSON() string {
	return r.raw
}

// The VM's network configuration
type APIVmDeleteResponseDataNetworkInfo struct {
	GuestIP     string                                 `json:"guest_ip,required"`
	GuestMac    string                                 `json:"guest_mac,required"`
	SSHPort     int64                                  `json:"ssh_port,required"`
	Tap0IP      string                                 `json:"tap0_ip,required"`
	Tap0Name    string                                 `json:"tap0_name,required"`
	VmNamespace string                                 `json:"vm_namespace,required"`
	JSON        apiVmDeleteResponseDataNetworkInfoJSON `json:"-"`
}

// apiVmDeleteResponseDataNetworkInfoJSON contains the JSON metadata for the struct
// [APIVmDeleteResponseDataNetworkInfo]
type apiVmDeleteResponseDataNetworkInfoJSON struct {
	GuestIP     apijson.Field
	GuestMac    apijson.Field
	SSHPort     apijson.Field
	Tap0IP      apijson.Field
	Tap0Name    apijson.Field
	VmNamespace apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmDeleteResponseDataNetworkInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmDeleteResponseDataNetworkInfoJSON) RawJSON() string {
	return r.raw
}

// Whether the VM is running, paused, or not started.
type APIVmDeleteResponseDataState string

const (
	APIVmDeleteResponseDataStateNotStarted APIVmDeleteResponseDataState = "Not started"
	APIVmDeleteResponseDataStateRunning    APIVmDeleteResponseDataState = "Running"
	APIVmDeleteResponseDataStatePaused     APIVmDeleteResponseDataState = "Paused"
)

func (r APIVmDeleteResponseDataState) IsKnown() bool {
	switch r {
	case APIVmDeleteResponseDataStateNotStarted, APIVmDeleteResponseDataStateRunning, APIVmDeleteResponseDataStatePaused:
		return true
	}
	return false
}

type APIVmBranchResponse struct {
	Data        APIVmBranchResponseData `json:"data,required"`
	DurationNs  int64                   `json:"duration_ns,required"`
	OperationID string                  `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                   `json:"time_start,required"`
	JSON      apiVmBranchResponseJSON `json:"-"`
}

// apiVmBranchResponseJSON contains the JSON metadata for the struct
// [APIVmBranchResponse]
type apiVmBranchResponseJSON struct {
	Data        apijson.Field
	DurationNs  apijson.Field
	OperationID apijson.Field
	TimeStart   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
	// The VM's local IP address on the VM subnet
	IPAddress string `json:"ip_address,required"`
	// The VM's network configuration
	NetworkInfo APIVmBranchResponseDataNetworkInfo `json:"network_info,required"`
	// Whether the VM is running, paused, or not started.
	State APIVmBranchResponseDataState `json:"state,required"`
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
	IPAddress   apijson.Field
	NetworkInfo apijson.Field
	State       apijson.Field
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

type APIVmCommitResponse struct {
	Data        APIVmCommitResponseData `json:"data,required"`
	DurationNs  int64                   `json:"duration_ns,required"`
	OperationID string                  `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                   `json:"time_start,required"`
	JSON      apiVmCommitResponseJSON `json:"-"`
}

// apiVmCommitResponseJSON contains the JSON metadata for the struct
// [APIVmCommitResponse]
type apiVmCommitResponseJSON struct {
	Data        apijson.Field
	DurationNs  apijson.Field
	OperationID apijson.Field
	TimeStart   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmCommitResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmCommitResponseJSON) RawJSON() string {
	return r.raw
}

type APIVmCommitResponseData struct {
	ID   string                      `json:"id,required"`
	JSON apiVmCommitResponseDataJSON `json:"-"`
}

// apiVmCommitResponseDataJSON contains the JSON metadata for the struct
// [APIVmCommitResponseData]
type apiVmCommitResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmCommitResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmCommitResponseDataJSON) RawJSON() string {
	return r.raw
}

type APIVmGetSSHKeyResponse struct {
	Data        string `json:"data,required"`
	DurationNs  int64  `json:"duration_ns,required"`
	OperationID string `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                      `json:"time_start,required"`
	JSON      apiVmGetSSHKeyResponseJSON `json:"-"`
}

// apiVmGetSSHKeyResponseJSON contains the JSON metadata for the struct
// [APIVmGetSSHKeyResponse]
type apiVmGetSSHKeyResponseJSON struct {
	Data        apijson.Field
	DurationNs  apijson.Field
	OperationID apijson.Field
	TimeStart   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmGetSSHKeyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmGetSSHKeyResponseJSON) RawJSON() string {
	return r.raw
}

type APIVmUpdateStateResponse struct {
	Data        APIVmUpdateStateResponseData `json:"data,required"`
	DurationNs  int64                        `json:"duration_ns,required"`
	OperationID string                       `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                        `json:"time_start,required"`
	JSON      apiVmUpdateStateResponseJSON `json:"-"`
}

// apiVmUpdateStateResponseJSON contains the JSON metadata for the struct
// [APIVmUpdateStateResponse]
type apiVmUpdateStateResponseJSON struct {
	Data        apijson.Field
	DurationNs  apijson.Field
	OperationID apijson.Field
	TimeStart   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmUpdateStateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmUpdateStateResponseJSON) RawJSON() string {
	return r.raw
}

type APIVmUpdateStateResponseData struct {
	// The ID of the VM.
	ID string `json:"id,required"`
	// The IDs of direct children branched from this VM.
	Children []string `json:"children,required"`
	// The VM's cluster ID
	ClusterID string `json:"cluster_id,required"`
	// The VM's local IP address on the VM subnet
	IPAddress string `json:"ip_address,required"`
	// The VM's network configuration
	NetworkInfo APIVmUpdateStateResponseDataNetworkInfo `json:"network_info,required"`
	// Whether the VM is running, paused, or not started.
	State APIVmUpdateStateResponseDataState `json:"state,required"`
	// The parent VM's ID, if present. If None, then this VM is a root VM.
	ParentID string                           `json:"parent_id,nullable"`
	JSON     apiVmUpdateStateResponseDataJSON `json:"-"`
}

// apiVmUpdateStateResponseDataJSON contains the JSON metadata for the struct
// [APIVmUpdateStateResponseData]
type apiVmUpdateStateResponseDataJSON struct {
	ID          apijson.Field
	Children    apijson.Field
	ClusterID   apijson.Field
	IPAddress   apijson.Field
	NetworkInfo apijson.Field
	State       apijson.Field
	ParentID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmUpdateStateResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmUpdateStateResponseDataJSON) RawJSON() string {
	return r.raw
}

// The VM's network configuration
type APIVmUpdateStateResponseDataNetworkInfo struct {
	GuestIP     string                                      `json:"guest_ip,required"`
	GuestMac    string                                      `json:"guest_mac,required"`
	SSHPort     int64                                       `json:"ssh_port,required"`
	Tap0IP      string                                      `json:"tap0_ip,required"`
	Tap0Name    string                                      `json:"tap0_name,required"`
	VmNamespace string                                      `json:"vm_namespace,required"`
	JSON        apiVmUpdateStateResponseDataNetworkInfoJSON `json:"-"`
}

// apiVmUpdateStateResponseDataNetworkInfoJSON contains the JSON metadata for the
// struct [APIVmUpdateStateResponseDataNetworkInfo]
type apiVmUpdateStateResponseDataNetworkInfoJSON struct {
	GuestIP     apijson.Field
	GuestMac    apijson.Field
	SSHPort     apijson.Field
	Tap0IP      apijson.Field
	Tap0Name    apijson.Field
	VmNamespace apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmUpdateStateResponseDataNetworkInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmUpdateStateResponseDataNetworkInfoJSON) RawJSON() string {
	return r.raw
}

// Whether the VM is running, paused, or not started.
type APIVmUpdateStateResponseDataState string

const (
	APIVmUpdateStateResponseDataStateNotStarted APIVmUpdateStateResponseDataState = "Not started"
	APIVmUpdateStateResponseDataStateRunning    APIVmUpdateStateResponseDataState = "Running"
	APIVmUpdateStateResponseDataStatePaused     APIVmUpdateStateResponseDataState = "Paused"
)

func (r APIVmUpdateStateResponseDataState) IsKnown() bool {
	switch r {
	case APIVmUpdateStateResponseDataStateNotStarted, APIVmUpdateStateResponseDataStateRunning, APIVmUpdateStateResponseDataStatePaused:
		return true
	}
	return false
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

type APIVmUpdateStateParams struct {
	PatchRequest PatchRequestParam `json:"patch_request,required"`
}

func (r APIVmUpdateStateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PatchRequest)
}
