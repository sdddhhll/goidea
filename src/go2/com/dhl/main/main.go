package main

import (
	"go2/com/dhl/MainInterface"
	"go2/com/dhl/customer"
)

func main(){
	var mi MainInterface.MainInterface
	var css []customer.Customer
	mi.RenderMain()
	mi.MainCommand(&css)
}
