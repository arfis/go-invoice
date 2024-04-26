package server

import (
	"encoding/json"
	"fmt"
	messageQueue "github.com/arfis/go-invoice/gateway/internal/message"
	"github.com/arfis/go-invoice/gateway/internal/services/invoice"
	"github.com/arfis/go-invoice/gateway/pkg/model"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	invoiceService        invoice.InvoiceService
	commandMessageFactory messageQueue.CommandMessage
)

var producer = messageQueue.GetProducerInstance()

type RestApiServer struct{}

func (ra *RestApiServer) StartWebServer(port uint, terminateChan chan int) error {
	r := mux.NewRouter()
	r.HandleFunc("/ping", pingHandler).Methods("GET")
	r.HandleFunc("/invoice", createInvoiceHandler).Methods("POST")
	r.HandleFunc("/invoice", getInvoicesHandler).Methods("GET")
	r.HandleFunc("/invoice/{id}/send", sendInvoiceHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
	terminateChan <- 1
	return nil
}

func createInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	var input model.Invoice
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	invoiceService.Create("invoice", input)
}

func getInvoicesHandler(w http.ResponseWriter, r *http.Request) {
	result, err := invoiceService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, result)
}

func sendInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var body struct {
		EmailTo string `json:"emailTo"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	command := commandMessageFactory.Send(id, body.EmailTo)
	producer.SendOperation(command)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
