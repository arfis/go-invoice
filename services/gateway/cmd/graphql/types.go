package customgraph

import (
	"github.com/arfis/go-invoice/gateway/pkg/model"
	"github.com/graphql-go/graphql"
)

var CompanyType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Company",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"companyNumber": &graphql.Field{
				Type: graphql.String,
			},
			"taxNumber": &graphql.Field{
				Type: graphql.String,
			},
			"companyTaxNumber": &graphql.Field{
				Type: graphql.String,
			},
			"street": &graphql.Field{
				Type: graphql.String,
			},
			"city": &graphql.Field{
				Type: graphql.String,
			},
			"country": &graphql.Field{
				Type: graphql.String,
			},
			"zipcode": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var InvoiceItemType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "InvoiceItem",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"invoiceId": &graphql.Field{
				Type: graphql.Int,
			},
			"code": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var InvoiceType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Invoice",
		Description: "Details of an invoice",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"createdDate": &graphql.Field{
				Type: graphql.DateTime,
			},
			"validUntil": &graphql.Field{
				Type: graphql.DateTime,
			},
			"daysToPay": &graphql.Field{
				Type: graphql.Int,
			},
			"emailTo": &graphql.Field{
				Type: graphql.String,
			},
			"ownerId": &graphql.Field{
				Type: graphql.String,
			},
			"receivingCompany": &graphql.Field{
				Type: CompanyType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Assuming you have a way to fetch Company by ID
					// This is just a placeholder
					return invoiceService.GetCompanyByID(p.Source.(model.Invoice).ReceivingCompanyID)
				},
			},
			"issuingCompany": &graphql.Field{
				Type: CompanyType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Assuming you have a way to fetch Company by ID
					return invoiceService.GetCompanyByID(p.Source.(model.Invoice).IssuingCompanyID)
				},
			},
			"invoiceItems": &graphql.Field{
				Type: graphql.NewList(InvoiceItemType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Fetch Invoice Items related to the invoice
					return invoiceService.GetInvoiceItemsByInvoiceID(p.Source.(model.Invoice).ID)
				},
			},
		},
	},
)
