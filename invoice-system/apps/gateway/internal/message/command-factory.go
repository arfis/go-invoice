package messageQueue

type CommandMessage struct {
	Command string                 `json:"command"`
	Data    map[string]interface{} `json:"data"`
}

func (commandMessage *CommandMessage) PrintPdf(invoiceId string) CommandMessage {
	return CommandMessage{
		Command: "PrintPDF",
		Data:    map[string]interface{}{"InvoiceId": invoiceId},
	}
}

func (commandMessage *CommandMessage) Send(invoiceId string, mailTo string) CommandMessage {
	return CommandMessage{
		Command: "Send",
		Data:    map[string]interface{}{"InvoiceId": invoiceId, "MailTo": mailTo},
	}
}
