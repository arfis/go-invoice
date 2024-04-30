package messageQueue

import sharedCommands "lib/commands"

type CommandMessage struct {
	Command sharedCommands.Command `json:"command"`
	Data    map[string]interface{} `json:"data"`
}

func (commandMessage *CommandMessage) CreatePdf(invoiceId string) CommandMessage {
	return CommandMessage{
		Command: sharedCommands.CreatePDF,
		Data:    map[string]interface{}{"InvoiceId": invoiceId},
	}
}

func (commandMessage *CommandMessage) Send(invoiceId string, mailTo string) CommandMessage {
	return CommandMessage{
		Command: sharedCommands.SendInvoice,
		Data:    map[string]interface{}{"InvoiceId": invoiceId, "MailTo": mailTo},
	}
}
