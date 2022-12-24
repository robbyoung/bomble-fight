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
	clearStorage()
	api := Api()

	if api == nil {
		t.Fatalf("Api() returned nil")
	}

	if api.application == nil {
		t.Fatalf("Api object has no application")
	}
}

func TestApi_CreatePlayer(t *testing.T) {
	clearStorage()
	api := Api()

	appEnv := common.AppEnv{
		Render: render.New(),
	}
	response := spec.NewMockHttpResponseWriter()
	req := &http.Request{
		Body: io.NopCloser(strings.NewReader("{\"Name\":\"Percy\"}")),
	}
	api.CreatePlayer(response, req, appEnv)

	spec.ExpectEqualInts(t, 200, response.StatusCode, "Unexpected status code")
	spec.ExpectEqualInts(t, 1, len(storage.players), "Unexpected stored combatant count")
}

func TestApi_CreatePlayer_InvalidBody(t *testing.T) {
	clearStorage()
	api := Api()

	appEnv := common.AppEnv{
		Render: render.New(),
	}
	response := spec.NewMockHttpResponseWriter()
	req := &http.Request{
		Body: io.NopCloser(strings.NewReader("{\"Name\":8}")),
	}
	api.CreatePlayer(response, req, appEnv)

	spec.ExpectEqualInts(t, 400, response.StatusCode, "Unexpected status code")
	spec.ExpectEqualInts(t, 0, len(storage.players), "Unexpected stored combatant count")
}

func TestApi_CreatePlayer_EmptyName(t *testing.T) {
	clearStorage()
	api := Api()

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
	spec.ExpectEqualInts(t, 0, len(storage.players), "Unexpected stored combatant count")
}
