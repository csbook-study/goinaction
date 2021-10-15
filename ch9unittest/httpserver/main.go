// This sample code implement a simple web service.
package main

import (
	"log"
	"net/http"

	"github.com/huxiangyu99/goinaction/ch9unittest/httpserver/handlers"
)

// main is the entry point for the application.
func main() {
	handlers.Routes()

	// http://localhost:4000/sendjson
	log.Println("listener : Started : Listening on :4000")
	http.ListenAndServe(":4000", nil)
}
