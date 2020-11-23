package main

import "fmt"
import "net/http"
func main() {
	var myHandler = MyHandler{}
	http.Handle("/",&myHandler)
	http.ListenAndServe(":8882",nil)
}

type  MyHandler struct {

}

func (m *MyHandler) ServeHTTP (writer http.ResponseWriter,req *http.Request){
	fmt.Fprintln(writer,"正在处理您的请求...")
}