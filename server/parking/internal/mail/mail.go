package mail

import (
	"log"
	"net/smtp"
)

const (
	appError      = 800
	freeSpaceBusy = 801
	onGrass       = 802
)

func Send(email string, errCode int, body string) error {
	from := email
	pass := "wbahjdjqghjhsd"
	to := "romaactapov@gmail.com"
	var subject string

	if errCode == appError {
		subject = "Ошибка в приложении!"
	} else if errCode == freeSpaceBusy {
		subject = "Свободное место занято!"
	} else {
		subject = "Парковка в неположенном месте!"
	}

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail("smtp.mail.ru:25",
		smtp.PlainAuth("", from, pass, "smtp.mail.ru"),
		from, []string{to}, []byte(msg))
	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}

	log.Printf("successful send email to %s\n", email)
	return nil
}
