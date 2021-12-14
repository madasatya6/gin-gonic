package gomail

import (
	"bytes"
	"html/template"
	"path/filepath"
	"strings"

	gomail "gopkg.in/gomail.v2"
)

var CONFIG_SMTP_HOST = "smtp.gmail.com"
var CONFIG_SMTP_PORT = 587
var CONFIG_SENDER_NAME = "UD MADA JAYA <madasatya6@gmail.com>"
var CONFIG_AUTH_EMAIL = "madasatya6@gmail.com"
var CONFIG_AUTH_PASSWORD = "ranggaMsba643043055"

func SendGomail(subject string, toEmail string, emailCc []string, titleCc string, viewLocation string, data map[string]interface{}, fileAttachment []string) error {
	
	dir := filepath.Join("app/views/", viewLocation)
	t := template.New(dir)

	t, err := t.ParseFiles(dir)
	if err != nil {
		return err 
	}
	
	var tpl bytes.Buffer
	if err := t.ExecuteTemplate(&tpl, filepath.Base(dir), data); err != nil {
		return err 
	}

	tmplHTML := tpl.String()

    mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", toEmail)
	
	if len(emailCc) > 0 {
		mailer.SetAddressHeader("Cc", strings.Join(emailCc, ","), titleCc)
	}
	
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", tmplHTML)
	
	if fileAttachment != nil {
		mailer.Attach(strings.Join(fileAttachment, ","))
	}
	
	dialer := gomail.NewDialer(CONFIG_SMTP_HOST, CONFIG_SMTP_PORT, CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD)
	
	//kirim email dengan smtp relat atau no auth/tanpa otentikasi
	//dialer := &DialAndSend(mailer)
	
	err = dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}
	
	return nil
}
