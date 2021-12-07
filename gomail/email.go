package gomail

import (
	"fmt"
	"net/mail"
	"net/smtp"
	"state/config/env"
)

// SendMail , sends an email to the provided address(es)
func SendMail(to string, subject string, body string, errCh chan error) {
	addr := env.GetString("mail.server")
	
	auth := smtp.CRAMMD5Auth(env.GetString("mail.auth.user"),
		env.GetString("mail.auth.password"))
	
	from := mail.Address{Name: env.GetString("mail.from.name"),
		Address: env.GetString("mail.from.address")}
	
	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to
	headers["Subject"] = subject
	
	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body
	err := smtp.SendMail(addr, auth, from.Address, []string{to}, []byte(message))
	if err != nil {
		errCh <- err
	}
	errCh <- nil
}
