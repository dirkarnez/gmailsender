package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	email := ""

	// Sender data.
	from := email
	password := "" //gmail needs app password, not gmail account password

	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(fmt.Sprintf("To: %s\nSubject: Hi!\nThis is the email body.\n", email))

	// client, err := smtp.Dial(smtpHost + ":" + smtpPort)
	// client.Auth()

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, smtp.PlainAuth("", email, password, smtpHost), from, to, message)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
