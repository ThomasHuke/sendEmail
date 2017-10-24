package main

import (
	e "github.com/thomashuke/sendEmail"
)

func main() {
	user := "ceo@coastroad.net"
	host := "smtp.exmail.qq.com:25"
	to := "##;##"
	password := "##"
	subject := "coastroad官方邮件"
	body := `
    <!DOCTYPE html>
    <html>
    <head>
    <!-- 新 Bootstrap 核心 CSS 文件 -->
<link href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
<!-- jQuery文件。务必在bootstrap.min.js 之前引入 -->
<script src="https://cdn.bootcss.com/jquery/2.1.1/jquery.min.js"></script>

<!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
<script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
    </head>
    <body>
    <div class="container">
    <div class="row">
    <div class="col-md-4 col-sm-12">
    <p class="text-info ">
    尊敬的用户您好，我是coastroad的CEO ThomasHuke 。发送这个邮件的目的是通知您来我公司面试。
    </p>
    <a href="https://github.com/thomashuke" class="text-info bg-primary">github</a>
    </div>
    <div class="col-md-4 col-sm-12">
    <a href="http://weibo.com/u/1949695097" class="text-info bg-primary">微博</a>
    </div>
    <div class="col-md-4 col-sm-12">
    <a href="https://twitter.com/thomashukec" class="text-info bg-primary">twitter</a>
    </div>
    </div>
    </div>
    </body>
    </html>
    `
	mailType := "html"
	e.Run(user, to, password, host, subject, body, mailType)
}
