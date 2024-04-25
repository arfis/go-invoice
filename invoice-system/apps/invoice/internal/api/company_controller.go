package api

import (
	"encoding/json"
	"fmt"
	"github.com/arfis/go-invoice/invoice/internal/model"
	"github.com/arfis/go-invoice/invoice/internal/service"
	"io"
	"log"
	"net/http"
)

type CompanyController struct{}

func (cc *CompanyController) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/company", handleCompanies)
}

func handleCompanies(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\nGot request for company")
	if r.Method == "GET" {
		invoices, err := service.GetInvoices()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(invoices)
	} else if r.Method == "POST" {

		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read the request body: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Log the raw body to ensure what we receive is correct

		// Decode the body
		var input model.Company
		if err := json.Unmarshal(bodyBytes, &input); err != nil {
			http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Process the decoded invoice
		company, err := service.CreateCompany(input)
		if err != nil {
			http.Error(w, "Error creating invoice: "+err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("\n\n\nRETURNING raw JSON: %v", company.ID)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(company)
	}
}
