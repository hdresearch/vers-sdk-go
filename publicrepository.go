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
	"github.com/hdresearch/vers-sdk-go/internal/requestconfig"
	"github.com/hdresearch/vers-sdk-go/option"
)

// PublicRepositoryService contains methods and other services that help with
// interacting with the vers API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPublicRepositoryService] method instead.
type PublicRepositoryService struct {
	Options []option.RequestOption
}

// NewPublicRepositoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPublicRepositoryService(opts ...option.RequestOption) (r *PublicRepositoryService) {
	r = &PublicRepositoryService{}
	r.Options = opts
	return
}

func (r *PublicRepositoryService) List(ctx context.Context, opts ...option.RequestOption) (res *ListPublicRepositoriesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/public/repositories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *PublicRepositoryService) Get(ctx context.Context, orgName string, repoName string, opts ...option.RequestOption) (res *PublicRepositoryInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	if orgName == "" {
		err = errors.New("missing required org_name parameter")
		return nil, err
	}
	if repoName == "" {
		err = errors.New("missing required repo_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/public/repositories/%s/%s", orgName, repoName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *PublicRepositoryService) GetTag(ctx context.Context, orgName string, repoName string, tagName string, opts ...option.RequestOption) (res *RepoTagInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	if orgName == "" {
		err = errors.New("missing required org_name parameter")
		return nil, err
	}
	if repoName == "" {
		err = errors.New("missing required repo_name parameter")
		return nil, err
	}
	if tagName == "" {
		err = errors.New("missing required tag_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/public/repositories/%s/%s/tags/%s", orgName, repoName, tagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *PublicRepositoryService) ListTags(ctx context.Context, orgName string, repoName string, opts ...option.RequestOption) (res *ListRepoTagsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if orgName == "" {
		err = errors.New("missing required org_name parameter")
		return nil, err
	}
	if repoName == "" {
		err = errors.New("missing required repo_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/public/repositories/%s/%s/tags", orgName, repoName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Response body for GET /api/v1/public/repositories
type ListPublicRepositoriesResponse struct {
	Repositories []PublicRepositoryInfo             `json:"repositories" api:"required"`
	JSON         listPublicRepositoriesResponseJSON `json:"-"`
}

// listPublicRepositoriesResponseJSON contains the JSON metadata for the struct
// [ListPublicRepositoriesResponse]
type listPublicRepositoriesResponseJSON struct {
	Repositories apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ListPublicRepositoriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listPublicRepositoriesResponseJSON) RawJSON() string {
	return r.raw
}

// Public repository information (includes owner org name for namespacing)
type PublicRepositoryInfo struct {
	// When the repository was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Full reference: org_name/repo_name
	FullName string `json:"full_name" api:"required"`
	// The repository name
	Name string `json:"name" api:"required"`
	// The owning organization's name (namespace)
	OrgName string `json:"org_name" api:"required"`
	// The repository's unique identifier
	RepoID string `json:"repo_id" api:"required" format:"uuid"`
	// Optional description
	Description string                   `json:"description" api:"nullable"`
	JSON        publicRepositoryInfoJSON `json:"-"`
}

// publicRepositoryInfoJSON contains the JSON metadata for the struct
// [PublicRepositoryInfo]
type publicRepositoryInfoJSON struct {
	CreatedAt   apijson.Field
	FullName    apijson.Field
	Name        apijson.Field
	OrgName     apijson.Field
	RepoID      apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PublicRepositoryInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r publicRepositoryInfoJSON) RawJSON() string {
	return r.raw
}
