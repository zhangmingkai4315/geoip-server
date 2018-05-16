package web

import (
	"encoding/json"
	"net"
	"net/http"
	"strings"

	"github.com/zhangmingkai4315/geoip-server/cache"
)

// Response define the basic http json response
type Response struct {
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
}

func returnStatusOkResponse(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(Response{
		Error:  "",
		Data:   data,
		Status: 200,
	})
}

func returnBadRequestResponse(w http.ResponseWriter, err string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(Response{
		Error:  err,
		Status: 400,
	})
}

func returnServerFailResponse(w http.ResponseWriter, err string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(Response{
		Error:  err,
		Status: 500,
	})
}

//
func validateIPAddressFromList(iplist []string) cache.IPInfoList {
	var ipListInfo cache.IPInfoList
	for _, ip := range iplist {
		ipinfo := &cache.IPInfo{IP: ip}
		if !IsIPv4(ip) {
			ipinfo.Error = "valid error"
		}
		ipListInfo = append(ipListInfo, ipinfo)
	}
	return ipListInfo
}

// IsIPv4 check if the string is an IP version 4.
func IsIPv4(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && strings.Contains(str, ".")
}

// IsIPv6 check if the string is an IP version 6.
func IsIPv6(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && strings.Contains(str, ":")
}

func apiHelp() string {
	help := `IP To Location Query Api Server
-------------------------------
/api/geoip2/:address (default english)
	- method get 
	- return json data
/api/geoip2/:address/en
	- method get 
	- return json data
/api/geoip2/:address/zh-ch
	- method get 
	- return json data

/api/geoip2/iplist (default english)
	- method post
	- example ['1.2.4.8','8.8.8.8'...]
	- max 1000 items in one query
	- return list object
/api/geoip2/iplist/zh-ch
	- method post
	- example ['1.2.4.8','8.8.8.8'...]
	- max 1000 items in one query
	- return list object
/api/geoip2/iplist/en
	- method post
	- example ['1.2.4.8','8.8.8.8'...]
	- max 1000 items in one query
	- return list object
	`
	return help
}
