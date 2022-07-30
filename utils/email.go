package utils

import (
	"net/smtp"
)

// Sends an email using gmail SMTP server
func SendEmail(from, to, password, subject, body string) error {

	receivers := []string{to}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	message := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	return smtp.SendMail(smtpHost+":"+smtpPort,
		auth,
		from, receivers, []byte(message))
}
