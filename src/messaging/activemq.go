package messaging

import (
	"context"
	"log"
	"time"

	"github.com/Azure/go-amqp"
)

func ConnectToActiveMQ() (*amqp.Conn, error) {
	// Define the connection URL
	uri := `broker="0.0.0.0",component=addresses,address="pedidos",subcomponent=queues,routing-type="anycast",queue="resposta"`

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to the ActiveMQ server
	client, err := amqp.Dial(ctx, uri, nil)

	if err != nil {
		log.Printf("Failed to connect to ActiveMQ: %v", err)
		return nil, err
	}

	log.Println("Successfully connected to ActiveMQ")
	return client, nil
}
