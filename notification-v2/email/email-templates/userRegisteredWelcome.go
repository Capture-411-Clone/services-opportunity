package emailtemplates

import (
	"bytes"
	"html/template"
)

type UserRegisteredWelcomeData struct {
	UserFullName string
}

var welcomeTitle = "Welcome to the Capture411 Community!"

func UserRegisteredWelcomeTemplate(data UserRegisteredWelcomeData) (string, string, string, error) {
	var plainTextTmpl = `<p>Dear <strong>{{.UserFullName}}</strong>,</p>

Welcome to Capture411!

We're absolutely thrilled to have you join our community. Thank you for choosing us to be a part of your journey. At Capture411, we are dedicated to providing you with the best experience possible, helping you capture and embrace every moment that matters.

Explore Our Features: Dive into our platform and discover tools designed to enhance your experience.
Stay Updated: Keep an eye on your inbox for exciting news, updates, and exclusive offers.
We're Here for You: If you have any questions or need assistance, our support team is just an email away at support@capture411.com.

Warm regards,
Christina Mee CEO Capture411
`

	plainTpl, err := template.New("UserRegisteredWelcome").Parse(plainTextTmpl)
	if err != nil {
		return "", "", "", err
	}

	var plainBuf bytes.Buffer
	if err = plainTpl.Execute(&plainBuf, data); err != nil {
		return "", "", "", err
	}

	// ---------------------------------HTML----------------------------------

	var tmpl = `<p>Dear <strong>{{.UserFullName}}</strong>,</p>

<p>Welcome to <strong>Capture411!</strong></p>

<p>We're absolutely thrilled to have you join our community. Thank you for choosing us to be a part of your journey. At Capture411, we are dedicated to providing you with the best experience possible, helping you capture and embrace every moment that matters.</p>
<p>
<ul>
	<li>Explore Our Features: Dive into our platform and discover tools designed to enhance your experience.</li>
	<li>Stay Updated: Keep an eye on your inbox for exciting news, updates, and exclusive offers.</li>
	<li>We're Here for You: If you have any questions or need assistance, our support team is just an email away at support@capture411.com.</li>
</ul>
</p>
<p>Warm regards,</p>
<p>Christina Mee CEO Capture411</p>
`

	tpl, err := template.New("UserRegisteredWelcome").Parse(tmpl)
	if err != nil {
		return "", "", "", err
	}

	var buf bytes.Buffer
	if err = tpl.Execute(&buf, data); err != nil {
		return "", "", "", err
	}

	return welcomeTitle, plainBuf.String(), buf.String(), nil

}

func UserRegisteredWelcomeShortTemplate(data UserRegisteredWelcomeData) (string, string, error) {
	var tmpl = `Dear {{.UserFullName}} - Welcome to Capture411! We're thrilled to have you join our community. Explore our features, stay updated, and remember that we're here for you.`

	tpl, err := template.New("UserRegisteredWelcome").Parse(tmpl)
	if err != nil {
		return "", "", err
	}

	var buf bytes.Buffer
	if err = tpl.Execute(&buf, data); err != nil {
		return "", "", err
	}

	return welcomeTitle, buf.String(), nil
}
