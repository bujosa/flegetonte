package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func main() {
	// Load the environment variables from the .env file.
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		os.Exit(1)
	}

	var firstName, lastName, email, toEmail, subject string
	fmt.Print("Enter your first name: ")
	fmt.Scanln(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scanln(&lastName)
	fmt.Print("Enter your email: ")
	fmt.Scanln(&email)
	fmt.Print("Enter the email of the recipient: ")
	fmt.Scanln(&toEmail)
	fmt.Print("Enter the subject of the email: ")
	fmt.Scanln(&subject)

	bodyTemplate := `
        <html>
            <body>
                <p>My Name is {{.FirstName}} {{.LastName}},</p>
                <p>This is a test email.</p>
            </body>
        </html>
    `
	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", subject)

	tmpl, err := template.New("email").Parse(bodyTemplate)
	if err != nil {
		fmt.Println("Error parsing email template:", err)
		os.Exit(1)
	}
	bodyData := struct {
		FirstName string
		LastName  string
	}{
		FirstName: firstName,
		LastName:  lastName,
	}
	var bodyContent bytes.Buffer
	err = tmpl.Execute(&bodyContent, bodyData)
	if err != nil {
		fmt.Println("Error generating email body:", err)
		os.Exit(1)
	}
	m.SetBody("text/html", bodyContent.String())

	d := gomail.NewDialer("smtp-relay.brevo.com", 587, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASSWORD"))

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	fmt.Println("Email sent successfully!")
}
