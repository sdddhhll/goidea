package main

import (
	"flag"
	"os"
)

func main(){
	var command string
	var config Configurtor;
	var cmdMap map[string]Commandor
	cmdMap = make(map[string]Commandor, 0)
	GenCommandorList(&cmdMap,&config)

	args := os.Args
	if len(args) == 1 {
		//说明没有跟任何参数
		RunControllerWithNoArgs("help",&cmdMap,&config)
	}else if args[1] == "-h" || args[1] == "-help" ||args[1] == "help" || args[1] == "?" {
		RunController("help",&cmdMap,&config)
	}else if args[1] == "about"{
		RunController("about",&cmdMap,&config)
	}else {
		flag.StringVar(&command, "c", "help", "dhl -c COMMAND")
		flag.Parse()
		RunController(command, &cmdMap, &config)
	}
}

