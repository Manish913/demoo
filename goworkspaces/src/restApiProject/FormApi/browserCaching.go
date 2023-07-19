package main

import (
	"io"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
)

// This example shows how to use a WebService filter that passed the Http headers to disable browser cacheing.
//
// GET http://localhost:8080/hello

func main() {
	ws := new(restful.WebService)
	ws.Filter(restful.NoBrowserCacheFilter)
	ws.Route(ws.GET("/hello").To(hello))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8060", nil))
}

func hello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "world")
}
