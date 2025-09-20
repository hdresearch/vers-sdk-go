// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/hdresearch/vers-sdk-go/internal/apijson"
	"github.com/hdresearch/vers-sdk-go/internal/apiquery"
	"github.com/hdresearch/vers-sdk-go/internal/param"
	"github.com/hdresearch/vers-sdk-go/internal/requestconfig"
	"github.com/hdresearch/vers-sdk-go/option"
)

// APIRootfService contains methods and other services that help with interacting
// with the vers API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIRootfService] method instead.
type APIRootfService struct {
	Options []option.RequestOption
}

// NewAPIRootfService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPIRootfService(opts ...option.RequestOption) (r *APIRootfService) {
	r = &APIRootfService{}
	r.Options = opts
	return
}

// List all available rootfs names on the server.
func (r *APIRootfService) List(ctx context.Context, opts ...option.RequestOption) (res *APIRootfListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/rootfs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete an existing rootfs from the server.
func (r *APIRootfService) Delete(ctx context.Context, rootfsID string, opts ...option.RequestOption) (res *APIRootfDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if rootfsID == "" {
		err = errors.New("missing required rootfs_id parameter")
		return
	}
	path := fmt.Sprintf("api/rootfs/%s", rootfsID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Upload a rootfs tar archive to the server. The archive should contain the
// Dockerfile and all necessary dependencies.
func (r *APIRootfService) Upload(ctx context.Context, rootfsID string, body APIRootfUploadParams, opts ...option.RequestOption) (res *APIRootfUploadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if rootfsID == "" {
		err = errors.New("missing required rootfs_id parameter")
		return
	}
	path := fmt.Sprintf("api/rootfs/%s", rootfsID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type APIRootfListResponse struct {
	Data          APIRootfListResponseData          `json:"data,required"`
	DurationNs    int64                             `json:"duration_ns,required"`
	OperationCode APIRootfListResponseOperationCode `json:"operation_code,required"`
	OperationID   string                            `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                    `json:"time_start,required"`
	JSON      apiRootfListResponseJSON `json:"-"`
}

// apiRootfListResponseJSON contains the JSON metadata for the struct
// [APIRootfListResponse]
type apiRootfListResponseJSON struct {
	Data          apijson.Field
	DurationNs    apijson.Field
	OperationCode apijson.Field
	OperationID   apijson.Field
	TimeStart     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *APIRootfListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiRootfListResponseJSON) RawJSON() string {
	return r.raw
}

type APIRootfListResponseData struct {
	RootfsNames []string                     `json:"rootfs_names,required"`
	JSON        apiRootfListResponseDataJSON `json:"-"`
}

// apiRootfListResponseDataJSON contains the JSON metadata for the struct
// [APIRootfListResponseData]
type apiRootfListResponseDataJSON struct {
	RootfsNames apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIRootfListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiRootfListResponseDataJSON) RawJSON() string {
	return r.raw
}

type APIRootfListResponseOperationCode string

const (
	APIRootfListResponseOperationCodeListClusters     APIRootfListResponseOperationCode = "list_clusters"
	APIRootfListResponseOperationCodeGetCluster       APIRootfListResponseOperationCode = "get_cluster"
	APIRootfListResponseOperationCodeCreateCluster    APIRootfListResponseOperationCode = "create_cluster"
	APIRootfListResponseOperationCodeDeleteCluster    APIRootfListResponseOperationCode = "delete_cluster"
	APIRootfListResponseOperationCodeUpdateCluster    APIRootfListResponseOperationCode = "update_cluster"
	APIRootfListResponseOperationCodeGetClusterSSHKey APIRootfListResponseOperationCode = "get_cluster_ssh_key"
	APIRootfListResponseOperationCodeListVms          APIRootfListResponseOperationCode = "list_vms"
	APIRootfListResponseOperationCodeGetVm            APIRootfListResponseOperationCode = "get_vm"
	APIRootfListResponseOperationCodeUpdateVm         APIRootfListResponseOperationCode = "update_vm"
	APIRootfListResponseOperationCodeBranchVm         APIRootfListResponseOperationCode = "branch_vm"
	APIRootfListResponseOperationCodeCommitVm         APIRootfListResponseOperationCode = "commit_vm"
	APIRootfListResponseOperationCodeDeleteVm         APIRootfListResponseOperationCode = "delete_vm"
	APIRootfListResponseOperationCodeGetVmSSHKey      APIRootfListResponseOperationCode = "get_vm_ssh_key"
	APIRootfListResponseOperationCodeUploadRootfs     APIRootfListResponseOperationCode = "upload_rootfs"
	APIRootfListResponseOperationCodeDeleteRootfs     APIRootfListResponseOperationCode = "delete_rootfs"
	APIRootfListResponseOperationCodeListRootfs       APIRootfListResponseOperationCode = "list_rootfs"
)

func (r APIRootfListResponseOperationCode) IsKnown() bool {
	switch r {
	case APIRootfListResponseOperationCodeListClusters, APIRootfListResponseOperationCodeGetCluster, APIRootfListResponseOperationCodeCreateCluster, APIRootfListResponseOperationCodeDeleteCluster, APIRootfListResponseOperationCodeUpdateCluster, APIRootfListResponseOperationCodeGetClusterSSHKey, APIRootfListResponseOperationCodeListVms, APIRootfListResponseOperationCodeGetVm, APIRootfListResponseOperationCodeUpdateVm, APIRootfListResponseOperationCodeBranchVm, APIRootfListResponseOperationCodeCommitVm, APIRootfListResponseOperationCodeDeleteVm, APIRootfListResponseOperationCodeGetVmSSHKey, APIRootfListResponseOperationCodeUploadRootfs, APIRootfListResponseOperationCodeDeleteRootfs, APIRootfListResponseOperationCodeListRootfs:
		return true
	}
	return false
}

type APIRootfDeleteResponse struct {
	Data          APIRootfDeleteResponseData          `json:"data,required"`
	DurationNs    int64                               `json:"duration_ns,required"`
	OperationCode APIRootfDeleteResponseOperationCode `json:"operation_code,required"`
	OperationID   string                              `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                      `json:"time_start,required"`
	JSON      apiRootfDeleteResponseJSON `json:"-"`
}

// apiRootfDeleteResponseJSON contains the JSON metadata for the struct
// [APIRootfDeleteResponse]
type apiRootfDeleteResponseJSON struct {
	Data          apijson.Field
	DurationNs    apijson.Field
	OperationCode apijson.Field
	OperationID   apijson.Field
	TimeStart     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *APIRootfDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiRootfDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type APIRootfDeleteResponseData struct {
	RootfsName string                         `json:"rootfs_name,required"`
	JSON       apiRootfDeleteResponseDataJSON `json:"-"`
}

// apiRootfDeleteResponseDataJSON contains the JSON metadata for the struct
// [APIRootfDeleteResponseData]
type apiRootfDeleteResponseDataJSON struct {
	RootfsName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIRootfDeleteResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiRootfDeleteResponseDataJSON) RawJSON() string {
	return r.raw
}

type APIRootfDeleteResponseOperationCode string

const (
	APIRootfDeleteResponseOperationCodeListClusters     APIRootfDeleteResponseOperationCode = "list_clusters"
	APIRootfDeleteResponseOperationCodeGetCluster       APIRootfDeleteResponseOperationCode = "get_cluster"
	APIRootfDeleteResponseOperationCodeCreateCluster    APIRootfDeleteResponseOperationCode = "create_cluster"
	APIRootfDeleteResponseOperationCodeDeleteCluster    APIRootfDeleteResponseOperationCode = "delete_cluster"
	APIRootfDeleteResponseOperationCodeUpdateCluster    APIRootfDeleteResponseOperationCode = "update_cluster"
	APIRootfDeleteResponseOperationCodeGetClusterSSHKey APIRootfDeleteResponseOperationCode = "get_cluster_ssh_key"
	APIRootfDeleteResponseOperationCodeListVms          APIRootfDeleteResponseOperationCode = "list_vms"
	APIRootfDeleteResponseOperationCodeGetVm            APIRootfDeleteResponseOperationCode = "get_vm"
	APIRootfDeleteResponseOperationCodeUpdateVm         APIRootfDeleteResponseOperationCode = "update_vm"
	APIRootfDeleteResponseOperationCodeBranchVm         APIRootfDeleteResponseOperationCode = "branch_vm"
	APIRootfDeleteResponseOperationCodeCommitVm         APIRootfDeleteResponseOperationCode = "commit_vm"
	APIRootfDeleteResponseOperationCodeDeleteVm         APIRootfDeleteResponseOperationCode = "delete_vm"
	APIRootfDeleteResponseOperationCodeGetVmSSHKey      APIRootfDeleteResponseOperationCode = "get_vm_ssh_key"
	APIRootfDeleteResponseOperationCodeUploadRootfs     APIRootfDeleteResponseOperationCode = "upload_rootfs"
	APIRootfDeleteResponseOperationCodeDeleteRootfs     APIRootfDeleteResponseOperationCode = "delete_rootfs"
	APIRootfDeleteResponseOperationCodeListRootfs       APIRootfDeleteResponseOperationCode = "list_rootfs"
)

func (r APIRootfDeleteResponseOperationCode) IsKnown() bool {
	switch r {
	case APIRootfDeleteResponseOperationCodeListClusters, APIRootfDeleteResponseOperationCodeGetCluster, APIRootfDeleteResponseOperationCodeCreateCluster, APIRootfDeleteResponseOperationCodeDeleteCluster, APIRootfDeleteResponseOperationCodeUpdateCluster, APIRootfDeleteResponseOperationCodeGetClusterSSHKey, APIRootfDeleteResponseOperationCodeListVms, APIRootfDeleteResponseOperationCodeGetVm, APIRootfDeleteResponseOperationCodeUpdateVm, APIRootfDeleteResponseOperationCodeBranchVm, APIRootfDeleteResponseOperationCodeCommitVm, APIRootfDeleteResponseOperationCodeDeleteVm, APIRootfDeleteResponseOperationCodeGetVmSSHKey, APIRootfDeleteResponseOperationCodeUploadRootfs, APIRootfDeleteResponseOperationCodeDeleteRootfs, APIRootfDeleteResponseOperationCodeListRootfs:
		return true
	}
	return false
}

type APIRootfUploadResponse struct {
	Data          APIRootfUploadResponseData          `json:"data,required"`
	DurationNs    int64                               `json:"duration_ns,required"`
	OperationCode APIRootfUploadResponseOperationCode `json:"operation_code,required"`
	OperationID   string                              `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                      `json:"time_start,required"`
	JSON      apiRootfUploadResponseJSON `json:"-"`
}

// apiRootfUploadResponseJSON contains the JSON metadata for the struct
// [APIRootfUploadResponse]
type apiRootfUploadResponseJSON struct {
	Data          apijson.Field
	DurationNs    apijson.Field
	OperationCode apijson.Field
	OperationID   apijson.Field
	TimeStart     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *APIRootfUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiRootfUploadResponseJSON) RawJSON() string {
	return r.raw
}

type APIRootfUploadResponseData struct {
	RootfsName string                         `json:"rootfs_name,required"`
	JSON       apiRootfUploadResponseDataJSON `json:"-"`
}

// apiRootfUploadResponseDataJSON contains the JSON metadata for the struct
// [APIRootfUploadResponseData]
type apiRootfUploadResponseDataJSON struct {
	RootfsName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIRootfUploadResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiRootfUploadResponseDataJSON) RawJSON() string {
	return r.raw
}

type APIRootfUploadResponseOperationCode string

const (
	APIRootfUploadResponseOperationCodeListClusters     APIRootfUploadResponseOperationCode = "list_clusters"
	APIRootfUploadResponseOperationCodeGetCluster       APIRootfUploadResponseOperationCode = "get_cluster"
	APIRootfUploadResponseOperationCodeCreateCluster    APIRootfUploadResponseOperationCode = "create_cluster"
	APIRootfUploadResponseOperationCodeDeleteCluster    APIRootfUploadResponseOperationCode = "delete_cluster"
	APIRootfUploadResponseOperationCodeUpdateCluster    APIRootfUploadResponseOperationCode = "update_cluster"
	APIRootfUploadResponseOperationCodeGetClusterSSHKey APIRootfUploadResponseOperationCode = "get_cluster_ssh_key"
	APIRootfUploadResponseOperationCodeListVms          APIRootfUploadResponseOperationCode = "list_vms"
	APIRootfUploadResponseOperationCodeGetVm            APIRootfUploadResponseOperationCode = "get_vm"
	APIRootfUploadResponseOperationCodeUpdateVm         APIRootfUploadResponseOperationCode = "update_vm"
	APIRootfUploadResponseOperationCodeBranchVm         APIRootfUploadResponseOperationCode = "branch_vm"
	APIRootfUploadResponseOperationCodeCommitVm         APIRootfUploadResponseOperationCode = "commit_vm"
	APIRootfUploadResponseOperationCodeDeleteVm         APIRootfUploadResponseOperationCode = "delete_vm"
	APIRootfUploadResponseOperationCodeGetVmSSHKey      APIRootfUploadResponseOperationCode = "get_vm_ssh_key"
	APIRootfUploadResponseOperationCodeUploadRootfs     APIRootfUploadResponseOperationCode = "upload_rootfs"
	APIRootfUploadResponseOperationCodeDeleteRootfs     APIRootfUploadResponseOperationCode = "delete_rootfs"
	APIRootfUploadResponseOperationCodeListRootfs       APIRootfUploadResponseOperationCode = "list_rootfs"
)

func (r APIRootfUploadResponseOperationCode) IsKnown() bool {
	switch r {
	case APIRootfUploadResponseOperationCodeListClusters, APIRootfUploadResponseOperationCodeGetCluster, APIRootfUploadResponseOperationCodeCreateCluster, APIRootfUploadResponseOperationCodeDeleteCluster, APIRootfUploadResponseOperationCodeUpdateCluster, APIRootfUploadResponseOperationCodeGetClusterSSHKey, APIRootfUploadResponseOperationCodeListVms, APIRootfUploadResponseOperationCodeGetVm, APIRootfUploadResponseOperationCodeUpdateVm, APIRootfUploadResponseOperationCodeBranchVm, APIRootfUploadResponseOperationCodeCommitVm, APIRootfUploadResponseOperationCodeDeleteVm, APIRootfUploadResponseOperationCodeGetVmSSHKey, APIRootfUploadResponseOperationCodeUploadRootfs, APIRootfUploadResponseOperationCodeDeleteRootfs, APIRootfUploadResponseOperationCodeListRootfs:
		return true
	}
	return false
}

type APIRootfUploadParams struct {
	// The path of the Dockerfile contained within the tar archive
	Dockerfile param.Field[string] `query:"dockerfile"`
}

// URLQuery serializes [APIRootfUploadParams]'s query parameters as `url.Values`.
func (r APIRootfUploadParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
