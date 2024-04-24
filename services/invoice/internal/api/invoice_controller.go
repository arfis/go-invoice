package api

import (
	"encoding/json"
	"github.com/arfis/go-invoice/invoice/internal/model"
	"github.com/arfis/go-invoice/invoice/internal/service"
	"io"
	"log"
	"net/http"
)

func RegisterInvoiceRoutes(router *http.ServeMux) {
	router.HandleFunc("/invoice", handleInvoices)
	router.HandleFunc("/ping", pingHandler)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) // Send an HTTP 200 OK status
	w.Write([]byte("pong"))      // Respond with "pong"
}

func handleInvoices(w http.ResponseWriter, r *http.Request) {
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
		log.Printf("Received raw JSON: %s", string(bodyBytes))

		// Decode the body
		var input model.Invoice
		if err := json.Unmarshal(bodyBytes, &input); err != nil {
			http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Process the decoded invoice
		invoice, err := service.CreateInvoice(input)
		if err != nil {
			http.Error(w, "Error creating invoice: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(invoice)
	}
}
