package email

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/smtp"

	"github.com/Capture-411/services-opportunity/config"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SMTPLoginData struct {
	username, password string
}

func SMTPLogin(username, password string) smtp.Auth {
	return &SMTPLoginData{username, password}
}

func (data *SMTPLoginData) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (data *SMTPLoginData) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(data.username), nil
		case "Password:":
			return []byte(data.password), nil
		default:
			return nil, errors.New("smtp: unknown from server")
		}
	}
	return nil, nil
}

func NewMailer() (*smtp.Client, error) {
	host := config.AppConfig.SMTPHost
	port := config.AppConfig.SMTPPort
	username := config.AppConfig.SMTPUsername
	password := config.AppConfig.SMTPPassword

	// SMTP authentication
	// auth := smtp.PlainAuth("", "apikey", "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", host)

	addr := fmt.Sprintf("%s:%d", host, port)

	conn, err := smtp.Dial(addr)
	if err != nil {
		log.Fatalf("Could not dial to SMTP server: %v", err)
	}

	// Upgrade the connection to TLS using START TLS
	tlsConfig := &tls.Config{
		ServerName:         host,
		InsecureSkipVerify: false, // set to true if you don't want to verify TLS certificate
	}
	if err := conn.StartTLS(tlsConfig); err != nil {
		log.Fatalf("Could not start TLS: %v", err)
	}

	//auth := smtp.PlainAuth("", username, password, host)
	if err := conn.Auth(SMTPLogin(username, password)); err != nil {
		return nil, fmt.Errorf("could not authenticate: %v", err)
	}

	return conn, nil
}

func Send(to, subject, plainTextMessage, htmlMessage string) error {
	fromEmail := config.AppConfig.SMTPFrom
	apiKey := config.AppConfig.SMTPPassword // assuming SMTPPassword holds the API key
	client := sendgrid.NewSendClient(apiKey)

	// Set up email
	from := mail.NewEmail("", fromEmail)
	toEmail := mail.NewEmail("", to)
	message := mail.NewSingleEmail(from, subject, toEmail, plainTextMessage, htmlMessage)

	// Send email
	response, err := client.Send(message)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return fmt.Errorf("could not send email: %v", err)
	}

	// Handle response for logging
	if response.StatusCode >= 400 {
		log.Printf("SendGrid returned status code %d with body: %s", response.StatusCode, response.Body)
		return errors.New("failed to send email with SendGrid API")
	}

	log.Printf("Email sent successfully with status code %d", response.StatusCode)
	return nil
}

// func Send(to, subject, plainTextMessage, htmlMessage string) error {
// 	client, err := NewMailer()
// 	if err != nil {
// 		fmt.Printf("Could not connect to email server: %v", err)
// 		return fmt.Errorf("could not connect to email server: %v", err)
// 	}

// 	from := config.AppConfig.SMTPFrom

// 	var buf bytes.Buffer

// 	// Set the headers
// 	headers := make(map[string]string)
// 	headers["To"] = to
// 	headers["From"] = from
// 	headers["Subject"] = subject
// 	headers["MIME-Version"] = "1.0"
// 	boundary := "myboundary" // You can generate a unique boundary
// 	headers["Content-Type"] = "multipart/alternative; boundary=" + boundary

// 	// Write the headers to the buffer
// 	for k, v := range headers {
// 		fmt.Fprintf(&buf, "%s: %s\r\n", k, v)
// 	}
// 	fmt.Fprintf(&buf, "\r\n")

// 	// Start the multipart section
// 	writer := multipart.NewWriter(&buf)
// 	writer.SetBoundary(boundary)

// 	// Plain text part
// 	textPartHeader := make(textproto.MIMEHeader)
// 	textPartHeader.Set("Content-Type", "text/plain; charset=\"utf-8\"")
// 	textPartHeader.Set("Content-Transfer-Encoding", "7bit")
// 	textPart, err := writer.CreatePart(textPartHeader)
// 	if err != nil {
// 		return fmt.Errorf("could not create plain text part: %v", err)
// 	}
// 	_, err = textPart.Write([]byte(plainTextMessage))
// 	if err != nil {
// 		return fmt.Errorf("could not write plain text message: %v", err)
// 	}

// 	// HTML part
// 	htmlPartHeader := make(textproto.MIMEHeader)
// 	htmlPartHeader.Set("Content-Type", "text/html; charset=\"utf-8\"")
// 	htmlPartHeader.Set("Content-Transfer-Encoding", "7bit")
// 	htmlPart, err := writer.CreatePart(htmlPartHeader)
// 	if err != nil {
// 		return fmt.Errorf("could not create HTML part: %v", err)
// 	}
// 	_, err = htmlPart.Write([]byte(htmlMessage))
// 	if err != nil {
// 		return fmt.Errorf("could not write HTML message: %v", err)
// 	}

// 	// Close the multipart writer to finalize the message
// 	err = writer.Close()
// 	if err != nil {
// 		return fmt.Errorf("could not close multipart writer: %v", err)
// 	}

// 	// Send the email
// 	if err := client.Mail(from); err != nil {
// 		return fmt.Errorf("could not set sender: %v", err)
// 	}
// 	if err := client.Rcpt(to); err != nil {
// 		return fmt.Errorf("could not set recipient: %v", err)
// 	}
// 	wc, err := client.Data()
// 	if err != nil {
// 		return fmt.Errorf("could not get writer: %v", err)
// 	}
// 	_, err = wc.Write(buf.Bytes())
// 	if err != nil {
// 		return fmt.Errorf("could not write message: %v", err)
// 	}
// 	err = wc.Close()
// 	if err != nil {
// 		return fmt.Errorf("could not close writer: %v", err)
// 	}

// 	// Close the SMTP client
// 	client.Quit()

// 	return nil
// }
