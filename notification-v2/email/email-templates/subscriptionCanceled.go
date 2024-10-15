package emailtemplates

import (
	"bytes"
	"html/template"
)

type SubscriptionCanceledData struct {
	UserFullName  string
	EffectiveDate string
}

var subscriptionCanceledTitle = "Important Update Subscription Downgrade Information"

func SubscriptionCanceledTemplate(data SubscriptionCanceledData) (string, string, error) {

	var tmpl = `<p>Dear <strong>{{.UserFullName}}</strong>,</p>

<p>We are truly sorry to see you go. Your subscription has been successfully canceled as per your request. This cancellation will take effect immediately, starting from the next billing cycle on <strong>{{.EffectiveDate}}</strong>. You will not be charged further.</p>

<p>While we're sad to part ways, we hope to have the opportunity to serve you again in the future. If there was anything about our service that did not meet your expectations, or if there's anything we can do to improve our offering, please let us know. Your feedback is invaluable as we strive to enhance our services.</p>

<p>We would love to welcome you back anytime, should you decide to use our services again. Please remember that we are always here to help with any needs you may have in the future.</p>

<p>Thank you for having been a part of our community. We wish you all the best on your future endeavors.</p>

<p>Best regards,<br>
Your Team at Capture411</p>
`

	tpl, err := template.New("SubscriptionCanceled").Parse(tmpl)
	if err != nil {
		return "", "", err
	}

	var buf bytes.Buffer
	if err = tpl.Execute(&buf, data); err != nil {
		return "", "", err
	}

	return subscriptionCanceledTitle, buf.String(), nil

}

func SubscriptionCanceledShortTemplate(data SubscriptionCanceledData) (string, string, error) {
	var tmpl = `Dear {{.UserFullName}} - We are truly sorry to see you go. Your subscription has been successfully canceled as per your request. This cancellation will take effect immediately, starting from the next billing cycle on {{.EffectiveDate}}. You will not be charged further. While we're sad to part ways, we hope to have the opportunity to serve you again in the future.`

	tpl, err := template.New("SubscriptionCanceledShort").Parse(tmpl)
	if err != nil {
		return "", "", err
	}

	var buf bytes.Buffer
	if err = tpl.Execute(&buf, data); err != nil {
		return "", "", err
	}
	return subscriptionCanceledTitle, buf.String(), nil
}
