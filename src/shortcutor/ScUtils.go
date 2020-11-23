package main

import "fmt"

func PrintSeparateStar(){
	PrintSeparate("*")
}

func PrintSeparatePoint(){
	PrintSeparate(".")
}

func PrintSeparate(str string){
	PrintSeparateWithLength(str,69)
}

func PrintSeparateWithLength(str string,col int){
	for i:=0;i<col;i++{
		fmt.Print(str)
	}
}

func PrintSeparateWithText(text string){
	PrintSeparateWithLength("*",30)
	fmt.Print(text)
	PrintSeparateWithLength("*",30)
	fmt.Println()
}



