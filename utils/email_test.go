package utils

import (
	"testing"
	"text/template"
)

// Change the variables for the test to work
func TestSendEmail(t *testing.T) {
	from := "testfrom@gmail.com"
	to := "testto@gmail.com"
	password := "testpassword"
	subject := "Test subject"
	body := "Test body"
	err := SendEmail(from, to, password, subject, body)
	if err != nil {
		t.Errorf("SendEmail() error: %v", err)
	}
}

func TestSendEmailWithTemplate(t *testing.T) {
	from := "testfrom@gmail.com"
	to := "testto@gmail.com"
	password := "testpassword"
	subject := "Test subject"
	data := struct {
		Name string
	}{
		Name: "Test",
	}

	tpl, err := template.ParseFiles("/PATH/TO/template.html")

	if err != nil {
		t.Errorf("SendEmailWithTemplate() error: %v", err)
	}

	err = SendEmailWithTemplate(from, to, password, subject, tpl, data)
	if err != nil {
		t.Errorf("SendEmailWithTemplate() error: %v", err)
	}
}
