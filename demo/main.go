package main

import (
	"fmt"
	"net/http"
	"github.com/csxuejin/iptool"
)

var (
	client = iptool.NewIPTool()
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your public ip is: %v\n", iptool.GetIPAddress(&r.Header))
}

func main() {
	http.HandleFunc("/", IndexPage)
	http.ListenAndServe(":8080", nil)
}

