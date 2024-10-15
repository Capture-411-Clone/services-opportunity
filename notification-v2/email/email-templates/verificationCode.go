package emailtemplates

import (
	"bytes"
	"html/template"
)

type VerificationCodeData struct {
	Code             string
	ResetPasswordURL string
}

var verificationCodeTitle = "Capture411 Verification Code"

func VerificationCodeTemplate(data VerificationCodeData) (string, string, string, error) {

	var plainTextTmpl = `
Capture411

Please verify you're really you by entering this 6-digit code when you sign in. Just a heads up, this code will expire in 5 minutes for security reasons.

Here is your verification code:

{{.Code}}

Enter this verification code in the app to complete your steps.

If you didn't request this code, you can safely ignore this email or reset your password here:
{{.ResetPasswordURL}}

If you have any questions, feel free to reach out to our support team at support@capture411.com.

© 2023 Capture411. All rights reserved.
1234 Memory Lane, Photo City, PC 56789
`

	plainTpl, err := template.New("VerificationCodePlain").Parse(plainTextTmpl)
	if err != nil {
		return "", "", "", err
	}

	var plainBuf bytes.Buffer
	if err = plainTpl.Execute(&plainBuf, data); err != nil {
		return "", "", "", err
	}

	var tmpl = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8" />
    <title>Verify Your Email Address</title>
  </head>
  <body style="margin: 0; padding: 0; background-color: #f4f4f4">
    <table
      align="center"
      border="0"
      cellpadding="0"
      cellspacing="0"
      width="600"
      style="
        border-collapse: collapse;
        background-color: #ffffff;
        margin-top: 50px;
        margin-bottom: 50px;
      "
    >
      <!-- Logo Section -->
      <tr>
        <td align="center" style="padding: 40px 0 30px 0">
          <img
            src="https://capture411.com/logo/Capture411_Original.svg"
            alt="Capture411 Logo"
            width="200"
            style="display: block"
          />
        </td>
      </tr>

      <!-- Title -->
      <tr>
        <td
          align="center"
          style="
            padding: 0 30px 20px 30px;
            font-family: Arial, sans-serif;
            color: #333333;
          "
        >
          <p style="font-size: 18px; margin: 0">
            Please verify you're really you by entering this 6-digit code when
            you sign in. Just a heads up, this code will expire in 5 minutes for
            security reasons.
          </p>
        </td>
      </tr>

      <!-- Message -->
      <tr>
        <td
          align="center"
          style="
            padding: 0 30px 20px 30px;
            font-family: Arial, sans-serif;
            color: #555555;
          "
        >
          <h1 style="font-size: 26px; line-height: 24px; margin: 0">
            Here is your verification code:
          </h1>
        </td>
      </tr>

      <!-- Verification Code -->
      <tr>
        <td align="center" style="padding: 20px 30px">
          <div
            style="
              display: inline-block;
              padding: 15px 25px;
              background-color: #4caf50;
              color: #ffffff;
              font-size: 24px;
              font-family: Arial, sans-serif;
              border-radius: 4px;
            "
          >
            <strong>{{.Code}}</strong>
          </div>
        </td>
      </tr>

      <!-- Instructions -->
      <tr>
        <td
          align="center"
          style="
            padding: 20px 30px 30px 30px;
            font-family: Arial, sans-serif;
            color: #555555;
          "
        >
          <p style="font-size: 16px; line-height: 24px; margin: 0">
            Enter this verification code in the app to complete your steps.
          </p>
        </td>
      </tr>

      <!-- Reset Password -->
      <tr>
        <td
          align="center"
          style="
            padding: 0 30px 30px 30px;
            font-family: Arial, sans-serif;
            color: #555555;
          "
        >
          <p style="font-size: 16px; line-height: 24px; margin: 0">
            If you didn't request this code, you can safely ignore this email or
            reset your password here:
          </p>

          <a
            href="{{.ResetPasswordURL}}"
            style="
              display: inline-block;
              padding: 10px 22px;
              background-color: black;
              color: #ffffff;
              font-size: 24px;
              font-family: Arial, sans-serif;
              border-radius: 4px;
              text-decoration: none;
              margin-top: 20px;
            "
            >Reset Password</a
          >
        </td>

        <!-- Support Message -->
      </tr>

      <tr>
        <td
          align="center"
          style="
            padding: 0 30px 30px 30px;
            font-family: Arial, sans-serif;
            color: #555555;
          "
        >
          <p style="font-size: 16px; line-height: 24px; margin: 0">
            If you have any questions, feel free to reach out to our support
            team at
            <a
              href="mailto:support@capture411.com"
              style="color: #4caf50; text-decoration: none"
              >support@capture411.com</a
            >.
          </p>
        </td>
      </tr>

      <!-- Footer -->
      <tr>
        <td
          align="center"
          style="
            padding: 30px;
            font-family: Arial, sans-serif;
            color: #999999;
            background-color: #f4f4f4;
          "
        >
          <p style="font-size: 12px; line-height: 18px; margin: 0">
            &copy; 2023 Capture411. All rights reserved.<br />
            1234 Memory Lane, Photo City, PC 56789
          </p>
        </td>
      </tr>
    </table>
  </body>
</html>
`

	tpl, err := template.New("VerificationCode").Parse(tmpl)
	if err != nil {
		return "", "", "", err
	}

	var buf bytes.Buffer
	if err = tpl.Execute(&buf, data); err != nil {
		return "", "", "", err
	}

	return verificationCodeTitle, plainBuf.String(), buf.String(), nil
}
