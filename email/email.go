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

// Sender ...
type Sender struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
	FromAddress  string
	FromName     string
}

// Write ...
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

// WriteHTMLEmail ...
func (s *Sender) WriteHTMLEmail(subject, body string) string {
	return s.Write("text/html", subject, body)
}

// WritePlainEmail ...
func (s *Sender) WritePlainEmail(subject, body string) string {
	return s.Write("text/plain", subject, body)
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
