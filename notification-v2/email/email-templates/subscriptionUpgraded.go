package emailtemplates

import (
	"bytes"
	"html/template"
)

type SubscriptionUpgradedData struct {
	UserFullName         string
	PrevSubscriptionType string
	NewSubscrptionType   string
}

var subscriptionUpgradedTitle = "Exciting News Your Subscription Has Been Upgraded!"

func SubscriptionUpgradedTemplate(data SubscriptionUpgradedData) (string, string, error) {

	var tmpl = `<p>Dear <strong>{{.UserFullName}}</strong>,</p>
<p>We are excited to inform you that your subscription has been successfully upgraded. Here are the details of your new subscription plan:</p>

<p><strong style="font-weight: bold;">Previous Subscription:</strong> {{.PrevSubscriptionType}}</p>
<p><strong style="font-weight: bold;">Upgraded Subscription:</strong> {{.NewSubscrptionType}}</p>

<p>This upgrade will take effect immediately, and you can start enjoying all the new features and benefits right away!</p>

<p>If you have any questions or require further information, please feel free to contact our customer support team. We are here to ensure you get the most out of your subscription.</p>

<p>Thank you for choosing to enhance your experience with us. We look forward to continuing to serve you with even more great features and services.</p>

<p>Best regards,<br>
Your Team at Capture411</p>
`

	tpl, err := template.New("subscriptionUpgraded").Parse(tmpl)
	if err != nil {
		return "", "", err
	}

	var buf bytes.Buffer
	if err = tpl.Execute(&buf, data); err != nil {
		return "", "", err
	}

	return subscriptionUpgradedTitle, buf.String(), nil

}

func SubscriptionUpgradedShortTemplate(data SubscriptionUpgradedData) (string, string, error) {
	var tmpl = `Dear {{.UserFullName}} - We are excited to inform you that your subscription has been successfully upgraded. Here are the details of your new subscription plan: Previous Subscription: {{.PrevSubscriptionType}}, Upgraded Subscription: {{.NewSubscrptionType}}. This upgrade will take effect immediately, and you can start enjoying all the new features and benefits right away!`

	tpl, err := template.New("subscriptionUpgradedShort").Parse(tmpl)
	if err != nil {
		return "", "", err
	}

	var buf bytes.Buffer
	if err = tpl.Execute(&buf, data); err != nil {
		return "", "", err
	}
	return subscriptionUpgradedTitle, buf.String(), nil
}
