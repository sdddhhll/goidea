package main

import (
	"fmt"
)

func RunController(command string,cmdMap *map[string]Commandor,config *Configurtor){
	switch command {
		case "help":
			ShowHelp(config)
		case "about":
			ShowDesc()
			AboutProgram()
		default:
			RunProgram(command,cmdMap)
	}
}

func RunControllerWithNoArgs(command string,cmdMap *map[string]Commandor,config *Configurtor){
	RunController(command,cmdMap,config)
	for {
		fmt.Println("请输入要运行的程序(quit|exit|bye退出本程序,about查看说明):")
		fmt.Scanf("%s\n",&command)
		switch command {
		case "":
		case "quit":
			fmt.Println("再见...")
			return
		case "exit":
			fmt.Println("再见...")
			return
		case "bye":
			fmt.Println("再见...")
			return
		default:
			RunController(command,cmdMap,config)
		}
	}

}
