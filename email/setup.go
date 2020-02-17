package email

import (
	"fmt"
	"html/template"
	"os"

	"github.com/mjah/jwt-auth/queue"
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

var (
	globalSender            *Sender
	welcomeEmailQueue       *queue.Queue
	confirmEmailQueue       *queue.Queue
	resetPasswordEmailQueue *queue.Queue
	welcomeTemplate         *template.Template
	confirmTemplate         *template.Template
	resetPasswordTemplate   *template.Template
)

// Setup ...
func Setup() error {
	// Set sender
	SetSender(&Sender{
		SMTPHost:     viper.GetString("email.smtp_host"),
		SMTPPort:     viper.GetInt("email.smtp_port"),
		SMTPUsername: viper.GetString("email.smtp_username"),
		SMTPPassword: viper.GetString("email.smtp_password"),
		FromAddress:  viper.GetString("email.from_address"),
		FromName:     viper.GetString("email.from_name"),
	})

	// Load templates
	if err := loadTemplates(); err != nil {
		return err
	}

	if err := declareQueues(); err != nil {
		return err
	}

	if err := addConsumers(); err != nil {
		return err
	}

	return nil
}

// SetSender ...
func SetSender(sender *Sender) {
	globalSender = sender
}

// GetSender ...
func GetSender() *Sender {
	return globalSender
}

func loadTemplates() error {
	// Get working directory
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	// Parse files
	welcomeTemplate, err = template.ParseFiles(fmt.Sprintf("%s/email/templates/welcome.html", dir))
	if err != nil {
		return err
	}

	confirmTemplate, err = template.ParseFiles(fmt.Sprintf("%s/email/templates/confirm.html", dir))
	if err != nil {
		return err
	}

	resetPasswordTemplate, err = template.ParseFiles(fmt.Sprintf("%s/email/templates/resetpassword.html", dir))
	if err != nil {
		return err
	}

	return nil
}

func declareQueues() error {
	q, err := queue.New("email_welcome_queue")
	if err != nil {
		return err
	}
	welcomeEmailQueue = q

	q, err = queue.New("email_confirm_queue")
	if err != nil {
		return err
	}
	confirmEmailQueue = q

	q, err = queue.New("email_resetpassword_queue")
	if err != nil {
		return err
	}
	resetPasswordEmailQueue = q

	return nil
}

func addConsumers() error {
	if err := welcomeEmailQueue.Consume(ProcessWelcomeEmailFromQueue); err != nil {
		return err
	}

	if err := confirmEmailQueue.Consume(ProcessConfirmEmailFromQueue); err != nil {
		return err
	}

	if err := resetPasswordEmailQueue.Consume(ProcessResetPasswordEmailFromQueue); err != nil {
		return err
	}

	return nil
}
