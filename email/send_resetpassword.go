package email

import (
	"bytes"
	"encoding/json"
	"html/template"
)

// ResetPasswordEmailParams ...
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

// Send ...
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

// AddToQueue ...
func (params *ResetPasswordEmailParams) AddToQueue() error {
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return err
	}
	return resetPasswordEmailQueue.Produce(paramsJSON)
}

// ProcessResetPasswordEmailFromQueue ...
func ProcessResetPasswordEmailFromQueue(paramsJSON []byte) error {
	params := ResetPasswordEmailParams{}
	err := json.Unmarshal(paramsJSON, &params)
	if err != nil {
		return err
	}
	return params.Send()
}
