package email

import (
	"fmt"

	gomail "gopkg.in/mail.v2"
)

type gomailService struct {
	dialer    *gomail.Dialer
	emailFrom string
}

func NewEmailService(emailFrom string, dialer *gomail.Dialer) EmailService {
	return &gomailService{
		dialer,
		emailFrom,
	}
}

func (svc *gomailService) SendPasswordResetEmail(personalName string, emailTo string, resetLink string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", svc.emailFrom)
	m.SetHeader("To", emailTo)
	m.SetHeader("Subject", "Set or reset your password")

	// Format the body
	body := fmt.Sprintf(`Hello <b>%s</b>,<h3>We have received a request to set or reset your password. Please click below link to reset your password.</h3>%s<br><p>If you did not request to set a new password or change your password, you can disregard this email and your password will remain the same.</p><p>Should you experience any issues accessing your Discovery App account, please contact us immediately.</p>Sincerely,<br>Discovery App`, personalName, resetLink)

	m.SetBody("text/html", body)

	err := svc.dialer.DialAndSend(m)

	return err
}
