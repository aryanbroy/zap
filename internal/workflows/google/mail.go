package google

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"google.golang.org/api/gmail/v1"
)

func SendMail(token string) error {
	log.Println("Sending mail...")
	url := "https://gmail.googleapis.com/gmail/v1/users/me/messages/send"

	data := makeBody(
		"aryanbroy003@gmail.com",
		"aryanbroy003@gmail.com",
		"Test message",
		"Testing a message here",
	)

	payload := map[string]string{"raw": data.Raw}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln("Error marshaling json data")
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		log.Fatalln("Error making post request to mail server")
		return err
	}

	log.Println("Setting token...")
	fmt.Println("Token: ", token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln("Error getting response")
		return err
	}

	defer res.Body.Close()

	bodyData, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Error reading body")
		return err
	}

	fmt.Println(string(bodyData))

	log.Println("Mail sent successfuly!")
	return nil
}

func makeBody(to string, from string, subject string, message string) gmail.Message {

	var msg gmail.Message

	messageStr := []byte(fmt.Sprintf(
		"From: %v\r\n"+
			"To: %v\r\n"+
			"Subject: %v\r\n\r\n"+
			"%v", from, to, subject, message))

	msg.Raw = base64.URLEncoding.EncodeToString(messageStr)
	return msg
}
