package internal

import (
	"fmt"

	"github.com/mailjet/mailjet-apiv3-go/v4"
	"github.com/rs/zerolog"
)

func (a *Application) SendEmail(log zerolog.Logger, subject string, toEmail string, toName string, plainTextContent, htmlContent string) error {
	log = log.With().
		Str("component", "send_email").
		Str("to", toEmail).
		Str("subject", subject).
		Logger()

	if a.Config.DevMode {
		fmt.Printf("=== EMAIL ===\nTo: %s <%s>\nSubject: %s\n\n%s\n\n", toName, toEmail, subject, plainTextContent)

		return nil
	}

	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "support@hsorectf.com",
				Name:  "HSOreCTF Support",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: toEmail,
					Name:  toName,
				},
			},
			Subject:  subject,
			TextPart: plainTextContent,
			HTMLPart: htmlContent,
			ReplyTo: &mailjet.RecipientV31{
				Email: "support@hsorectf.com",
				Name:  "HSOreCTF Support",
			},
		},
	}

	messages := mailjet.MessagesV31{Info: messagesInfo}

	res, err := a.MailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Err(err).Msg("failed to send email")
		return err
	}

	if res.ResultsV31[0].Status != "success" {
		log.Error().
			Str("status", res.ResultsV31[0].Status).
			Str("to", toEmail).
			Msg("error sending email")
		return fmt.Errorf("error sending email: %s", res.ResultsV31[0].Status)
	}

	log.Info().
		Str("status_code", res.ResultsV31[0].Status).
		Str("to", toEmail).
		Msg("successfully sent email")
	return nil
}