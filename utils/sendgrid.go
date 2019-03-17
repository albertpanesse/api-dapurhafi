package utils

import (
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(toName, toEmail, fromName, fromEmail, subject, mailCntPlain, mailCntHtml string) (status bool, code int, body string) {
	from := mail.NewEmail(fromName, fromEmail)
	to := mail.NewEmail(toName, toEmail)
	message := mail.NewSingleEmail(from, subject, to, mailCntPlain, mailCntHtml)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	response, err := client.Send(message)
	if err != nil {
		status = false
		code = -1
		body = ""
	} else {
		status = true
		code = response.StatusCode
		body = response.Body
	}
	
	return
}
