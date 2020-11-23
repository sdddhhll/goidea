package main
import "fmt"
import "net"

func process(conn net.Conn){
	defer conn.Close()
	for {
		buf := make([]byte,8096)
		fmt.Println("接收客户端发送的数据...")
		n,err := conn.Read(buf[:4])
		if err != nil || n!=4 {
			fmt.Println("出现错误",err)
			return
		}
		fmt.Println("读到数据：",buf[:4])
	}
}

func main(){
	lis,err := net.Listen("tcp","localhost:8888")
	if err != nil {
		fmt.Println("服务器启动监听失败")
		return
	}
	
	for {
		fmt.Println("等待客户端发送消息")
		conn,err := lis.Accept()
		if err != nil {
			fmt.Println("Accept错误=",err)
			return
		}
		go process(conn)
	}
}