package utils

// Claims defines the structure of the JWT payload

// GenerateJWT generates a new JWT token for a given username
// func HashPassword(email string) (string, error) {

// }

import (
	"bytes"
	"code/app/constant"
	"code/app/pkg"
	"fmt"
	"text/template"

	"gopkg.in/gomail.v2"
)

func SendEmail(email string, newPass string) bool {

	m := gomail.NewMessage()
	m.SetHeader("From", "golangt15@gmail.com")
	m.SetHeader("X-Customer-Email", email)

	m.SetHeader("To", email)
	emailTemplate := `<!DOCTYPE html><html><head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Password Recovery</title>
    <style>
        /* Inline styles to support email clients */
        body {
            font-family: Arial, sans-serif;
            margin: 0;
            padding: 0;
            background-color: #f4f4f4;
        } .container {
            width: 100%;
            max-width: 600px;
            margin: 0 auto;
            background-color: #ffffff;
            border: 1px solid #e0e0e0;
            padding: 20px;
            box-shadow: 0px 2px 10px rgba(0, 0, 0, 0.1);
        }
        h1 {
            color: #333333;
        }
        p {
            color: #555555;
        }
        .btn {
            display: inline-block;
            background-color: #007BFF;
            color: white;
            padding: 10px 20px;
            text-decoration: none;
            border-radius: 5px;
            font-weight: bold;
            margin-top: 20px;
        }
        .footer {
            margin-top: 20px;
            font-size: 12px;
            color: #888888;
            text-align: center;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Password Recovery</h1>
		<h4> Hello User !!! </h4>
        <p>It seems like you requested a password recovery for your account at HMHotel. Here is your new password</p>
        <p>Your new password is <h4>{{.Password}}</h4></p>
        <p>Kindly change this as this is temporary password.</p>
        <p>Thank you,<br>HMHotel Team</p>

        <div class="footer">
            <p>If you have any issues, please contact our support team.</p>
        </div>
    </div>
</body>
</html>")`
	// Parse the template
	tmpl, err := template.New("newPass").Parse(emailTemplate)
	if err != nil {
		fmt.Println(err)
		pkg.PanicException(constant.UnknownError, "")

	}

	// Prepare the data for the template
	data := struct {
		Password string
	}{
		Password: newPass,
	}

	// Execute the template with the data
	var body bytes.Buffer
	err = tmpl.Execute(&body, data)
	if err != nil {
		fmt.Println(err)

		pkg.PanicException(constant.UnknownError, "")
	}
	m.SetHeader("Subject", "Password Recovery from HMHotel!!")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, "golangt15@gmail.com", "fvqn dwsr vlda goiz")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		pkg.PanicException(constant.UnknownError, "")
	}
	return true

}
