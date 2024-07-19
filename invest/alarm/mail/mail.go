package mail

import (
	"log"
	"net/smtp"
)

func SendEmail(server SMTP, subject, target, body string) {

	to := []string{target} // Ensure this is a valid email address

	// Combine headers and body
	message := []byte("From: " + server.User() + "\r\n" +
		"To: " + to[0] + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body)

	// Authentication
	auth := smtp.PlainAuth("", server.User(), server.Password(), server.Server())

	// Send email
	err := smtp.SendMail(server.Server()+":"+server.Port(), auth, server.User(), to, message)
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}
	log.Println("Email sent successfully!")
}

type SMTP interface {
	Server() string
	Port() string
	User() string
	Password() string
}
