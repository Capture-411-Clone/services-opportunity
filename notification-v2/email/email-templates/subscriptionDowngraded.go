package emailtemplates

import (
	"bytes"
	"html/template"
)

type SubscriptionDowngradedData struct {
	UserFullName           string
	PrevSubscriptionType   string
	NewSubscrptionType     string
	DowngradeEffectiveDate string
}

var subscriptionDowngradedTitle = "Important Update Subscription Downgrade Information"

func SubscriptionDowngradedTemplate(data SubscriptionDowngradedData) (string, string, error) {

	var tmpl = `<p>Dear <strong>{{.UserFullName}}</strong>,</p>

<p>We are writing to confirm the changes to your subscription with us. Here are the details of your subscription downgrade:</p>

<p><strong style="font-weight: bold;">Previous Subscription:</strong> {{.PrevSubscriptionType}}</p>
<p><strong style="font-weight: bold;">New Subscription:</strong> {{.NewSubscrptionType}}</p>

<p>Please be aware that this change will take effect at the start of your next billing cycle, on <strong style="font-weight: bold;">{{.DowngradeEffectiveDate}}</strong>.</p>

<p>If you have any questions about this change or any other matter, please do not hesitate to contact our customer support team. We are here to help!</p>

<p>Thank you for continuing to choose our services.</p>

<p>Best regards,<br>
Your Team at Capture411</p>
`

	tpl, err := template.New("subscriptionDowngraded").Parse(tmpl)
	if err != nil {
		return "", "", err
	}

	var buf bytes.Buffer
	if err = tpl.Execute(&buf, data); err != nil {
		return "", "", err
	}

	return subscriptionDowngradedTitle, buf.String(), nil

}

func SubscriptionDowngradedShortTemplate(data SubscriptionDowngradedData) (string, string, error) {
	var tmpl = `Dear {{.UserFullName}} - We are writing to confirm the changes to your subscription with us. Here are the details of your subscription downgrade: Previous Subscription: {{.PrevSubscriptionType}}, New Subscription: {{.NewSubscrptionType}}. Please be aware that this change will take effect at the start of your next billing cycle, on {{.DowngradeEffectiveDate}}.`

	tpl, err := template.New("subscriptionDowngradedShort").Parse(tmpl)
	if err != nil {
		return "", "", err
	}

	var buf bytes.Buffer
	if err = tpl.Execute(&buf, data); err != nil {
		return "", "", err
	}
	return subscriptionDowngradedTitle, buf.String(), nil
}
