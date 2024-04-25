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

func (invoiceService *InvoiceService) Create(resourceType string, input interface{}) (interface{}, error) {
	// Retrieve the URL for the service
	invoiceServiceURL := os.Getenv("INVOICE_SERVICE_URL")
	if invoiceServiceURL == "" {
		log.Fatal("INVOICE_SERVICE_URL not set")
	}

	// Marshal input data to JSON
	jsonData, err := json.Marshal(input)
	if err != nil {
		log.Printf("Error marshalling input to JSON: %v", err)
		return nil, err
	}
	log.Printf("Sending JSON data: %s", jsonData)

	// Create request body
	bodyBuffer := bytes.NewBuffer(jsonData)

	// Create HTTP client
	client := &http.Client{}

	// Create request based on resource type
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", invoiceServiceURL, resourceType), bodyBuffer)
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
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %s", err)
		return nil, err
	}

	// Process response based on resource type
	switch resourceType {
	case "invoice":
		return transformToInvoice(responseBody)
	case "company":
		return transformToCompany(responseBody)
	// Add cases for other resource types if needed
	default:
		return nil, fmt.Errorf("Unsupported resource type: %s", resourceType)
	}
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

func transformToCompany(data []byte) (*model.Company, error) {
	var company *model.Company

	log.Printf("!!BEFORE UNMARSHAL request body: %q", string(data))

	errUnmarshal := json.Unmarshal(data, &company)
	if errUnmarshal != nil {
		log.Printf("Error unmarshalling JSON: %v", errUnmarshal)
		return nil, errUnmarshal
	}
	return company, errUnmarshal
}
