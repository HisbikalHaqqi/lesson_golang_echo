package main

import (
	"gopkg.in/gomail.v2"
	"log"
)

/*
	SETTING PROPERTY EMAIL
*/
const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "PT Hisbikal Membangun Negeri <hisbikalhaqiqi36@gmail.com>"
const CONFIG_AUTH_EMAIL = "hisbikalhaqiqi36@gmail.com"
const CONFIG_AUTH_PASSWORD = "hisbikal_5321"

func main() {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", "hisbikalhaqiqi32@gmail.com", "hisbikalhaqiqi37@gmail.com")
	mailer.SetAddressHeader("Cc", "neminsurya@gmail.com", "Tra Lala La")
	mailer.SetHeader("Subject", "Undangan Interview")
	mailer.SetBody("text/html", "Hello, <b>have a nice day</b>")
	mailer.Attach("./node.jpg")

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}