package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main1(){
	//打开文件
	f,e := os.Open("E:\\inspur\\ideas\\learnIDEA\\goidea\\src\\go3\\com\\dhl\\data.json")
	if e!=nil {
		fmt.Println("打开文件时，出现错误")
	}

	fmt.Printf("file=%v",f)
	err := f.Close()
	if err !=nil {
		fmt.Println("关闭文件时，出现错误")
	}


}

func main2(){

	//当函数退出时，要及时的关闭file
	file,err := os.Open("E:\\inspur\\ideas\\learnIDEA\\goidea\\src\\go3\\com\\dhl\\data.json")
	if err != nil {
		fmt.Println("打开文件时，出现错误：",err)
	}
	//创建一个  *reader  带缓冲的

	reader := bufio.NewReader(file)
	var con string
	for { //io.EOF :表示
		content, reerr := reader.ReadString('\n')
		if reerr == nil {//
			con += content
		}else  {
			break
		}
	}
	fmt.Println(con)
	defer file.Close()//最后需要关闭defer，防止内存泄漏
}


func main3() {
	//下列文件会自动打开和关闭文件，无需我们手动处理，这样显得比较简洁，Open与Close被封装到了ReadFile内部
	bytes, err1 := ioutil.ReadFile("E:\\inspur\\ideas\\learnIDEA\\goidea\\src\\go3\\com\\dhl\\data.json")
	if err1 != nil {
		fmt.Println("读取文件时，出现错误：", err1)
	}
	fmt.Printf("%s",bytes)
	fmt.Printf("%v",bytes) //输出数字数组（字节数组）
}




//文件打开模式



func main() {
	s, _ := os.Executable()
	dir := filepath.Dir(s)
	fmt.Println(dir)
}



func main5() {
	file, e := os.OpenFile("d:/abc.txt", os.O_CREATE|os.O_RDWR, 0666)//0666win下面没有意义
	if e != nil {
		fmt.Println("读取文件出现错误",e)
		return
	}
	writer := bufio.NewWriter(file)//因为Writer是带缓存的，因此写入的数据还没有落盘因此在调用writestring方法时 其实内容是先写入到缓存
	//所以需要调用Flush方法，将缓冲的数据真正的写入到文件中，否则文件中会没有数据！！！
	for i:=0;i<5;i++{
		writer.WriteString("Good,Golday!\n")
	}
	writer.Flush()
	fmt.Println("写入完成...")
	defer file.Close()

}


func main6() {
	file, e := os.OpenFile("d:/abc.txt", os.O_WRONLY, 0666)
	if e!=nil{
		fmt.Println("读取文件出现错误",e)
		return
	}
	writer := bufio.NewWriter(file)
	for i:=0;i<10;i++{
		writer.WriteString("您好 世界")
	}

	writer.Flush()
	defer file.Close()

}


func main7() {
	file, e := os.OpenFile("d:/abc.txt", os.O_WRONLY|os.O_APPEND, 0666)
	if e!=nil{
		fmt.Println("读取文件出现错误",e)
		return
	}
	writer := bufio.NewWriter(file)
		writer.WriteString("hello,everyone")

	writer.Flush()
	defer file.Close()

}



func main8() {
	file, e := os.OpenFile("d:/abc.txt", os.O_RDONLY, 0666)
	if e!=nil{
		fmt.Println("读取文件出现错误",e)
		return
	}
	//reader := bufio.NewReader(file)
	//var str string = ""
	//for   {
	//	s, e := reader.ReadString('\n')
	//	if e != nil {
	//		fmt.Println(e)
	//		break
	//	}
	//	fmt.Println(s)
	//	str += s
	//}
	//fmt.Println(str)

	bytes, e := ioutil.ReadFile("d:/abc.txt")
	fmt.Printf("%s",bytes)

	defer file.Close()

}




func main9() {

	//当函数退出时，要及时的关闭file
	file,err := os.Open("E:\\inspur\\ideas\\learnIDEA\\goidea\\src\\go3\\com\\dhl\\data.json")
	if err != nil {
		fmt.Println("打开文件时，出现错误：",err)
	}
	//创建一个  *reader  带缓冲的

	reader := bufio.NewReader(file)
	var con string
	for { //io.EOF :表示
		content, reerr := reader.ReadString('\n')
		if reerr == nil {//
			con += content
		}else  {
			con += content//这一行必须，否则会缺少最后一行
			break
		}
	}
	fmt.Println(con)
	defer file.Close()//最后需要关闭defer，防止内存泄漏

}




func main10() {

	//当函数退出时，要及时的关闭file
	file,err := os.Open("d:/abc.txt")
	if err != nil {
		fmt.Println("打开文件时，出现错误：",err)
	}
	//创建一个  *reader  带缓冲的

	reader := bufio.NewReader(file)
	var con string
	for { //io.EOF :表示
		content, reerr := reader.ReadString('\n')
		if reerr == nil {//
			con += content
		}else  {
			con += content//这一行必须，否则会缺少最后一行
			break
		}
	}
	fmt.Println(con)
	defer file.Close()//最后需要关闭defer，防止内存泄漏

}



