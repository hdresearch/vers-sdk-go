// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vers

import (
	"context"
	"net/http"

	"github.com/hdresearch/vers-sdk-go/internal/apijson"
	"github.com/hdresearch/vers-sdk-go/internal/requestconfig"
	"github.com/hdresearch/vers-sdk-go/option"
)

// APITelemetryService contains methods and other services that help with
// interacting with the vers API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPITelemetryService] method instead.
type APITelemetryService struct {
	Options []option.RequestOption
}

// NewAPITelemetryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPITelemetryService(opts ...option.RequestOption) (r *APITelemetryService) {
	r = &APITelemetryService{}
	r.Options = opts
	return
}

// Get telemetry information
func (r *APITelemetryService) GetInfo(ctx context.Context, opts ...option.RequestOption) (res *TelemetryDto, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/telemetry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TelemetryDto struct {
	ID                  string           `json:"id,required"`
	FsMibCurrent        int64            `json:"fs_mib_current,required"`
	FsMibMax            int64            `json:"fs_mib_max,required"`
	MemMibCurrent       int64            `json:"mem_mib_current,required"`
	MemMibMax           int64            `json:"mem_mib_max,required"`
	VcpuCurrent         int64            `json:"vcpu_current,required"`
	VcpuMax             int64            `json:"vcpu_max,required"`
	VmNetworkCountInUse int64            `json:"vm_network_count_in_use,required"`
	VmNetworkCountTotal int64            `json:"vm_network_count_total,required"`
	JSON                telemetryDtoJSON `json:"-"`
}

// telemetryDtoJSON contains the JSON metadata for the struct [TelemetryDto]
type telemetryDtoJSON struct {
	ID                  apijson.Field
	FsMibCurrent        apijson.Field
	FsMibMax            apijson.Field
	MemMibCurrent       apijson.Field
	MemMibMax           apijson.Field
	VcpuCurrent         apijson.Field
	VcpuMax             apijson.Field
	VmNetworkCountInUse apijson.Field
	VmNetworkCountTotal apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *TelemetryDto) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r telemetryDtoJSON) RawJSON() string {
	return r.raw
}
