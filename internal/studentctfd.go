package internal

import (
	"context"
	htmltemplate "html/template"
	"strings"
	texttemplate "text/template"

	"github.com/rs/zerolog"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func (a *Application) sendCTFdPasswordEmail(ctx context.Context, studentName, email string, ctfdPassword string) error {
	log := zerolog.Ctx(ctx).With().Str("action", "sendStudentCTFdPasswordEmail").Logger()

	templateData := map[string]any{
		"StudentName":  studentName,
		"CTFdPassword": ctfdPassword,
	}

	var plainTextContent, htmlContent strings.Builder
	texttemplate.Must(texttemplate.ParseFS(emailTemplates, "emailtemplates/ctfdaccount.txt")).Execute(&plainTextContent, templateData)
	htmltemplate.Must(htmltemplate.ParseFS(emailTemplates, "emailtemplates/ctfdaccount.html")).Execute(&htmlContent, templateData)

	subject := "Oresec HSOreCTF CTFd Password"
	err := a.SendEmail(log, subject,
		mail.NewEmail(studentName, email),
		plainTextContent.String(),
		htmlContent.String())
	if err != nil {
		log.Err(err).Msg("failed to send email")
		return err
	}
	log.Info().Msg("sent email")
	return nil

}
