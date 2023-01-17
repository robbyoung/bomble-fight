package players

import (
	"bomble-fight/internal/common"
	"bomble-fight/internal/spec"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/unrolled/render"
)

func TestApi_New(t *testing.T) {
	api := Api()

	if api == nil {
		t.Fatalf("Api() returned nil")
	}
}

func TestApi_CreatePlayer(t *testing.T) {
	api := TestApi()

	appEnv := common.AppEnv{
		Render: render.New(),
	}
	response := spec.NewMockHttpResponseWriter()
	req := &http.Request{
		Body: io.NopCloser(strings.NewReader("{\"Name\":\"Percy\"}")),
	}
	api.CreatePlayer(response, req, appEnv)

	spec.ExpectEqualInts(t, 200, response.StatusCode, "Unexpected status code")
}

func TestApi_CreatePlayer_InvalidBody(t *testing.T) {
	api := TestApi()

	appEnv := common.AppEnv{
		Render: render.New(),
	}
	response := spec.NewMockHttpResponseWriter()
	req := &http.Request{
		Body: io.NopCloser(strings.NewReader("{\"Name\":8}")),
	}
	api.CreatePlayer(response, req, appEnv)

	spec.ExpectEqualInts(t, 400, response.StatusCode, "Unexpected status code")
}

func TestApi_CreatePlayer_EmptyName(t *testing.T) {
	api := TestApi()

	appEnv := common.AppEnv{
		Render: render.New(),
	}
	response := spec.NewMockHttpResponseWriter()
	req := &http.Request{
		Body: io.NopCloser(strings.NewReader("{\"Name\":\"\"}")),
	}
	api.CreatePlayer(response, req, appEnv)

	spec.ExpectEqualInts(t, 400, response.StatusCode, "Unexpected status code")
	spec.ExpectIncludedString(t, "Player name can't be empty", response.Text, "Unexpected response text")
}
