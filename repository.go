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

// RepositoryService contains methods and other services that help with interacting
// with the vers API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRepositoryService] method instead.
type RepositoryService struct {
	Options []option.RequestOption
}

// NewRepositoryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRepositoryService(opts ...option.RequestOption) (r *RepositoryService) {
	r = &RepositoryService{}
	r.Options = opts
	return
}

func (r *RepositoryService) New(ctx context.Context, body RepositoryNewParams, opts ...option.RequestOption) (res *CreateRepositoryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/repositories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *RepositoryService) List(ctx context.Context, opts ...option.RequestOption) (res *ListRepositoriesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/repositories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *RepositoryService) Delete(ctx context.Context, repoName string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if repoName == "" {
		err = errors.New("missing required repo_name parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/repositories/%s", repoName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

func (r *RepositoryService) NewTag(ctx context.Context, repoName string, body RepositoryNewTagParams, opts ...option.RequestOption) (res *CreateRepoTagResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if repoName == "" {
		err = errors.New("missing required repo_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/repositories/%s/tags", repoName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *RepositoryService) DeleteTag(ctx context.Context, repoName string, tagName string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if repoName == "" {
		err = errors.New("missing required repo_name parameter")
		return err
	}
	if tagName == "" {
		err = errors.New("missing required tag_name parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/repositories/%s/tags/%s", repoName, tagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

func (r *RepositoryService) Fork(ctx context.Context, body RepositoryForkParams, opts ...option.RequestOption) (res *ForkRepositoryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/repositories/fork"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *RepositoryService) Get(ctx context.Context, repoName string, opts ...option.RequestOption) (res *RepositoryInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	if repoName == "" {
		err = errors.New("missing required repo_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/repositories/%s", repoName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *RepositoryService) GetTag(ctx context.Context, repoName string, tagName string, opts ...option.RequestOption) (res *RepoTagInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	if repoName == "" {
		err = errors.New("missing required repo_name parameter")
		return nil, err
	}
	if tagName == "" {
		err = errors.New("missing required tag_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/repositories/%s/tags/%s", repoName, tagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *RepositoryService) ListTags(ctx context.Context, repoName string, opts ...option.RequestOption) (res *ListRepoTagsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if repoName == "" {
		err = errors.New("missing required repo_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/repositories/%s/tags", repoName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *RepositoryService) SetVisibility(ctx context.Context, repoName string, body RepositorySetVisibilityParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if repoName == "" {
		err = errors.New("missing required repo_name parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/repositories/%s/visibility", repoName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

func (r *RepositoryService) UpdateTag(ctx context.Context, repoName string, tagName string, body RepositoryUpdateTagParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if repoName == "" {
		err = errors.New("missing required repo_name parameter")
		return err
	}
	if tagName == "" {
		err = errors.New("missing required tag_name parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/repositories/%s/tags/%s", repoName, tagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

// Request body for creating a tag within a repository: POST
// /api/v1/repositories/{repo_name}/tags
type CreateRepoTagRequestParam struct {
	// The commit ID this tag should point to
	CommitID param.Field[string] `json:"commit_id" api:"required" format:"uuid"`
	// The tag name (e.g. "latest", "v1.0")
	TagName param.Field[string] `json:"tag_name" api:"required"`
	// Optional description of what this tag represents
	Description param.Field[string] `json:"description"`
}

func (r CreateRepoTagRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Response body for POST /api/v1/repositories/{repo_name}/tags
type CreateRepoTagResponse struct {
	// The commit ID this tag points to
	CommitID string `json:"commit_id" api:"required" format:"uuid"`
	// Full reference in image_name:tag format
	Reference string `json:"reference" api:"required"`
	// The ID of the newly created tag
	TagID string                    `json:"tag_id" api:"required" format:"uuid"`
	JSON  createRepoTagResponseJSON `json:"-"`
}

// createRepoTagResponseJSON contains the JSON metadata for the struct
// [CreateRepoTagResponse]
type createRepoTagResponseJSON struct {
	CommitID    apijson.Field
	Reference   apijson.Field
	TagID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreateRepoTagResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r createRepoTagResponseJSON) RawJSON() string {
	return r.raw
}

// Request body for POST /api/v1/repositories
type CreateRepositoryRequestParam struct {
	// The name of the repository (alphanumeric, hyphens, underscores, dots, 1-64
	// chars)
	Name param.Field[string] `json:"name" api:"required"`
	// Optional description of the repository
	Description param.Field[string] `json:"description"`
}

func (r CreateRepositoryRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Response body for POST /api/v1/repositories
type CreateRepositoryResponse struct {
	// The name of the repository
	Name string `json:"name" api:"required"`
	// The ID of the newly created repository
	RepoID string                       `json:"repo_id" api:"required" format:"uuid"`
	JSON   createRepositoryResponseJSON `json:"-"`
}

// createRepositoryResponseJSON contains the JSON metadata for the struct
// [CreateRepositoryResponse]
type createRepositoryResponseJSON struct {
	Name        apijson.Field
	RepoID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreateRepositoryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r createRepositoryResponseJSON) RawJSON() string {
	return r.raw
}

// Request body for POST /api/v1/repositories/fork
type ForkRepositoryRequestParam struct {
	// The organization that owns the source public repository
	SourceOrg param.Field[string] `json:"source_org" api:"required"`
	// The source repository name
	SourceRepo param.Field[string] `json:"source_repo" api:"required"`
	// The tag to fork (e.g. "latest", "v1.0")
	SourceTag param.Field[string] `json:"source_tag" api:"required"`
	// Name for the new repository in your org (defaults to source_repo if omitted)
	RepoName param.Field[string] `json:"repo_name"`
	// Tag name in the new repo (defaults to source_tag if omitted)
	TagName param.Field[string] `json:"tag_name"`
}

func (r ForkRepositoryRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Response body for POST /api/v1/repositories/fork
type ForkRepositoryResponse struct {
	// The new commit in your org (snapshot of the forked VM)
	CommitID string `json:"commit_id" api:"required" format:"uuid"`
	// Full reference: repo_name:tag_name
	Reference string `json:"reference" api:"required"`
	// The new repository name in your org
	RepoName string `json:"repo_name" api:"required"`
	// The tag name pointing to the forked commit
	TagName string `json:"tag_name" api:"required"`
	// The new VM that was created from the fork
	VmID string                     `json:"vm_id" api:"required"`
	JSON forkRepositoryResponseJSON `json:"-"`
}

// forkRepositoryResponseJSON contains the JSON metadata for the struct
// [ForkRepositoryResponse]
type forkRepositoryResponseJSON struct {
	CommitID    apijson.Field
	Reference   apijson.Field
	RepoName    apijson.Field
	TagName     apijson.Field
	VmID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ForkRepositoryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r forkRepositoryResponseJSON) RawJSON() string {
	return r.raw
}

// Response body for GET /api/v1/repositories/{repo_name}/tags
type ListRepoTagsResponse struct {
	// The repository name
	Repository string `json:"repository" api:"required"`
	// List of tags in this repository
	Tags []RepoTagInfo            `json:"tags" api:"required"`
	JSON listRepoTagsResponseJSON `json:"-"`
}

// listRepoTagsResponseJSON contains the JSON metadata for the struct
// [ListRepoTagsResponse]
type listRepoTagsResponseJSON struct {
	Repository  apijson.Field
	Tags        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListRepoTagsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listRepoTagsResponseJSON) RawJSON() string {
	return r.raw
}

// Response body for GET /api/v1/repositories
type ListRepositoriesResponse struct {
	// List of all repositories in the user's organization
	Repositories []RepositoryInfo             `json:"repositories" api:"required"`
	JSON         listRepositoriesResponseJSON `json:"-"`
}

// listRepositoriesResponseJSON contains the JSON metadata for the struct
// [ListRepositoriesResponse]
type listRepositoriesResponseJSON struct {
	Repositories apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ListRepositoriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listRepositoriesResponseJSON) RawJSON() string {
	return r.raw
}

// Tag information within a repository context
type RepoTagInfo struct {
	// The commit ID this tag currently points to
	CommitID string `json:"commit_id" api:"required" format:"uuid"`
	// When the tag was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Full reference in image_name:tag format
	Reference string `json:"reference" api:"required"`
	// The tag's unique identifier
	TagID string `json:"tag_id" api:"required" format:"uuid"`
	// The tag name
	TagName string `json:"tag_name" api:"required"`
	// When the tag was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Optional description
	Description string          `json:"description" api:"nullable"`
	JSON        repoTagInfoJSON `json:"-"`
}

// repoTagInfoJSON contains the JSON metadata for the struct [RepoTagInfo]
type repoTagInfoJSON struct {
	CommitID    apijson.Field
	CreatedAt   apijson.Field
	Reference   apijson.Field
	TagID       apijson.Field
	TagName     apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RepoTagInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r repoTagInfoJSON) RawJSON() string {
	return r.raw
}

// Repository information returned in list and get operations
type RepositoryInfo struct {
	// When the repository was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether this repository is publicly visible
	IsPublic bool `json:"is_public" api:"required"`
	// The repository name
	Name string `json:"name" api:"required"`
	// The repository's unique identifier
	RepoID string `json:"repo_id" api:"required" format:"uuid"`
	// Optional description
	Description string             `json:"description" api:"nullable"`
	JSON        repositoryInfoJSON `json:"-"`
}

// repositoryInfoJSON contains the JSON metadata for the struct [RepositoryInfo]
type repositoryInfoJSON struct {
	CreatedAt   apijson.Field
	IsPublic    apijson.Field
	Name        apijson.Field
	RepoID      apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RepositoryInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r repositoryInfoJSON) RawJSON() string {
	return r.raw
}

// Request body for PATCH /api/v1/repositories/{repo_name}/visibility
type SetRepositoryVisibilityRequestParam struct {
	// Whether the repository should be publicly visible
	IsPublic param.Field[bool] `json:"is_public" api:"required"`
}

func (r SetRepositoryVisibilityRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Request body for PATCH /api/v1/repositories/{repo_name}/tags/{tag_name}
type UpdateRepoTagRequestParam struct {
	// Optional new commit ID to move the tag to
	CommitID param.Field[string] `json:"commit_id" format:"uuid"`
	// Optional new description for the tag. Send `null` to clear.
	Description param.Field[string] `json:"description"`
}

func (r UpdateRepoTagRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RepositoryNewParams struct {
	// Request body for POST /api/v1/repositories
	CreateRepositoryRequest CreateRepositoryRequestParam `json:"create_repository_request" api:"required"`
}

func (r RepositoryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateRepositoryRequest)
}

type RepositoryNewTagParams struct {
	// Request body for creating a tag within a repository: POST
	// /api/v1/repositories/{repo_name}/tags
	CreateRepoTagRequest CreateRepoTagRequestParam `json:"create_repo_tag_request" api:"required"`
}

func (r RepositoryNewTagParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateRepoTagRequest)
}

type RepositoryForkParams struct {
	// Request body for POST /api/v1/repositories/fork
	ForkRepositoryRequest ForkRepositoryRequestParam `json:"fork_repository_request" api:"required"`
}

func (r RepositoryForkParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ForkRepositoryRequest)
}

type RepositorySetVisibilityParams struct {
	// Request body for PATCH /api/v1/repositories/{repo_name}/visibility
	SetRepositoryVisibilityRequest SetRepositoryVisibilityRequestParam `json:"set_repository_visibility_request" api:"required"`
}

func (r RepositorySetVisibilityParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SetRepositoryVisibilityRequest)
}

type RepositoryUpdateTagParams struct {
	// Request body for PATCH /api/v1/repositories/{repo_name}/tags/{tag_name}
	UpdateRepoTagRequest UpdateRepoTagRequestParam `json:"update_repo_tag_request" api:"required"`
}

func (r RepositoryUpdateTagParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.UpdateRepoTagRequest)
}
