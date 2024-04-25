package customgraph

import (
	"fmt"
	"github.com/arfis/go-invoice/gateway/pkg/model"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"log"
)

var CompanyType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Company",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: UUIDType,
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

var UUIDType = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "UUID",
	Description: "A UUID scalar type represented as a string",
	Serialize: func(value interface{}) interface{} {
		fmt.Println("SERIALIZE")
		switch value := value.(type) {
		case uuid.UUID:
			return value.String()
		default:
			return value
		}
	},
	ParseValue: func(value interface{}) interface{} {
		if str, ok := value.(string); ok {
			return str // Return the string representation directly
		}
		log.Printf("Value is not a string: %v", value)
		return nil
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		if str, ok := valueAST.(*ast.StringValue); ok {
			return str.Value // Return the string value directly
		}
		log.Printf("Value is not a string literal: %v", valueAST)
		return nil
	},
})
