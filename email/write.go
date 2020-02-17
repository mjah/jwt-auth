package email

import (
	"bytes"
	"fmt"
	"mime/quotedprintable"
)

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
