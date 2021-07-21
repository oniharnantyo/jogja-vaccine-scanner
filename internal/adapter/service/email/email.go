package email

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"

	"github.com/oniharnantyo/jogja-vaccine-scanner/config"
)

type Email interface {
	Send(content string) error
}

type EmailService struct {
	Config *config.Config
}

func (s *EmailService) Send(content string) error {
	if !s.Config.SMTP.Enable {
		return nil
	}

	smtpConfig := s.Config.SMTP

	d := gomail.NewDialer(smtpConfig.Host, smtpConfig.Port, smtpConfig.Sender, smtpConfig.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", smtpConfig.Sender)
	m.SetHeader("To", s.Config.Emails...)
	m.SetHeader("Subject", s.Config.Subject)
	m.SetBody("text/html", content)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func (s *EmailService) generateMIMEContent(content string) []byte {
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + s.Config.Subject + "!\n"
	msg := []byte(subject + mime + "\n" + content)

	return msg
}

func NewEmailService(config *config.Config) Email {
	return &EmailService{Config: config}
}
