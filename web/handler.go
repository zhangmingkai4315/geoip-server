package web

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/zhangmingkai4315/geoip-server/cache"
)

func apiHelp() string {
	help := `IP To Location Query Api Server
-------------------------------
/api/geoip2/:address  
	- method get 
	- return json data

/api/geoip2/ 
	- method post
	- example ['1.2.4.8','8.8.8.8'...]
	- max 1000 items in one query
	- return 
		- [{...},{...},...]
	`
	return help
}

func httpIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, apiHelp())
	return
}

func httpIPQueryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// fmt.Fprintf(w, "query ip is :%s", ps.ByName("ipaddress"))
	ipaddress := ps.ByName("ipaddress")
	ipinfo := &cache.IPInfo{IP: ipaddress}
	err := ipinfo.GetInfo("en")
	if err != nil {
		log.Printf("[Error] IP:%s Error:%s\n", ipaddress, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(ipinfo)
}
