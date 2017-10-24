# sendEmail

- smtp业务，发布一个可定制Html的SMTP业务，使用golang进行编写主要使用了net/smtp包
- body 内容内嵌一个HTML（当然也可以是一个文本，不过建议使用HTML这样可以编写出漂亮的文字）
- to 也就是发送的人，使用“;”进行分割

## use

```go

sendEamil.Run(user, to, password, host, subject, body, mailType string)

```
其中：
- user: 用户名
- to 发送人
- password 密码
- host 你使用 smtp服务器的host
- subject 你发送邮件时的主题
- body 你的发送内容
- mailType 你发送的时候用的什么类型的body
    - html 小写字符
    - another 如果不是html的话 一律解释为plain

## 捐赠

[如果您希望帮助我更好的开源，也可以资助我一些资金](https://www.github.com/thomashuke/donate)
