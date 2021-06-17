package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Inside HelloServer handler")
	// fmt.Fprint和fmt.Fprintf都是用来写入http.ResponseWriter的不错的函数（他们实现了io.Writer）
	fmt.Fprintf(w, "hello, " + r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
