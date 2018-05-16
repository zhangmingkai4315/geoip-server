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
	router.GET("/api/geoip2/:ipaddress", httpIPQueryHandler("en"))
	router.GET("/api/geoip2/:ipaddress/zh-cn", httpIPQueryHandler("zh-ch"))
	router.GET("/api/geoip2/:ipaddress/en", httpIPQueryHandler("en"))

	router.POST("/api/geoip2/iplist", httpIPListQueryHandler("en"))
	router.POST("/api/geoip2/iplist/zh-cn", httpIPListQueryHandler("zh-ch"))
	router.POST("/api/geoip2/iplist/en", httpIPListQueryHandler("en"))
	log.Fatal(http.ListenAndServe(url, router))
}
