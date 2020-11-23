package main
import "fmt"
import "net"
import "links/com/dhl/common/message"
import "encoding/binary"
import "encoding/json"
func login(userId int,userPwd string) error{
	fmt.Println("用户ID",userId)
	fmt.Println("用户密码",userPwd)
	
	fmt.Println("准备发送数据")
	
	conn,err := net.Dial("tcp","localhost:8888")
	defer conn.Close()
	if err != nil {
		fmt.Println("连接服务器失败")
		return err
	}
	var mes message.Message 
	mes.Type=message.LoginMesType
	
	var loginMes message.LoginMessage 
	loginMes.UserId=userId
	loginMes.UserPwd=userPwd
	loginMes.UserName="测试"
	mes.Data = loginMes
	
	data,err := json.Marshal(mes)
	if err != nil{
		fmt.Println("转换json数据出错")
		return err
	}
	var bytes []byte = make([]byte,4)
	var lens uint32 = uint32(len(data))
	binary.BigEndian.PutUint32(bytes, lens)
	n,err := conn.Write(bytes)
	if err != nil || n != 4{
		fmt.Println("发送json数据长度出错")
		return err
	}
	fmt.Printf("客户端发送数据长度：%d",n)
	return nil
}