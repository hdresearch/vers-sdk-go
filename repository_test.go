// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vers_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/hdresearch/vers-sdk-go"
	"github.com/hdresearch/vers-sdk-go/internal/testutil"
	"github.com/hdresearch/vers-sdk-go/option"
)

func TestRepositoryNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := vers.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Repositories.New(context.TODO(), vers.RepositoryNewParams{
		CreateRepositoryRequest: vers.CreateRepositoryRequestParam{
			Name:        vers.F("name"),
			Description: vers.F("description"),
		},
	})
	if err != nil {
		var apierr *vers.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRepositoryList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := vers.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Repositories.List(context.TODO())
	if err != nil {
		var apierr *vers.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRepositoryDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := vers.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Repositories.Delete(context.TODO(), "repo_name")
	if err != nil {
		var apierr *vers.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRepositoryNewTagWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := vers.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Repositories.NewTag(
		context.TODO(),
		"repo_name",
		vers.RepositoryNewTagParams{
			CreateRepoTagRequest: vers.CreateRepoTagRequestParam{
				CommitID:    vers.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				TagName:     vers.F("tag_name"),
				Description: vers.F("description"),
			},
		},
	)
	if err != nil {
		var apierr *vers.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRepositoryDeleteTag(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := vers.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Repositories.DeleteTag(
		context.TODO(),
		"repo_name",
		"tag_name",
	)
	if err != nil {
		var apierr *vers.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRepositoryForkWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := vers.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Repositories.Fork(context.TODO(), vers.RepositoryForkParams{
		ForkRepositoryRequest: vers.ForkRepositoryRequestParam{
			SourceOrg:  vers.F("source_org"),
			SourceRepo: vers.F("source_repo"),
			SourceTag:  vers.F("source_tag"),
			RepoName:   vers.F("repo_name"),
			TagName:    vers.F("tag_name"),
		},
	})
	if err != nil {
		var apierr *vers.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRepositoryGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := vers.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Repositories.Get(context.TODO(), "repo_name")
	if err != nil {
		var apierr *vers.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRepositoryGetTag(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := vers.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Repositories.GetTag(
		context.TODO(),
		"repo_name",
		"tag_name",
	)
	if err != nil {
		var apierr *vers.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRepositoryListTags(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := vers.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Repositories.ListTags(context.TODO(), "repo_name")
	if err != nil {
		var apierr *vers.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRepositorySetVisibility(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := vers.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Repositories.SetVisibility(
		context.TODO(),
		"repo_name",
		vers.RepositorySetVisibilityParams{
			SetRepositoryVisibilityRequest: vers.SetRepositoryVisibilityRequestParam{
				IsPublic: vers.F(true),
			},
		},
	)
	if err != nil {
		var apierr *vers.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRepositoryUpdateTagWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := vers.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Repositories.UpdateTag(
		context.TODO(),
		"repo_name",
		"tag_name",
		vers.RepositoryUpdateTagParams{
			UpdateRepoTagRequest: vers.UpdateRepoTagRequestParam{
				CommitID:    vers.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Description: vers.F("description"),
			},
		},
	)
	if err != nil {
		var apierr *vers.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
