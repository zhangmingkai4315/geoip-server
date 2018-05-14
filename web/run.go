package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func apiHelp() string {
	help := `
IP to Location query api server
-------------------------------
/geoip2/:address  
	- method get 
	- return json data

/geoip2/ 
	- method post
	- example ['1.2.4.8','8.8.8.8'...]
	- max 1000 items in one query
	- return 
		- [{...},{...},{...},null...]

	`
	return help
}

func httpIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, apiHelp())
	return
}

func httpIPQueryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// fmt.Fprintf(w, "query ip is :%s", ps.ByName("ipaddress"))

	return
}

// Start server and listen new request
func Start(url string) {
	router := httprouter.New()
	router.GET("/", httpIndex)
	router.GET("/ip/:ipaddress", httpIPQueryHandler)
	log.Fatal(http.ListenAndServe(url, router))
}
