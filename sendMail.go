package sendMail

import (
	"fmt"
)

/*
Run 函数暴露在外，是为了接入 user to password subject body mailType等信息。
通常来说user就是你的用户名，password是该邮件的密码。host是smtp提供的一个address subject是邮件的主题
to是收件人，body是内容，maiType是发送邮件的类型通常是Html和Plain
*/
func Run(user, to, password, host, subject, body, mailType string) {
	fmt.Println("send email")
	err := sendToMail(user, password, host, to, subject, body, mailType)
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}
