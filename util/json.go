package util

import (
	"net/http"
	"encoding/json"
)

func ConvertObjectToJson(object interface{}, writer http.ResponseWriter) []byte  {
	bytes, err := json.Marshal(object)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	return bytes
}

func WriteJsonResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}