package controller

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"
)

func SendMail(usuario string, saldoTotal string, debitoTotal string, creditoTotal string, meses []string) {
	auth := smtp.PlainAuth(
		"",
		"nepomucenoalmontejuan@gmail.com",
		"giwwpqkftlpgjyeh",
		"smtp.gmail.com",
	)

	tmpl, err := template.ParseFiles("Assets/template.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	data := struct {
		Name string
	}{
		Name: "Juan Nepomuceno",
	}

	var bodyContent bytes.Buffer
	err = tmpl.Execute(&bodyContent, data)
	if err != nil {
		fmt.Println(err)
		return
	}

	content := fmt.Sprintf("To: galigaribaldi0@gmail.com\r\n"+
		"Subject: Correo con Plantilla\r\n"+
		"Content-Type: text/html; charset=utf-8\r\n"+
		"\r\n"+
		"%s", bodyContent.String())
	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"galigaribaldi0@gmail.com",
		[]string{"galigaribaldi0@gmail.com"},
		[]byte(content),
	)
	if err != nil {
		fmt.Println(err)
	}
}
