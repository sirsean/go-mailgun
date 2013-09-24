# go-mailgun

go-mailgun is a simple library that lets you send email messages via Mailgun.

# Configuration

Before sending any emails, you'll need to set the Mailgun API endpoint and API key.

```
mailgun.ApiEndpoint = "ENDPOINT"
mailgun.ApiKey = "KEY"
```

# Sending Messages

When you send a message, you have to specify the sender's name and email address, the recipient's email address, the subject, and the body.

It's probably a good idea to do this inside a goroutine.

```
go func() {
    mailgun.Send(mailgun.Message{
        FromName: "Foo Bar",
        FromAddress: "foo@bar.test",
        ToAddress: "recipient@bar.test",
        Subject: "This is an example message",
        Body: "It's pretty easy to send messages via Mailgun!",
    })
}()
```

go-mailgun currently doesn't support CC or BCC or attachments, or any of the other cool stuff you can do with emails. Mailgun does support all those things, so it wouldn't be hard to add those features if you want them.

# Example App

See the example app.

# License

Licensed under the Apache 2.0 license.
