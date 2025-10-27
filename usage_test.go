// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vers_test

import (
	"context"
	"os"
	"testing"

	"github.com/hdresearch/vers-sdk-go"
	"github.com/hdresearch/vers-sdk-go/internal/testutil"
	"github.com/hdresearch/vers-sdk-go/option"
)

func TestUsage(t *testing.T) {
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
	newVmResponse, err := client.Vm.NewRoot(context.TODO(), vers.VmNewRootParams{
		NewRootRequest: vers.NewRootRequestParam{
			VmConfig: vers.F(vers.NewRootRequestVmConfigParam{}),
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", newVmResponse.ID)
}
