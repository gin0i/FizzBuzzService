package handlers

import (
	"encoding/json"
	"net/http"
	"FizzBuzzService/managers"
)


func HandleGetBiggestRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		result, err := managers.BiggestRequest()
		if err != nil {
			switch err.Error() {
			case "No request records":
				http.Error(w, "No requests records yet", http.StatusNotFound)
			default:
				http.Error(w, "Failed to get biggest request", http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(*result)
	}
}
