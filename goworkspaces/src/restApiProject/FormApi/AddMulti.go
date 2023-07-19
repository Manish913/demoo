package main

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/gorilla/schema"
	"log"
	"net/http"
)

type Adds struct {
	A int
	B int
}

var decd *schema.Decoder

func main() {
	decd = schema.NewDecoder()
	ws := new(restful.WebService)
	ws.Route(ws.GET("/gt"))
	ws.Route(ws.POST("/ps"))
	restful.Add(ws)
	log.Println(http.ListenAndServe(":8888", nil))
}
func Add(req *restful.Request, res *restful.Response) {
	er := req.Request.ParseForm()
	if er != nil {
		res.WriteErrorString(http.StatusBadRequest, er.Error())
	}
	add := new(Adds)
	//decode take two arguments one interface second map[string][]string
	/*
		Decode decodes a map[string][]string to a struct.
		The first parameter must be a pointer to a struct.
		The second parameter is a map, typically url.Values from an HTTP request. Keys are "paths" in dotted notation to the struct fields and nested structs.
		See the package documentation for a full explanation of the mechanics.
	*/
	er = decd.Decode(add, req.Request.PostForm)
	if er != nil {
		res.WriteErrorString(http.StatusBadRequest, er.Error())
	}
}
