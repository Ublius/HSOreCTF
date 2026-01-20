package internal

import (
	"context"
	"encoding/base64"
	"fmt"
	htmltemplate "html/template"
	"strings"
	texttemplate "text/template"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mailjet/mailjet-apiv3-go/v4"
	"github.com/rs/zerolog"
	qrcode "github.com/skip2/go-qrcode"
)

func (a *Application) getStudentQRCodeURL(email string) (string, error) {
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		Issuer:  string(IssuerStudentQRCode),
		Subject: email,
	})

	signedTok, err := tok.SignedString(a.Config.ReadSecretKey())
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/volunteer/scan?tok=%s", a.Config.Domain, signedTok), nil
}

func (a *Application) getStudentQRCodeImage(email string) ([]byte, error) {
	url, err := a.getStudentQRCodeURL(email)
	if err != nil {
		return nil, err
	}
	return qrcode.Encode(url, qrcode.Medium, 256)
}

func (a *Application) sendQRCodeEmail(ctx context.Context, studentName, email string) error {
	subject := "OreSec HSOreCTF Ticket"
	log := zerolog.Ctx(ctx).With().
		Str("component", "send_email").
		Interface("to", email).
		Str("subject", subject).
		Logger()

	log.Info().Msg("sending email")

	qrcodeBytes, err := a.getStudentQRCodeImage(email)
	if err != nil {
		a.Log.Err(err).Msg("failed to get student QR code image")
		return err
	}
	// Base64 encode the qrcodeBytes
	qrcodeBase64 := base64.StdEncoding.EncodeToString(qrcodeBytes)

	templateData := map[string]any{
		"StudentName":  studentName,
		"QRCodeBase64": qrcodeBase64,
	}

	var plainTextContent, htmlContent strings.Builder
	texttemplate.Must(texttemplate.ParseFS(emailTemplates, "emailtemplates/ticket.txt")).Execute(&plainTextContent, templateData)
	htmltemplate.Must(htmltemplate.ParseFS(emailTemplates, "emailtemplates/ticket.html")).Execute(&htmlContent, templateData)

	log.Debug().Any("pt", plainTextContent.String()).Any("html", htmlContent.String()).Msg("sending email")

	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "support@hsorectf.com",
				Name:  "HSOreCTF Support",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: email,
					Name:  "",
				},
			},
			Subject:  subject,
			TextPart: plainTextContent.String(),
			HTMLPart: htmlContent.String(),
			ReplyTo: &mailjet.RecipientV31{
				Email: "support@hsorectf.com",
				Name:  "HSOreCTF Support",
			},
			Attachments: &mailjet.AttachmentsV31{
				mailjet.AttachmentV31{
					Filename:      "ticket.png",
					ContentType:   "image/png",
					Base64Content: qrcodeBase64,
				},
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
			Str("status_code", res.ResultsV31[0].Status).
			Msg("error sending email")
		return fmt.Errorf("error sending email: %s", res.ResultsV31[0].Status)
	}

	log.Info().
		Str("status_code", res.ResultsV31[0].Status).
		Msg("successfully sent email")
	return nil
}
