# courier-swu-shim-go

Golang shim to send Courier messages with a SWU-esque API

## Example using shim

```golang
var recipientName *string
var cc []string

authToken := os.Getenv("COURIER_AUTH_TOKEN")
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

shim := swushim.CreateClient(authToken, nil)
err := shim.SendEmail(recipientEmail, recipientName, templateID, cc, bccTeam, tmplParams)
```

## Legacy Callsite

```golang
func (swu *swuMailer) sendEmailWithAttachment(recipientEmail string, recipientName *string, templateId string, cc []string, bccTeam bool, attachments map[string]*bytes.Reader, tmplParams map[string]interface{}) {
  // assembles recipients into expected email types, reads attachments into []byte, makes POST to SWU API
}
func (swu *swuMailer) sendEmail(recipientEmail string, recipientName *string, templateId string, cc []string, bccTeam bool, tmplParams map[string]interface{}) {
	swu.sendEmailWithAttachment(recipientEmail, recipientName, templateId, cc, bccTeam, nil, tmplParams)
}
```

## Testing

Create `env.sh` locally:

```bash
export COURIER_AUTH_TOKEN=TODO
export SHIM_TEMPLATE_ID=TODO
export SHIM_EMAIL_TO=TODO
```

Then run:

```bash
source ./env.sh
go test
```

## License
The package is available as open source under the terms of the MIT License.

[MIT License](http://www.opensource.org/licenses/mit-license.php)
