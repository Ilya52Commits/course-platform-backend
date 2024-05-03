package scripts

import (
	"net/smtp"
	"strconv"
)

// SendMail /* Функция для отправки письма на почту */
func SendMail(mailAddressTo string, code int) error {
	// Настройка почты отправителя
	auth := smtp.PlainAuth("", "krasnenkov.ilia@gmail.com", "ymvv xxjf sjer picd", "smtp.gmail.com")

	to := []string{mailAddressTo} // Присвоение почты получателя

	// Составление письма
	msg := []byte("To: " + mailAddressTo + "\r\n" +
		"Subject: Подтверждение почты\r\n" +
		"\r\n" +
		"Введите данный код: " + strconv.Itoa(code) + "\r\n")

	// Отправка сообщения и присвоение ошибки, если таковая возникнет
	if err := smtp.SendMail("smtp.gmail.com:587", auth, "krasnenkov.ilia@gmail.com", to, msg); err != nil {
		return err
	}

	// Возвращаем ноль
	return nil
}
