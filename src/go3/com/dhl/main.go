package main

import (
	"encoding/json"
	"flag"
	"fmt"
)

type User struct {
	Name string
	Age int
	Addr string
	Gender string
}
func main21(){
	var str string = "{\"Name\":\"张三\",\"Age\":44,\"Gender\":\"男\",\"Addr\":\"济南市\"}"
	var user User
	e := json.Unmarshal([]byte(str), &user)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(user)//{张三 44 济南市 男}
}

//json反序列化为map
func main22(){
	var str string = "{\"name\":\"张三\",\"age\":44,\"gender\":\"男\",\"addr\":\"济南市\"}"
	var vmap map[string]interface{}
	e := json.Unmarshal([]byte(str), &vmap)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(vmap)
}

//json反序列化为切片
func main33(){
	var str string = "[{\"Name\":\"张三\",\"Age\":\"44\",\"Gender\":\"男\",\"Addr\":\"济南市\"},{\"Name\":\"张三2\",\"Age\":\"442\",\"Gender\":\"男2\",\"Addr\":\"济南市2\"}]"
	var list []map[string]interface{}
	json.Unmarshal([]byte(str),&list)
	fmt.Println(list)
}

func main(){
	var name string
	var age int64
	var note string
	flag.StringVar(&name,"n","","不能为空")
	flag.StringVar(&note,"t","","可以为空")
	flag.Int64Var(&age,"a",0,"不能为空")
	flag.Parse()
	fmt.Println("name=",name)
	fmt.Println("age=",age)
	fmt.Println("note=",note)

}