func main11(){
	file, e := os.OpenFile("d:/abc.txt", os.O_RDWR, 06666)
	if e != nil {
		fmt.Println("读取文件发生错误:" , e)
		return
	}

	reader := bufio.NewReader(file)
	var ss string
	for{
		s, erre := reader.ReadString('\n')
		if erre !=nil {
			ss += s
			break
		}else {
			ss += s
		}
	}
	fmt.Println(ss)
	writer := bufio.NewWriter(file)
	i, e := writer.WriteString("\r\n这是一个清空文件写入新数据的测试3")
	if e !=nil {
		fmt.Println("错误")
		return
	}
	writer.Flush()
	fmt.Println(i)
	defer file.Close()

}


func main12(){
	fmt.Println("开始")
	bytes, e := ioutil.ReadFile("d:/abc.txt")
	if e != nil {
		fmt.Println("读取文件发生错误",e)
		return
	}
	e2 := ioutil.WriteFile("d:/abc1.txt",bytes,0666)
	if e2 != nil {
		fmt.Println("写入文件发生错误",e2)
		return
	}
	fmt.Println("结束")

}

func main13(){
	file, e := os.Open("d:/abc.txt")
	if e != nil {
		fmt.Println("错误，读取文件出现问题")
		return
	}
	_, e1 := file.Stat()
	if e1 != nil {
		fmt.Println("文件不存在")
	}else {
		fmt.Println("文件存在")
	}
}

func main14(){
	srcFile, e := os.Open("d:/abc.txt")
	if e != nil {
		fmt.Println("读取文件出现错误...")
		return
	}
	srcReader := bufio.NewReader(srcFile)

	distFile, e := os.OpenFile("c:/321.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if e != nil {
		fmt.Println("读取写入文件出现错误...")
		return
	}
	disWriter := bufio.NewWriter(distFile)
	_, e1 := io.Copy(disWriter, srcReader)
	disWriter.Flush() //如果没有这句话 复制后的
	if e1!= nil {
		fmt.Println("失败")
	}
	defer srcFile.Close()
	defer distFile.Close()

}
func main15(){
	srcFile, e := os.Open("C:\\Users\\sdddhhll\\Pictures\\dhl2.jpg")
	if e != nil {
		fmt.Println("读取文件出现错误...")
		return
	}
	srcReader := bufio.NewReader(srcFile)

	distFile, e := os.OpenFile("c:/321.jpg", os.O_WRONLY|os.O_CREATE, 0666)
	if e != nil {
		fmt.Println("读取写入文件出现错误...")
		return
	}
	disWriter := bufio.NewWriter(distFile)
	_, e1 := io.Copy(disWriter, srcReader)
	disWriter.Flush() //如果没有这句话 复制后的
	if e1!= nil {
		fmt.Println("失败")
	}
	defer srcFile.Close()
	defer distFile.Close()

}

func main16(){
	file, e := os.Open("D:\\abc.txt")
	if e!=nil {
		fmt.Println("错误...")
		return
	}
	reader := bufio.NewReader(file)
	var enCount int64 = 0
	var enNum int64 = 0
	var enSpace int64 = 0
	var other int64 = 0

	for {
	s, e3 := reader.ReadString('\n')
	for _,v := range s {
		fmt.Println(v)
		switch  {
		case v >= '0' && v<='9':
			enNum ++
		case v>='a' && v<='z':
			enCount++
		case v == ' ':
			enSpace ++
		default:
			other ++
		}
	}
		if e3 !=nil {
			break
		}
	}


	fmt.Println("英文字符：",enCount)
	fmt.Println("数字字符：",enNum)
	fmt.Println("空格字符：",enSpace)
	fmt.Println("其他字符：",other)
}

func main111(){
	strings := os.Args
	for index,str := range strings {
		fmt.Println("%d是%s",index,str)
	}
}

type Student struct {
	Name string `json:"username"`
	Age int64 `json:"userage"`
	Gender string
	Node string
}

func main222(){
	student := Student{
		Name:   "张三",
		Age:    10,
		Gender: "男",
		Node:   "这是一个测试的结构体",
	}
	bytes, e := json.Marshal(&student)
	if e!=nil {
		fmt.Println("错误")
		return
	}
	//fmt.Printf("%v",bytes)
	fmt.Printf("%s",bytes)
}



func main333(){
	//将map进行序列化

	var a map[string]interface{}
	//使用map之前需要make
	a = make(map[string]interface{})
	a["name"]="李四"
	a["age"]=12
	a["addr"]="高新区"
	a["node"]="备注"
	bytes, e := json.Marshal(a)
	if e != nil {
		fmt.Println("ddd")
	}

	fmt.Printf("%s",bytes)
}



func main444(){
	var a []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "张三"
	m1["age"] = 66

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"]="王五"
	m2["age"]=33

	a = append(a,m1)
	a = append(a,m2)
	bytes, e := json.Marshal(a)
	if e != nil {
		fmt.Println("ddd")
	}

	fmt.Printf("%s",bytes)
}










