package internal

import (
	"context"
	"errors"
	"fmt"
	htmltemplate "html/template"
	"net/http"
	"strings"
	texttemplate "text/template"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/Ublius/HSOreCTF/database"
)
