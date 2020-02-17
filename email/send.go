package email

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/mjah/jwt-auth/logger"
	"github.com/spf13/viper"
)

// WelcomeEmailDetails ...
type WelcomeEmailDetails struct {
	ReceipientEmail string
	UserFirstName   string
	EmailFromName   string
}

// ConfirmEmailDetails ...
type ConfirmEmailDetails struct {
	ReceipientEmail  string
	UserFirstName    string
	ConfirmationLink string
	EmailFromName    string
}

// ResetPasswordEmailDetails ...
type ResetPasswordEmailDetails struct {
	ReceipientEmail   string
	UserFirstName     string
	ResetPasswordLink string
	EmailFromName     string
}

// Send ...
func (s *Sender) Send(dest []string, content string) error {
	header := make(map[string]string)
	header["From"] = s.FromAddress
	header["To"] = strings.Join(dest, ",")

	message := ""

	for key, value := range header {
		message += fmt.Sprintf("%s: %s\n", key, value)
	}

	message += content

	logger.Log().Info(message + "\n")

	err := smtp.SendMail(s.SMTPHost+":"+strconv.Itoa(s.SMTPPort),
		smtp.PlainAuth("", s.SMTPUsername, s.SMTPPassword, s.SMTPHost),
		s.SMTPUsername, dest, []byte(message))
	if err != nil {
		return err
	}

	return nil
}

// SendTestEmail ...
func SendTestEmail() {
	sender := &Sender{
		SMTPHost:     viper.GetString("email.smtp_host"),
		SMTPPort:     viper.GetInt("email.smtp_port"),
		SMTPUsername: viper.GetString("email.smtp_username"),
		SMTPPassword: viper.GetString("email.smtp_password"),
		FromAddress:  viper.GetString("email.from_address"),
	}

	emailContent := sender.WriteHTMLEmail("Test Email", "This is a test email.")

	receipient := viper.GetString("email.test_receipient")
	err := sender.Send([]string{receipient}, emailContent)
	if err != nil {
		logger.Log().Error("SMTP Error: %s", err)
	} else {
		logger.Log().Info("Test email successfully sent.")
	}
}

// Send ...
func (details *WelcomeEmailDetails) Send() error {
	buf := new(bytes.Buffer)
	err := welcomeTemplate.Execute(buf, details)
	if err != nil {
		return err
	}

	emailContent := globalSender.WriteHTMLEmail("Welcome!", buf.String())

	err = globalSender.Send([]string{details.ReceipientEmail}, emailContent)
	if err != nil {
		return err
	}

	return nil
}

// Send ...
func (details *ConfirmEmailDetails) Send() error {
	buf := new(bytes.Buffer)
	err := confirmTemplate.Execute(buf, details)
	if err != nil {
		return err
	}

	emailContent := globalSender.WriteHTMLEmail("Confirm your account.", buf.String())

	err = globalSender.Send([]string{details.ReceipientEmail}, emailContent)
	if err != nil {
		return err
	}

	return nil
}

// Send ...
func (details *ResetPasswordEmailDetails) Send() error {
	buf := new(bytes.Buffer)
	err := resetPasswordTemplate.Execute(buf, details)
	if err != nil {
		return err
	}

	emailContent := globalSender.WriteHTMLEmail("Reset your password.", buf.String())

	err = globalSender.Send([]string{details.ReceipientEmail}, emailContent)
	if err != nil {
		return err
	}

	return nil
}
