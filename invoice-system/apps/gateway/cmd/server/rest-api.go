package server

import (
	"encoding/json"
	"fmt"
	"github.com/arfis/go-invoice/gateway/internal/services/invoice"
	"github.com/arfis/go-invoice/gateway/pkg/model"
	"log"
	"net/http"
)

var invoiceService invoice.InvoiceService

type RestApiServer struct{}

func (ra *RestApiServer) StartWebServer(port uint, terminateChan chan int) error {
	handler := http.NewServeMux()
	handler.HandleFunc("/ping", pingHandler) // Set up the ping handler
	handler.HandleFunc("/invoice", invoiceHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
	terminateChan <- 1
	return nil
}

func invoiceHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		{
			var input model.Invoice
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			invoiceService.Create("invoice", input)
		}

	case "GET":
		{
			result, error := invoiceService.GetAll()
			if error != nil {

			}
			fmt.Fprint(w, result)
		}
	}

}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) // Send an HTTP 200 OK status
	w.Write([]byte("pong"))      // Respond with "pong"
}
