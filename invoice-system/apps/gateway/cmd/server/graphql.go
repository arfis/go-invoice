package server

import (
	"fmt"
	customgraph "github.com/arfis/go-invoice/gateway/cmd/graphql"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
)

type GraphQLServer struct{}

func (ra *GraphQLServer) StartWebServer(port uint, terminateChan chan int) error {
	rootQuery := customgraph.GetRootQuery()
	mutation := customgraph.GetMutation()
	schemaConfig := graphql.SchemaConfig{Query: rootQuery, Mutation: mutation}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Failed to create GraphQL schema: %v", err)
		return err
	}

	http.Handle("/graphql", handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	}))

	http.HandleFunc("/sandbox", sandboxHandler)

	log.Println("Starting GraphQL server on http://localhost:8081/graphql")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

	terminateChan <- 1
	return nil
}

func sandboxHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "cmd/server/index.html")
}
