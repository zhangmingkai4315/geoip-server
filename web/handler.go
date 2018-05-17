package web

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/zhangmingkai4315/geoip-server/cache"
)

func httpIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, apiHelp())
	return
}

// httpIPQueryHandler will receive one ipaddress then return  referenced ipinfo
// lang only support en and zh-cn right now.
func httpIPQueryHandler(lang string) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		ipaddress := ps.ByName("ipaddress")
		if !IsIPv4(ipaddress) {
			returnBadRequestResponse(w, "IP address not valid")
			return
		}
		ipinfo := &cache.IPInfo{IP: ipaddress}
		err := ipinfo.GetIPInfo(lang)
		if err != nil {
			log.Printf("[Error] IP:%s Error:%s\n", ipaddress, err)
			returnServerFailResponse(w, err.Error())
		}
		returnStatusOkResponse(w, ipinfo)
	}
}

// httpIPListQueryHandler will receive posted iplist data and return all referenced ipinfo list data
// lang only support en and zh-cn right now.
func httpIPListQueryHandler(lang string) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var postList []string
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			returnBadRequestResponse(w, "post data error or oversize")
			return
		}
		if err := r.Body.Close(); err != nil {
			returnServerFailResponse(w, "can't close http reader")
			return
		}

		if err := json.Unmarshal(body, &postList); err != nil {
			returnBadRequestResponse(w, "post data error or json marshal fail")
			return
		}
		postListStruct := validateIPAddressFromList(postList)
		postListStruct.GetIPInfo(lang)
		returnStatusOkResponse(w, &postListStruct)
	}
}

func httpIPASNQueryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ipaddress := ps.ByName("ipaddress")
	if !IsIPv4(ipaddress) {
		returnBadRequestResponse(w, "IP address not valid")
		return
	}
	ipinfo := &cache.IPInfo{IP: ipaddress}
	err := ipinfo.GetASNInfo()
	if err != nil {
		log.Printf("[Error] IP:%s Error:%s\n", ipaddress, err)
		returnServerFailResponse(w, err.Error())
	}
	returnStatusOkResponse(w, ipinfo)
}

func httpIPASNListQueryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var postList []string
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		returnBadRequestResponse(w, "post data error or oversize")
		return
	}
	if err := r.Body.Close(); err != nil {
		returnServerFailResponse(w, "can't close http reader")
		return
	}

	if err := json.Unmarshal(body, &postList); err != nil {
		returnBadRequestResponse(w, "post data error or json marshal fail")
		return
	}
	postListStruct := validateIPAddressFromList(postList)
	postListStruct.GetASNInfo()
	returnStatusOkResponse(w, &postListStruct)
}

func httpIPAllQueryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ipaddress := ps.ByName("ipaddress")
	if !IsIPv4(ipaddress) {
		returnBadRequestResponse(w, "IP address not valid")
		return
	}
	ipinfo := &cache.IPInfo{IP: ipaddress}
	err := ipinfo.GetAllInfo()
	if err != nil {
		log.Printf("[Error] IP:%s Error:%s\n", ipaddress, err)
		returnServerFailResponse(w, err.Error())
	}
	returnStatusOkResponse(w, ipinfo)
}

func httpIPAllListQueryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var postList []string
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		returnBadRequestResponse(w, "post data error or oversize")
		return
	}
	if err := r.Body.Close(); err != nil {
		returnServerFailResponse(w, "can't close http reader")
		return
	}
	if err := json.Unmarshal(body, &postList); err != nil {
		returnBadRequestResponse(w, "post data error or json marshal fail")
		return
	}
	postListStruct := validateIPAddressFromList(postList)
	postListStruct.GetAllInfo()
	returnStatusOkResponse(w, &postListStruct)
}
