package main

import (
	"fmt"
	"github.com/sirsean/go-mailgun/mailgun"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	mg_client := mailgun.NewClient("YOUR-API-KEY-HERE", "YOUR-DOMAIN-HERE.mailgun.org")

	message1 := mailgun.Message{
		FromName:    "Spacely Sprockets",
		FromAddress: "spacely@spacelysprockets.com",
		ToAddress:   "george@thejetsons.com",
		Subject:     "Go Mailgun sample message",
		Body:        "It's *way* easy to send messages via the Go Mailgun API!",
	}

	fmt.Println("Attempting to send to ", mg_client.Endpoint(message1))

	body, err := mg_client.Send(message1)
	if err != nil {
		fmt.Println("Got an error:", err)
	} else {
		fmt.Println(body)
	}

	// Also try to send a MIME message

	mime_file := "mailgun/test_data/message.mime"
	file, err := os.Open(mime_file)
	if err != nil {
		log.Fatal("Error when opening '" + mime_file + "'")
	}

	mimeContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Error when reading '" + mime_file + "'")
	}

	message2 := mailgun.MimeMessage{
		ToAddress:   "foo@testorghere.com",
		Content: mimeContents}

	fmt.Println("Attempting to send to ", mg_client.Endpoint(message2))

	body, err = mg_client.Send(message2)
	if err != nil {
		fmt.Println("Got an error:", err)
	} else {
		fmt.Println(body)
	}
}
