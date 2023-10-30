package email

import (
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

func SendEmail(m *gomail.Message) {
	d := gomail.NewDialer("smtp-relay.brevo.com", 587, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	fmt.Println("Email sent successfully!")
}
