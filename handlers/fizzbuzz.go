package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"FizzBuzzService/managers"
)

type RawFizzBuzzRequest struct {
	MulA  int    `json:"int1"`
	MulB  int    `json:"int2"`
	Limit int    `json:"limit"`
	StrA  string `json:"str1"`
	StrB  string `json:"str2"`
}

func checkRanges(req RawFizzBuzzRequest) error {
	if req.MulA <= 0 || req.MulB <= 0 {
		return errors.New("Multiple value must be > 0")
	}

	if req.Limit <= 0 || req.Limit > 65000 {
		return errors.New("Limit value must be > 0 and < 65000")
	}

	return nil
}

func HandleFizzBuzz(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var received RawFizzBuzzRequest
		if r.Body == nil {
			http.Error(w, "Please send a request body", http.StatusBadRequest)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&received)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = checkRanges(received)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := managers.FizzBuzz(received.MulA, received.MulB, received.StrA, received.StrB, received.Limit)
		if err != nil {
			switch err.Error() {
			default:
				http.Error(w, "Failed to fizzbuzz", http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}
