package mailgun

import (
    "fmt"
    "net/http"
    "strings"
    "io/ioutil"
    "net/url"
)

var ApiEndpoint, ApiKey string

type Message struct {
    FromName string
    FromAddress string
    ToAddress string
    Subject string
    Body string
}

func (m Message) From() string {
    return fmt.Sprintf("%s <%s>", m.FromName, m.FromAddress)
}

func Send(message Message) error {
    client := &http.Client{}

    values := make(url.Values)
    values.Set("from", message.From())
    values.Set("to", message.ToAddress)
    values.Set("subject", message.Subject)
    values.Set("text", message.Body)

    request, _ := http.NewRequest("POST", ApiEndpoint, strings.NewReader(values.Encode()))
    request.Header.Set("content-type", "application/x-www-form-urlencoded")
    request.SetBasicAuth("api", ApiKey)

    response, e1 := client.Do(request)
    if e1 != nil {
        fmt.Println("Failed to send request")
        fmt.Println(e1)
        return e1
    }
    defer response.Body.Close()

    body, e2 := ioutil.ReadAll(response.Body)
    if e2 != nil {
        fmt.Println("Failed to read response")
        fmt.Println(e2)
        return e2
    }

    fmt.Println(string(body))
    return nil
}
