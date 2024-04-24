package service

import (
	"fmt"
	"github.com/arfis/go-invoice/invoice/internal/db"
	"github.com/arfis/go-invoice/invoice/internal/model"
	"log"
)

type InvoiceError struct {
	Message string
}

func (e *InvoiceError) Error() string {
	return e.Message
}

func DropInvoices() {
	dbConnection := db.GetConnection()
	dbConnection.Migrator().DropTable(&model.Invoice{})
	dbConnection.Migrator().DropTable(&model.Company{})
	err := dbConnection.AutoMigrate(&model.Invoice{})
	if err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	err2 := dbConnection.AutoMigrate(&model.InvoiceItem{})
	if err2 != nil {
		log.Fatalf("AutoMigrate failed: %v", err2)
	}

}

func GetInvoices() (*[]model.Invoice, error) {
	dbConnection := db.GetConnection()

	DropInvoices()
	//err := dbConnection.AutoMigrate(&model.Invoice{})

	var invoices []model.Invoice

	query := dbConnection.Model(&model.Invoice{})

	result := query.Find(&invoices)
	if result.Error != nil {
		log.Fatalf("Error when getting product: %v", result.Error)
		return nil, result.Error
	}
	fmt.Printf("\n Invoices from DB: %d \n %s", len(invoices), invoices)

	return &invoices, nil
}

func CreateInvoice(invoice model.Invoice) (*model.Invoice, error) {
	// Validate input data
	//if invoiceInput.Code == "" || invoiceInput.OwnerId == "" {
	//	//http.Error(w, , http.StatusBadRequest)
	//	return nil, &InvoiceError{Message: "Wrong input parameters"}
	//}
	//
	//// Create a new Invoice instance from the input
	//invoice := model.Invoice{
	//	Code:    invoiceInput.Code,
	//	Price:   invoiceInput.Price,
	//	OwnerId: invoiceInput.OwnerId,
	//}

	fmt.Printf("??!!BEFORE CREATING request body: %+v", invoice)
	result := db.GetConnection().Create(&invoice)
	if result.Error != nil {
		return nil, &InvoiceError{Message: "Db connection failed"}
	}

	return &invoice, nil
}

//func mapInvoiceToResponse(invoice model.Invoice) model.InvoiceResponse {
//	return model.InvoiceResponse{
//		Code:  invoice.Code,
//		Price: invoice.Price,
//	}
//}
