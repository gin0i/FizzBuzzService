package routers

import (
	"FizzBuzzService/handlers"
	"github.com/gorilla/mux"
)


func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc(
		"/fizzbuzz/api/launch",
		handlers.HandleFizzBuzz).Methods("POST")

	router.HandleFunc(
		"/fizzbuzz/api/stat",
		handlers.HandleGetBiggestRequest).Methods("GET")


	return router
}
