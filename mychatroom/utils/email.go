package utils

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"

	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
)

// 给用户发送一个邮件，内容是验证码
func SendEmail(toEmail, verificationCode string) error {
	//读取配置，发送邮箱和发送人,邮箱IMAP密码
	sendEmail := viper.GetString("server.email")
	sender := viper.GetString("server.sender")
	mailPassword := viper.GetString("server.mailpassword")
	//邮箱初始化设置
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender, sendEmail)
	e.To = []string{toEmail}
	e.Bcc = []string{sendEmail}
	e.Cc = []string{sendEmail}
	e.Subject = "请查收验证码"
	e.Text = []byte(" 您收到了验证码 ")
	e.HTML = []byte("<h1> 邮件内容是 </h1>" + "<b>" + verificationCode + "</b>")
	return e.Send("smtp.qq.com:587", smtp.PlainAuth("", sendEmail, mailPassword, "smtp.qq.com"))
}

// 生成邮箱验证码
func GetVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	res := ""
	for i := 0; i < 6; i++ {
		res += strconv.Itoa(rand.Intn(10))
	}
	return res
}
