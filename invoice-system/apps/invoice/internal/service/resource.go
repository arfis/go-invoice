package service

import (
	"fmt"
	"github.com/arfis/go-invoice/invoice/internal/db"
	"github.com/arfis/go-invoice/invoice/internal/model"
	"github.com/google/uuid"
	"log"
)

type InvoiceError struct {
	Message string
}

func (e *InvoiceError) Error() string {
	return e.Message
}

func AutoMigrate() {
	dbConnection := db.GetConnection()
	err := dbConnection.AutoMigrate(&model.Invoice{})
	if err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}
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

func GetInvoiceById(id string) (*model.Invoice, error) {
	uuid, error := stringToUUID(id)

	if error != nil {
		return nil, error
	}
	invoice := model.Invoice{ID: uuid}
	result := db.GetConnection().First(&invoice)
	if result.Error != nil {
		return nil, &InvoiceError{Message: "Db connection failed"}
	}

	return &invoice, nil
}

func CreateInvoice(invoice model.Invoice) (*model.Invoice, error) {
	result := db.GetConnection().Create(&invoice)
	if result.Error != nil {
		return nil, &InvoiceError{Message: "Db connection failed"}
	}

	return &invoice, nil
}

func CreateCompany(company model.Company) (*model.Company, error) {

	result := db.GetConnection().Create(&company)
	if result.Error != nil {
		return nil, &InvoiceError{Message: "Db connection failed"}
	}

	return &company, nil
}

func stringToUUID(str string) (*uuid.UUID, error) {
	// Parse the string to a UUID
	id, err := uuid.Parse(str)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID format: %v", err)
	}
	return &id, nil
}
