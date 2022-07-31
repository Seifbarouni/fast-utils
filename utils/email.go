package utils

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"
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

// Sends an email with a given template
func SendEmailWithTemplate(from, to, password, subject string, tpl *template.Template, data any) error {
	receivers := []string{to}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", subject, mimeHeaders)))

	err := tpl.Execute(&body, data)
	if err != nil {
		return err
	}

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, receivers, body.Bytes())

	return err
}
