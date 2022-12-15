package combatants

import (
	"bomble-fight/internal/common"
	"bomble-fight/pkg/status"
	"encoding/json"
	"net/http"
	"strconv"
)

type api struct {
	application ICombatantApplication
}

func newApi(app *application) *api {
	return &api{
		application: app,
	}
}

func (a *api) GenerateCombatants(w http.ResponseWriter, req *http.Request, appEnv common.AppEnv) {
	var requestBody generateCombatantsRequest
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&requestBody)

	if err != nil {
		response := status.Response{
			Status:  strconv.Itoa(http.StatusBadRequest),
			Message: "malformed post body",
		}
		appEnv.Render.JSON(w, http.StatusBadRequest, response)
		return
	}

	result := a.application.GenerateCombatants(requestBody.Count)

	responseObject := make(map[string]interface{})
	responseObject["combatants"] = result
	appEnv.Render.JSON(w, http.StatusOK, responseObject)
}
