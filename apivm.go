// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vers

import (
	"context"
	"net/http"

	"github.com/hdresearch/vers-sdk-go/internal/apijson"
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

// List all VMs.
func (r *APIVmService) List(ctx context.Context, opts ...option.RequestOption) (res *APIVmListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/vm"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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
	// Human-readable name assigned to the VM.
	Alias string `json:"alias,nullable"`
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
	Alias       apijson.Field
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
	IPAddress   apijson.Field
	NetworkInfo apijson.Field
	State       apijson.Field
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
