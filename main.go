package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	//往文件里面写东西的函数
	b, _ := ioutil.ReadFile("./hello.txt")
	_, _ = fmt.Fprintln(w, string(b))
}

func main() {
	http.HandleFunc("/hello", SayHello)
	err := http.ListenAndServe(":9090", nil) //这里冒号别搞丢了
	if err != nil {
		fmt.Printf("http serve failed: %v\n", err)
	}
}
