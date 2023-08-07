package email

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"gopkg.in/gomail.v2"
)

type User struct {
	Email string `json:"email"`
}

func Send(html string) {
	users := getToUser()

	wg := sync.WaitGroup{}

	for _, user := range users {
		wg.Add(1)
		go func(user User) {
			defer wg.Done()
			sendEmail(html, user.Email)
		}(user)
	}
	wg.Wait()
}

func getToUser() []User {
	var users []User
	toJson := os.Getenv("EMAIL_TO")

	err := json.Unmarshal([]byte(toJson), &users)
	if err != nil {
		log.Fatalf("Parse users from %s error: %s", toJson, err)
	}
	return users
}

func sendEmail(content string, to string) {
	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("EMAIL_PASSWORD")

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "每日一读")
	m.SetBody("text/html", content)
	d := gomail.NewDialer("smtp.qq.com", 25, from, password)

	err := d.DialAndSend(m)
	if err != nil {
		log.Printf("Send email fail, error: %s", err)
	} else {
		log.Printf("Send email %s success", to)
	}

}
