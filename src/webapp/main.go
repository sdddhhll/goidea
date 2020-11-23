package main

import "fmt"
import "net/http"
func main() {
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8081",nil)
}

func handler(writer http.ResponseWriter,req *http.Request){
	fmt.Fprintln(writer,"dhl")
}


