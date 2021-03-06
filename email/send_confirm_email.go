package email

import (
	"bytes"
	"encoding/json"
	"html/template"
)

// ConfirmEmailParams holds parameters for the template.
type ConfirmEmailParams struct {
	ReceipientEmail  string
	UserFirstName    string
	ConfirmationLink string
	EmailFromName    string
}

var (
	confirmSubject  = "Confirm your account."
	confirmTemplate = `Hi {{ .UserFirstName }},

Please confirm your account using the following link:
{{ .ConfirmationLink }}

Thanks,
{{ .EmailFromName }}
`
)

// Send generates the email from the parameters and is sent.
func (params *ConfirmEmailParams) Send() error {
	tmpl, err := template.New("ConfirmTemplate").Parse(confirmTemplate)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, params); err != nil {
		return err
	}

	emailContent := GetSender().WriteHTMLEmail(confirmSubject, buf.String())
	return GetSender().Send([]string{params.ReceipientEmail}, emailContent)
}

// AddToQueue adds the parameters to queue.
func (params *ConfirmEmailParams) AddToQueue() error {
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return err
	}
	return confirmEmailQueue.Produce(paramsJSON)
}

// ProcessConfirmEmailFromQueue processes the parameters from the queue.
func ProcessConfirmEmailFromQueue(paramsJSON []byte) error {
	params := ConfirmEmailParams{}
	err := json.Unmarshal(paramsJSON, &params)
	if err != nil {
		return err
	}
	return params.Send()
}
