package sendto

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/AnhducNA/go-ecommerce/global"
	"go.uber.org/zap"
)

// Email configuration
// To use Gmail SMTP:
// 1. Enable 2-Step Verification in your Google Account
// 2. Generate an App Password:
//    - Go to your Google Account settings
//    - Navigate to Security
//    - Under "2-Step Verification", click on "App passwords"
//    - Select "Mail" and your device
//    - Copy the generated 16-character password
// 3. Set the following environment variables:
//    - SMTP_HOST (default: smtp.gmail.com)
//    - SMTP_PORT (default: 587)
//    - SMTP_USERNAME (your Gmail address)
//    - SMTP_PASSWORD (your App Password)

const (
	smtpHost     = "smtp.gmail.com"
	smtpPort     = "587"
	smtpUsername = "daiphatmanagement@gmail.com"
	smtpPassword = "sfat pkxn unpb ryeb"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress
	To      []string
	Subject string
	Body    string
}

func BuildMessage(mail Mail) string {
	// Format recipients list properly
	toHeader := mail.To[0]
	for i := 1; i < len(mail.To); i++ {
		toHeader += ", " + mail.To[i]
	}

	// Build message with proper MIME headers and formatting
	msg := fmt.Sprintf("From: %s <%s>\r\n", mail.From.Name, mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", toHeader)
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += "MIME-Version: 1.0\r\n"
	msg += "Content-Type: text/html; charset=UTF-8\r\n"
	msg += "\r\n" // Empty line separating headers from body
	msg += mail.Body

	return msg
}

func SendTextEmailOTP(to []string, from string, otp string) error {
	// Fix the Printf formatting
	fmt.Printf("SendTextEmailOTP config: %s %s\n", smtpUsername, smtpPassword)

	if smtpUsername == "" || smtpPassword == "" {
		return fmt.Errorf("SMTP_USERNAME and SMTP_PASSWORD environment variables must be set")
	}

	contentEmail := Mail{
		From:    EmailAddress{Address: from, Name: "Go CRM"},
		To:      to,
		Subject: "OTP Verification",
		Body:    fmt.Sprintf("Your OTP is %s. Please enter it to verify your account", otp),
	}

	messageMail := BuildMessage(contentEmail)
	fmt.Printf("messageMail:: %s\n", messageMail)
	// Set up authentication information.
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)
	err := smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth, from, to,
		[]byte(messageMail),
	)
	if err != nil {
		global.Logger.Error("SendTextEmailOTP error::", zap.Error(err))
		return err
	}
	return nil
}

func SendTemplateEmailOTP(
	to []string, from string, templateName string, dataTemplate map[string]interface{},
) error {
	htmlBody, err := getMailTemplate(templateName, dataTemplate)
	if err != nil {
		return err
	}
	return Send(to, from, htmlBody)
}

func getMailTemplate(templateName string, templateData map[string]interface{}) (string, error) {
	htmlTemplate := new(bytes.Buffer)
	t := template.Must(template.New(templateName).ParseFiles("templates-email/" + templateName))
	err := t.Execute(htmlTemplate, templateData)
	if err != nil {
		return "", err
	}
	return htmlTemplate.String(), nil
}

func Send(to []string, from string, htmlTemplate string) error {
	contentEmail := Mail{
		From:    EmailAddress{Address: from, Name: "Go CRM"},
		To:      to,
		Subject: "OTP Verification",
		Body:    fmt.Sprintf("Your OTP is %s. Please enter it to verify your account", htmlTemplate),
	}
	messageMail := BuildMessage(contentEmail)

	// send smtp
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)
	err := smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth, from, to,
		[]byte(messageMail),
	)
	if err != nil {
		global.Logger.Error("SendTemplateEmailOTP error::", zap.Error(err))
		return err
	}
	return nil
}
