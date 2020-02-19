package email

import (
	"github.com/mjah/jwt-auth/queue"
	"github.com/spf13/viper"
)

var (
	globalSender            *Sender
	welcomeEmailQueue       *queue.Queue
	confirmEmailQueue       *queue.Queue
	resetPasswordEmailQueue *queue.Queue
)

// Setup configures the sender, and declares the queues and consumers.
func Setup() error {
	SetSender(&Sender{
		SMTPHost:     viper.GetString("email.smtp_host"),
		SMTPPort:     viper.GetInt("email.smtp_port"),
		SMTPUsername: viper.GetString("email.smtp_username"),
		SMTPPassword: viper.GetString("email.smtp_password"),
		FromAddress:  viper.GetString("email.from_address"),
		FromName:     viper.GetString("email.from_name"),
	})

	if err := declareQueues(); err != nil {
		return err
	}

	if err := addConsumers(); err != nil {
		return err
	}

	return nil
}

// SetSender sets the sender.
func SetSender(sender *Sender) {
	globalSender = sender
}

// GetSender returns a pointer to the sender.
func GetSender() *Sender {
	return globalSender
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
