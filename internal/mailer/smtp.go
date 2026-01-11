package mailer

import (
	"DSCMailer/internal/utils"
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

func SendSMTP(to string, subject string, body string, attachmentPath string) error {

	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	senderName := os.Getenv("SMTP_SENDER_NAME")

	auth := smtp.PlainAuth(
		"",
		user,
		pass,
		host,
	)

	msg, err := utils.BuildMessage(
		user,
		senderName,
		to,
		subject,
		body,
		attachmentPath,
	)
	if err != nil {
		return err
	}

	addr := fmt.Sprintf("%s:%s", host, port)

	return smtp.SendMail(
		addr,
		auth,
		user,
		[]string{strings.TrimSpace(to)},
		msg,
	)
}
