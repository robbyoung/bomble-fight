package bomble

import (
	"bomble-fight/internal/common"

	"github.com/unrolled/render"
)

// CreateContextForTestSetup initialises an application context struct
// for testing purposes
func CreateContextForTestSetup() common.AppEnv {
	testVersion := "0.1.0"
	appEnv := common.AppEnv{
		Render:  render.New(),
		Version: testVersion,
		Env:     "LOCAL",
		Port:    "3001",
	}
	return appEnv
}
