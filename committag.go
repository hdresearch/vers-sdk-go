// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/hdresearch/vers-sdk-go/internal/apijson"
	"github.com/hdresearch/vers-sdk-go/internal/param"
	"github.com/hdresearch/vers-sdk-go/internal/requestconfig"
	"github.com/hdresearch/vers-sdk-go/option"
)

// CommitTagService contains methods and other services that help with interacting
// with the vers API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCommitTagService] method instead.
type CommitTagService struct {
	Options []option.RequestOption
}

// NewCommitTagService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCommitTagService(opts ...option.RequestOption) (r *CommitTagService) {
	r = &CommitTagService{}
	r.Options = opts
	return
}

func (r *CommitTagService) New(ctx context.Context, body CommitTagNewParams, opts ...option.RequestOption) (res *CreateTagResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/commit_tags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *CommitTagService) Update(ctx context.Context, tagName string, body CommitTagUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if tagName == "" {
		err = errors.New("missing required tag_name parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/commit_tags/%s", tagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

func (r *CommitTagService) List(ctx context.Context, opts ...option.RequestOption) (res *ListTagsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/commit_tags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *CommitTagService) Delete(ctx context.Context, tagName string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if tagName == "" {
		err = errors.New("missing required tag_name parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/commit_tags/%s", tagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

func (r *CommitTagService) Get(ctx context.Context, tagName string, opts ...option.RequestOption) (res *TagInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	if tagName == "" {
		err = errors.New("missing required tag_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/commit_tags/%s", tagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Request body for POST /api/v1/commit_tags
type CreateTagRequestParam struct {
	// The commit ID this tag should point to
	CommitID param.Field[string] `json:"commit_id" api:"required" format:"uuid"`
	// The name of the tag (alphanumeric, hyphens, underscores, dots, 1-64 chars)
	TagName param.Field[string] `json:"tag_name" api:"required"`
	// Optional description of what this tag represents
	Description param.Field[string] `json:"description"`
}

func (r CreateTagRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Response body for POST /api/v1/commit_tags
type CreateTagResponse struct {
	// The commit ID this tag points to
	CommitID string `json:"commit_id" api:"required" format:"uuid"`
	// The ID of the newly created tag
	TagID string `json:"tag_id" api:"required" format:"uuid"`
	// The name of the tag
	TagName string                `json:"tag_name" api:"required"`
	JSON    createTagResponseJSON `json:"-"`
}

// createTagResponseJSON contains the JSON metadata for the struct
// [CreateTagResponse]
type createTagResponseJSON struct {
	CommitID    apijson.Field
	TagID       apijson.Field
	TagName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreateTagResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r createTagResponseJSON) RawJSON() string {
	return r.raw
}

// Response body for GET /api/v1/commit_tags
type ListTagsResponse struct {
	// List of all tags in the user's organization
	Tags []TagInfo            `json:"tags" api:"required"`
	JSON listTagsResponseJSON `json:"-"`
}

// listTagsResponseJSON contains the JSON metadata for the struct
// [ListTagsResponse]
type listTagsResponseJSON struct {
	Tags        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListTagsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listTagsResponseJSON) RawJSON() string {
	return r.raw
}

// Tag information returned in list and get operations
type TagInfo struct {
	// The commit ID this tag currently points to
	CommitID string `json:"commit_id" api:"required" format:"uuid"`
	// When the tag was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The tag's unique identifier
	TagID string `json:"tag_id" api:"required" format:"uuid"`
	// The name of the tag
	TagName string `json:"tag_name" api:"required"`
	// When the tag was last updated (moved to different commit or description changed)
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Optional description of what this tag represents
	Description string      `json:"description" api:"nullable"`
	JSON        tagInfoJSON `json:"-"`
}

// tagInfoJSON contains the JSON metadata for the struct [TagInfo]
type tagInfoJSON struct {
	CommitID    apijson.Field
	CreatedAt   apijson.Field
	TagID       apijson.Field
	TagName     apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TagInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tagInfoJSON) RawJSON() string {
	return r.raw
}

// Request body for PATCH /api/v1/commit_tags/{tag_name}
//
// For `description`:
//
// - Field absent from JSON → don't change the description
// - Field present as `null` → clear the description
// - Field present as `"text"` → set the description to "text"
type UpdateTagRequestParam struct {
	// Optional new commit ID to move the tag to
	CommitID param.Field[string] `json:"commit_id" format:"uuid"`
	// Optional new description for the tag. Send `null` to clear an existing
	// description.
	Description param.Field[string] `json:"description"`
}

func (r UpdateTagRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CommitTagNewParams struct {
	// Request body for POST /api/v1/commit_tags
	CreateTagRequest CreateTagRequestParam `json:"create_tag_request" api:"required"`
}

func (r CommitTagNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateTagRequest)
}

type CommitTagUpdateParams struct {
	// Request body for PATCH /api/v1/commit_tags/{tag_name}
	//
	// For `description`:
	//
	// - Field absent from JSON → don't change the description
	// - Field present as `null` → clear the description
	// - Field present as `"text"` → set the description to "text"
	UpdateTagRequest UpdateTagRequestParam `json:"update_tag_request" api:"required"`
}

func (r CommitTagUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.UpdateTagRequest)
}
