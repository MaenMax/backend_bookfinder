package model

import "github.com/graphql-go/graphql"

// Define the Book struct to represent book information (For RESTful API)
type Book struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	DatePublished string `json:"date_published"`
	BookCoverURL  string `json:"book_cover_url"`
}

// Define GraphQL type for Book (for graphql API)
var GraphQLBook = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type: graphql.Int,
		},
		"Title": &graphql.Field{
			Type: graphql.String,
		},
		"Author": &graphql.Field{
			Type: graphql.String,
		},
		"DatePublished": &graphql.Field{
			Type: graphql.String,
		},
		"BookCoverURL": &graphql.Field{
			Type: graphql.String,
		},
	},
})
