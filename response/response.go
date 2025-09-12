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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.WriteHeader(status)

	if errMsg != "" {
		errResponse := Error{
			Message: errMsg,
		}
		err := json.NewEncoder(w).Encode(errResponse)
		if err != nil {
			return
		}
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, `{"status":500,"error":"failed to encode response"}`, http.StatusInternalServerError)
	}
}

func Ok(w http.ResponseWriter, data interface{}) {
	writeResponse(w, http.StatusOK, data, "")
}

func NoContent(w http.ResponseWriter) {
	writeResponse(w, http.StatusNoContent, nil, "")
}

func NotFound(w http.ResponseWriter, errMsg string) {
	writeResponse(w, http.StatusNotFound, nil, errMsg)
}

func BadRequest(w http.ResponseWriter, errMsg string) {
	writeResponse(w, http.StatusBadRequest, nil, errMsg)
}
