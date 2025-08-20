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
	ID                   string           `json:"id,required"`
	CPUCoresAvailable    int64            `json:"cpu_cores_available,required"`
	CPUCoresMargin       int64            `json:"cpu_cores_margin,required"`
	CPUCoresTotal        int64            `json:"cpu_cores_total,required"`
	CPUCoresUsed         int64            `json:"cpu_cores_used,required"`
	DiskDataMibAvailable int64            `json:"disk_data_mib_available,required"`
	DiskDataMibTotal     int64            `json:"disk_data_mib_total,required"`
	DiskVmMibAvailable   int64            `json:"disk_vm_mib_available,required"`
	DiskVmMibTotal       int64            `json:"disk_vm_mib_total,required"`
	MemoryMibAvailable   int64            `json:"memory_mib_available,required"`
	MemoryMibMargin      int64            `json:"memory_mib_margin,required"`
	MemoryMibTotal       int64            `json:"memory_mib_total,required"`
	MemoryMibUsed        int64            `json:"memory_mib_used,required"`
	VmNetworkCountInUse  int64            `json:"vm_network_count_in_use,required"`
	VmNetworkCountTotal  int64            `json:"vm_network_count_total,required"`
	JSON                 telemetryDtoJSON `json:"-"`
}

// telemetryDtoJSON contains the JSON metadata for the struct [TelemetryDto]
type telemetryDtoJSON struct {
	ID                   apijson.Field
	CPUCoresAvailable    apijson.Field
	CPUCoresMargin       apijson.Field
	CPUCoresTotal        apijson.Field
	CPUCoresUsed         apijson.Field
	DiskDataMibAvailable apijson.Field
	DiskDataMibTotal     apijson.Field
	DiskVmMibAvailable   apijson.Field
	DiskVmMibTotal       apijson.Field
	MemoryMibAvailable   apijson.Field
	MemoryMibMargin      apijson.Field
	MemoryMibTotal       apijson.Field
	MemoryMibUsed        apijson.Field
	VmNetworkCountInUse  apijson.Field
	VmNetworkCountTotal  apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TelemetryDto) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r telemetryDtoJSON) RawJSON() string {
	return r.raw
}
