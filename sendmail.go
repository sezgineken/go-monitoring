// FOR LİNUX

/*
package main

import (

	"fmt"
	"os/exec"
	"strings"

)

	func sendMail(subject, body string) {
		cmd := exec.Command(
			"msmtp",
			"--from=default",
			"-t",
		)

		message := fmt.Sprintf("To: ensmartit@gmail.com\nSubject: %s\n\n%s", subject, body)
		cmd.Stdin = strings.NewReader(message)

		err := cmd.Run()
		if err != nil {
			fmt.Println("❌ Mail gönderilemedi:", err)
			return
		}

		fmt.Println("📧 Uyarı e-postası gönderildi.")
	}
*/
package main

import (
	"fmt"
	"net/smtp"
)

func sendMail(subject, body string) {
	from := "ensmartit@gmail.com"
	password := "amogitqmfaspisuy"
	to := "ensmartit@gmail.com"

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
