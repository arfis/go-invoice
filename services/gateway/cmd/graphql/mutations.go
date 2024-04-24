package customgraph

import (
	"github.com/arfis/go-invoice/gateway/cmd/invoice"
	"github.com/arfis/go-invoice/gateway/pkg/model"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

var invoiceService invoice.InvoiceService

func GetMutation() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createInvoice": &graphql.Field{
				Type: InvoiceType,
				Args: graphql.FieldConfigArgument{
					"emailTo": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"ownerId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"receivingCompanyId": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"issuingCompanyId": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					receivingCompanyId, rcOk := params.Args["receivingCompanyId"].(*uuid.UUID)
					issuingCompanyId, icOk := params.Args["issuingCompanyId"].(*uuid.UUID)

					invoice := model.Invoice{
						EmailTo: params.Args["emailTo"].(string),
						OwnerId: params.Args["ownerId"].(string),
					}

					if rcOk {
						invoice.ReceivingCompanyID = receivingCompanyId
					}

					if icOk {
						invoice.IssuingCompanyID = issuingCompanyId
					}

					// Assume CreateInvoice is a function to save the invoice to the database
					return invoiceService.Create(invoice)
				},
			},
			"updateInvoice": &graphql.Field{
				Type: InvoiceType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"emailTo": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					// Include other fields as necessary
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id := params.Args["id"].(string)
					// Create an update map
					update := map[string]interface{}{}
					if emailTo, ok := params.Args["emailTo"].(string); ok {
						update["emailTo"] = emailTo
					}
					// Add other fields to update map as necessary
					return invoiceService.Update(id, update)
				},
			},
			"deleteInvoice": &graphql.Field{
				Type: graphql.Int,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id := params.Args["id"].(string)
					return invoiceService.Delete(id)
				},
			},
			// Additional mutations can be added here
		},
	})
}
