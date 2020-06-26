# courier-swu-shim-go

Golang shim to send Courier messages with a SWU-esque API

## Using this shim

```golang
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

shim := swushim.CreateClient(&swushim.CourierClientOptions{
  // AuthToken: &authToken, // OPTIONAL: will default to COURIER_AUTH_TOKEN environment variable
  TeamEmails: []string{"team@example.com"}, // email addresses to be used when bccTeam=true
})
err := shim.SendEmail(recipientEmail, recipientName, templateID, cc, bccTeam, tmplParams)
```

## License
The package is available as open source under the terms of the MIT License.

[MIT License](http://www.opensource.org/licenses/mit-license.php)
