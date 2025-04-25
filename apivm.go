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
func (r *APIVmService) Get(ctx context.Context, vmID string, opts ...option.RequestOption) (res *Vm, err error) {
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
func (r *APIVmService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Vm, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/vm"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a VM.
func (r *APIVmService) Delete(ctx context.Context, vmID string, body APIVmDeleteParams, opts ...option.RequestOption) (res *Vm, err error) {
	opts = append(r.Options[:], opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("api/vm/%s", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Commit a VM.
func (r *APIVmService) Commit(ctx context.Context, vmID string, body APIVmCommitParams, opts ...option.RequestOption) (res *Vm, err error) {
	opts = append(r.Options[:], opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("api/vm/%s/commit", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Branch a VM.
func (r *APIVmService) NewBranch(ctx context.Context, vmID string, body APIVmNewBranchParams, opts ...option.RequestOption) (res *Vm, err error) {
	opts = append(r.Options[:], opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("api/vm/%s/branch", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Execute a command in a VM.
func (r *APIVmService) Execute(ctx context.Context, vmID string, body APIVmExecuteParams, opts ...option.RequestOption) (res *APIVmExecuteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("api/vm/%s/execute", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
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

type APIVmExecuteResponse struct {
	ID            string                            `json:"id,required"`
	CommandResult APIVmExecuteResponseCommandResult `json:"command_result,required"`
	JSON          apiVmExecuteResponseJSON          `json:"-"`
}

// apiVmExecuteResponseJSON contains the JSON metadata for the struct
// [APIVmExecuteResponse]
type apiVmExecuteResponseJSON struct {
	ID            apijson.Field
	CommandResult apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *APIVmExecuteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmExecuteResponseJSON) RawJSON() string {
	return r.raw
}

type APIVmExecuteResponseCommandResult struct {
	ExitCode int64                                 `json:"exit_code,required"`
	Stderr   string                                `json:"stderr,required"`
	Stdout   string                                `json:"stdout,required"`
	JSON     apiVmExecuteResponseCommandResultJSON `json:"-"`
}

// apiVmExecuteResponseCommandResultJSON contains the JSON metadata for the struct
// [APIVmExecuteResponseCommandResult]
type apiVmExecuteResponseCommandResultJSON struct {
	ExitCode    apijson.Field
	Stderr      apijson.Field
	Stdout      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmExecuteResponseCommandResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmExecuteResponseCommandResultJSON) RawJSON() string {
	return r.raw
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

type APIVmCommitParams struct {
	Body interface{} `json:"body,required"`
}

func (r APIVmCommitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type APIVmNewBranchParams struct {
	Body interface{} `json:"body,required"`
}

func (r APIVmNewBranchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type APIVmExecuteParams struct {
	Command param.Field[string] `json:"command,required"`
}

func (r APIVmExecuteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
