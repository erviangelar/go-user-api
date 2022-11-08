package helper

import (
	"github.com/erviangelar/go-user-api/common/config"
	"gopkg.in/gomail.v2"
)

type EmailMessage struct {
	To          string
	Subject     string
	MessageType string
	Message     string
	Attachment  string
}

func SendEmail(config *config.Configurations, message *EmailMessage) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", config.Smpt.Sender)
	msg.SetHeader("To", message.To)
	msg.SetHeader("Subject", message.Subject)
	msg.SetBody(message.MessageType, message.Message)
	msg.Attach(message.Attachment)

	n := gomail.NewDialer(config.Smpt.Host, config.Smpt.Port, config.Smpt.Username, config.Smpt.Password)

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}
	return nil
}
