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

func TestAPIRootfList(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.API.Rootfs.List(context.TODO())
	if err != nil {
		var apierr *vers.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAPIRootfDelete(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.API.Rootfs.Delete(context.TODO(), "rootfs_id")
	if err != nil {
		var apierr *vers.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAPIRootfUploadWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.API.Rootfs.Upload(
		context.TODO(),
		"rootfs_id",
		vers.APIRootfUploadParams{
			Dockerfile: vers.F(vers.APIRootfUploadParamsDockerfile{
				Dockerfile: vers.F("dockerfile"),
			}),
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
