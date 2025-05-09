// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

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
	opts = append(r.Options[:], opts...)
	path := "api/rootfs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete an existing rootfs from the server.
func (r *APIRootfService) Delete(ctx context.Context, rootfsID string, opts ...option.RequestOption) (res *APIRootfDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
	if rootfsID == "" {
		err = errors.New("missing required rootfs_id parameter")
		return
	}
	path := fmt.Sprintf("api/rootfs/%s", rootfsID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type APIRootfListResponse struct {
	Data        APIRootfListResponseData `json:"data,required"`
	DurationNs  int64                    `json:"duration_ns,required"`
	OperationID string                   `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                    `json:"time_start,required"`
	JSON      apiRootfListResponseJSON `json:"-"`
}

// apiRootfListResponseJSON contains the JSON metadata for the struct
// [APIRootfListResponse]
type apiRootfListResponseJSON struct {
	Data        apijson.Field
	DurationNs  apijson.Field
	OperationID apijson.Field
	TimeStart   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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

type APIRootfDeleteResponse struct {
	Data        APIRootfDeleteResponseData `json:"data,required"`
	DurationNs  int64                      `json:"duration_ns,required"`
	OperationID string                     `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                      `json:"time_start,required"`
	JSON      apiRootfDeleteResponseJSON `json:"-"`
}

// apiRootfDeleteResponseJSON contains the JSON metadata for the struct
// [APIRootfDeleteResponse]
type apiRootfDeleteResponseJSON struct {
	Data        apijson.Field
	DurationNs  apijson.Field
	OperationID apijson.Field
	TimeStart   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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

type APIRootfUploadResponse struct {
	Data        APIRootfUploadResponseData `json:"data,required"`
	DurationNs  int64                      `json:"duration_ns,required"`
	OperationID string                     `json:"operation_id,required"`
	// Unix epoch time (secs)
	TimeStart int64                      `json:"time_start,required"`
	JSON      apiRootfUploadResponseJSON `json:"-"`
}

// apiRootfUploadResponseJSON contains the JSON metadata for the struct
// [APIRootfUploadResponse]
type apiRootfUploadResponseJSON struct {
	Data        apijson.Field
	DurationNs  apijson.Field
	OperationID apijson.Field
	TimeStart   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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

type APIRootfUploadParams struct {
	// The path of the Dockerfile contained within the tar archive
	Dockerfile param.Field[APIRootfUploadParamsDockerfile] `query:"dockerfile,required"`
}

// URLQuery serializes [APIRootfUploadParams]'s query parameters as `url.Values`.
func (r APIRootfUploadParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The path of the Dockerfile contained within the tar archive
type APIRootfUploadParamsDockerfile struct {
	Dockerfile param.Field[string] `query:"dockerfile"`
}

// URLQuery serializes [APIRootfUploadParamsDockerfile]'s query parameters as
// `url.Values`.
func (r APIRootfUploadParamsDockerfile) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
