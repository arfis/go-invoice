package customgraph

import (
	"fmt"
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
						Type: UUIDType,
					},
					"issuingCompanyId": &graphql.ArgumentConfig{
						Type: UUIDType,
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
					return invoiceService.Create("invoice", invoice)
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
			"createCompany": &graphql.Field{
				Type: CompanyType,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"companyNumber": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"taxNumber": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"companyTaxNumber": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"street": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"city": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"country": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"zipcode": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					// Retrieve arguments from ResolveParams
					name := params.Args["name"].(string)
					companyNumber := params.Args["companyNumber"].(string)
					taxNumber, _ := params.Args["taxNumber"].(string)
					companyTaxNumber, _ := params.Args["companyTaxNumber"].(string)
					street := params.Args["street"].(string)
					city := params.Args["city"].(string)
					country := params.Args["country"].(string)
					zipcode := params.Args["zipcode"].(string)

					// Create a new company object
					newCompany := &model.Company{
						// Assuming you have a Company struct defined
						// with corresponding fields
						Name:             name,
						CompanyNumber:    companyNumber,
						TaxNumber:        taxNumber,
						CompanyTaxNumber: companyTaxNumber,
						Street:           street,
						City:             city,
						Country:          country,
						Zipcode:          zipcode,
					}

					// Perform logic to create the company, such as storing it in a database
					// For demonstration purposes, let's assume a function `createCompany` is available
					createdCompany, err := invoiceService.Create("company", newCompany)
					fmt.Println("LOCAL TEST")
					fmt.Println(createdCompany)
					fmt.Printf("Retrieved company %s", createdCompany)
					if err != nil {
						return nil, err
					}

					// Return the newly created company
					return createdCompany, nil
				},
			},
			// Additional mutations can be added here
		},
	})
}
