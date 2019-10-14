package iptool

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

var (
	client = NewIPTool()
)

func Test_GetIPAddress(t *testing.T) {
	assertion := assert.New(t)

	publicIP := "106.38.115.25"
	privateIP := "127.0.0.1"
	//  get ip form X-Forwarded-For header
	{
		header := &http.Header{}
		header.Set(headerXForwardedFor, strings.Join([]string{publicIP, privateIP}, ","))
		assertion.Equal(publicIP, client.GetIPAddress(header))
	}

	// if there are multi public ips, only get the first one
	{
		newPublicIP := "106.38.115.26"
		header := &http.Header{}
		header.Set(headerXForwardedFor, strings.Join([]string{newPublicIP, publicIP}, ","))
		assertion.Equal(newPublicIP, client.GetIPAddress(header))
	}

	//  get ip form X-Real-Ip header
	{
		header := &http.Header{}
		header.Set(headerXRealIP, strings.Join([]string{publicIP, privateIP}, ","))
		assertion.Equal(publicIP, client.GetIPAddress(header))
	}

	// if there is no public ip
	{
		header := &http.Header{}
		header.Set(headerXRealIP, strings.Join([]string{privateIP, privateIP}, ","))
		assertion.Empty(client.GetIPAddress(header))
	}
}

func Test_GetIPGeoInfo(t *testing.T) {
	assertion := assert.New(t)

	r, err := client.GetIPGeoInfo("106.38.115.25")
	assertion.Nil(err)
	assertion.Equal(r.Data.Country, "中国")
	assertion.Equal(r.Data.Region, "北京")
	assertion.Equal(r.Data.City, "北京")
	fmt.Printf("res is %#v\n", r)
}
