package main
import (
	"fmt"
    "net/http"
    "io/ioutil"
    "log"
    "encoding/json"
    "strings"
    //"os"
    "time"
    "bytes"
    "net/url"
)
var accessToken = "EAAE0VZC5ZBvm4BABdcp1eZBS3zhUG6HJ7E0DZBJIlCVgpJSzWmi6LMcyxWznuwXvoTh8Uw8rWgLuda8E3S2aMPpEMi5ZAZCJzuA82wBFNYq2iWyNrA95G37QPXoZBZBZBQ0RAAZAok8x39ZCUUQJZCIAj0Ri66RZCMbvvt8aRSdy1xhBvYwZDZD"
var verifyToken = "ZXaSWb29a223c1a0Gbc354YHFda9c525e96Lda3f6f4"

const FacebookEndPoint = "https://graph.facebook.com/v2.10/me/messages"

type ReceivedMessage struct {
    Object string  `json:"object"`
    Entry  []Entry `json:"entry"`
}

type Entry struct {
    // ID        int64       `json:"id"`
    // Time      int64       `json:"time"`
    Messaging []Messaging `json:"messaging"`
}

type Messaging struct {
    Sender    Sender    `json:"sender"`
    Recipient Recipient `json:"recipient"`
    // Timestamp int64     `json:"timestamp"`
    Message Message `json:"message"`
}

type Sender struct {
    ID string `json:"id"`
}

type Recipient struct {
    ID string `json:"id"`
}

type Message struct {
    MID  string `json:"mid"`
    Seq  int64  `json:"seq"`
    Text string `json:"text"`
}

type Payload struct {
    TemplateType string  `json:"template_type"`
    Text         string  `json:"text"`
    Buttons      Buttons `json:"buttons"`
}

type Buttons struct {
    Type  string `json:"type"`
    Url   string `json:"url"`
    Title string `json:"title"`
}

type Attachment struct {
    Type    string  `json:"type"`
    Payload Payload `json:"payload"`
}

type ButtonMessageBody struct {
    Attachment Attachment `json:"attachment"`
}

type ButtonMessage struct {
    Recipient         Recipient         `json:"recipient"`
    ButtonMessageBody ButtonMessageBody `json:"message"`
}

type SendMessage struct {
    Recipient Recipient `json:"recipient"`
    Message   struct {
        Text string `json:"text"`
    } `json:"message"`
}
func main() {
    http.HandleFunc("/webhook", webhookHandler)
    http.ListenAndServe(":8085", nil)
    fmt.Println("Success!!!")
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        verifyTokenAction(w, r)
    }
    if r.Method == "POST" {
        webhookPostAction(w, r)
    }
}

func verifyTokenAction(w http.ResponseWriter, r *http.Request) {
    if r.URL.Query().Get("hub.verify_token") == verifyToken {
        log.Print("verify token success.")
        fmt.Fprintf(w, r.URL.Query().Get("hub.challenge"))
    } else {
        log.Print("Error: verify token failed.")
        fmt.Fprintf(w, "Error, wrong validation token")
    }
}

func webhookPostAction(w http.ResponseWriter, r *http.Request) {
    var receivedMessage ReceivedMessage
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Print(err)
    }
    if err = json.Unmarshal(body, &receivedMessage); err != nil {
        log.Print(err)
    }
    messagingEvents := receivedMessage.Entry[0].Messaging
    for _, event := range messagingEvents {
        senderID := event.Sender.ID
        if &event.Message != nil && event.Message.Text != "" {
            message := getReplyMessage(event.Message.Text)
            sendTextMessage(senderID, message)
        }
    }
    fmt.Fprintf(w, "Success")
}
func getReplyMessage(receivedMessage string) string {
    var message string
    receivedMessage = strings.ToUpper(receivedMessage)
    log.Print(" Received message: " + receivedMessage)

    if strings.Contains(receivedMessage, "HELLO") {
        message = "Hi, my name is Annie. Nice to meet you"
    }

    return message
}

func sendTextMessage(senderID string, text string) {
    recipient := new(Recipient)
    recipient.ID = senderID
    sendMessage := new(SendMessage)
    sendMessage.Recipient = *recipient
    sendMessage.Message.Text = text
    sendMessageBody, err := json.Marshal(sendMessage)
    if err != nil {
        log.Print(err)
    }
    req, err := http.NewRequest("POST", FacebookEndPoint, bytes.NewBuffer(sendMessageBody))
    if err != nil {
        log.Print(err)
    }
    fmt.Println("%T", req)
    fmt.Println("%T", err)

    values := url.Values{}
    values.Add("access_token", accessToken)
    req.URL.RawQuery = values.Encode()
    req.Header.Add("Content-Type", "application/json; charset=UTF-8")
    client := &http.Client{Timeout: time.Duration(30 * time.Second)}
    res, err := client.Do(req)
    if err != nil {
        log.Print(err)
    }
    defer res.Body.Close()
    var result map[string]interface{}
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Print(err)
    }
    if err := json.Unmarshal(body, &result); err != nil {
        log.Print(err)
    }
    log.Print(result)
}