/*
   Author: Maen Abu Hammour
   Date: October 5, 2023
   Description: This Go file is the application's entry point. It initializes HTTP handlers for both GraphQL and RESTful APIs, starts the servers on port 8080, and logs their status. If any errors occur during server setup, it logs the error and terminates the application.
*/

package main

import (
	"book-info-service/src/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Builds HTTP handlers for both: GraphQL and RESTFul APIs
	handlers.BuildHandlers()

	fmt.Println("GraphQL server is running on :8080/books-graphql")
	fmt.Println("RESTful server is running on :8080/books")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
