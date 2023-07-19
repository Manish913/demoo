package main

import (
	"github.com/emicklei/go-restful/v3"
	"net/http"
	"strings"
)

type GT struct {
	Name   string
	Id     string
	Salary int64
}
type Detail struct {
	dt map[string]GT
}

func Rt(path, doc, method string, handle func(req *restful.Request, res *restful.Response)) *restful.RouteBuilder {
	rt := &restful.RouteBuilder{}
	rt.Doc(doc)
	rt.To(handle)
	rt.Path(path)
	rt.Consumes(restful.MIME_JSON)
	//rt.Produces(restful.MIME_JSON)
	rt.Method(strings.ToUpper(method))

	return rt
}
func Hii(req *restful.Request, res *restful.Response) {
	res.WriteAsJson("Hey manish kumar")
}
func GETS(req *restful.Request, res *restful.Response) {
	//dt := Detail{}
	//log.Println("", dt)
	res.WriteAsJson("what's up")

}
func main() {
	rs := new(restful.WebService)
	rt := Rt("/gt", "Getting ", "GET", Hii)
	rs.Route(rt)
	//rl := Rt("/gt/{user-id}", "POST", "Post", GETS)
	//rs.Route(rl)
	rst := restful.DefaultContainer
	add := "localhost:8050"
	http.ListenAndServe(add, rst)
}
