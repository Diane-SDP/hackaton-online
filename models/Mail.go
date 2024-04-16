package models

import (
	"fmt"
	"net/smtp"
)

func SendMail(tosend string, messagetosend string,subject string) {

	// Sender data.
	from := "tracktheure@gmail.com"
	// password := "TrackTheur_12"
	password := "qfcw xnhv rbwo nkip"

	// Receiver email address.
	to := []string{
		tosend,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("Subject: "+subject+"\r\n" +
	"\r\n" +
	messagetosend+"\r\n")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
