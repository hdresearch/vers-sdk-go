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

// Update VM state (pause/resume)
func (r *APIVmService) Update(ctx context.Context, vmID string, body APIVmUpdateParams, opts ...option.RequestOption) (res *APIVmUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("api/vm/%s", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all VMs.
func (r *APIVmService) List(ctx context.Context, opts ...option.RequestOption) (res *[]APIVmListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/vm"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete the specified VM.
func (r *APIVmService) Delete(ctx context.Context, vmID string, body APIVmDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("api/vm/%s", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Creates a branch of the specified VM.
func (r *APIVmService) NewBranch(ctx context.Context, vmID string, body APIVmNewBranchParams, opts ...option.RequestOption) (res *APIVmNewBranchResponse, err error) {
	opts = append(r.Options[:], opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("api/vm/%s/branch", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Execute a command on the specified VM.
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

type APIVmGetResponse struct {
	// The ID of the VM.
	ID string `json:"id,required"`
	// The IDs of direct children branched from this VM.
	Children []string `json:"children,required"`
	// The VM's local IP address on the VM subnet
	IPAddress string `json:"ip_address,required"`
	// The VM's network configuration
	NetworkInfo APIVmGetResponseNetworkInfo `json:"network_info,required"`
	// Whether the VM is running, paused, or not started.
	State APIVmGetResponseState `json:"state,required"`
	// The parent VM's ID, if present. If None, then this VM is a root VM.
	ParentID string               `json:"parent_id,nullable"`
	JSON     apiVmGetResponseJSON `json:"-"`
}

// apiVmGetResponseJSON contains the JSON metadata for the struct
// [APIVmGetResponse]
type apiVmGetResponseJSON struct {
	ID          apijson.Field
	Children    apijson.Field
	IPAddress   apijson.Field
	NetworkInfo apijson.Field
	State       apijson.Field
	ParentID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmGetResponseJSON) RawJSON() string {
	return r.raw
}

// The VM's network configuration
type APIVmGetResponseNetworkInfo struct {
	GuestIP     string                          `json:"guest_ip,required"`
	GuestMac    string                          `json:"guest_mac,required"`
	Tap0IP      string                          `json:"tap0_ip,required"`
	Tap0Name    string                          `json:"tap0_name,required"`
	VmNamespace string                          `json:"vm_namespace,required"`
	JSON        apiVmGetResponseNetworkInfoJSON `json:"-"`
}

// apiVmGetResponseNetworkInfoJSON contains the JSON metadata for the struct
// [APIVmGetResponseNetworkInfo]
type apiVmGetResponseNetworkInfoJSON struct {
	GuestIP     apijson.Field
	GuestMac    apijson.Field
	Tap0IP      apijson.Field
	Tap0Name    apijson.Field
	VmNamespace apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmGetResponseNetworkInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmGetResponseNetworkInfoJSON) RawJSON() string {
	return r.raw
}

// Whether the VM is running, paused, or not started.
type APIVmGetResponseState string

const (
	APIVmGetResponseStateNotStarted APIVmGetResponseState = "Not started"
	APIVmGetResponseStateRunning    APIVmGetResponseState = "Running"
	APIVmGetResponseStatePaused     APIVmGetResponseState = "Paused"
)

func (r APIVmGetResponseState) IsKnown() bool {
	switch r {
	case APIVmGetResponseStateNotStarted, APIVmGetResponseStateRunning, APIVmGetResponseStatePaused:
		return true
	}
	return false
}

type APIVmUpdateResponse struct {
	// The ID of the VM.
	ID string `json:"id,required"`
	// The IDs of direct children branched from this VM.
	Children []string `json:"children,required"`
	// The VM's local IP address on the VM subnet
	IPAddress string `json:"ip_address,required"`
	// The VM's network configuration
	NetworkInfo APIVmUpdateResponseNetworkInfo `json:"network_info,required"`
	// Whether the VM is running, paused, or not started.
	State APIVmUpdateResponseState `json:"state,required"`
	// The parent VM's ID, if present. If None, then this VM is a root VM.
	ParentID string                  `json:"parent_id,nullable"`
	JSON     apiVmUpdateResponseJSON `json:"-"`
}

// apiVmUpdateResponseJSON contains the JSON metadata for the struct
// [APIVmUpdateResponse]
type apiVmUpdateResponseJSON struct {
	ID          apijson.Field
	Children    apijson.Field
	IPAddress   apijson.Field
	NetworkInfo apijson.Field
	State       apijson.Field
	ParentID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The VM's network configuration
type APIVmUpdateResponseNetworkInfo struct {
	GuestIP     string                             `json:"guest_ip,required"`
	GuestMac    string                             `json:"guest_mac,required"`
	Tap0IP      string                             `json:"tap0_ip,required"`
	Tap0Name    string                             `json:"tap0_name,required"`
	VmNamespace string                             `json:"vm_namespace,required"`
	JSON        apiVmUpdateResponseNetworkInfoJSON `json:"-"`
}

// apiVmUpdateResponseNetworkInfoJSON contains the JSON metadata for the struct
// [APIVmUpdateResponseNetworkInfo]
type apiVmUpdateResponseNetworkInfoJSON struct {
	GuestIP     apijson.Field
	GuestMac    apijson.Field
	Tap0IP      apijson.Field
	Tap0Name    apijson.Field
	VmNamespace apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmUpdateResponseNetworkInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmUpdateResponseNetworkInfoJSON) RawJSON() string {
	return r.raw
}

// Whether the VM is running, paused, or not started.
type APIVmUpdateResponseState string

const (
	APIVmUpdateResponseStateNotStarted APIVmUpdateResponseState = "Not started"
	APIVmUpdateResponseStateRunning    APIVmUpdateResponseState = "Running"
	APIVmUpdateResponseStatePaused     APIVmUpdateResponseState = "Paused"
)

func (r APIVmUpdateResponseState) IsKnown() bool {
	switch r {
	case APIVmUpdateResponseStateNotStarted, APIVmUpdateResponseStateRunning, APIVmUpdateResponseStatePaused:
		return true
	}
	return false
}

type APIVmListResponse struct {
	// The ID of the VM.
	ID string `json:"id,required"`
	// The IDs of direct children branched from this VM.
	Children []string `json:"children,required"`
	// The VM's local IP address on the VM subnet
	IPAddress string `json:"ip_address,required"`
	// The VM's network configuration
	NetworkInfo APIVmListResponseNetworkInfo `json:"network_info,required"`
	// Whether the VM is running, paused, or not started.
	State APIVmListResponseState `json:"state,required"`
	// The parent VM's ID, if present. If None, then this VM is a root VM.
	ParentID string                `json:"parent_id,nullable"`
	JSON     apiVmListResponseJSON `json:"-"`
}

// apiVmListResponseJSON contains the JSON metadata for the struct
// [APIVmListResponse]
type apiVmListResponseJSON struct {
	ID          apijson.Field
	Children    apijson.Field
	IPAddress   apijson.Field
	NetworkInfo apijson.Field
	State       apijson.Field
	ParentID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmListResponseJSON) RawJSON() string {
	return r.raw
}

// The VM's network configuration
type APIVmListResponseNetworkInfo struct {
	GuestIP     string                           `json:"guest_ip,required"`
	GuestMac    string                           `json:"guest_mac,required"`
	Tap0IP      string                           `json:"tap0_ip,required"`
	Tap0Name    string                           `json:"tap0_name,required"`
	VmNamespace string                           `json:"vm_namespace,required"`
	JSON        apiVmListResponseNetworkInfoJSON `json:"-"`
}

// apiVmListResponseNetworkInfoJSON contains the JSON metadata for the struct
// [APIVmListResponseNetworkInfo]
type apiVmListResponseNetworkInfoJSON struct {
	GuestIP     apijson.Field
	GuestMac    apijson.Field
	Tap0IP      apijson.Field
	Tap0Name    apijson.Field
	VmNamespace apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmListResponseNetworkInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmListResponseNetworkInfoJSON) RawJSON() string {
	return r.raw
}

// Whether the VM is running, paused, or not started.
type APIVmListResponseState string

const (
	APIVmListResponseStateNotStarted APIVmListResponseState = "Not started"
	APIVmListResponseStateRunning    APIVmListResponseState = "Running"
	APIVmListResponseStatePaused     APIVmListResponseState = "Paused"
)

func (r APIVmListResponseState) IsKnown() bool {
	switch r {
	case APIVmListResponseStateNotStarted, APIVmListResponseStateRunning, APIVmListResponseStatePaused:
		return true
	}
	return false
}

type APIVmNewBranchResponse struct {
	// The ID of the VM.
	ID string `json:"id,required"`
	// The IDs of direct children branched from this VM.
	Children []string `json:"children,required"`
	// The VM's local IP address on the VM subnet
	IPAddress string `json:"ip_address,required"`
	// The VM's network configuration
	NetworkInfo APIVmNewBranchResponseNetworkInfo `json:"network_info,required"`
	// Whether the VM is running, paused, or not started.
	State APIVmNewBranchResponseState `json:"state,required"`
	// The parent VM's ID, if present. If None, then this VM is a root VM.
	ParentID string                     `json:"parent_id,nullable"`
	JSON     apiVmNewBranchResponseJSON `json:"-"`
}

// apiVmNewBranchResponseJSON contains the JSON metadata for the struct
// [APIVmNewBranchResponse]
type apiVmNewBranchResponseJSON struct {
	ID          apijson.Field
	Children    apijson.Field
	IPAddress   apijson.Field
	NetworkInfo apijson.Field
	State       apijson.Field
	ParentID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmNewBranchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmNewBranchResponseJSON) RawJSON() string {
	return r.raw
}

// The VM's network configuration
type APIVmNewBranchResponseNetworkInfo struct {
	GuestIP     string                                `json:"guest_ip,required"`
	GuestMac    string                                `json:"guest_mac,required"`
	Tap0IP      string                                `json:"tap0_ip,required"`
	Tap0Name    string                                `json:"tap0_name,required"`
	VmNamespace string                                `json:"vm_namespace,required"`
	JSON        apiVmNewBranchResponseNetworkInfoJSON `json:"-"`
}

// apiVmNewBranchResponseNetworkInfoJSON contains the JSON metadata for the struct
// [APIVmNewBranchResponseNetworkInfo]
type apiVmNewBranchResponseNetworkInfoJSON struct {
	GuestIP     apijson.Field
	GuestMac    apijson.Field
	Tap0IP      apijson.Field
	Tap0Name    apijson.Field
	VmNamespace apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIVmNewBranchResponseNetworkInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiVmNewBranchResponseNetworkInfoJSON) RawJSON() string {
	return r.raw
}

// Whether the VM is running, paused, or not started.
type APIVmNewBranchResponseState string

const (
	APIVmNewBranchResponseStateNotStarted APIVmNewBranchResponseState = "Not started"
	APIVmNewBranchResponseStateRunning    APIVmNewBranchResponseState = "Running"
	APIVmNewBranchResponseStatePaused     APIVmNewBranchResponseState = "Paused"
)

func (r APIVmNewBranchResponseState) IsKnown() bool {
	switch r {
	case APIVmNewBranchResponseStateNotStarted, APIVmNewBranchResponseStateRunning, APIVmNewBranchResponseStatePaused:
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

type APIVmUpdateParams struct {
	Body APIVmUpdateParamsBodyUnion `json:"body,required"`
}

func (r APIVmUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type APIVmUpdateParamsBody struct {
	Action param.Field[APIVmUpdateParamsBodyAction] `json:"action,required"`
}

func (r APIVmUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r APIVmUpdateParamsBody) implementsAPIVmUpdateParamsBodyUnion() {}

// Satisfied by [APIVmUpdateParamsBodyAction], [APIVmUpdateParamsBodyAction],
// [APIVmUpdateParamsBody].
type APIVmUpdateParamsBodyUnion interface {
	implementsAPIVmUpdateParamsBodyUnion()
}

type APIVmUpdateParamsBodyAction struct {
	Action param.Field[APIVmUpdateParamsBodyActionAction] `json:"action,required"`
}

func (r APIVmUpdateParamsBodyAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r APIVmUpdateParamsBodyAction) implementsAPIVmUpdateParamsBodyUnion() {}

type APIVmUpdateParamsBodyActionAction string

const (
	APIVmUpdateParamsBodyActionActionPause APIVmUpdateParamsBodyActionAction = "pause"
)

func (r APIVmUpdateParamsBodyActionAction) IsKnown() bool {
	switch r {
	case APIVmUpdateParamsBodyActionActionPause:
		return true
	}
	return false
}

type APIVmDeleteParams struct {
	// Optionally delete all child VMs recursively
	Recursive param.Field[bool] `query:"recursive"`
}

// URLQuery serializes [APIVmDeleteParams]'s query parameters as `url.Values`.
func (r APIVmDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
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
