package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

/**
	运行程序
 */
func(com *Commandor) RunProgram(){
	PrintSeparateStar()
	fmt.Println()
	fmt.Println("\t\t命令：",com.Cmdstr)
	fmt.Println("\t\t程序：",com.Cmddesc)
	fmt.Println("\t\t路径：",com.Exepath)
	fmt.Println("\t\t参数：",com.Datapath)
	fmt.Println("\t\t状态：","正在启动...")
	var command *exec.Cmd
	if com.Datapath == "" {
		command = exec.Command(com.Exepath)
	}else {
		command = exec.Command(com.Exepath, com.Datapath)
	}
	start := command.Start()
	if start == nil {
		fmt.Println("\t\t结果：","启动成功!!!")
	}else {
		fmt.Println("\t\t日志：启动失败",start)
	}
	PrintSeparateStar()
	fmt.Println()
}
/**
生成集合
 */
func GenCommandorList(cmdMap *map[string]Commandor,conf *Configurtor){
	file, e := os.Open(GetRunConfig())
	//file, e := os.Open("E:\\inspur\\ideas\\learnIDEA\\goidea\\src\\shortcutor\\dhl.conf")

	if e != nil {
		fmt.Println("出现错误:",e)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var comLines string
	for{
		cmdLine, readE := reader.ReadString('\n')
		cmdLine = strings.Replace(cmdLine,"\n","",1)
		cmdLine = strings.Replace(cmdLine,"\r","",1)
		GenCommandor(&cmdLine,cmdMap)
		comLines += cmdLine
		if readE == io.EOF {
			conf.Content=comLines
			return
		} else if readE != nil {
			fmt.Println("出现错误:",readE)
			continue
		}
	}
}
/**
根据每行的数据生成对象
 */
func GenCommandor(cmdLine *string,cmdMap *map[string]Commandor) {
	defer  func() {
		err := recover()
		if err != nil {
			fmt.Println("处理出现错误，请检查您的配置文件",err)
		}
	}()
	splits := strings.Split(*cmdLine, "||")
	keys := strings.Split(splits[0], ",")
	for _,key := range keys{
		var com Commandor=Commandor{key,splits[1],splits[2],splits[3]}
		(*cmdMap)[key]=com
	}
	*cmdLine = splits[0] +">>>>>>>>>>>>>>>"+ splits[1]+"\r\n"
}
/**
获取当前目录
 */
func GetRunDir() string{
	s, _ := os.Executable()
	dir := filepath.Dir(s)
	return dir
}1
/**
生成配置文件
 */
func GetRunConfig() string{
	return GetRunDir()+"/dhl.conf"
}

/**
显示帮助文件
 */
func ShowHelp(conf *Configurtor){
	PrintSeparateWithText("系统帮助")
	fmt.Print(conf.Content)
	PrintSeparateWithText("帮助结束")
}

func ShowDesc(){
	PrintSeparateWithText("使用说明")
	fmt.Println("您需要修改dhl.conf文件,实现自定义快捷启动")
	fmt.Println("您可以配置环境变量，在命令行中输入：dhl -c 自定义命令  来快速启动")
	fmt.Println("您可以双击直接运行本exe")
	fmt.Println("您可以输入[dhl] help或者[dhl] about随时查看帮助或者作者联系信息")
	fmt.Println("\t\t\t\t\t作者：丁怀雷")
	PrintSeparateWithText("说明结束")
}

func AboutProgram(){
	PrintSeparateWithText("作者信息")
	fmt.Println("\t\t程序：快捷启动")
	fmt.Println("\t\t作者：丁怀雷	")
	fmt.Println("\t\t邮箱：ddhhll.love@163.com")
	fmt.Println("\t\t说明：基于golang，快捷启动")
	PrintSeparateWithText("信息结束")
}

func RunProgram(command string,cmdMap *map[string]Commandor){
	var com Commandor = (*cmdMap)[command]
	com.RunProgram()
}
