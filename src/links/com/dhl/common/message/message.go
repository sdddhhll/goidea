package message

const(
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"
)

type LoginMessage struct{
	UserId int //用户id
	UserPwd string //用户密码 
	UserName string //用户名
}

type Message struct{
	Type string //消息类型
	Data interface{} //消息内容
}

type LoginResMes struct{
	Code int //返回状态码 500表示用户未注册，200表示成功
	Error string 
}

