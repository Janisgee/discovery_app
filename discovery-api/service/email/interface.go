package email

type EmailService interface {
	SendPasswordResetEmail(personalName string, emailTo string, resetLink string) error
}
