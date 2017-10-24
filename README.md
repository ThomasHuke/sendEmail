# sendEmail

- smtp业务，发布一个可定制Html的SMTP业务，使用golang进行编写主要使用了net/smtp包
- body 内容内嵌一个HTML（当然也可以是一个文本，不过建议使用HTML这样可以编写出漂亮的文字）
- to 也就是发送的人，使用“;”进行分割
