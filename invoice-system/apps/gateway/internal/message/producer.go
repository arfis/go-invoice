package messageQueue

import (
	"fmt"
	"github.com/arfis/go-invoice/gateway/internal/util"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
	"os"
	"sync"
)

var (
	once             sync.Once
	onceErr          error
	producerInstance *Producer
)

type Producer struct {
	instance *kafka.Producer
}

func GetProducerInstance() *Producer {
	once.Do(func() {
		kafkaUrl := os.Getenv("KAFKA_URL")
		if kafkaUrl == "" {
			log.Fatal("KAFKA_URL was not set") // It might be better to handle this more gracefully.
			return
		}

		p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaUrl})
		if err != nil {
			log.Fatalf("Failed to create producer: %v", err) // Consider handling this error differently.
			return
		}
		fmt.Printf("Creating instance %s\n", p.String())

		if producerInstance == nil {
			producerInstance = &Producer{instance: p}
		}

		// Start a background goroutine to handle delivery reports
		go func() {
			for e := range p.Events() {
				switch ev := e.(type) {
				case *kafka.Message:
					if ev.TopicPartition.Error != nil {
						fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
					} else {
						fmt.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n",
							*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
					}
				}
			}
		}()
	})

	return producerInstance
}

func (producer *Producer) SendOperation(command CommandMessage) {
	instance := producer.instance
	if instance == nil {
		fmt.Println("!Producer instance is not created")
		return
	}
	topic := "operation"
	operationAsBytes, err := util.ConvertStructToBytes(command)
	if err != nil {
		fmt.Println("There was a problem converting operation to bytes:", err)
		return
	}

	instance.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          operationAsBytes,
	}, nil)
}

// GenerateOperations should be called after ensuring producer is initialized and no error has occurred.
func (producer *Producer) GenerateOperations() {
	messageQueue := CommandMessage{}
	commands := []CommandMessage{
		messageQueue.CreatePdf("2"),
		messageQueue.Send("3", "michalsevcikk@gmail.com"),
		messageQueue.Send("5", "michalsevcikk@gmail.com"),
	}
	for _, command := range commands {
		producer.SendOperation(command)
	}
}

// Close should be called when the application is shutting down or when you're sure no more messages will be produced.
func (producer *Producer) CloseProducer() {
	if producer.instance != nil {
		producer.instance.Close()
	}
}
