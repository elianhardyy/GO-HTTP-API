package utils

import (
	"fmt"
	"net/smtp"
)

func SendVerificationEmail(toEmail, token string) error {
	from := "elian.hardiawan2001@gmail.com"
	password := "qpvuzzwqxzcwejhr"
	smtpServer := "smtp.gmail.com"
	port := 587

	subject := "Email Verification"
	//verificationLink := fmt.Sprintf("http://localhost:7000/verify-email?token=%s",token)
	body := fmt.Sprintf("This is your token : %s",token)

	auth := smtp.PlainAuth("",from,password,smtpServer)
	to := []string{toEmail}
	msg := []byte("To: "+toEmail+"\r\n"+
		"Subject: "+subject+"\r\n"+
		"\r\n"+
		body+"\r\n")
	err := smtp.SendMail(fmt.Sprintf("%s:%d",smtpServer,port),auth,from,to,msg)
	if err != nil {
		return err
	}
	return nil
	
}