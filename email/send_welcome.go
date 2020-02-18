package email

import (
	"bytes"
	"encoding/json"
	"html/template"
)

// WelcomeEmailParams ...
type WelcomeEmailParams struct {
	ReceipientEmail string
	UserFirstName   string
	EmailFromName   string
}

var (
	welcomeSubject  = "Welcome!"
	welcomeTemplate = `Welcome {{ .UserFirstName }},

Thanks for signing up with us. We hope you enjoy using our services.

Cheers,
{{ .EmailFromName }}
`
)

// Send ...
func (params *WelcomeEmailParams) Send() error {
	tmpl, err := template.New("WelcomeTemplate").Parse(welcomeTemplate)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, params); err != nil {
		return err
	}

	emailContent := GetSender().WriteHTMLEmail(welcomeSubject, buf.String())
	return GetSender().Send([]string{params.ReceipientEmail}, emailContent)
}

// AddToQueue ...
func (params *WelcomeEmailParams) AddToQueue() error {
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return err
	}
	return welcomeEmailQueue.Produce(paramsJSON)
}

// ProcessWelcomeEmailFromQueue ...
func ProcessWelcomeEmailFromQueue(paramsJSON []byte) error {
	params := WelcomeEmailParams{}
	err := json.Unmarshal(paramsJSON, &params)
	if err != nil {
		return err
	}
	return params.Send()
}
