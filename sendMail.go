package sendMail

import (
	"fmt"
)

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
