package MainInterface

import (
	"fmt"
	"go2/com/dhl/customer"
)
type MainInterface struct {

}

func (mi *MainInterface) RenderMain(){
	fmt.Println("--------------------客户信息管理--------------------")
	fmt.Println()
	fmt.Print(`
					1：添加客户
					2：修改客户
					3：删除客户
					4：客户列表
					5：退    出
			`)
	fmt.Println()
}

func (mi *MainInterface) RenderAdd(css *[]customer.Customer){
	fmt.Println("--------------------添加--客户--------------------")
	fmt.Println()

	var cus customer.Customer
	cus.ID = cus.GetMaxCustomer(css).ID + 1
	fmt.Print("姓名：")
	fmt.Scanf("%s\n",&cus.Name)
	fmt.Print("姓别：")
	fmt.Scanf("%s\n",&cus.Gender)
	fmt.Print("年龄：")
	fmt.Scanf("%d\n",&cus.Age)
	fmt.Print("电话：")
	fmt.Scanf("%s\n",&cus.Tel)
	fmt.Print("邮箱：")
	fmt.Scanf("%s\n",&cus.Mail)
	cus.AddCustomer(css)
	fmt.Println()
	fmt.Println("--------------------添加--完成--------------------")
	fmt.Println()
	mi.SubCommand(css)
}


func (mi *MainInterface) RenderModify(css *[]customer.Customer){
	fmt.Println("--------------------修改--客户--------------------")
	fmt.Println()
	var iId int64
	var cus customer.Customer
	fmt.Print("请输入客户号：")
	fmt.Scanf("%d\n",&iId)
	id, b := cus.GetCustomerById(css, iId)
	if !b {
		fmt.Println()
		fmt.Println("--------------------修改--失败--------------------")
		return
	}

	fmt.Printf("姓名(%s)：",id.Name)
	fmt.Scanf("%s\n",&cus.Name)
	fmt.Printf("姓别(%s)：",id.Gender)
	fmt.Scanf("%s\n",&cus.Gender)
	fmt.Printf("年龄(%d)：",id.Age)
	fmt.Scanf("%d\n",&cus.Age)
	fmt.Printf("电话(%s)：",id.Tel)
	fmt.Scanf("%s\n",&cus.Tel)
	fmt.Printf("邮箱(%s)：",id.Mail)
	fmt.Scanf("%s\n",&cus.Mail)
	if cus.Name != "" {
		id.Name = cus.Name
	}

	if cus.Gender !="" {
		id.Gender = cus.Gender
	}

	if cus.Age != 0 {
		id.Age = cus.Age
	}

	if cus.Tel != "" {
		id.Tel = cus.Tel
	}

	if cus.Mail != "" {
		id.Mail = cus.Mail
	}

	fmt.Println()
	fmt.Println("--------------------修改--完成--------------------")
	fmt.Println()
	mi.SubCommand(css)
}


func (mi *MainInterface) RenderDelete(css *[]customer.Customer){
	fmt.Println("--------------------删除--客户--------------------")
	fmt.Println()
	var iId int64
	var cus customer.Customer
	fmt.Print("请输入要删除的客户号：")
	fmt.Scanf("%d\n",&iId)
	_, b := cus.GetCustomerById(css, iId)
	if !b {
		fmt.Println()
		fmt.Println("--------------------客户不存在--------------------")
		mi.RenderDelete(css)
		return
	}
	cus.DeleteCustomer(css,iId)

	fmt.Println()
	fmt.Println("--------------------删除--完成--------------------")
	fmt.Println()
	mi.SubCommand(css)
}


func (mi *MainInterface) RenderList(css *[]customer.Customer){
	fmt.Println("--------------------客户--列表--------------------")
	fmt.Println()
	fmt.Printf("客户号\t姓名\t性别\t年龄\t电话\t邮箱")
	fmt.Println()
	for _,v := range *css {
		fmt.Printf("%d\t%s\t%s\t%d\t%s\t%s",v.ID,v.Name,v.Gender,v.Age,v.Tel,v.Mail)
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("--------------------列表--完成--------------------")
	fmt.Println()
	mi.SubCommand(css)
}


func (mi *MainInterface) RenderExit(css *[]customer.Customer){
	fmt.Println("--------------------退出  软件--------------------")
	fmt.Println()
	mi.ExitCommand(css)

	fmt.Println()
	fmt.Println("--------------------退出  完成--------------------")
	fmt.Println()
}

func(mi *MainInterface) MainCommand(css *[]customer.Customer){
	var iB byte
	fmt.Println("--------------------请选择（1-5）：--------------------")
	fmt.Scanf("%d\n",&iB)

	switch iB {
	case 1:
		mi.RenderAdd(css)
	case 2:
		mi.RenderModify(css)
	case 3:
		mi.RenderDelete(css)
	case 4:
		mi.RenderList(css)
	case 5:
		mi.RenderExit(css)
	default:
		fmt.Println("输入有误，请重新选择")
		mi.RenderMain()
		mi.MainCommand(css)
	}
}

func(mi *MainInterface) SubCommand(css *[]customer.Customer){
	var iB byte
	fmt.Println("-----------------输入1返回主界面-----------------")
	fmt.Scanf("%d\n",&iB)

	switch iB {
	case 1:
		mi.RenderMain()
		mi.MainCommand(css)
	default:
		fmt.Println("输入有误，请重新输入")
		mi.SubCommand(css)
	}
}


func(mi *MainInterface) ExitCommand(css *[]customer.Customer){
	var iB byte
	fmt.Println("-----------------确定退出系统吗(1:返回主界面  2：退出系统)-----------------")
	fmt.Scanf("%d\n",&iB)

	switch iB {
	case 1:
		mi.RenderMain()
		mi.MainCommand(css)
	case 2:
	default:
		fmt.Println("输入有误，请重新输入")
		mi.SubCommand(css)
	}
}