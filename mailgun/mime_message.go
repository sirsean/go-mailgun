package mailgun

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

type MimeMessage struct {
	ToAddress string
	Content   []byte
}

func (message MimeMessage) IsValid() (validity bool) {
	if message.ToAddress == "" || string(message.Content) == "" {
		return false
	}

	return true
}

func (message MimeMessage) MimeReader() (b io.Reader, boundary string) {
	buffer := new(bytes.Buffer)
	mimeWriter := multipart.NewWriter(buffer)
	boundary = mimeWriter.Boundary()

	go func() {
		defer mimeWriter.Close()
		mimeWriter.WriteField("to", message.ToAddress)

		messageField, err := mimeWriter.CreateFormFile("message", "message.mime")
		if err != nil {
			log.Fatal("Could not create MIME part for the 'message' field!")
		}
		messageField.Write(message.Content)
	}()

	return buffer, boundary
}

func (message MimeMessage) GetRequest(mailgun Client) (request *http.Request) {
	mimeReader, boundary := message.MimeReader()
	request, _ = http.NewRequest("POST", mailgun.Endpoint(message), mimeReader)
	request.Header.Set("content-type", fmt.Sprintf("multipart/form-data; boundary=%s", boundary))
	return
}

func (message MimeMessage) Endpoint() string {
	return "messages.mime"
}
