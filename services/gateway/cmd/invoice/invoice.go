package invoice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/arfis/go-invoice/gateway/pkg/model"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"os"
)

type Handler interface {
	GetAll() ([]*model.Invoice, error)
}

type InvoiceService struct{}

func (InvoiceService *InvoiceService) GetInvoiceItemsByInvoiceID(id *uuid.UUID) (*[]model.InvoiceItem, error) {
	var a, b = 2, 3
	fmt.Printf("%d %d", a, b)
	return nil, nil
}

func (invoiceService *InvoiceService) GetCompanyByID(id *uuid.UUID) (model.Company, error) {
	return model.Company{Street: "Hurbanova 14", City: "Hlohovec", CompanyNumber: "123451", CompanyTaxNumber: "SK123451", TaxNumber: "123129321", Zipcode: "92001"}, nil
}

func (invoiceService *InvoiceService) Update(id string, updateData map[string]interface{}) (*model.Invoice, error) {
	return nil, nil
}

func (invoiceService *InvoiceService) Delete(id string) (bool, error) {
	return true, nil
}

func (invoiceService *InvoiceService) Create(input model.Invoice) (*model.Invoice, error) {
	invoiceServiceURL := os.Getenv("INVOICE_SERVICE_URL")

	if invoiceServiceURL == "" {
		log.Fatal("INVOICE_SERVICE_URL not set")
	}

	log.Printf("!!Raw request body: %+v", input)

	jsonData, err := json.Marshal(input)
	if err != nil {
		log.Printf("Error marshalling input to JSON: %v", err)
		return nil, err
	}
	log.Printf("Sending JSON data: %s", jsonData)

	bodyBuffer := bytes.NewBuffer(jsonData)

	client := &http.Client{}

	req, err := http.NewRequest("POST", invoiceServiceURL+"/invoice", bodyBuffer)
	if err != nil {
		log.Fatalf("Error creating request: %s", err)
		return nil, err
	}

	// Set the content type to application/json; this is important!
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %s", err)
		return nil, err
	}

	defer resp.Body.Close()

	// Optionally, read the response body
	// Note: In production code, you should handle the response properly.
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %s", err)
		return nil, err
	}
	return transformToInvoice(responseBody)
}

func (invoiceService *InvoiceService) GetById(id string) (*model.Invoice, error) {
	invoiceServiceURL := os.Getenv("INVOICE_SERVICE_URL")
	var invoice *model.Invoice

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/invoice?id=%s", invoiceServiceURL, id), nil)
	if err != nil {
		fmt.Printf("Got error %s", err.Error())
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, errRead := io.ReadAll(resp.Body) // Read response body
	if errRead != nil {
		log.Fatalf("Error reading response body: %s", errRead.Error())
		return nil, nil
	}

	// Read the response body
	errUnmarshal := json.Unmarshal(body, &invoice)
	if errUnmarshal != nil {
		log.Fatal("Error unmarshalling JSON: ", err)
	}

	return invoice, nil
}

func (invoiceService *InvoiceService) GetAll() ([]*model.Invoice, error) {
	invoiceServiceURL := os.Getenv("INVOICE_SERVICE_URL")
	var invoices []*model.Invoice

	if invoiceServiceURL == "" {
		log.Fatal("INVOICE_SERVICE_URL not set")
	}

	// Create a new HTTP client
	client := &http.Client{}
	req, err := http.NewRequest("GET", invoiceServiceURL+"/invoice?id=2", nil)
	if err != nil {
		fmt.Printf("Got error %s", err.Error())
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, errRead := io.ReadAll(resp.Body) // Read response body
	if errRead != nil {
		log.Fatalf("Error reading response body: %s", errRead.Error())
		return nil, nil
	}

	// Read the response body
	errUnmarshal := json.Unmarshal(body, &invoices)
	if errUnmarshal != nil {
		log.Fatal("Error unmarshalling JSON: ", err)
	}

	return invoices, nil
}

func transformToInvoice(data []byte) (*model.Invoice, error) {
	var invoice *model.Invoice

	log.Printf("!!BEFORE UNMARSHAL request body: %+v", data)

	errUnmarshal := json.Unmarshal(data, &invoice)
	if errUnmarshal != nil {
		log.Fatal("Error unmarshalling JSON: ", errUnmarshal)
	}
	return invoice, errUnmarshal
}
