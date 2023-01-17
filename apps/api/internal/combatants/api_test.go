package combatants

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

func TestApi_GenerateCombatants(t *testing.T) {
	api := TestApi()

	appEnv := common.AppEnv{
		Render: render.New(),
	}
	response := spec.NewMockHttpResponseWriter()
	req := &http.Request{
		Body: io.NopCloser(strings.NewReader("{\"Count\":2}")),
	}
	api.GenerateCombatants(response, req, appEnv)

	spec.ExpectEqualInts(t, 200, response.StatusCode, "Unexpected status code")
}

func TestApi_GenerateCombatants_InvalidBody(t *testing.T) {
	api := TestApi()

	appEnv := common.AppEnv{
		Render: render.New(),
	}
	response := spec.NewMockHttpResponseWriter()
	req := &http.Request{
		Body: io.NopCloser(strings.NewReader("{\"Count\":\"invalid\"}")),
	}
	api.GenerateCombatants(response, req, appEnv)

	spec.ExpectEqualInts(t, 400, response.StatusCode, "Unexpected status code")
}
