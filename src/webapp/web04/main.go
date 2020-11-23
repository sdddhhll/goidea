package main

import (
	"fmt"
	"net/http"
)

func main() {
	var myHandler = MyHandler{}
	http.Handle("/zmx",&myHandler)
	http.ListenAndServe(":8080",nil)
}

type MyHandler struct {

}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"单独处理请求....")
}
