package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type Email struct{
	*SMTPInfo
}

//定义SMTPinfo 结构体，用于传递发送邮箱所必须的信息。
type SMTPInfo struct{
	Host string
	Port int
	IsSSL bool
	UserName string
	PassWord string
	From string
}

func NewEmail(info *SMTPInfo) *Email{
	return &Email{SMTPInfo:info}
}

func (e *Email) SendMail(to []string,subject,body string)error{
	m := gomail.NewMessage()
	m.SetHeader("From",e.From)	//发件人 From
	m.SetHeader("To",to...)	//收件人 TO
	m.SetHeader("Subject",subject)  //邮件主题 Subject
	m.SetBody("text/html",body)  //邮件正文 Body

	//gomail.NewDialer 创建一个新的 SMTP拨号实例。设置对应的拨号信息，用于连接 SMTP服务器最后调用 DialAndSend方法，
	//打开与 SMTP 服务器的连接并发送电子邮件。
	dialer := gomail.NewDialer(e.Host,e.Port,e.UserName,e.PassWord)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSL}
	return dialer.DialAndSend(m)
}