package messageQueue

import (
	"encoding/json"
	"fmt"
	"github.com/arfis/go-invoice/invoice/internal/commands"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	sharedCommands "lib/commands"
	"log"
	"os"
	"time"
)

type CommandMessage struct {
	Command     sharedCommands.Command     `json:"command"`
	CommandStep sharedCommands.CommandStep `json:"commandStep"`
	Data        map[string]interface{}     `json:"data"`
}

func StartListening() {
	kafkaUrl := os.Getenv("KAFKA_URL")

	fmt.Println("Starting listening")
	if kafkaUrl == "" {
		log.Fatal("KAFKA WAS NOT SET!")
	}

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaUrl,
		"group.id":          "invoiceMicroservice",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}

	defer c.Close()

	for {
		err := c.Subscribe("operation", nil)
		if err != nil {
			log.Printf("Subscribe failed: %v, retrying...", err)
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}

	for {
		msg, err := c.ReadMessage(-1) // Block until a message is received

		var cmdMsg CommandMessage
		if err := json.Unmarshal(msg.Value, &cmdMsg); err != nil {
			log.Printf("Failed to decode message: %v\n", err)
			continue
		}

		if err == nil {
			switch cmdMsg.Command {
			case sharedCommands.CreatePDF:
				{
					invoiceId := cmdMsg.Data["InvoiceId"].(string)
					pdfPrinter := commands.PrintPdf{InvoiceId: invoiceId}
					pdfPrinter.Execute()
				}

			case sharedCommands.SendInvoice:
				{
					invoiceId := cmdMsg.Data["InvoiceId"].(string)
					mailTo := cmdMsg.Data["MailTo"].(string)
					sendInvoice := commands.SendInvoice{InvoiceId: invoiceId, MailTo: mailTo}
					sendInvoice.Execute()
				}
			}
			fmt.Printf("Received message: %s\n", string(msg.Value))
		} else {
			// Handle errors
			if kafkaErr, ok := err.(kafka.Error); ok && kafkaErr.Code() == kafka.ErrAllBrokersDown {
				log.Fatalf("All brokers are down: %v", err)
			} else {
				log.Printf("Error reading message: %v\n", err)
			}
		}
	}
}
