package main
import "os"
import "fmt"
import "io/ioutil"
import "encoding/json"
type  Monster struct{
	Name string
	Age int
	Skill string
} 

func(mon *Monster) Store(){
	file,error := os.OpenFile("c:/777.txt",os.O_CREATE|os.O_RDWR,0666)
	if error != nil {
		fmt.Println("打开文件出现错误")
	}
	b,e := json.Marshal(mon)
	if e != nil {
		fmt.Println("序列化出错")
	} 
	_,err := file.Write(b)
	if err != nil {
		fmt.Println("写入文件出错")
	}
	file.Close()
}

func (mon *Monster) ReStore(){
	b,e := ioutil.ReadFile("c:/777.txt")
	if e != nil {
		fmt.Println("打开文件出现错误")
	}
	err := json.Unmarshal(b,mon)
	if err != nil {
		fmt.Println("转换为对象出现错误")
	}
}
