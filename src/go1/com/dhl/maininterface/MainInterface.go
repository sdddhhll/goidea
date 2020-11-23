package maininterface

import (
	"fmt"
	"go1/com/dhl/FamilyAccount"
)

type MainInterface struct {

}


func (mi *MainInterface)RenderMain(){
	fmt.Println("---------------主界面---------------")
	fmt.Println()
	fmt.Print(`
		1：收支明细
		2：登记收入
		3：登记支出
		4：退出系统
`)
	fmt.Println()
	fmt.Println("------------------------------------")
}

func (mi *MainInterface) RenderDetail(fas *[]FamilyAccount.FamilyAccount){
	fmt.Println("---------------账户明细--------------")
	fmt.Println()
	fmt.Printf("收支说明\t账户余额\t本次收支\t收支备注")
	fmt.Println()
	for _,val:= range *fas {
		fmt.Printf("%s\t%f\t%f\t%s",val.RevenueName,val.TotleAccount,val.RevenueAccount,val.RevenueNote)
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("------------------------------------")
}

func (mi *MainInterface) GetDetailCommand(fas *[]FamilyAccount.FamilyAccount){
	fmt.Println("输入1返回主界面")
	var iByte byte
	fmt.Scanf("%d\n",&iByte)
	switch iByte {
	case 1:
		mi.RenderMain()
		mi.GetMainCommand(fas)
	default:
		fmt.Println("命令输入有误,请重新输入")
		mi.GetDetailCommand(fas)
	}
}

func (mi *MainInterface) GotIncome(fas *[]FamilyAccount.FamilyAccount){
	var RevenueAccount float64
	var RevenueNote string
	fmt.Println("输入收入金额")
	fmt.Scanf("%f",&RevenueAccount)
	fmt.Println("请输入收入说明")
	fmt.Scanf("%s",&RevenueNote)
	var fa FamilyAccount.FamilyAccount
	lastAccount := fa.GetPreFamilyAccount(fas)
	fa.Income(RevenueAccount,RevenueNote,&lastAccount)
	*fas = append(*fas, fa)
	fmt.Println("添加成功")
	mi.RenderMain()
	mi.GetMainCommand(fas)
}

func (mi *MainInterface) GotOutcome(fas *[]FamilyAccount.FamilyAccount){
	var RevenueAccount float64
	var RevenueNote string
	fmt.Println("输入支出金额")
	fmt.Scanf("%f",&RevenueAccount)
	fmt.Println("请输入支出说明")
	fmt.Scanf("%s",&RevenueNote)
	var fa FamilyAccount.FamilyAccount
	lastAccount := fa.GetPreFamilyAccount(fas)
	fa.Outcome(RevenueAccount,RevenueNote,&lastAccount)
	*fas = append(*fas, fa)
	fmt.Println("添加成功")
	mi.RenderMain()
	mi.GetMainCommand(fas)
}

func (mi *MainInterface) GotExit(fas *[]FamilyAccount.FamilyAccount){
	var comm string
	fmt.Println("确定退出吗？(y/n)")
	fmt.Scanf("%s",&comm)
		if comm == "y" {
			fmt.Println("退出成功")
		} else{
			mi.GetMainCommand(fas)
		}
}

func (mi *MainInterface) GetMainCommand(fas *[]FamilyAccount.FamilyAccount){
	mi.RenderMain()
	fmt.Println()
	fmt.Println("请输入要执行的命令")
	var iByte byte
	fmt.Scanf("%d\n",&iByte)
	switch iByte {
	case 1:
		mi.RenderDetail(fas)
		mi.GetDetailCommand(fas)
	case 2:
		mi.GotIncome(fas)
	case 3:
		mi.GotOutcome(fas)
	case 4:
		mi.GotExit(fas)
	default:
		fmt.Println("输入有误请重新输入")
		mi.GetMainCommand(fas)
	}
}
