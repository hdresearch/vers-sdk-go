// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/hdresearch/vers-sdk-go/internal/apijson"
	"github.com/hdresearch/vers-sdk-go/internal/param"
	"github.com/hdresearch/vers-sdk-go/internal/requestconfig"
	"github.com/hdresearch/vers-sdk-go/option"
)

// OrchestratorVmService contains methods and other services that help with
// interacting with the vers API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrchestratorVmService] method instead.
type OrchestratorVmService struct {
	Options []option.RequestOption
}

// NewOrchestratorVmService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrchestratorVmService(opts ...option.RequestOption) (r *OrchestratorVmService) {
	r = &OrchestratorVmService{}
	r.Options = opts
	return
}

func (r *OrchestratorVmService) Delete(ctx context.Context, vmID string, opts ...option.RequestOption) (res *VmDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("vm/%s", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

func (r *OrchestratorVmService) Branch(ctx context.Context, vmID string, opts ...option.RequestOption) (res *NewVmResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("vm/%s/branch", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

func (r *OrchestratorVmService) Commit(ctx context.Context, vmID string, opts ...option.RequestOption) (res *VmCommitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("vm/%s/commit", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

func (r *OrchestratorVmService) NewRoot(ctx context.Context, body OrchestratorVmNewRootParams, opts ...option.RequestOption) (res *NewVmResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "vm/new_root"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *OrchestratorVmService) RestoreFromCommit(ctx context.Context, body OrchestratorVmRestoreFromCommitParams, opts ...option.RequestOption) (res *NewVmResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "vm/from_commit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *OrchestratorVmService) UpdateState(ctx context.Context, vmID string, body OrchestratorVmUpdateStateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("vm/%s/state", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
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

type NewVmResponse struct {
	// The VM ID, a (v4) UUID
	ID   string            `json:"id,required"`
	JSON newVmResponseJSON `json:"-"`
}

// newVmResponseJSON contains the JSON metadata for the struct [NewVmResponse]
type newVmResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NewVmResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r newVmResponseJSON) RawJSON() string {
	return r.raw
}

// The response body for POST /api/vm/{vm_id}/commit
type VmCommitResponse struct {
	// The commit ID, a (v4) UUID
	CommitID string `json:"commit_id,required"`
	// The host architecture, eg: "x86_64" (currently implemented with `uname -m“)
	HostArchitecture string               `json:"host_architecture,required"`
	JSON             vmCommitResponseJSON `json:"-"`
}

// vmCommitResponseJSON contains the JSON metadata for the struct
// [VmCommitResponse]
type vmCommitResponseJSON struct {
	CommitID         apijson.Field
	HostArchitecture apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *VmCommitResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vmCommitResponseJSON) RawJSON() string {
	return r.raw
}

// Response body for DELETE /api/vm/{vm_id}
type VmDeleteResponse struct {
	DeletedIDs []string             `json:"deleted_ids,required"`
	JSON       vmDeleteResponseJSON `json:"-"`
}

// vmDeleteResponseJSON contains the JSON metadata for the struct
// [VmDeleteResponse]
type vmDeleteResponseJSON struct {
	DeletedIDs  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VmDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vmDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Request body for POST /api/vm/from_commit
type VmFromCommitRequestParam struct {
	CommitID param.Field[string] `json:"commit_id,required" format:"uuid"`
}

func (r VmFromCommitRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

type OrchestratorVmNewRootParams struct {
	NewRootRequest NewRootRequestParam `json:"new_root_request,required"`
}

func (r OrchestratorVmNewRootParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.NewRootRequest)
}

type OrchestratorVmRestoreFromCommitParams struct {
	// Request body for POST /api/vm/from_commit
	VmFromCommitRequest VmFromCommitRequestParam `json:"vm_from_commit_request,required"`
}

func (r OrchestratorVmRestoreFromCommitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.VmFromCommitRequest)
}

type OrchestratorVmUpdateStateParams struct {
	// Request body for PATCH /api/vm/{vm_id}/state
	VmUpdateStateRequest VmUpdateStateRequestParam `json:"vm_update_state_request,required"`
}

func (r OrchestratorVmUpdateStateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.VmUpdateStateRequest)
}
