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
	clearStorage()
	api := Api()

	if api == nil {
		t.Fatalf("Api() returned nil")
	}

	if api.application == nil {
		t.Fatalf("Api object has no application")
	}
}

func TestApi_GenerateCombatants(t *testing.T) {
	clearStorage()
	api := Api()

	appEnv := common.AppEnv{
		Render: render.New(),
	}
	response := spec.NewMockHttpResponseWriter()
	req := &http.Request{
		Body: io.NopCloser(strings.NewReader("{\"Count\":2}")),
	}
	api.GenerateCombatants(response, req, appEnv)

	spec.ExpectEqualInts(t, 200, response.StatusCode, "Unexpected status code")
	spec.ExpectEqualInts(t, 2, len(storage.combatants), "Unexpected stored combatant count")
}

func TestApi_GenerateCombatants_InvalidBody(t *testing.T) {
	clearStorage()
	api := Api()

	appEnv := common.AppEnv{
		Render: render.New(),
	}
	response := spec.NewMockHttpResponseWriter()
	req := &http.Request{
		Body: io.NopCloser(strings.NewReader("{\"Count\":\"invalid\"}")),
	}
	api.GenerateCombatants(response, req, appEnv)

	spec.ExpectEqualInts(t, 400, response.StatusCode, "Unexpected status code")
	spec.ExpectEqualInts(t, 0, len(storage.combatants), "Unexpected stored combatant count")
}
