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
	return res, err
}

func (r *VmService) Delete(ctx context.Context, vmID string, body VmDeleteParams, opts ...option.RequestOption) (res *VmDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/vm/%s", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

func (r *VmService) Branch(ctx context.Context, vmOrCommitID string, body VmBranchParams, opts ...option.RequestOption) (res *NewVmsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if vmOrCommitID == "" {
		err = errors.New("missing required vm_or_commit_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/vm/%s/branch", vmOrCommitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *VmService) BranchByCommit(ctx context.Context, commitID string, body VmBranchByCommitParams, opts ...option.RequestOption) (res *NewVmsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if commitID == "" {
		err = errors.New("missing required commit_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/vm/branch/by_commit/%s", commitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *VmService) BranchByTag(ctx context.Context, tagName string, body VmBranchByTagParams, opts ...option.RequestOption) (res *NewVmsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if tagName == "" {
		err = errors.New("missing required tag_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/vm/branch/by_tag/%s", tagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *VmService) BranchByVm(ctx context.Context, vmID string, body VmBranchByVmParams, opts ...option.RequestOption) (res *NewVmsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/vm/branch/by_vm/%s", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *VmService) Commit(ctx context.Context, vmID string, params VmCommitParams, opts ...option.RequestOption) (res *VmCommitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/vm/%s/commit", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

func (r *VmService) NewRoot(ctx context.Context, params VmNewRootParams, opts ...option.RequestOption) (res *NewVmResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/vm/new_root"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

func (r *VmService) Exec(ctx context.Context, vmID string, body VmExecParams, opts ...option.RequestOption) (res *VmExecResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/vm/%s/exec", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *VmService) ExecStream(ctx context.Context, vmID string, body VmExecStreamParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/vm/%s/exec/stream", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

func (r *VmService) ExecStreamAttach(ctx context.Context, vmID string, body VmExecStreamAttachParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/vm/%s/exec/stream/attach", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

func (r *VmService) GetLogs(ctx context.Context, vmID string, query VmGetLogsParams, opts ...option.RequestOption) (res *VmExecLogResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/vm/%s/logs", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *VmService) GetMetadata(ctx context.Context, vmID string, opts ...option.RequestOption) (res *VmMetadataResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/vm/%s/metadata", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *VmService) GetSSHKey(ctx context.Context, vmID string, opts ...option.RequestOption) (res *VmSSHKeyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/vm/%s/ssh_key", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *VmService) ResizeDisk(ctx context.Context, vmID string, params VmResizeDiskParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/vm/%s/disk", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, nil, opts...)
	return err
}

func (r *VmService) RestoreFromCommit(ctx context.Context, body VmRestoreFromCommitParams, opts ...option.RequestOption) (res *NewVmResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/vm/from_commit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *VmService) Status(ctx context.Context, vmID string, opts ...option.RequestOption) (res *Vm, err error) {
	opts = slices.Concat(r.Options, opts)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/vm/%s/status", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *VmService) UpdateState(ctx context.Context, vmID string, params VmUpdateStateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/vm/%s/state", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, nil, opts...)
	return err
}

type NewRootRequestParam struct {
	// Struct representing configuration options common to all VMs
	VmConfig param.Field[NewRootRequestVmConfigParam] `json:"vm_config" api:"required"`
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
	KernelName param.Field[string]            `json:"kernel_name"`
	Labels     param.Field[map[string]string] `json:"labels"`
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
	VmID string            `json:"vm_id" api:"required"`
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

type NewVmsResponse struct {
	Vms  []NewVmResponse    `json:"vms" api:"required"`
	JSON newVmsResponseJSON `json:"-"`
}

// newVmsResponseJSON contains the JSON metadata for the struct [NewVmsResponse]
type newVmsResponseJSON struct {
	Vms         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NewVmsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r newVmsResponseJSON) RawJSON() string {
	return r.raw
}

type Vm struct {
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	OwnerID   string    `json:"owner_id" api:"required" format:"uuid"`
	// The state of a VM
	State  VmState           `json:"state" api:"required"`
	VmID   string            `json:"vm_id" api:"required" format:"uuid"`
	Labels map[string]string `json:"labels" api:"nullable"`
	JSON   vmJSON            `json:"-"`
}

// vmJSON contains the JSON metadata for the struct [Vm]
type vmJSON struct {
	CreatedAt   apijson.Field
	OwnerID     apijson.Field
	State       apijson.Field
	VmID        apijson.Field
	Labels      apijson.Field
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
	VmStateBooting  VmState = "booting"
	VmStateRunning  VmState = "running"
	VmStatePaused   VmState = "paused"
	VmStateSleeping VmState = "sleeping"
	VmStateDead     VmState = "dead"
)

func (r VmState) IsKnown() bool {
	switch r {
	case VmStateBooting, VmStateRunning, VmStatePaused, VmStateSleeping, VmStateDead:
		return true
	}
	return false
}

// The response body for POST /api/vm/{vm_id}/commit
type VmCommitResponse struct {
	// The UUID of the newly-created commit
	CommitID string               `json:"commit_id" api:"required" format:"uuid"`
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
	VmID string               `json:"vm_id" api:"required"`
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

// Response for exec log tail requests.
type VmExecLogResponse struct {
	// Returned log entries.
	Entries []VmExecLogResponseEntry `json:"entries" api:"required"`
	// True when the end of file was reached.
	Eof bool `json:"eof" api:"required"`
	// Next byte offset to continue from.
	NextOffset int64                 `json:"next_offset" api:"required"`
	JSON       vmExecLogResponseJSON `json:"-"`
}

// vmExecLogResponseJSON contains the JSON metadata for the struct
// [VmExecLogResponse]
type vmExecLogResponseJSON struct {
	Entries     apijson.Field
	Eof         apijson.Field
	NextOffset  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VmExecLogResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vmExecLogResponseJSON) RawJSON() string {
	return r.raw
}

// Individual log entry describing emitted stdout/stderr chunk.
type VmExecLogResponseEntry struct {
	// Base64-encoded bytes from stdout/stderr chunk.
	DataB64 string `json:"data_b64" api:"required"`
	// Streams available for exec logging.
	Stream    VmExecLogResponseEntriesStream `json:"stream" api:"required"`
	Timestamp string                         `json:"timestamp" api:"required"`
	ExecID    string                         `json:"exec_id" api:"nullable" format:"uuid"`
	JSON      vmExecLogResponseEntryJSON     `json:"-"`
}

// vmExecLogResponseEntryJSON contains the JSON metadata for the struct
// [VmExecLogResponseEntry]
type vmExecLogResponseEntryJSON struct {
	DataB64     apijson.Field
	Stream      apijson.Field
	Timestamp   apijson.Field
	ExecID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VmExecLogResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vmExecLogResponseEntryJSON) RawJSON() string {
	return r.raw
}

// Streams available for exec logging.
type VmExecLogResponseEntriesStream string

const (
	VmExecLogResponseEntriesStreamStdout VmExecLogResponseEntriesStream = "stdout"
	VmExecLogResponseEntriesStreamStderr VmExecLogResponseEntriesStream = "stderr"
)

func (r VmExecLogResponseEntriesStream) IsKnown() bool {
	switch r {
	case VmExecLogResponseEntriesStreamStdout, VmExecLogResponseEntriesStreamStderr:
		return true
	}
	return false
}

// Request body for POST /api/vm/{vm_id}/exec
type VmExecRequestParam struct {
	// Command and arguments to execute.
	Command param.Field[[]string] `json:"command" api:"required"`
	// Optional environment variables to set for the process.
	Env param.Field[map[string]string] `json:"env"`
	// Optional exec identifier for tracking.
	ExecID param.Field[string] `json:"exec_id" format:"uuid"`
	// Optional stdin payload passed to the command.
	Stdin param.Field[string] `json:"stdin"`
	// Timeout in seconds (0 = no timeout).
	TimeoutSecs param.Field[int64] `json:"timeout_secs"`
	// Optional working directory for the command.
	WorkingDir param.Field[string] `json:"working_dir"`
}

func (r VmExecRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Response body for POST /api/vm/{vm_id}/exec
type VmExecResponse struct {
	// Exit code returned by the command.
	ExitCode int64 `json:"exit_code" api:"required"`
	// UTF-8 decoded stderr (lossy).
	Stderr string `json:"stderr" api:"required"`
	// UTF-8 decoded stdout (lossy).
	Stdout string `json:"stdout" api:"required"`
	// Exec identifier associated with this run.
	ExecID string             `json:"exec_id" api:"nullable" format:"uuid"`
	JSON   vmExecResponseJSON `json:"-"`
}

// vmExecResponseJSON contains the JSON metadata for the struct [VmExecResponse]
type vmExecResponseJSON struct {
	ExitCode    apijson.Field
	Stderr      apijson.Field
	Stdout      apijson.Field
	ExecID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VmExecResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vmExecResponseJSON) RawJSON() string {
	return r.raw
}

// Request body for POST /api/vm/{vm_id}/exec/stream/attach
type VmExecStreamAttachRequestParam struct {
	// Identifier of the exec stream session to reattach to.
	ExecID param.Field[string] `json:"exec_id" api:"required" format:"uuid"`
	// Optional cursor to resume from (exclusive). If omitted, the full retained
	// backlog is replayed.
	Cursor param.Field[int64] `json:"cursor"`
	// Start streaming after the latest retained chunk (ignores cursor).
	FromLatest param.Field[bool] `json:"from_latest"`
}

func (r VmExecStreamAttachRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Request body for POST /api/v1/vm/from_commit
type VmFromCommitRequestParam struct {
	// The commit ID to restore from
	CommitID param.Field[string] `json:"commit_id" format:"uuid"`
	// A repository reference in "repo_name:tag_name" format
	Ref param.Field[string] `json:"ref"`
	// The tag name to restore from (legacy org-scoped tag)
	TagName param.Field[string] `json:"tag_name"`
}

func (r VmFromCommitRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VmFromCommitRequestParam) implementsVmFromCommitRequestUnionParam() {}

// Request body for POST /api/v1/vm/from_commit
//
// Satisfied by [VmFromCommitRequestCommitIDParam],
// [VmFromCommitRequestTagNameParam], [VmFromCommitRequestRefParam],
// [VmFromCommitRequestParam].
type VmFromCommitRequestUnionParam interface {
	implementsVmFromCommitRequestUnionParam()
}

// The commit ID to restore from
type VmFromCommitRequestCommitIDParam struct {
	// The commit ID to restore from
	CommitID param.Field[string] `json:"commit_id" api:"required" format:"uuid"`
}

func (r VmFromCommitRequestCommitIDParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VmFromCommitRequestCommitIDParam) implementsVmFromCommitRequestUnionParam() {}

// The tag name to restore from (legacy org-scoped tag)
type VmFromCommitRequestTagNameParam struct {
	// The tag name to restore from (legacy org-scoped tag)
	TagName param.Field[string] `json:"tag_name" api:"required"`
}

func (r VmFromCommitRequestTagNameParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VmFromCommitRequestTagNameParam) implementsVmFromCommitRequestUnionParam() {}

// A repository reference in "repo_name:tag_name" format
type VmFromCommitRequestRefParam struct {
	// A repository reference in "repo_name:tag_name" format
	Ref param.Field[string] `json:"ref" api:"required"`
}

func (r VmFromCommitRequestRefParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VmFromCommitRequestRefParam) implementsVmFromCommitRequestUnionParam() {}

// Response for GET /api/v1/vm/{vm_id}/metadata
type VmMetadataResponse struct {
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	IP        string    `json:"ip" api:"required"`
	OwnerID   string    `json:"owner_id" api:"required" format:"uuid"`
	// The state of a VM
	State           VmMetadataResponseState `json:"state" api:"required"`
	VmID            string                  `json:"vm_id" api:"required" format:"uuid"`
	DeletedAt       time.Time               `json:"deleted_at" api:"nullable" format:"date-time"`
	GrandparentVmID string                  `json:"grandparent_vm_id" api:"nullable" format:"uuid"`
	ParentCommitID  string                  `json:"parent_commit_id" api:"nullable" format:"uuid"`
	JSON            vmMetadataResponseJSON  `json:"-"`
}

// vmMetadataResponseJSON contains the JSON metadata for the struct
// [VmMetadataResponse]
type vmMetadataResponseJSON struct {
	CreatedAt       apijson.Field
	IP              apijson.Field
	OwnerID         apijson.Field
	State           apijson.Field
	VmID            apijson.Field
	DeletedAt       apijson.Field
	GrandparentVmID apijson.Field
	ParentCommitID  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *VmMetadataResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vmMetadataResponseJSON) RawJSON() string {
	return r.raw
}

// The state of a VM
type VmMetadataResponseState string

const (
	VmMetadataResponseStateBooting  VmMetadataResponseState = "booting"
	VmMetadataResponseStateRunning  VmMetadataResponseState = "running"
	VmMetadataResponseStatePaused   VmMetadataResponseState = "paused"
	VmMetadataResponseStateSleeping VmMetadataResponseState = "sleeping"
	VmMetadataResponseStateDead     VmMetadataResponseState = "dead"
)

func (r VmMetadataResponseState) IsKnown() bool {
	switch r {
	case VmMetadataResponseStateBooting, VmMetadataResponseStateRunning, VmMetadataResponseStatePaused, VmMetadataResponseStateSleeping, VmMetadataResponseStateDead:
		return true
	}
	return false
}

// Request body for PATCH /api/vm/{vm_id}/disk
type VmResizeDiskRequestParam struct {
	// The new disk size in MiB. Must be strictly greater than the current size.
	FsSizeMib param.Field[int64] `json:"fs_size_mib" api:"required"`
}

func (r VmResizeDiskRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Response body for GET /api/vm/{vm_id}/ssh_key
type VmSSHKeyResponse struct {
	// The SSH port that will be DNAT'd to the VM's netns (and, in turn, to its TAP
	// device)
	SSHPort int64 `json:"ssh_port" api:"required"`
	// Private SSH key in stringified OpenSSH format
	SSHPrivateKey string               `json:"ssh_private_key" api:"required"`
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
	State param.Field[VmUpdateStateRequestState] `json:"state" api:"required"`
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

type VmBranchParams struct {
	// Number of VMs to branch (optional; default 1)
	Count param.Field[int64] `query:"count"`
	// If true, keep VM paused after commit. Only applicable when branching a VM ID.
	KeepPaused param.Field[bool] `query:"keep_paused"`
	// If true, immediately return an error if VM is booting instead of waiting. Only
	// applicable when branching a VM ID.
	SkipWaitBoot param.Field[bool] `query:"skip_wait_boot"`
}

// URLQuery serializes [VmBranchParams]'s query parameters as `url.Values`.
func (r VmBranchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VmBranchByCommitParams struct {
	// Number of VMs to branch (optional; default 1)
	Count param.Field[int64] `query:"count"`
}

// URLQuery serializes [VmBranchByCommitParams]'s query parameters as `url.Values`.
func (r VmBranchByCommitParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VmBranchByTagParams struct {
	// Number of VMs to branch (optional; default 1)
	Count param.Field[int64] `query:"count"`
}

// URLQuery serializes [VmBranchByTagParams]'s query parameters as `url.Values`.
func (r VmBranchByTagParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VmBranchByVmParams struct {
	// Number of VMs to branch (optional; default 1)
	Count param.Field[int64] `query:"count"`
	// If true, keep VM paused after commit
	KeepPaused param.Field[bool] `query:"keep_paused"`
	// If true, immediately return an error if VM is booting instead of waiting
	SkipWaitBoot param.Field[bool] `query:"skip_wait_boot"`
}

// URLQuery serializes [VmBranchByVmParams]'s query parameters as `url.Values`.
func (r VmBranchByVmParams) URLQuery() (v url.Values) {
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
	// If provided, chelsea will use the requested commit UUID. Otherwise, it will
	// generate a UUID itself.
	CommitID param.Field[string] `json:"commit_id" format:"uuid"`
	// Optional description for the commit.
	Description param.Field[string] `json:"description"`
	// Optional human-readable name for the commit. Defaults to auto-generated name if
	// not provided.
	Name param.Field[string] `json:"name"`
}

func (r VmCommitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [VmCommitParams]'s query parameters as `url.Values`.
func (r VmCommitParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VmNewRootParams struct {
	NewRootRequest NewRootRequestParam `json:"new_root_request" api:"required"`
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

type VmExecParams struct {
	// Request body for POST /api/vm/{vm_id}/exec
	VmExecRequest VmExecRequestParam `json:"vm_exec_request" api:"required"`
}

func (r VmExecParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.VmExecRequest)
}

type VmExecStreamParams struct {
	// Request body for POST /api/vm/{vm_id}/exec
	VmExecRequest VmExecRequestParam `json:"vm_exec_request" api:"required"`
}

func (r VmExecStreamParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.VmExecRequest)
}

type VmExecStreamAttachParams struct {
	// Request body for POST /api/vm/{vm_id}/exec/stream/attach
	VmExecStreamAttachRequest VmExecStreamAttachRequestParam `json:"vm_exec_stream_attach_request" api:"required"`
}

func (r VmExecStreamAttachParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.VmExecStreamAttachRequest)
}

type VmGetLogsParams struct {
	// Maximum number of log entries to return
	MaxEntries param.Field[int64] `query:"max_entries"`
	// Byte offset into the log file (default: 0)
	Offset param.Field[int64] `query:"offset"`
	// Filter by 'stdout' or 'stderr'
	Stream param.Field[string] `query:"stream"`
}

// URLQuery serializes [VmGetLogsParams]'s query parameters as `url.Values`.
func (r VmGetLogsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VmResizeDiskParams struct {
	// Request body for PATCH /api/vm/{vm_id}/disk
	VmResizeDiskRequest VmResizeDiskRequestParam `json:"vm_resize_disk_request" api:"required"`
	// If true, return an error immediately if the VM is still booting. Default: false
	SkipWaitBoot param.Field[bool] `query:"skip_wait_boot"`
}

func (r VmResizeDiskParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.VmResizeDiskRequest)
}

// URLQuery serializes [VmResizeDiskParams]'s query parameters as `url.Values`.
func (r VmResizeDiskParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VmRestoreFromCommitParams struct {
	// Request body for POST /api/v1/vm/from_commit
	VmFromCommitRequest VmFromCommitRequestUnionParam `json:"vm_from_commit_request" api:"required"`
}

func (r VmRestoreFromCommitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.VmFromCommitRequest)
}

type VmUpdateStateParams struct {
	// Request body for PATCH /api/vm/{vm_id}/state
	VmUpdateStateRequest VmUpdateStateRequestParam `json:"vm_update_state_request" api:"required"`
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
