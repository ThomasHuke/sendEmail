package sendMail

import (
	"net/smtp"
	"strings"
)

func htmlH(mailType string) bool {
	if mailType == "html" {
		return true
	} else if mailType == "Html" {
		return true
	} else {
		return false
	}
}

func sendToMail(user, password, host, to, subject, body, mailType string) error {
	hp := strings.Split(host, ":")                    // 分割字符串
	auth := smtp.PlainAuth("", user, password, hp[0]) //返回一个auth的interface的实例。这个interface内容是nil
	//实际上就是一个实现了next和start的一个struct。这里取hp[0]是因为传入的只是一个hostname并不需要端口，故。
	var contentType string
	if htmlH(mailType) { //mailtype是邮件的类型。
		contentType = "Content-Type: text/" + mailType + "; charset=UTF-8" //如果是HTML那么头文件就是
		// Content-Type: text/html; charset=UTF-8
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8" //
		//Content-Type:text/plain; charset=UTF-8
	}
	// 将字符串 转换为 []byte 类型,因为在senMail中我们要传入一个[]byte类型的参数
	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";") //使用；作为to中的分隔符。当然返回的是一个slice
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	return err
}
