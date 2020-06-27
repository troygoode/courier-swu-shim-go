package swushim

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/trycourier/courier-go/v2"
)

// CourierClientOptions let you configure the Courier Client
type CourierClientOptions struct {
	AuthToken  *string
	BaseURL    *string
	TeamEmails []string
}

// DefaultOptions specify Courier's recommended default client configuration
func DefaultOptions() CourierClientOptions {
	return CourierClientOptions{
		AuthToken:  nil, // wil lget from env var
		BaseURL:    nil, // will use default
		TeamEmails: nil, // for use with bccTeam boolean
	}
}

// SWUShim stores the reference to the Courier client for later use
type SWUShim struct {
	options CourierClientOptions
	Courier *courier.Client
}

// CreateClient is a shim that maps existing SWU calls to Courier
func CreateClient(options *CourierClientOptions) *SWUShim {
	var opt CourierClientOptions
	if options == nil {
		opt = DefaultOptions()
	} else {
		opt = *options
	}

	var authToken string
	if opt.AuthToken == nil {
		authToken = os.Getenv(("COURIER_AUTH_TOKEN"))
	} else {
		authToken = *opt.AuthToken
	}

	return &SWUShim{
		options: opt,
		Courier: courier.CreateClient(authToken, opt.BaseURL),
	}
}

// SendEmail sends an email via Courier
func (shim *SWUShim) SendEmail(recipientEmail string, recipientName *string, templateID string, cc []string, bccTeam bool, tmplParams map[string]interface{}) (string, error) {
	return shim.SendEmailWithAttachment(recipientEmail, recipientName, templateID, cc, bccTeam, nil, tmplParams)
}

// SendEmailWithAttachment sends an email via Courier, with attachment(s)
func (shim *SWUShim) SendEmailWithAttachment(recipientEmail string, recipientName *string, templateID string, cc []string, bccTeam bool, attachments map[string]*bytes.Reader, tmplParams map[string]interface{}) (string, error) {
	var files []attachment

	if attachments != nil {
		files = make([]attachment, 0, len(attachments))
		for k, v := range attachments {
			files = append(files, attachment{
				filename:    k,
				file:        v,
				contentType: nil,
			})
		}
	}

	return shim.sendEmailNotification(recipientEmail, recipientEmail, recipientName, templateID, cc, bccTeam, files, tmplParams)
}

// Attachment specifies the details of an email file attachment
type attachment struct {
	filename    string
	contentType *string
	file        *bytes.Reader
}

// SendEmailNotification sends an email via Courier
func (shim *SWUShim) sendEmailNotification(recipientID string, recipientEmail string, recipientName *string, templateID string, cc []string, bccTeam bool, attachments []attachment, tmplParams map[string]interface{}) (string, error) {
	ctx := context.Background()
	eventID := templateID

	to := recipientEmail
	if recipientName != nil {
		to = fmt.Sprintf("\"%s\" <%s>", *recipientName, recipientEmail)
	}

	profile := make(map[string]interface{})
	profile["email"] = to

	data := make(map[string]interface{})
	for k, v := range tmplParams {
		data[k] = v
	}
	if cc != nil && len(cc) > 0 {
		data["cc"] = strings.Join(cc, ",")
	}
	if bccTeam == true && shim.options.TeamEmails != nil && len(shim.options.TeamEmails) > 0 {
		data["bcc"] = strings.Join(shim.options.TeamEmails, ",")
	}

	body := make(map[string]interface{})
	body["profile"] = profile
	body["data"] = data

	if attachments != nil && len(attachments) > 0 {
		override := make(map[string]interface{})
		overrideMailgun := make(map[string]interface{})
		override["mailgun"] = overrideMailgun
		body["override"] = override

		files := make([]map[string]interface{}, 0, len(attachments))
		for index := range attachments {
			contentBytes, attachmentErr := ioutil.ReadAll(attachments[index].file)
			if attachmentErr != nil {
				return "", attachmentErr
			}
			contentBase64 := base64.StdEncoding.EncodeToString(contentBytes)

			f := make(map[string]interface{})
			f["filename"] = attachments[index].filename
			f["contentType"] = attachments[index].contentType
			f["data"] = contentBase64
			files = append(files, f)
		}
		overrideMailgun["attachments"] = files
	}

	return shim.Courier.SendMap(ctx, eventID, recipientID, body)
}
