package combatants

import (
	"bomble-fight/internal/common"
	"bomble-fight/pkg/status"
	"encoding/json"
	"net/http"
	"strconv"
)

type CombatantApi struct {
	application ICombatantApplication
}

func NewCombatantApi() *CombatantApi {
	return &CombatantApi{
		application: createApplication(),
	}
}

func (a *CombatantApi) GenerateCombatants(w http.ResponseWriter, req *http.Request, appEnv common.AppEnv) {
	var count int
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&count)

	if err != nil {
		response := status.Response{
			Status:  strconv.Itoa(http.StatusBadRequest),
			Message: "malformed post body",
		}
		appEnv.Render.JSON(w, http.StatusBadRequest, response)
		return
	}

	result := a.application.GenerateCombatants(count)

	responseObject := make(map[string]interface{})
	responseObject["combatants"] = result
	appEnv.Render.JSON(w, http.StatusOK, responseObject)
}
