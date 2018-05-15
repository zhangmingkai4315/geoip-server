package web

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func apiHelp() string {
	help := `

IP to Location query api server
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
	fmt.Fprintf(w, "query ip is :%s", ps.ByName("ipaddress"))
	// ipaddress := ps.ByName("address")
	// ipinfo := &cache.IPInfo{IP: ipaddress}
	// infostring, err := ipinfo.GetInfo()
	// if err != nil {

	// }
	return
}
