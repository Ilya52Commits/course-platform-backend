package scripts

import (
	"net/smtp"
	"strconv"
)

/* Функция для отправки письма на почту */
func SendMail(mailAddresTo string, code int) error {
	// Настройка почты отправителя
	auth := smtp.PlainAuth("", "krasnenkov.ilia@gmail.com", "ymvv xxjf sjer picd", "smtp.gmail.com")

	to := []string{mailAddresTo} // Присвоение почты получателя

	// Составление письма
	msg := []byte("To: " + mailAddresTo + "\r\n" +
		"Subject: Подтверждение почты\r\n" +
		"\r\n" +
		"Введите данный код: " + strconv.Itoa(code) + "\r\n")

	// Отправка сообщения и присвоение ошибки, если таковая возникнит
	if err := smtp.SendMail("smtp.gmail.com:587", auth, "krasnenkov.ilia@gmail.com", to, msg); err != nil {
		return err
	}

	// Возвращаем ноль
	return nil
}
