package lib

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

type RequestMail struct {
	from      string
	to        []string
	subject   string
	body      string
}

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

func NewRequest(from string, to []string, subject string) *RequestMail {
	return &RequestMail{
		to      :      to,
		from    :      from,
		subject : subject,
	}
}

func (r *RequestMail) parseTemplate(fileName string, data interface{}) error {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		logs.Println(err)
		fmt.Println(err)
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		logs.Println(err)
		fmt.Println(err)
	}
	r.body = buffer.String()
	return nil
}

func (r *RequestMail) sendMail() bool {
	body := "To 	 : "  + r.to[0]   + "\r\n" +
			"From 	 : "  + r.from    + "\r\n" +
			"Subject : "  + r.subject + "\r\n" + MIME + "\r\n" + r.body

	SMTP   := fmt.Sprintf("%s:%s", GetEnv("SMTP_SERVER"), GetEnv("SMTP_PORT"))
	if err := smtp.SendMail(SMTP, smtp.PlainAuth("", GetEnv("SMTP_EMAIL"), GetEnv("SMTP_PASSWORD"), GetEnv("SMTP_SERVER")), r.from, r.to, []byte(body)); err != nil {
		logs.Println(err)
		fmt.Println(err)
		return false
	}
	return true
}

func (r *RequestMail) Send(templateName string, items interface{}) bool {
	err := r.parseTemplate(templateName, items)
	if err != nil {
		logs.Println("Email has been sent to ", r.to)
		fmt.Println("Email has been sent to ", r.to)
	}
	if ok := r.sendMail(); ok {
		logs.Println("Email has been sent to ", r.to)
		fmt.Println("Email has been sent to ", r.to)
	} else {
		logs.Println("Email has been sent to ", r.to)
		fmt.Println("Failed to send the email to ", r.to)
		return false
	}
	return true
}

