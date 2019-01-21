package main

import (
	"log"
	"net/http"

	"FizzBuzzService/routers"
	"FizzBuzzService/models/request"
	"FizzBuzzService/storage"
	"FizzBuzzService/config"

	"github.com/gorilla/handlers"
	"github.com/codegangsta/negroni"
)



func main() {
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "content-type",  "Accept"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST"})
	allowCredentials := handlers.AllowCredentials()


	storage.Storage().AutoMigrate(&request.Request{})


	router := routers.NewRouter()
	negroniTool := negroni.Classic()
	negroniTool.UseHandler(handlers.CORS(allowedHeaders, allowedMethods, allowCredentials)(router))

	log.Println("Starting service...")
	log.Fatal(http.ListenAndServe(":"+config.GetListenPort(), negroniTool))
}
