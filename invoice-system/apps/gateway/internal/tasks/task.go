package tasks

type Task struct{}

type PrintPdf struct {
	InvoiceId string
}

type SendInvoice struct {
	InvoiceId string
	MailTo    string
}
