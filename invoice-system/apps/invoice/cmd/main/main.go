package main

import (
	"fmt"
	"github.com/arfis/go-invoice/invoice/internal/api"
	db2 "github.com/arfis/go-invoice/invoice/internal/db"
	messageQueue "github.com/arfis/go-invoice/invoice/internal/message"
	"github.com/arfis/go-invoice/invoice/internal/service"
	"net/http"
)

// this is still nil
var (
	db                *db2.Database
	invoiceController *api.InvoiceController
	companyController *api.CompanyController
)

func main() {

	// Register the GET handler
	mux := http.NewServeMux()

	// Registering routes related to invoices
	//service.DropInvoices()
	service.AutoMigrate()
	invoiceController.RegisterRoutes(mux)
	companyController.RegisterRoutes(mux)
	messageQueue.StartListening()

	fmt.Println("Invoice Server starting on port 8080...")
	err := http.ListenAndServe("0.0.0.0:8080", mux) // This will block and keep the program running
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

}
