// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/hdresearch/vers-sdk-go/internal/apijson"
	"github.com/hdresearch/vers-sdk-go/internal/requestconfig"
	"github.com/hdresearch/vers-sdk-go/option"
)

// OrchestratorNodeService contains methods and other services that help with
// interacting with the vers API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrchestratorNodeService] method instead.
type OrchestratorNodeService struct {
	Options []option.RequestOption
}

// NewOrchestratorNodeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrchestratorNodeService(opts ...option.RequestOption) (r *OrchestratorNodeService) {
	r = &OrchestratorNodeService{}
	r.Options = opts
	return
}

func (r *OrchestratorNodeService) ListVms(ctx context.Context, nodeID string, opts ...option.RequestOption) (res *VmListAllResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if nodeID == "" {
		err = errors.New("missing required node_id parameter")
		return
	}
	path := fmt.Sprintf("node/%s/vms", nodeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Response body for GET /api/vm
type VmListAllResponse struct {
	// A list of nodes, each of which is a "root VM" with one or more children
	Vms  []VmListAllResponseVm `json:"vms,required"`
	JSON vmListAllResponseJSON `json:"-"`
}

// vmListAllResponseJSON contains the JSON metadata for the struct
// [VmListAllResponse]
type vmListAllResponseJSON struct {
	Vms         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VmListAllResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vmListAllResponseJSON) RawJSON() string {
	return r.raw
}

// Represents a tree node for a VM
type VmListAllResponseVm struct {
	// The VM ID, a (v4) UUID.
	VmID string `json:"vm_id,required"`
	// The VM's parent ID
	ParentID string                  `json:"parent_id,nullable"`
	JSON     vmListAllResponseVmJSON `json:"-"`
}

// vmListAllResponseVmJSON contains the JSON metadata for the struct
// [VmListAllResponseVm]
type vmListAllResponseVmJSON struct {
	VmID        apijson.Field
	ParentID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VmListAllResponseVm) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vmListAllResponseVmJSON) RawJSON() string {
	return r.raw
}
