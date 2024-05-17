package service

import (
	"fmt"
	"os"

	"net/smtp"
)

func getEmailPassword() string {
	if os.Getenv("EMAIL_SECRET_KEY") == "" {
		return "Empty"
	}
	return os.Getenv("EMAIL_SECRET_KEY")
}

var PASSWORD = os.Getenv("EMAIL_SECRET_KEY")

func SetupEmail(body string,to string)  {

	
	from := "andremendes0113@gmail.com"
	
	password := getEmailPassword()  // Sua senha do Elastic Email
	
	subject := "Login Notification"
	

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := SendMail(from, password, "smtp.elasticemail.com", 2525, []string{to}, []byte(msg))
	if err != nil {
		fmt.Println("Erro ao enviar o email:", err)
	} else {
		fmt.Println("Email enviado com sucesso!")


	}

}


func SendMail(from, password, host string, port int, to []string, message []byte) error {
	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(fmt.Sprintf("%s:%d", host, port), auth, from, to, message)
	if err != nil {
		return err
	}

	return nil
}
