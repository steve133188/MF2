package Util

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"net/smtp"
)

// func encodeRFC2047(String string) string {
// 	// use mail's rfc2047 to encode any string
// 	addr := mail.Address{String, ""}
// 	return strings.Trim(addr.String(), " <>")
// }

func SendEmail(address string, HashPassword string) error {
	from := GoDotEnvVariable("SENDER_EMAIL")
	password := GoDotEnvVariable("SENDER_PASSWORD")

	to := address
	title := "Reset Your Password"

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	var body bytes.Buffer

	t, err1 := template.ParseFiles("Util/template.html")
	if err1 != nil {
		fmt.Println(err1.Error())
	}

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	header := make(map[string]string)
	header["From"] = from
	header["To"] = to
	header["Subject"] = title
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	// randomPassword := generatePassword(2, 2, 2, 20)

	exErr := t.Execute(&body, struct {
		Password string
	}{
		Password: HashPassword,
	})

	if exErr != nil {
		fmt.Println(exErr)
	}
	// body.Write([]byte(message))
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body.Bytes()))

	// message += body.String()

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
	return err
}
