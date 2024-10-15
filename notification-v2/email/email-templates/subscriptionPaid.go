package emailtemplates

import (
	"bytes"
	"html/template"
)

type SubscriptionPaidData struct {
	UserFullName     string
	SubscriptionType string
	CreditsAllocated string
	AmountPaid       string
	SubscriptionDate string
}

var subscriptionPaidTitle = "Confirmation of Your Subscription Details"

func SubscriptionPaidTemplate(data SubscriptionPaidData) (string, string, error) {

	var tmpl = `<p>Dear <strong>{{.UserFullName}}</strong>,</p>
<p>Thank you for subscribing! We are pleased to confirm your subscription details as follows:</p>

<p><strong style="font-weight: bold;">Subscription Type:</strong> {{.SubscriptionType}}</p>	<p><strong style="font-weight: bold;">Credits Allocated:</strong> {{.CreditsAllocated}}</p>
<p><strong style="font-weight: bold;">Amount Paid:</strong> {{.AmountPaid}}</p>
<p><strong style="font-weight: bold;">Subscription Date:</strong> {{.SubscriptionDate}}</p>

<p>This email is to ensure you have all the necessary details about your subscription. If you have any questions or require further assistance, feel free to reach out to our support team.</p>

<p style="font-weight: bold;">Please note: This is an automated message. Replies to this email will not be monitored or answered.</p>

<p>Thank you for choosing us!</p>

<p>Best regards,<br>
Your Team at Capture411</p>
`

	tpl, err := template.New("subscriptionPaid").Parse(tmpl)
	if err != nil {
		return "", "", err
	}

	var buf bytes.Buffer
	if err = tpl.Execute(&buf, data); err != nil {
		return "", "", err
	}

	return subscriptionPaidTitle, buf.String(), nil

}

func SubscriptionPaidShortTemplate(data SubscriptionPaidData) (string, string, error) {
	var tmpl = `Dear {{.UserFullName}} - Thank you for subscribing! We are pleased to confirm your subscription details as follows: Subscription Type: {{.SubscriptionType}}, Credits Allocated: {{.CreditsAllocated}}, Amount Paid: {{.AmountPaid}}, Subscription Date: {{.SubscriptionDate}}. This email is to ensure you have all the necessary details about your subscription.`

	tpl, err := template.New("subscriptionPaidShort").Parse(tmpl)
	if err != nil {
		return "", "", err
	}

	var buf bytes.Buffer
	if err = tpl.Execute(&buf, data); err != nil {
		return "", "", err
	}
	return subscriptionPaidTitle, buf.String(), nil
}
