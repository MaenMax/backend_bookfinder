/*
   Author: Maen Abu Hammour
   Date: October 5, 2023
   Description: This Go code defines a GraphQL API for managing book information within a backend service. It provides a set of GraphQL queries and mutations, including operations to retrieve a list of books, retrieve a single book by ID, create new books, update existing books, and delete books. The code integrates with a storage system and utilizes GraphQL's schema definition to enable structured interactions with book data, making it suitable for building a book information service with GraphQL capabilities.
*/

package graphql

import (
	"book-info-service/src/model"
	"book-info-service/src/storage"

	l4g "github.com/alecthomas/log4go"
	"github.com/graphql-go/graphql"
)

var (
	RootQuery    *graphql.Object
	RootMutation *graphql.Object
)

func init() {

	// Define the root Query
	RootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"books": &graphql.Field{
				Type: graphql.NewList(model.GraphQLBook),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					book, err := storage.DB.GetAllBooks()
					return book, err
				},
			},
			"book": &graphql.Field{
				Type: model.GraphQLBook,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, ok := params.Args["id"].(int)
					if ok {
						book, err := storage.DB.GetBookByID(id)
						if err == nil {
							return book, nil
						}
					}
					return nil, nil
				},
			},
		},
	})

	// Define the root Mutation for CRUD operations
	RootMutation = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createBook": &graphql.Field{
				Type: model.GraphQLBook,
				Args: graphql.FieldConfigArgument{
					"Title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"Author": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"DatePublished": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"BookCoverURL": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					title := params.Args["Title"].(string)
					author := params.Args["Author"].(string)
					datePublished := params.Args["DatePublished"].(string)
					bookCoverURL := params.Args["BookCoverURL"].(string)

					book := model.Book{
						Title:         title,
						Author:        author,
						DatePublished: datePublished,
						BookCoverURL:  bookCoverURL,
					}

					id, err := storage.DB.CreateBook(book)
					book.ID = id
					return book, err
				},
			},
			"updateBook": &graphql.Field{
				Type: model.GraphQLBook,
				Args: graphql.FieldConfigArgument{
					"ID": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"Title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Author": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"DatePublished": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"BookCoverURL": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id := params.Args["ID"].(int)
					book, err := storage.DB.GetBookByID(id)
					if err != nil {
						return nil, l4g.Error("Book with ID %d not found", id)
					}

					if title, ok := params.Args["Title"].(string); ok {
						book.Title = title
					}
					if author, ok := params.Args["Author"].(string); ok {
						book.Author = author
					}
					if datePublished, ok := params.Args["DatePublished"].(string); ok {
						book.DatePublished = datePublished
					}
					if bookCoverURL, ok := params.Args["BookCoverURL"].(string); ok {
						book.BookCoverURL = bookCoverURL
					}

					err = storage.DB.UpdateBook(id, book)
					return book, err
				},
			},
			"deleteBook": &graphql.Field{
				Type: model.GraphQLBook,
				Args: graphql.FieldConfigArgument{
					"ID": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id := params.Args["ID"].(int)
					book, err := storage.DB.DeleteBook(id)
					return book, err
				},
			},
		},
	})

}
