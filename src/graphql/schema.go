/*
   Author: Maen Abu Hammour
   Date: October 5, 2023
   Description: This file is responsible for defining and initializing the GraphQL schema for a backend service. It imports necessary libraries for logging and GraphQL operations, then creates a GraphQL schema using the RootQuery and RootMutation objects defined elsewhere in the code. If any errors occur during schema creation, it utilizes the log4go library for logging and exits with an error message. This file plays a crucial role in configuring and setting up the GraphQL schema for the application's queries and mutations.
*/

package graphql

import (
	l4g "github.com/alecthomas/log4go"
	"github.com/graphql-go/graphql"
	gql "github.com/graphql-go/graphql"
)

// Define the schema with Query and Mutation
var Schema gql.Schema

func init() {
	var err error
	Schema, err = gql.NewSchema(graphql.SchemaConfig{
		Query:    RootQuery,
		Mutation: RootMutation,
	})
	if err != nil {
		l4g.Exitf("Error creating GraphQL schema: %v", err)
	}
}
