package main

import (
	"fmt"
	"github.com/arfis/go-invoice/gateway/cmd/server"
	"github.com/arfis/go-invoice/gateway/internal/services/invoice"
)

// this is still nil
var (
	invoiceService invoice.InvoiceService
)
var terminateChan = make(chan int, 2)

func main() {
	// this automatically does an object an assigns an address
	//mux := http.NewServeMux()

	// Registering routes related to invoices
	//api.RegisterInvoiceRoutes(mux)
	//invoiceService.Handle()

	//db.TrySelect()
	//db.Test()

	fmt.Println("Server starting on port 81212")

	fmt.Println("Server starting on port 1.")

	//var tf util.TestingFunc

	//tf("123", 3)
	//var producer = messageQueue.GetProducerInstance()
	var graphQlServer = server.GraphQLServer{}
	var restApiServer = server.RestApiServer{}

	go Startup(&graphQlServer, 8081)
	go Startup(&restApiServer, 8080)

	//messageQueue.StartListening()

	for i := 0; i < 2; i++ {
		<-terminateChan
	}
	//
	//
	//

	//
	//// Add other routes
	//
	//
	//// If you're using a wrapper function, apply it
	//
	//// Start the server
	//err := http.ListenAndServe("0.0.0.0:8080", handler) // Use the wrapped handler
	//if err != nil {
	//	fmt.Println("Error starting server:", err)
	//}
}

func Startup(startupServer server.StartupHandler, port uint) {
	startupServer.StartWebServer(port, terminateChan)
}
