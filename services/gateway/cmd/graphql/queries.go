package customgraph

import (
	"errors"
	"github.com/graphql-go/graphql"
)

func GetRootQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"invoices": &graphql.Field{
				Type: graphql.NewList(InvoiceType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return invoiceService.GetAll()
				},
			},
			"invoice": &graphql.Field{
				Type: InvoiceType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(string)
					if !ok {
						return nil, errors.New("id is required and should be an integer")
					}
					return invoiceService.GetById(id)
				},
			},
		},
	})
}
