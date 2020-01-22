package email

import (
	"bytes"
	"fmt"
	"mime/quotedprintable"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

// Sender ...
type Sender struct {
	SMTPHost string
	SMTPPort int
	Username string
	Password string
}

// Setup ...
func Setup(smtpHost string, smtpPort int, username, password string) *Sender {
	return &Sender{
		SMTPHost: smtpHost,
		SMTPPort: smtpPort,
		Username: username,
		Password: password,
	}
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
	header["From"] = s.Username
	header["To"] = strings.Join(dest, ",")

	message := ""

	for key, value := range header {
		message += fmt.Sprintf("%s: %s\n", key, value)
	}

	message += content

	err := smtp.SendMail(s.SMTPHost+":"+strconv.Itoa(s.SMTPPort),
		smtp.PlainAuth("", s.Username, s.Password, s.SMTPHost),
		s.Username, dest, []byte(message))
	if err != nil {
		return err
	}

	return nil
}

// Test ...
func Test() {
	smtpHost := viper.GetString("email.smtp_host")
	smtpPort := viper.GetInt("email.smtp_port")
	username := viper.GetString("email.username")
	password := viper.GetString("email.password")
	receipient := viper.GetString("email.test_receipient")

	testEmail := Setup(smtpHost, smtpPort, username, password)

	emailContent := testEmail.WriteHTMLEmail("Test Email", "This is a test email.")

	err := testEmail.Send([]string{receipient}, emailContent)
	if err != nil {
		fmt.Printf("SMTP Error: %s\n", err)
	} else {
		fmt.Println("Email successfully sent.")
	}
}
