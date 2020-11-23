package main

import (
	"go1/com/dhl/FamilyAccount"
	"go1/com/dhl/maininterface"
)


func main(){
var mi maininterface.MainInterface
var fas []FamilyAccount.FamilyAccount = make([]FamilyAccount.FamilyAccount,0)
mi.GetMainCommand(&fas)
}











