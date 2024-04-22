package scripts

import (
	"log"
	"net/smtp"
)

func SendMail(mailAddresTo string) {
	auth := smtp.PlainAuth("", "krasnenkov.ilia@gmail.com", "ymvv xxjf sjer picd", "smtp.gmail.com")

	to := []string{mailAddresTo}

	msg := []byte("To: " + mailAddresTo + "\r\n" +
		"Subject: Подтверждение почты\r\n" +
		"\r\n" +
		"Введите данный код: \r\n")

	err := smtp.SendMail("smtp.gmail.com:587", auth, "krasnenkov.ilia@gmail.com", to, msg)

	if err != nil {
		log.Fatal(err)
	}
}
