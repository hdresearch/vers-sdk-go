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

// CommitService contains methods and other services that help with interacting
// with the vers API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCommitService] method instead.
type CommitService struct {
	Options []option.RequestOption
}

// NewCommitService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCommitService(opts ...option.RequestOption) (r *CommitService) {
	r = &CommitService{}
	r.Options = opts
	return
}

func (r *CommitService) Update(ctx context.Context, commitID string, body CommitUpdateParams, opts ...option.RequestOption) (res *CommitInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	if commitID == "" {
		err = errors.New("missing required commit_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/commits/%s", commitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

func (r *CommitService) List(ctx context.Context, opts ...option.RequestOption) (res *ListCommitsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/commits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *CommitService) Delete(ctx context.Context, commitID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if commitID == "" {
		err = errors.New("missing required commit_id parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/commits/%s", commitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

func (r *CommitService) ListParents(ctx context.Context, commitID string, opts ...option.RequestOption) (res *[]CommitListParentsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if commitID == "" {
		err = errors.New("missing required commit_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/vm/commits/%s/parents", commitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *CommitService) ListPublic(ctx context.Context, opts ...option.RequestOption) (res *ListCommitsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/commits/public"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type CommitInfo struct {
	CommitID            string         `json:"commit_id" api:"required"`
	CreatedAt           string         `json:"created_at" api:"required"`
	IsPublic            bool           `json:"is_public" api:"required"`
	Name                string         `json:"name" api:"required"`
	OwnerID             string         `json:"owner_id" api:"required"`
	Description         string         `json:"description" api:"nullable"`
	GrandparentCommitID string         `json:"grandparent_commit_id" api:"nullable"`
	ParentVmID          string         `json:"parent_vm_id" api:"nullable"`
	JSON                commitInfoJSON `json:"-"`
}

// commitInfoJSON contains the JSON metadata for the struct [CommitInfo]
type commitInfoJSON struct {
	CommitID            apijson.Field
	CreatedAt           apijson.Field
	IsPublic            apijson.Field
	Name                apijson.Field
	OwnerID             apijson.Field
	Description         apijson.Field
	GrandparentCommitID apijson.Field
	ParentVmID          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CommitInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitInfoJSON) RawJSON() string {
	return r.raw
}

type ListCommitsResponse struct {
	Commits []CommitInfo            `json:"commits" api:"required"`
	Limit   int64                   `json:"limit" api:"required"`
	Offset  int64                   `json:"offset" api:"required"`
	Total   int64                   `json:"total" api:"required"`
	JSON    listCommitsResponseJSON `json:"-"`
}

// listCommitsResponseJSON contains the JSON metadata for the struct
// [ListCommitsResponse]
type listCommitsResponseJSON struct {
	Commits     apijson.Field
	Limit       apijson.Field
	Offset      apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListCommitsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listCommitsResponseJSON) RawJSON() string {
	return r.raw
}

// Request body for PATCH /commits/{commit_id}
type UpdateCommitRequestParam struct {
	IsPublic param.Field[bool] `json:"is_public" api:"required"`
	// Optional description for the commit.
	Description param.Field[string] `json:"description"`
	// Optional human-readable name for the commit.
	Name param.Field[string] `json:"name"`
}

func (r UpdateCommitRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CommitListParentsResponse struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether this commit is publicly accessible (readable/restorable by anyone).
	IsPublic bool   `json:"is_public" api:"required"`
	Name     string `json:"name" api:"required"`
	// api key id.
	OwnerID     string `json:"owner_id" api:"required" format:"uuid"`
	Description string `json:"description" api:"nullable"`
	// The commit that this commit's parent VM was started from, if any. Intended to
	// optimize traversing the commit tree.
	GrandparentCommitID string `json:"grandparent_commit_id" api:"nullable" format:"uuid"`
	// The VM that this commit was created from, if any.
	ParentVmID string                        `json:"parent_vm_id" api:"nullable" format:"uuid"`
	JSON       commitListParentsResponseJSON `json:"-"`
}

// commitListParentsResponseJSON contains the JSON metadata for the struct
// [CommitListParentsResponse]
type commitListParentsResponseJSON struct {
	ID                  apijson.Field
	CreatedAt           apijson.Field
	IsPublic            apijson.Field
	Name                apijson.Field
	OwnerID             apijson.Field
	Description         apijson.Field
	GrandparentCommitID apijson.Field
	ParentVmID          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CommitListParentsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitListParentsResponseJSON) RawJSON() string {
	return r.raw
}

type CommitUpdateParams struct {
	// Request body for PATCH /commits/{commit_id}
	UpdateCommitRequest UpdateCommitRequestParam `json:"update_commit_request" api:"required"`
}

func (r CommitUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.UpdateCommitRequest)
}
