package response

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message string `json:"error"`
}

func writeResponse(w http.ResponseWriter, status int, data interface{}, errMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if errMsg != "" {
		errResponse := Error{
			Message: errMsg,
		}
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, `{"status":500,"error":"failed to encode response"}`, http.StatusInternalServerError)
	}
}

func Ok(w http.ResponseWriter, data interface{}) {
	writeResponse(w, http.StatusOK, data, "")
}

func NoContent(w http.ResponseWriter, data interface{}) {
	writeResponse(w, http.StatusNoContent, data, "")
}

func NotFound(w http.ResponseWriter, data interface{}, errMsg string) {
	writeResponse(w, http.StatusNotFound, data, errMsg)
}
