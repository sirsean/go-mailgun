package main

import (
    "github.com/sirsean/go-mailgun/mailgun"
)

func main() {
    mailgun.ApiEndpoint = "https://api.mailgun.net/v2/YOURNAME.mailgun.org/messages"
    mailgun.ApiKey = "YOURKEY"

    go func() {
        err := mailgun.Send(mailgun.Message{
            FromName: "Foo Bar",
            FromAddress: "foo@bar.test",
            ToAddress: "recipient@bar.test",
            Subject: "This is an example message",
            Body: "It's pretty easy to send messages via Mailgun!",
        })
        if err != nil {
            // you can handle sending errors here
        }
    }()
}
