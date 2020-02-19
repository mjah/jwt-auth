// Package email provides email sending functionality.
package email

import (
	"bytes"
	"fmt"
	"mime/quotedprintable"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/mjah/jwt-auth/logger"
	"github.com/spf13/viper"
)

// Sender holds the SMTP details.
type Sender struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
	FromAddress  string
	FromName     string
}

// Write returns the string of the composed email.
func (s *Sender) Write(contentType, subject, body string) string {
	header := make(map[string]string)
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", contentType)
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	message := ""

	for key, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	var encodedMessage bytes.Buffer

	finalMessage := quotedprintable.NewWriter(&encodedMessage)
	finalMessage.Write([]byte(body))
	finalMessage.Close()

	message += "\r\n" + encodedMessage.String()

	return message
}

// WriteHTMLEmail returns the HTML composed email.
func (s *Sender) WriteHTMLEmail(subject, body string) string {
	return s.Write("text/html", subject, body)
}

// WritePlainEmail returns the plain text composed email.
func (s *Sender) WritePlainEmail(subject, body string) string {
	return s.Write("text/plain", subject, body)
}

// Send adds the from and to headers and sends the email.
func (s *Sender) Send(dest []string, content string) error {
	header := make(map[string]string)
	header["From"] = s.FromAddress
	header["To"] = strings.Join(dest, ",")

	message := ""
	for key, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	message += content

	if viper.GetString("environment") != "production" && viper.GetBool("log_email") {
		logger.Log().Info(message + "\n")
	}

	return smtp.SendMail(
		s.SMTPHost+":"+strconv.Itoa(s.SMTPPort),
		smtp.PlainAuth("", s.SMTPUsername, s.SMTPPassword, s.SMTPHost),
		s.SMTPUsername, dest, []byte(message),
	)
}

// SendTestEmail provides a simple way to test SMTP configuration.
func SendTestEmail() {
	sender := &Sender{
		SMTPHost:     viper.GetString("email.smtp_host"),
		SMTPPort:     viper.GetInt("email.smtp_port"),
		SMTPUsername: viper.GetString("email.smtp_username"),
		SMTPPassword: viper.GetString("email.smtp_password"),
		FromAddress:  viper.GetString("email.from_address"),
	}

	receipient := viper.GetString("email.test_receipient")
	emailContent := sender.WriteHTMLEmail("Test Email", "This is a test email.")

	if err := sender.Send([]string{receipient}, emailContent); err != nil {
		logger.Log().Error("SMTP Error: %s", err)
	} else {
		logger.Log().Info("Test email successfully sent.")
	}
}
