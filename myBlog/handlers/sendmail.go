package handlers

import (
	"net/smtp"
)

type sendResult struct {
	status  bool
	message string
}

// пользователь, через которого мы авторизуемся
var from = "yub@crtweb.com"

//Пароль от своей почты
var password = "пароль свой"

// сервер, через который отправляем электронную почту
var host = "smtp.gmail.com:587"

/*Настройка smtp gmail.com:
Хост для smtp: smtp.gmail.com
Тип шифрования для smtp: tls
Порт для smtp: 587*/

func send(toSend string, subject string, message string) sendResult {

	result := sendResult{
		status:  true,
		message: "Сообщение успешно отправлено. Я отвечу вам в самое ближайшее время",
	}
	// Создаем аутентификацию для SendMail ()
	// используя PlainText, но другие методы аутентификации приветствуются
	auth := smtp.PlainAuth("", from, password, host)

	// ПРИМЕЧАНИЕ: использование обратного ключа здесь `работает как heredoc, поэтому все
	// остальные строки переносятся в начало строки, в противном случае
	// форматирование неверно для стиля RFC 822
	sendMessage := `To: User <` + toSend + `>
From: "Other User" <` + from + `>
Subject: ` + subject + `

` + message

	err := smtp.SendMail(host, auth, from, []string{toSend}, []byte(sendMessage))
	if err != nil {
		result.status = false
		result.message = "Не удалось отправить сообщение"
	}

	return result
}
