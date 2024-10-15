package emailtemplates

import (
	"bytes"
	"html/template"
)

type CreditsAddedData struct {
	UserFullName     string
	CreditsPurchased string
	AmountPaid       string
	PurchaseDate     string
}

func CreditsAddedTemplate(data CreditsAddedData) (string, string, error) {
	var title = "Confirmation of Your Additional Credits Purchase"

	var tmpl = `<p>Dear <strong>{{.UserFullName}}</strong>,</p>

<p>Thank you for your recent purchase of additional credits! We appreciate your continued support and commitment to our services. Here are the details of your purchase:</p>

<p><strong style="font-weight: bold;">Credits Purchased:</strong> {{.CreditsPurchased}}</p>
<p><strong style="font-weight: bold;">Amount Paid:</strong> {{.AmountPaid}}</p>
<p><strong style="font-weight: bold;">Date of Purchase:</strong> {{.PurchaseDate}}</p>

<p>These additional credits are available for your use immediately. We hope they help you achieve all your desired tasks and enhance your experience with our services.</p>

<p>If you have any questions or require further assistance, please do not hesitate to reach out to our customer support team.</p>

<p>Thank you for choosing us!</p>

<p>Best regards,<br>
Your Team at Capture411</p>
`

	tpl, err := template.New("CreditsAdded").Parse(tmpl)
	if err != nil {
		return "", "", err
	}

	var buf bytes.Buffer
	if err = tpl.Execute(&buf, data); err != nil {
		return "", "", err
	}

	return title, buf.String(), nil

}

func CreditsAddedShortTemplate(data CreditsAddedData) (string, string, error) {
	var title = "Confirmation of Your Additional Credits Purchase"

	var tmpl = `Dear {{.UserFullName}} - Thank you for your recent purchase of additional credits! Here are the details of your purchase: Credits Purchased: {{.CreditsPurchased}}, Amount Paid: {{.AmountPaid}}, Date of Purchase: {{.PurchaseDate}}. These additional credits are available for your use immediately.`

	tpl, err := template.New("CreditsAddedShort").Parse(tmpl)
	if err != nil {
		return "", "", err
	}

	var buf bytes.Buffer
	if err = tpl.Execute(&buf, data); err != nil {
		return "", "", err
	}

	return title, buf.String(), nil
}
