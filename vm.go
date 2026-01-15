// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/hdresearch/vers-sdk-go/internal/apijson"
	"github.com/hdresearch/vers-sdk-go/internal/apiquery"
	"github.com/hdresearch/vers-sdk-go/internal/param"
	"github.com/hdresearch/vers-sdk-go/internal/requestconfig"
	"github.com/hdresearch/vers-sdk-go/option"
)

// VmService contains methods and other services that help with interacting with
// the vers API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVmService] method instead.
type VmService struct {
	Options []option.RequestOption
}

// NewVmService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVmService(opts ...option.RequestOption) (r *VmService) {
	r = &VmService{}
	r.Options = opts
	return
}

func (r *VmService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Vm, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/vms"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *VmService) Delete(ctx context.Context, vmID string, body VmDeleteParams, opts ...option.RequestOption) (res *VmDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/vm/%s", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

func (r *VmService) Commit(ctx context.Context, vmID string, body VmCommitParams, opts ...option.RequestOption) (res *VmCommitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/vm/%s/commit", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *VmService) NewRoot(ctx context.Context, params VmNewRootParams, opts ...option.RequestOption) (res *NewVmResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/vm/new_root"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

func (r *VmService) GetSSHKey(ctx context.Context, vmID string, opts ...option.RequestOption) (res *VmSSHKeyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/vm/%s/ssh_key", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *VmService) RestoreFromCommit(ctx context.Context, body VmRestoreFromCommitParams, opts ...option.RequestOption) (res *NewVmResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/vm/from_commit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *VmService) Status(ctx context.Context, vmID string, opts ...option.RequestOption) (res *Vm, err error) {
	opts = slices.Concat(r.Options, opts)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/vm/%s/status", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *VmService) UpdateState(ctx context.Context, vmID string, params VmUpdateStateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/vm/%s/state", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, nil, opts...)
	return
}

type NewRootRequestParam struct {
	// Struct representing configuration options common to all VMs
	VmConfig param.Field[NewRootRequestVmConfigParam] `json:"vm_config,required"`
}

func (r NewRootRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Struct representing configuration options common to all VMs
type NewRootRequestVmConfigParam struct {
	// The disk size, in MiB.
	FsSizeMib param.Field[int64] `json:"fs_size_mib"`
	// The filesystem base image name. Currently, must be 'default'
	ImageName param.Field[string] `json:"image_name"`
	// The kernel name. Currently, must be 'default.bin'
	KernelName param.Field[string] `json:"kernel_name"`
	// The RAM size, in MiB.
	MemSizeMib param.Field[int64] `json:"mem_size_mib"`
	// How many vCPUs to allocate to this VM (and its children)
	VcpuCount param.Field[int64] `json:"vcpu_count"`
}

func (r NewRootRequestVmConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Response body for new VM requests (new_root, from_commit, branch)
type NewVmResponse struct {
	VmID string            `json:"vm_id,required"`
	JSON newVmResponseJSON `json:"-"`
}

// newVmResponseJSON contains the JSON metadata for the struct [NewVmResponse]
type newVmResponseJSON struct {
	VmID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NewVmResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r newVmResponseJSON) RawJSON() string {
	return r.raw
}

type Vm struct {
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	OwnerID   string    `json:"owner_id,required" format:"uuid"`
	// The state of a VM
	State VmState `json:"state,required"`
	VmID  string  `json:"vm_id,required" format:"uuid"`
	JSON  vmJSON  `json:"-"`
}

// vmJSON contains the JSON metadata for the struct [Vm]
type vmJSON struct {
	CreatedAt   apijson.Field
	OwnerID     apijson.Field
	State       apijson.Field
	VmID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Vm) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vmJSON) RawJSON() string {
	return r.raw
}

// The state of a VM
type VmState string

const (
	VmStateBooting VmState = "booting"
	VmStateRunning VmState = "running"
	VmStatePaused  VmState = "paused"
)

func (r VmState) IsKnown() bool {
	switch r {
	case VmStateBooting, VmStateRunning, VmStatePaused:
		return true
	}
	return false
}

// The response body for POST /api/vm/{vm_id}/commit
type VmCommitResponse struct {
	// The UUID of the newly-created commit
	CommitID string               `json:"commit_id,required" format:"uuid"`
	JSON     vmCommitResponseJSON `json:"-"`
}

// vmCommitResponseJSON contains the JSON metadata for the struct
// [VmCommitResponse]
type vmCommitResponseJSON struct {
	CommitID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VmCommitResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vmCommitResponseJSON) RawJSON() string {
	return r.raw
}

// Response body for DELETE /api/vm/{vm_id}
type VmDeleteResponse struct {
	VmID string               `json:"vm_id,required"`
	JSON vmDeleteResponseJSON `json:"-"`
}

// vmDeleteResponseJSON contains the JSON metadata for the struct
// [VmDeleteResponse]
type vmDeleteResponseJSON struct {
	VmID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VmDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vmDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Request body for POST /api/v1/vm/from_commit
type VmFromCommitRequestParam struct {
	CommitID param.Field[string] `json:"commit_id,required" format:"uuid"`
}

func (r VmFromCommitRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Response body for GET /api/vm/{vm_id}/ssh_key
type VmSSHKeyResponse struct {
	// The SSH port that will be DNAT'd to the VM's netns (and, in turn, to its TAP
	// device)
	SSHPort int64 `json:"ssh_port,required"`
	// Private SSH key in stringified OpenSSH format
	SSHPrivateKey string               `json:"ssh_private_key,required"`
	JSON          vmSSHKeyResponseJSON `json:"-"`
}

// vmSSHKeyResponseJSON contains the JSON metadata for the struct
// [VmSSHKeyResponse]
type vmSSHKeyResponseJSON struct {
	SSHPort       apijson.Field
	SSHPrivateKey apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VmSSHKeyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vmSSHKeyResponseJSON) RawJSON() string {
	return r.raw
}

// Request body for PATCH /api/vm/{vm_id}/state
type VmUpdateStateRequestParam struct {
	// The requested state for the VM
	State param.Field[VmUpdateStateRequestState] `json:"state,required"`
}

func (r VmUpdateStateRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The requested state for the VM
type VmUpdateStateRequestState string

const (
	VmUpdateStateRequestStatePaused  VmUpdateStateRequestState = "Paused"
	VmUpdateStateRequestStateRunning VmUpdateStateRequestState = "Running"
)

func (r VmUpdateStateRequestState) IsKnown() bool {
	switch r {
	case VmUpdateStateRequestStatePaused, VmUpdateStateRequestStateRunning:
		return true
	}
	return false
}

type VmDeleteParams struct {
	// If true, return an error immediately if the VM is still booting. Default: false
	SkipWaitBoot param.Field[bool] `query:"skip_wait_boot"`
}

// URLQuery serializes [VmDeleteParams]'s query parameters as `url.Values`.
func (r VmDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VmCommitParams struct {
	// If true, keep VM paused after commit
	KeepPaused param.Field[bool] `query:"keep_paused"`
	// If true, return an error immediately if the VM is still booting. Default: false
	SkipWaitBoot param.Field[bool] `query:"skip_wait_boot"`
}

// URLQuery serializes [VmCommitParams]'s query parameters as `url.Values`.
func (r VmCommitParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VmNewRootParams struct {
	NewRootRequest NewRootRequestParam `json:"new_root_request,required"`
	// If true, wait for the newly-created VM to finish booting before returning.
	// Default: false.
	WaitBoot param.Field[bool] `query:"wait_boot"`
}

func (r VmNewRootParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.NewRootRequest)
}

// URLQuery serializes [VmNewRootParams]'s query parameters as `url.Values`.
func (r VmNewRootParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VmRestoreFromCommitParams struct {
	// Request body for POST /api/v1/vm/from_commit
	VmFromCommitRequest VmFromCommitRequestParam `json:"vm_from_commit_request,required"`
}

func (r VmRestoreFromCommitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.VmFromCommitRequest)
}

type VmUpdateStateParams struct {
	// Request body for PATCH /api/vm/{vm_id}/state
	VmUpdateStateRequest VmUpdateStateRequestParam `json:"vm_update_state_request,required"`
	// If true, error immediately if the VM is not finished booting. Defaults to false
	SkipWaitBoot param.Field[bool] `query:"skip_wait_boot"`
}

func (r VmUpdateStateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.VmUpdateStateRequest)
}

// URLQuery serializes [VmUpdateStateParams]'s query parameters as `url.Values`.
func (r VmUpdateStateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
