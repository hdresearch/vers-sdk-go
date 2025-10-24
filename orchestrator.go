// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vers

import (
	"github.com/hdresearch/vers-sdk-go/option"
)

// OrchestratorService contains methods and other services that help with
// interacting with the vers API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrchestratorService] method instead.
type OrchestratorService struct {
	Options []option.RequestOption
	Vm      *OrchestratorVmService
	Node    *OrchestratorNodeService
}

// NewOrchestratorService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrchestratorService(opts ...option.RequestOption) (r *OrchestratorService) {
	r = &OrchestratorService{}
	r.Options = opts
	r.Vm = NewOrchestratorVmService(opts...)
	r.Node = NewOrchestratorNodeService(opts...)
	return
}
