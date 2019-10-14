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
	ip := client.GetIPAddress(&r.Header)
	fmt.Fprintf(w, "Your public ip is: %v\n", ip)

	if ip != ""{
		res, err := client.GetIPGeoInfo(ip)
		if err != nil{
			return
		}

		fmt.Fprintf(w, "Geographical location information is: %#v\n", res.Data)
	}
}

func main() {
	http.HandleFunc("/", IndexPage)
	http.ListenAndServe(":8080", nil)
}

