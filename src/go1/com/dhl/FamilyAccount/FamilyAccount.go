package FamilyAccount

type FamilyAccount struct {
	RevenueName string
	TotleAccount float64
	RevenueAccount float64
	RevenueNote string
}

/**
收入
 */
func (fa *FamilyAccount) Income (RevenueAccount float64,RevenueNote string,lastFa *FamilyAccount)  {
	fa.RevenueName="收入"
	fa.TotleAccount=lastFa.TotleAccount + RevenueAccount
	fa.RevenueAccount = RevenueAccount
	fa.RevenueNote = RevenueNote
}

/**
支出
 */

func (fa *FamilyAccount) Outcome (RevenueAccount float64,RevenueNote string,lastFa *FamilyAccount)  {
	fa.RevenueName="支出"
	fa.TotleAccount=lastFa.TotleAccount - RevenueAccount
	fa.RevenueAccount = RevenueAccount
	fa.RevenueNote = RevenueNote
}

/**
获取最后一次账户明细
 */
func (fa *FamilyAccount) GetPreFamilyAccount(fas *[]FamilyAccount) FamilyAccount{
	if len(*fas) ==0 {
		return FamilyAccount{"",0,0,""}
	}else {
		return (*fas)[len(*fas) - 1]
	}
}

