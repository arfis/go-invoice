package commands

import "fmt"

type SendInvoice struct {
	InvoiceId string
	MailTo    string
}

func (sendInvoice *SendInvoice) Execute() {
	fmt.Printf("...............Sending to %s\n", sendInvoice.MailTo)
}
