package email

import (
	"os"

	"gopkg.in/gomail.v2"
)

func Send() error {
	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("EMAIL_PASSWORD")

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", os.Getenv("EMAIL_TO"))
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	d := gomail.NewDialer("smtp.qq.com", 25, from, password)
	return d.DialAndSend(m)
}
