package commands

import "fmt"

type PrintPdf struct {
	InvoiceId string
}

func (printPdf *PrintPdf) Execute() {
	fmt.Printf("...............Printing\n")
}
