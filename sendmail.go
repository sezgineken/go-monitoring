package main

import (
	"fmt"
	"net/smtp"
)

func sendMail(subject, body string) {
	from := "youremail@example.com"
	password := "your-app-password"
	to := "youremail@example.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" + body

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		smtp.PlainAuth("", from, password, "smtp.gmail.com"),
		from, []string{to}, []byte(msg),
	)

	if err != nil {
		fmt.Println("❌ Mail gönderilemedi:", err)
		return
	}
	fmt.Println("✅ Mail gönderildi!")
}
