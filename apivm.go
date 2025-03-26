// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vers

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

// Attempt to update VM state. Will return a 409 if you attempt to resume a VM
// which has children. Parent VMs must not be mutated, in order to ensure all
// children are derived from the same snapshot.
func (r *APIVmService) Update(ctx context.Context, vmID string, body APIVmUpdateParams, opts ...option.RequestOption) (res *Vm, err error) {
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
func (r *APIVmService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Vm, err error) {
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

type Vm struct {
	// The ID of the VM.
	ID string `json:"id,required"`
	// The IDs of direct children branched from this VM.
	Children []string `json:"children,required"`
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
	Tap0IP      string            `json:"tap0_ip,required"`
	Tap0Name    string            `json:"tap0_name,required"`
	VmNamespace string            `json:"vm_namespace,required"`
	JSON        vmNetworkInfoJSON `json:"-"`
}

// vmNetworkInfoJSON contains the JSON metadata for the struct [VmNetworkInfo]
type vmNetworkInfoJSON struct {
	GuestIP     apijson.Field
	GuestMac    apijson.Field
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

type APIVmUpdateParams struct {
	Body APIVmUpdateParamsBody `json:"body,required"`
}

func (r APIVmUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type APIVmUpdateParamsBody string

const (
	APIVmUpdateParamsBodyPause  APIVmUpdateParamsBody = "pause"
	APIVmUpdateParamsBodyResume APIVmUpdateParamsBody = "resume"
)

func (r APIVmUpdateParamsBody) IsKnown() bool {
	switch r {
	case APIVmUpdateParamsBodyPause, APIVmUpdateParamsBodyResume:
		return true
	}
	return false
}

type APIVmDeleteParams struct {
	// Optionally stop all child VMs recursively
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
