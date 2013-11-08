# go-mailgun

go-mailgun is a simple library that sends email messages via Mailgun.

# Prerequisites

Before sending any emails, an account on Mailgun is required, and a domain must
be configured and an API key created.

# Sending messages from Go

There are two ways to send email through Mailgun using this module.

1. Specify individual fields of the message
2. create a MIME email message and send that.

## Client configuration

A Mailgun client is configured by passing in the API key and the domain

```
	mg_client := mailgun.NewClient("key-9ic-qz3tvx1id18e3e5cf950fueirqh3", "mydomain.org.mailgun.org")
```

## Sending an Email using Individual Fields

This is a simple way to send mail through the API. When sending a message,
the sender's email address, the recipient's email address, the subject,
and the body must all be specified.

```
	message := mailgun.Message{
		FromName:    "Foo From",
		FromAddress: "foo@fakedomaingoeshere.com",
		ToAddress:   "bar@anotherfakedomainhere.com",
		Subject:     "test message",
		Body:        "This is the body of the message. It's not very interesting.",
	}

	mg_client.Send(message)
```

## Sending a pre-generated MIME message

This part of the API will take a MIME message and send it via Mailgun's
MIME endpoint. It does not do any validation on the MIME message passed
in.

```
	mime_message := mailgun.MimeMessage{
		ToAddress:   "sent_to_address@fakedomainhere.com",
		Content: mimeContents}

	mg_client.Send(message2)
```

# Example App

See the example app (from which most of the above code was ripped).

# License

Licensed under the Apache 2.0 license.
