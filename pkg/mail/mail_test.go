package mail

import (
	"testing"
)

func TestSend(t *testing.T) {
	options := &Options{
		MailHost: "smtp.163.com",
		MailPort: 465,
		MailUser: "15873458721@163.com",
		MailPass: "xx", //授权码
		MailTo:   "624796905@qq.com",
		Subject:  "subject",
		Body:     "body",
	}
	err := Send(options)
	if err != nil {
		t.Error("Mail Send error", err)
		return
	}
	t.Log("Mail Send success")
}
