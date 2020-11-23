package main
import "fmt"
import "os"

func main(){
	var key int
	var loop bool = true
	
	for loop {
		fmt.Println("---------------聊天系统---------------------")
		fmt.Println("\t\t\t1.登录聊天室")
		fmt.Println("\t\t\t2.注册用户")
		fmt.Println("\t\t\t3.退出系统")
		
		fmt.Println("请输入您的选择（1-3）：")
		fmt.Scanf("%d\n",&key)
		
		switch key {
			case 1:
				loop = false
			case 2:
				loop = false
			case 3:
				fmt.Println("退出系统...")
				os.Exit(0)
			default:
				fmt.Println("您的输入有误，请重新输入")
		}
	}
	
	if key == 1 {
		fmt.Println("欢迎登录聊天室，请输入您的用户号和密码")
		var userId int
		var userPwd string 
		fmt.Println("请输入用户ID:")
		fmt.Scanf("%d\n",&userId)
		fmt.Println("请输入用户密码:")
		fmt.Scanf("%s\n",&userPwd)
		err := login(userId,userPwd)
		if err != nil {
			fmt.Println("登录成功...")
		}
	}
	
	
}


























