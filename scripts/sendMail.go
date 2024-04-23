package scripts

import (
	"net/smtp"
	"strconv"
)

func SendMail(mailAddresTo string, code int) error {
	auth := smtp.PlainAuth("", "krasnenkov.ilia@gmail.com", "ymvv xxjf sjer picd", "smtp.gmail.com")

	to := []string{mailAddresTo}

	msg := []byte("To: " + mailAddresTo + "\r\n" +
		"Subject: Подтверждение почты\r\n" +
		"\r\n" +
		"Введите данный код: " + strconv.Itoa(code) + "\r\n")

	if err := smtp.SendMail("smtp.gmail.com:587", auth, "krasnenkov.ilia@gmail.com", to, msg); err != nil {
		return err
	}

	return nil
}
