/*
   Author: Maen Abu Hammour
   Date: October 5, 2023
   Description: This Go file in the "handlers" package configures HTTP handlers for GraphQL and RESTful APIs, mapping routes to request-handling functions for book data operations while also serving a Swagger JSON file for API documentation.
*/

package handlers

import (
	local_qraphql "book-info-service/src/graphql"
	"net/http"

	"github.com/graphql-go/handler"
)

// BuildHandlers initializes and configures HTTP handlers for both GraphQL and RESTful APIs.
// @Description Initialize and configure HTTP handlers.
// @Router / [post]
func BuildHandlers() error {

	// Create a GraphQL HTTP handler
	graphqlHandler := handler.New(&handler.Config{
		Schema: &local_qraphql.Schema,
		Pretty: true,
	})
	http.Handle("/books-graphql", graphqlHandler)

	// List of RESTful handlers (for the RESTful part of the application)
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetAllBooks(w)
		case http.MethodPost:
			CreateBook(w, r)
		default:
			http.Error(w, "Invalid HTTP Method", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetBookByIDHandler(w, r)
		case http.MethodPut:
			UpdateBook(w, r)
		case http.MethodDelete:
			DeleteBookHandler(w, r)
		default:
			http.Error(w, "Invalid HTTP Method", http.StatusMethodNotAllowed)
		}
	})

	// Serve the Swagger JSON file
	http.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "docs/swagger/doc.json")
	})

	return nil
}
