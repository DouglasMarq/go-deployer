package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DouglasMarq/go-deployer/constants"
	"github.com/DouglasMarq/go-deployer/internal/worker"
	"github.com/DouglasMarq/go-deployer/utils/api"
)

var (
	response constants.DeploymentResponse
)

func DeployHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.Message = "method not allowed"
		api.WriteToJSON(w,
			response,
			http.StatusMethodNotAllowed)
		return
	}

	var payload constants.DeploymentRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response.Message = "invalid body"
		api.WriteToJSON(w,
			response,
			http.StatusBadRequest)
		return
	}

	if (payload.ScriptType == "" || len(payload.ScriptType) == 0) ||
		(payload.FullPath == "" || len(payload.FullPath) == 0) {
		response.Message = "invalid body"
		api.WriteToJSON(w,
			response,
			http.StatusBadRequest)
		return
	}

	if err := worker.Execute(payload.ScriptType, payload.FullPath); err != nil {
		response.Message = "failed to execute script"
		response.Reason = err.Error()
		api.WriteToJSON(w,
			response,
			http.StatusInternalServerError)
		return
	}

	response.Message = "success"
	api.WriteToJSON(w, response, http.StatusOK)
	return
}
