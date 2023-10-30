package templates

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"gopkg.in/gomail.v2"
)

func PrimaryTemplate() *template.Template {
	// Parse the email template from the HTML file.
	tmpl, err := template.ParseFiles("./templates/primary_template.html")
	if err != nil {
		fmt.Println("Error parsing email template:", err)
		os.Exit(1)
	}

	return tmpl
}

func UsePrimaryTemplate(receipt string, subject string, firstName string, lastName string) *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_USER"))
	m.SetHeader("To", receipt)
	m.SetHeader("Subject", subject)

	tmpl := PrimaryTemplate()

	bodyData := struct {
		FirstName string
		LastName  string
	}{
		FirstName: firstName,
		LastName:  lastName,
	}
	var bodyContent bytes.Buffer
	err := tmpl.Execute(&bodyContent, bodyData)
	if err != nil {
		fmt.Println("Error generating email body:", err)
		os.Exit(1)
	}
	m.SetBody("text/html", bodyContent.String())

	return m
}
