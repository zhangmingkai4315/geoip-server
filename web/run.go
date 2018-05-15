package web

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Start server and listen new request
func Start(url string) {
	router := httprouter.New()
	router.GET("/", httpIndex)
	router.GET("/geoip2/:ipaddress", httpIPQueryHandler)
	log.Fatal(http.ListenAndServe(url, router))
}
