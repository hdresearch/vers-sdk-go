// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/hdresearch/vers-sdk-go/internal/apijson"
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
func (r *APIClusterService) Get(ctx context.Context, clusterID string, opts ...option.RequestOption) (res *APIClusterGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("api/cluster/%s", clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all clusters.
func (r *APIClusterService) List(ctx context.Context, opts ...option.RequestOption) (res *[]APIClusterListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/cluster"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a cluster.
func (r *APIClusterService) Delete(ctx context.Context, clusterID string, opts ...option.RequestOption) (res *APIClusterDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("api/cluster/%s", clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type APIClusterNewResponse struct {
	// The cluster's ID.
	ID string `json:"id,required"`
	// How many VMs are currently running on this cluster.
	VmCount int64 `json:"vm_count,required"`
	// The VMs that are children of the cluster, including the root VM.
	Vms []APIClusterNewResponseVm `json:"vms,required"`
	// The ID of the cluster's root VM.
	RootVmID string                    `json:"root_vm_id,nullable"`
	JSON     apiClusterNewResponseJSON `json:"-"`
}

// apiClusterNewResponseJSON contains the JSON metadata for the struct
// [APIClusterNewResponse]
type apiClusterNewResponseJSON struct {
	ID          apijson.Field
	VmCount     apijson.Field
	Vms         apijson.Field
	RootVmID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIClusterNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterNewResponseJSON) RawJSON() string {
	return r.raw
}

type APIClusterNewResponseVm struct {
	// The ID of the VM.
	ID string `json:"id,required"`
	// The IDs of direct children branched from this VM.
	Children []string `json:"children,required"`
	// The VM's local IP address on the VM subnet
	IPAddress string `json:"ip_address,required"`
	// The VM's network configuration
	NetworkInfo APIClusterNewResponseVmsNetworkInfo `json:"network_info,required"`
	// Whether the VM is running, paused, or not started.
	State APIClusterNewResponseVmsState `json:"state,required"`
	// The parent VM's ID, if present. If None, then this VM is a root VM.
	ParentID string                      `json:"parent_id,nullable"`
	JSON     apiClusterNewResponseVmJSON `json:"-"`
}

// apiClusterNewResponseVmJSON contains the JSON metadata for the struct
// [APIClusterNewResponseVm]
type apiClusterNewResponseVmJSON struct {
	ID          apijson.Field
	Children    apijson.Field
	IPAddress   apijson.Field
	NetworkInfo apijson.Field
	State       apijson.Field
	ParentID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIClusterNewResponseVm) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterNewResponseVmJSON) RawJSON() string {
	return r.raw
}

// The VM's network configuration
type APIClusterNewResponseVmsNetworkInfo struct {
	GuestIP     string                                  `json:"guest_ip,required"`
	GuestMac    string                                  `json:"guest_mac,required"`
	Tap0IP      string                                  `json:"tap0_ip,required"`
	Tap0Name    string                                  `json:"tap0_name,required"`
	VmNamespace string                                  `json:"vm_namespace,required"`
	JSON        apiClusterNewResponseVmsNetworkInfoJSON `json:"-"`
}

// apiClusterNewResponseVmsNetworkInfoJSON contains the JSON metadata for the
// struct [APIClusterNewResponseVmsNetworkInfo]
type apiClusterNewResponseVmsNetworkInfoJSON struct {
	GuestIP     apijson.Field
	GuestMac    apijson.Field
	Tap0IP      apijson.Field
	Tap0Name    apijson.Field
	VmNamespace apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIClusterNewResponseVmsNetworkInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterNewResponseVmsNetworkInfoJSON) RawJSON() string {
	return r.raw
}

// Whether the VM is running, paused, or not started.
type APIClusterNewResponseVmsState string

const (
	APIClusterNewResponseVmsStateNotStarted APIClusterNewResponseVmsState = "Not started"
	APIClusterNewResponseVmsStateRunning    APIClusterNewResponseVmsState = "Running"
	APIClusterNewResponseVmsStatePaused     APIClusterNewResponseVmsState = "Paused"
)

func (r APIClusterNewResponseVmsState) IsKnown() bool {
	switch r {
	case APIClusterNewResponseVmsStateNotStarted, APIClusterNewResponseVmsStateRunning, APIClusterNewResponseVmsStatePaused:
		return true
	}
	return false
}

type APIClusterGetResponse struct {
	// The cluster's ID.
	ID string `json:"id,required"`
	// How many VMs are currently running on this cluster.
	VmCount int64 `json:"vm_count,required"`
	// The VMs that are children of the cluster, including the root VM.
	Vms []APIClusterGetResponseVm `json:"vms,required"`
	// The ID of the cluster's root VM.
	RootVmID string                    `json:"root_vm_id,nullable"`
	JSON     apiClusterGetResponseJSON `json:"-"`
}

// apiClusterGetResponseJSON contains the JSON metadata for the struct
// [APIClusterGetResponse]
type apiClusterGetResponseJSON struct {
	ID          apijson.Field
	VmCount     apijson.Field
	Vms         apijson.Field
	RootVmID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIClusterGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterGetResponseJSON) RawJSON() string {
	return r.raw
}

type APIClusterGetResponseVm struct {
	// The ID of the VM.
	ID string `json:"id,required"`
	// The IDs of direct children branched from this VM.
	Children []string `json:"children,required"`
	// The VM's local IP address on the VM subnet
	IPAddress string `json:"ip_address,required"`
	// The VM's network configuration
	NetworkInfo APIClusterGetResponseVmsNetworkInfo `json:"network_info,required"`
	// Whether the VM is running, paused, or not started.
	State APIClusterGetResponseVmsState `json:"state,required"`
	// The parent VM's ID, if present. If None, then this VM is a root VM.
	ParentID string                      `json:"parent_id,nullable"`
	JSON     apiClusterGetResponseVmJSON `json:"-"`
}

// apiClusterGetResponseVmJSON contains the JSON metadata for the struct
// [APIClusterGetResponseVm]
type apiClusterGetResponseVmJSON struct {
	ID          apijson.Field
	Children    apijson.Field
	IPAddress   apijson.Field
	NetworkInfo apijson.Field
	State       apijson.Field
	ParentID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIClusterGetResponseVm) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterGetResponseVmJSON) RawJSON() string {
	return r.raw
}

// The VM's network configuration
type APIClusterGetResponseVmsNetworkInfo struct {
	GuestIP     string                                  `json:"guest_ip,required"`
	GuestMac    string                                  `json:"guest_mac,required"`
	Tap0IP      string                                  `json:"tap0_ip,required"`
	Tap0Name    string                                  `json:"tap0_name,required"`
	VmNamespace string                                  `json:"vm_namespace,required"`
	JSON        apiClusterGetResponseVmsNetworkInfoJSON `json:"-"`
}

// apiClusterGetResponseVmsNetworkInfoJSON contains the JSON metadata for the
// struct [APIClusterGetResponseVmsNetworkInfo]
type apiClusterGetResponseVmsNetworkInfoJSON struct {
	GuestIP     apijson.Field
	GuestMac    apijson.Field
	Tap0IP      apijson.Field
	Tap0Name    apijson.Field
	VmNamespace apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIClusterGetResponseVmsNetworkInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterGetResponseVmsNetworkInfoJSON) RawJSON() string {
	return r.raw
}

// Whether the VM is running, paused, or not started.
type APIClusterGetResponseVmsState string

const (
	APIClusterGetResponseVmsStateNotStarted APIClusterGetResponseVmsState = "Not started"
	APIClusterGetResponseVmsStateRunning    APIClusterGetResponseVmsState = "Running"
	APIClusterGetResponseVmsStatePaused     APIClusterGetResponseVmsState = "Paused"
)

func (r APIClusterGetResponseVmsState) IsKnown() bool {
	switch r {
	case APIClusterGetResponseVmsStateNotStarted, APIClusterGetResponseVmsStateRunning, APIClusterGetResponseVmsStatePaused:
		return true
	}
	return false
}

type APIClusterListResponse struct {
	// The cluster's ID.
	ID string `json:"id,required"`
	// How many VMs are currently running on this cluster.
	VmCount int64 `json:"vm_count,required"`
	// The VMs that are children of the cluster, including the root VM.
	Vms []APIClusterListResponseVm `json:"vms,required"`
	// The ID of the cluster's root VM.
	RootVmID string                     `json:"root_vm_id,nullable"`
	JSON     apiClusterListResponseJSON `json:"-"`
}

// apiClusterListResponseJSON contains the JSON metadata for the struct
// [APIClusterListResponse]
type apiClusterListResponseJSON struct {
	ID          apijson.Field
	VmCount     apijson.Field
	Vms         apijson.Field
	RootVmID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIClusterListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterListResponseJSON) RawJSON() string {
	return r.raw
}

type APIClusterListResponseVm struct {
	// The ID of the VM.
	ID string `json:"id,required"`
	// The IDs of direct children branched from this VM.
	Children []string `json:"children,required"`
	// The VM's local IP address on the VM subnet
	IPAddress string `json:"ip_address,required"`
	// The VM's network configuration
	NetworkInfo APIClusterListResponseVmsNetworkInfo `json:"network_info,required"`
	// Whether the VM is running, paused, or not started.
	State APIClusterListResponseVmsState `json:"state,required"`
	// The parent VM's ID, if present. If None, then this VM is a root VM.
	ParentID string                       `json:"parent_id,nullable"`
	JSON     apiClusterListResponseVmJSON `json:"-"`
}

// apiClusterListResponseVmJSON contains the JSON metadata for the struct
// [APIClusterListResponseVm]
type apiClusterListResponseVmJSON struct {
	ID          apijson.Field
	Children    apijson.Field
	IPAddress   apijson.Field
	NetworkInfo apijson.Field
	State       apijson.Field
	ParentID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIClusterListResponseVm) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterListResponseVmJSON) RawJSON() string {
	return r.raw
}

// The VM's network configuration
type APIClusterListResponseVmsNetworkInfo struct {
	GuestIP     string                                   `json:"guest_ip,required"`
	GuestMac    string                                   `json:"guest_mac,required"`
	Tap0IP      string                                   `json:"tap0_ip,required"`
	Tap0Name    string                                   `json:"tap0_name,required"`
	VmNamespace string                                   `json:"vm_namespace,required"`
	JSON        apiClusterListResponseVmsNetworkInfoJSON `json:"-"`
}

// apiClusterListResponseVmsNetworkInfoJSON contains the JSON metadata for the
// struct [APIClusterListResponseVmsNetworkInfo]
type apiClusterListResponseVmsNetworkInfoJSON struct {
	GuestIP     apijson.Field
	GuestMac    apijson.Field
	Tap0IP      apijson.Field
	Tap0Name    apijson.Field
	VmNamespace apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIClusterListResponseVmsNetworkInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterListResponseVmsNetworkInfoJSON) RawJSON() string {
	return r.raw
}

// Whether the VM is running, paused, or not started.
type APIClusterListResponseVmsState string

const (
	APIClusterListResponseVmsStateNotStarted APIClusterListResponseVmsState = "Not started"
	APIClusterListResponseVmsStateRunning    APIClusterListResponseVmsState = "Running"
	APIClusterListResponseVmsStatePaused     APIClusterListResponseVmsState = "Paused"
)

func (r APIClusterListResponseVmsState) IsKnown() bool {
	switch r {
	case APIClusterListResponseVmsStateNotStarted, APIClusterListResponseVmsStateRunning, APIClusterListResponseVmsStatePaused:
		return true
	}
	return false
}

type APIClusterDeleteResponse struct {
	// The cluster's ID.
	ID string `json:"id,required"`
	// How many VMs are currently running on this cluster.
	VmCount int64 `json:"vm_count,required"`
	// The VMs that are children of the cluster, including the root VM.
	Vms []APIClusterDeleteResponseVm `json:"vms,required"`
	// The ID of the cluster's root VM.
	RootVmID string                       `json:"root_vm_id,nullable"`
	JSON     apiClusterDeleteResponseJSON `json:"-"`
}

// apiClusterDeleteResponseJSON contains the JSON metadata for the struct
// [APIClusterDeleteResponse]
type apiClusterDeleteResponseJSON struct {
	ID          apijson.Field
	VmCount     apijson.Field
	Vms         apijson.Field
	RootVmID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIClusterDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type APIClusterDeleteResponseVm struct {
	// The ID of the VM.
	ID string `json:"id,required"`
	// The IDs of direct children branched from this VM.
	Children []string `json:"children,required"`
	// The VM's local IP address on the VM subnet
	IPAddress string `json:"ip_address,required"`
	// The VM's network configuration
	NetworkInfo APIClusterDeleteResponseVmsNetworkInfo `json:"network_info,required"`
	// Whether the VM is running, paused, or not started.
	State APIClusterDeleteResponseVmsState `json:"state,required"`
	// The parent VM's ID, if present. If None, then this VM is a root VM.
	ParentID string                         `json:"parent_id,nullable"`
	JSON     apiClusterDeleteResponseVmJSON `json:"-"`
}

// apiClusterDeleteResponseVmJSON contains the JSON metadata for the struct
// [APIClusterDeleteResponseVm]
type apiClusterDeleteResponseVmJSON struct {
	ID          apijson.Field
	Children    apijson.Field
	IPAddress   apijson.Field
	NetworkInfo apijson.Field
	State       apijson.Field
	ParentID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIClusterDeleteResponseVm) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterDeleteResponseVmJSON) RawJSON() string {
	return r.raw
}

// The VM's network configuration
type APIClusterDeleteResponseVmsNetworkInfo struct {
	GuestIP     string                                     `json:"guest_ip,required"`
	GuestMac    string                                     `json:"guest_mac,required"`
	Tap0IP      string                                     `json:"tap0_ip,required"`
	Tap0Name    string                                     `json:"tap0_name,required"`
	VmNamespace string                                     `json:"vm_namespace,required"`
	JSON        apiClusterDeleteResponseVmsNetworkInfoJSON `json:"-"`
}

// apiClusterDeleteResponseVmsNetworkInfoJSON contains the JSON metadata for the
// struct [APIClusterDeleteResponseVmsNetworkInfo]
type apiClusterDeleteResponseVmsNetworkInfoJSON struct {
	GuestIP     apijson.Field
	GuestMac    apijson.Field
	Tap0IP      apijson.Field
	Tap0Name    apijson.Field
	VmNamespace apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIClusterDeleteResponseVmsNetworkInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiClusterDeleteResponseVmsNetworkInfoJSON) RawJSON() string {
	return r.raw
}

// Whether the VM is running, paused, or not started.
type APIClusterDeleteResponseVmsState string

const (
	APIClusterDeleteResponseVmsStateNotStarted APIClusterDeleteResponseVmsState = "Not started"
	APIClusterDeleteResponseVmsStateRunning    APIClusterDeleteResponseVmsState = "Running"
	APIClusterDeleteResponseVmsStatePaused     APIClusterDeleteResponseVmsState = "Paused"
)

func (r APIClusterDeleteResponseVmsState) IsKnown() bool {
	switch r {
	case APIClusterDeleteResponseVmsStateNotStarted, APIClusterDeleteResponseVmsStateRunning, APIClusterDeleteResponseVmsStatePaused:
		return true
	}
	return false
}

type APIClusterNewParams struct {
	Body interface{} `json:"body,required"`
}

func (r APIClusterNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
