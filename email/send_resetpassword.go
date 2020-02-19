package email

import (
	"bytes"
	"encoding/json"
	"html/template"
)

// ResetPasswordEmailParams holds parameters for the template.
type ResetPasswordEmailParams struct {
	ReceipientEmail   string
	UserFirstName     string
	ResetPasswordLink string
	EmailFromName     string
}

var (
	resetPasswordSubject  = "Reset your password."
	resetPasswordTemplate = `Hi {{ .UserFirstName }},

We've received a request to reset your password. If you didn't make this request, just ignore this email. Otherwise, you can reset your password using this link:
{{ .ResetPasswordLink }}

Thanks,
{{ .EmailFromName }}
`
)

// Send generates the email from the parameters and is sent.
func (params *ResetPasswordEmailParams) Send() error {
	tmpl, err := template.New("ResetPasswordTemplate").Parse(resetPasswordTemplate)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, params); err != nil {
		return err
	}

	emailContent := GetSender().WriteHTMLEmail(resetPasswordSubject, buf.String())
	return GetSender().Send([]string{params.ReceipientEmail}, emailContent)
}

// AddToQueue adds the parameters to queue.
func (params *ResetPasswordEmailParams) AddToQueue() error {
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return err
	}
	return resetPasswordEmailQueue.Produce(paramsJSON)
}

// ProcessResetPasswordEmailFromQueue processes the parameters from the queue.
func ProcessResetPasswordEmailFromQueue(paramsJSON []byte) error {
	params := ResetPasswordEmailParams{}
	err := json.Unmarshal(paramsJSON, &params)
	if err != nil {
		return err
	}
	return params.Send()
}
