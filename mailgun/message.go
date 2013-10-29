package mailgun

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Message struct {
	FromName       string
	FromAddress    string
	ToAddress      string
	CCAddressList  []string
	BCCAddressList []string
	Subject        string
	Body           string
	AttachmentList []string
	InlineList     []string
}

func (m Message) From() string {
	return fmt.Sprintf("%s <%s>", m.FromName, m.FromAddress)
}

func (m Message) BCCAddresses() string {
	return strings.Join(m.BCCAddressList, ", ")
}

func (m Message) CCAddresses() string {
	return strings.Join(m.CCAddressList, ", ")
}

func (message Message) IsValid() (validity bool) {
	if message.ToAddress == "" ||
		message.FromAddress == "" ||
		message.Subject == "" ||
		message.Body == "" {
		return false
	}

	return true
}

func (m Message) URLValues() url.Values {
	values := make(url.Values)
	values.Set("to", m.ToAddress)
	values.Set("from", m.From())
	values.Set("subject", m.Subject)
	values.Set("text", m.Body)

	if m.CCAddresses() != "" {
		values.Set("cc", m.CCAddresses())
	}

	if m.BCCAddresses() != "" {
		values.Set("bcc", m.BCCAddresses())
	}

	for _, attachment := range m.AttachmentList {
		values.Add("attachment", attachment)
	}

	for _, inline := range m.InlineList {
		values.Add("inline", inline)
	}

	return values
}

func (message Message) GetRequest(mailgun Client) (request *http.Request) {
	request, _ = http.NewRequest("POST", mailgun.Endpoint(message), strings.NewReader(message.URLValues().Encode()))
	request.Header.Set("content-type", "application/x-www-form-urlencoded")
	return
}

func (message Message) Endpoint() string {
	return "messages"
}
