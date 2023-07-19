package main

import (
	"io"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
)

// This example shows how to create a (Route) Filter that performs Basic Authentication on the Http request.
//
// GET http://localhost:8080/secret
// and use admin,admin for the credentials

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/secret").Filter(basicAuthenticate).To(secret))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe("m2.plus:8082", nil))
}
func basicAuthenticate(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	// usr/pwd = admin/admin
	u, p, ok := req.Request.BasicAuth()
	if !ok || u != "admin" || p != "admin" {
		resp.AddHeader("WWW-Authenticate", "Basic realm=Protected Area")
		resp.WriteErrorString(401, "401: Not Authorized")
		return
	}
	chain.ProcessFilter(req, resp)
}
func secret(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "42")
	io.WriteString(resp, "manish kumar")
	resp.Write([]byte(" \nhello "))

}
