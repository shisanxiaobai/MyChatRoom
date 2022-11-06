package errmsg

const (
	SUCCESS = 200

	//code=1000...用户模块的错误
	ERROR_ACCOUNT_USED   = 1001
	ERROR_PASSWORD_WRONG = 1002
	ERROR_USER_NOT_EXIST = 1003
	ERROR_NOT_EMPTY      = 1004

	//Token的错误
	ERROR_TOKEN_INEXIST    = 1006
	ERROR_TOKEN_RUNTIME    = 1007
	ERROR_TOKEN_WRONG      = 1008
	ERROR_TOKEN_TYPE_WRONG = 1009
	ERROR_USER_NO_RIGHT    = 1010

	//email的错误
	ERROR_EMAIL_NOT_EMPTY    = 1011
	ERROR_EMAIL_EXIST        = 1012
	ERROR_EMAIL_FAIL_TO_SEND = 1013

	//验证码错误
	ERROR_CAPTCHA_WRONG = 1100

	//code=2000...数据库发生错误
	ERROR_MYSQL   = 2001
	ERROR_REDIS   = 2002
	ERROR_MONGODB = 2003

	//code=3000....room模块错误
	ERROR_ROOM_EXIST = 3001

	//参数不能为空
	ERROR_PARAM_EMPTY = 999

	ERROR = 500
)

var codeMsg = map[int]string{
	SUCCESS: "成功！",

	//code=1000...用户模块的错误
	ERROR_ACCOUNT_USED:   "帐号已存在",
	ERROR_PASSWORD_WRONG: "密码错误",
	ERROR_USER_NOT_EXIST: "用户不存在",
	ERROR_NOT_EMPTY:      "用户名或密码不能为空",

	//Token的错误
	ERROR_TOKEN_INEXIST:    "TOKEN不存在,请重新登陆",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期,请重新登陆",
	ERROR_TOKEN_WRONG:      "TOKEN不正确,请重新登陆",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误,请重新登录",
	ERROR_USER_NO_RIGHT:    "该用户无权限",

	//email的错误
	ERROR_EMAIL_NOT_EMPTY:    "email不能为空",
	ERROR_EMAIL_EXIST:        "当前邮箱已被注册",
	ERROR_EMAIL_FAIL_TO_SEND: "发送邮件失败",

	//验证码的错误
	ERROR_CAPTCHA_WRONG: "验证码不正确",

	//code=2000...数据库发生错误
	ERROR_MYSQL:   "mysql操作发生错误",
	ERROR_REDIS:   "redis操作发生错误",
	ERROR_MONGODB: "mongodb操作发生错误",

	//code=3000....room模块错误
	ERROR_ROOM_EXIST: "房间号已存在",

	//参数不能为空
	ERROR_PARAM_EMPTY: "参数不能为空",

	ERROR: "错误",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
