package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"正在处理请求")
}

func main() {
		http.HandleFunc("/dhl",handler)
		http.ListenAndServe(":8080",nil)
}
