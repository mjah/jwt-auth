package email

import (
	"encoding/json"
)

// AddToQueue ...
func (details *WelcomeEmailDetails) AddToQueue() error {
	detailsJSON, err := json.Marshal(details)
	if err != nil {
		return err
	}
	return welcomeEmailQueue.Produce(detailsJSON)
}

// AddToQueue ...
func (details *ConfirmEmailDetails) AddToQueue() error {
	detailsJSON, err := json.Marshal(details)
	if err != nil {
		return err
	}
	return confirmEmailQueue.Produce(detailsJSON)
}

// AddToQueue ...
func (details *ResetPasswordEmailDetails) AddToQueue() error {
	detailsJSON, err := json.Marshal(details)
	if err != nil {
		return err
	}
	return resetPasswordEmailQueue.Produce(detailsJSON)
}

// ProcessWelcomeEmailFromQueue ...
func ProcessWelcomeEmailFromQueue(detailsJSON []byte) error {
	details := WelcomeEmailDetails{}
	err := json.Unmarshal(detailsJSON, &details)
	if err != nil {
		return err
	}
	return details.Send()
}

// ProcessConfirmEmailFromQueue ...
func ProcessConfirmEmailFromQueue(detailsJSON []byte) error {
	details := ConfirmEmailDetails{}
	err := json.Unmarshal(detailsJSON, &details)
	if err != nil {
		return err
	}
	return details.Send()
}

// ProcessResetPasswordEmailFromQueue ...
func ProcessResetPasswordEmailFromQueue(detailsJSON []byte) error {
	details := ResetPasswordEmailDetails{}
	err := json.Unmarshal(detailsJSON, &details)
	if err != nil {
		return err
	}
	return details.Send()
}
