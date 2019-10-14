package iptool

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	headerXForwardedFor = "X-Forwarded-For"
	headerXRealIP       = "X-Real-Ip"

	TaobaoIPAPIHost = "http://ip.taobao.com"
	TaobaoIPAPIURL  = "/service/getIpInfo.php?ip=%s"
)

type Client struct {
	*http.Client
}

func NewIPTool() *Client {
	return &Client{
		Client : &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) GetIPAddress(header *http.Header) string {
	for _, h := range []string{headerXForwardedFor, headerXRealIP} {
		for _, ip := range strings.Split(header.Get(h), ",") {
			realIP := net.ParseIP(strings.TrimSpace(ip))
			if !realIP.IsGlobalUnicast() || c.IsPrivateSubnet(realIP) {
				continue
			}
			return ip
		}
	}

	return ""
}

func (c *Client) GetIPGeoInfo(ip string) (r *GetIPGeoOutput, err error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(TaobaoIPAPIHost+TaobaoIPAPIURL, ip), nil)
	if err != nil {
		log.Printf("err is %v\n", err)
		return
	}

	resp, err := c.Do(req)
	if err != nil {
		log.Printf("c.client.Do(): %v\n", err)
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ioutiol.ReadAll(): %v\n", err)
		return
	}

	if err = json.Unmarshal(data, &r); err != nil {
		log.Printf("json.Unmarshal(): %v\n", err)
		return
	}
	if r.Code != 0 {
		err = errors.New(fmt.Sprintf("response with wrong status codd %v", r.Code))
	}

	return
}

func (c *Client) IsPrivateSubnet(ipAddress net.IP) bool {
	if ipCheck := ipAddress.To4(); ipCheck != nil {
		for _, r := range privateRanges {
			if r.contains(ipAddress) {
				return true
			}
		}
	}

	return false
}

func (r ipRange) contains(ipAddress net.IP) bool {
	if bytes.Compare(ipAddress, r.start) >= 0 && bytes.Compare(ipAddress, r.end) < 0 {
		return true
	}
	return false
}

type ipRange struct {
	start, end net.IP
}

var privateRanges = []ipRange{
	{
		start: net.ParseIP("10.0.0.0"),
		end:   net.ParseIP("10.255.255.255"),
	},
	{
		start: net.ParseIP("100.64.0.0"),
		end:   net.ParseIP("100.127.255.255"),
	},
	{
		start: net.ParseIP("172.16.0.0"),
		end:   net.ParseIP("172.31.255.255"),
	},
	{
		start: net.ParseIP("192.0.0.0"),
		end:   net.ParseIP("192.0.0.255"),
	},
	{
		start: net.ParseIP("192.168.0.0"),
		end:   net.ParseIP("192.168.255.255"),
	},
	{
		start: net.ParseIP("198.18.0.0"),
		end:   net.ParseIP("198.19.255.255"),
	},
}
