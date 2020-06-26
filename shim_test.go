package swushim_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	swushim "github.com/troygoode/courier-swu-shim-go"
)

func Test_SendEmail(t *testing.T) {
	var recipientName *string
	var cc []string

	recipientEmail := os.Getenv("SHIM_EMAIL_TO")
	recipientName = nil
	templateID := os.Getenv("SHIM_TEMPLATE_ID")
	cc = nil
	bccTeam := false

	tmplParams := make(map[string]interface{})
	tmplParams["orgName"] = "Example"
	tmplParams["name"] = "Jane Doe"
	tmplParams["email"] = "inviter@example.com"
	tmplParams["inviteUrl"] = "Example"

	shim := swushim.CreateClient(nil)
	_, err := shim.SendEmail(recipientEmail, recipientName, templateID, cc, bccTeam, tmplParams)
	assert.Nil(t, err)
}

func Test_WithName(t *testing.T) {
	var cc []string

	recipientEmail := os.Getenv("SHIM_EMAIL_TO")
	recipientName := "Namely Namename"
	templateID := os.Getenv("SHIM_TEMPLATE_ID")
	cc = nil
	bccTeam := false

	tmplParams := make(map[string]interface{})
	tmplParams["orgName"] = "Example"
	tmplParams["name"] = "Jane Doe"
	tmplParams["email"] = "inviter@example.com"
	tmplParams["inviteUrl"] = "Example"

	shim := swushim.CreateClient(nil)
	_, err := shim.SendEmail(recipientEmail, &recipientName, templateID, cc, bccTeam, tmplParams)
	assert.Nil(t, err)
}

func Test_WithCC(t *testing.T) {
	var recipientName *string

	recipientEmail := os.Getenv("SHIM_EMAIL_TO")
	recipientName = nil
	templateID := os.Getenv("SHIM_TEMPLATE_ID")
	cc := []string{os.Getenv("SHIM_EMAIL_ALT")}
	bccTeam := false

	tmplParams := make(map[string]interface{})
	tmplParams["orgName"] = "Example"
	tmplParams["name"] = "CC_Test"
	tmplParams["email"] = "inviter@example.com"
	tmplParams["inviteUrl"] = "Example"

	shim := swushim.CreateClient(nil)
	_, err := shim.SendEmail(recipientEmail, recipientName, templateID, cc, bccTeam, tmplParams)
	assert.Nil(t, err)
}

func Test_WithBCC(t *testing.T) {
	var recipientName *string
	var cc []string

	recipientEmail := os.Getenv("SHIM_EMAIL_TO")
	recipientName = nil
	templateID := os.Getenv("SHIM_TEMPLATE_ID")
	cc = nil
	bccTeam := true

	tmplParams := make(map[string]interface{})
	tmplParams["orgName"] = "Example"
	tmplParams["name"] = "BCC_Test"
	tmplParams["email"] = "inviter@example.com"
	tmplParams["inviteUrl"] = "Example"

	shim := swushim.CreateClient(&swushim.CourierClientOptions{
		TeamEmails: []string{os.Getenv("SHIM_EMAIL_ALT")},
	})
	_, err := shim.SendEmail(recipientEmail, recipientName, templateID, cc, bccTeam, tmplParams)
	assert.Nil(t, err)
}

func Test_WithAttachment(t *testing.T) {
	var recipientName *string
	var cc []string

	recipientEmail := os.Getenv("SHIM_EMAIL_TO")
	recipientName = nil
	templateID := os.Getenv("SHIM_TEMPLATE_ID")
	cc = nil
	bccTeam := false

	logo, logoErr := ioutil.ReadFile("courier-logo.png")
	assert.Nil(t, logoErr)
	logoReader := bytes.NewReader(logo)

	attachments := make(map[string]*bytes.Reader)
	attachments["courier-logo.png"] = logoReader

	tmplParams := make(map[string]interface{})
	tmplParams["orgName"] = "Example"
	tmplParams["name"] = "Attachment_Test"
	tmplParams["email"] = "inviter@example.com"
	tmplParams["inviteUrl"] = "Example"

	shim := swushim.CreateClient(nil)
	_, err := shim.SendEmailWithAttachment(recipientEmail, recipientName, templateID, cc, bccTeam, attachments, tmplParams)
	assert.Nil(t, err)
}
