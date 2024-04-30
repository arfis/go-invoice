package sharedCommands

type Command string

const (
	CreatePDF   Command = "CREATE_PDF"
	SendInvoice Command = "SEND_INVOICE"
)

type CommandStep string

const (
	PdfCreated  CommandStep = "PDF_CREATED"
	InvoiceSent CommandStep = "INVOICE_SENT"
)
