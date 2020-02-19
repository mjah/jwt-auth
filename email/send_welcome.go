package email

import (
	"bytes"
	"encoding/json"
	"html/template"
)

// WelcomeEmailParams holds parameters for the template.
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

// Send generates the email from the parameters and is sent.
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

// AddToQueue adds the parameters to queue.
func (params *WelcomeEmailParams) AddToQueue() error {
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return err
	}
	return welcomeEmailQueue.Produce(paramsJSON)
}

// ProcessWelcomeEmailFromQueue processes the parameters from the queue.
func ProcessWelcomeEmailFromQueue(paramsJSON []byte) error {
	params := WelcomeEmailParams{}
	err := json.Unmarshal(paramsJSON, &params)
	if err != nil {
		return err
	}
	return params.Send()
}
