package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Azure/go-amqp"

	"go-challenge/src/messaging"
	"go-challenge/src/model"
	"go-challenge/src/repository"
	"go-challenge/src/repository/db"
)

func main() {
	fmt.Println("Starting execution!")

	db, err := db.ConnectToPostgres()

	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	addressAndQueue := "pedidos::resposta"

	// create connection
	connection, err := messaging.ConnectToActiveMQ()

	if err != nil {
		fmt.Println("Failed to connect to the ActiveMQ:", err)
		return
	}
	// create session
	session, err := connection.NewSession(context.TODO(), nil)

	if err != nil {
		fmt.Println("Failed to create session:", err)
		return
	}

	// create a new sender
	sender, err := session.NewSender(context.TODO(), addressAndQueue, nil)
	if err != nil {
		fmt.Println("Failed to create sender:", err)
		return
	}

	// create a new receiver
	receiver, err := session.NewReceiver(context.TODO(), addressAndQueue, nil)

	if err != nil {
		fmt.Println("Failed to create receiver:", err)
		return
	}

	isCreateProduct := false
	isSearchProduct := false
	isUpdateProduct := false
	isSendMessage := false
	isReceiveMessage := false

	if isCreateProduct {
		repository.CreateProduct(db, &model.Product{
			Product: "Pão de batata",
		})
	}

	if isSearchProduct {
		product, err := repository.GetProductByID(db, 1)

		if err != nil {
			fmt.Println("Failed to get product:", err)
			return
		}

		productJSON, err := json.Marshal(product)

		if err != nil {
			fmt.Println("Failed to get product:", err)
			return
		}

		log.Printf("Search result => : %s", productJSON)
	}

	if isUpdateProduct {
		// Update

		// Initialize the map
		updatedData := make(map[string]interface{})

		// Add some key-value pairs
		updatedData["product"] = "Pao de mandioca"
		updatedData["updated_at"] = "2024-09-28T12:14:00.000Z"

		product, err := repository.UpdateProductByID(db, 1, updatedData)

		if err != nil {
			fmt.Println("Failed to update product:", err)
			return
		}

		productJSON, err := json.Marshal(product)

		if err != nil {
			fmt.Println("Its not possible to serialize result:", err)
			return
		}

		log.Printf("Updated result => : %s", productJSON)
	}

	if isSendMessage {
		// send a message
		err = sender.Send(context.TODO(), amqp.NewMessage([]byte("Hello!")), nil)
		if err != nil {
			fmt.Println("Failed to send message:", err)
			return
		}
	}

	if isReceiveMessage {
		// receive the next message
		msg, err := receiver.Receive(context.TODO(), nil)
		if err != nil {
			// handle error
		}

		msgParsed := string(msg.GetData())

		fmt.Println("Message:", msgParsed)

		receiver.AcceptMessage(context.TODO(), msg)
	}

}
