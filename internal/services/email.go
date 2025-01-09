package services

import (
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

func SendOTPEmail(email string, otp string) error {
	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASS"))

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_USER"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Your Verification Code")
	m.SetBody("text/html", fmt.Sprintf("Your verification code is <b>%s</b>. It will expire in 10 minutes.", otp))

	return d.DialAndSend(m)
}
