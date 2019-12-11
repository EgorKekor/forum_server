package main

import (
	"db-forum/database"
	"db-forum/router"
	"flag"
	"log"

	"github.com/valyala/fasthttp"
)

func main() {
	flag.Parse()
	if err := database.InitDB(config.DB); err != nil {
		log.Println("can't init DB", err.Error())
		return
	}
	rt := router.CreateRouter()
	fasthttp.ListenAndServe(config.Port, rt.Handler)
}
