package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
	"strings"
)

func getRoute(path, doc, method string, handler func(*restful.Request, *restful.Response)) *restful.RouteBuilder {
	reout := &restful.RouteBuilder{}
	reout.Path(path)
	reout.To(handler)
	reout.Doc(doc)
	reout.Produces(restful.MIME_JSON)
	reout.Method(strings.ToUpper(method))
	return reout
}
func main() {
	//creating we service object.
	ws := &restful.WebService{}
	//creating routePath using restful.RouteBuilder object/struct
	helloRouteBuilder := getRoute("/hello", "testing hello api", "GET", Hello)

	/*helloRouteBuilder = &restful.RouteBuilder{}
	helloRouteBuilder.Path("/hello")
	helloRouteBuilder.To(Hello)
	helloRouteBuilder.Doc("testing hello api")
	helloRouteBuilder.Produces(restful.MIME_XML)
	helloRouteBuilder.Method("GET")
	param:= &restful.Parameter{}
	helloRouteBuilder.Param(param)*/

	//registering api/route
	ws.Route(helloRouteBuilder)

	/* creating new hellName route/api
	helloNameRouteBuilder := &restful.RouteBuilder{}
	helloNameRouteBuilder.Path("/hello/name")
	helloNameRouteBuilder.Method("GET")
	helloNameRouteBuilder.To(HelloName)
	helloNameRouteBuilder.Doc("testing hello api")
	helloNameRouteBuilder.Produces(restful.MIME_XML)*/

	//registering api/route\
	helloNameRouteBuilder := getRoute("/hello/name", "getiig aname", "GET", HelloName)
	ws.Route(helloNameRouteBuilder)
	cs := restful.NewContainer()
	cs.Add(ws)
	address := "localhost:8070"
	log.Println("starting http server on ", address)
	http.ListenAndServe(address, cs)

}

func HelloName(req *restful.Request, res *restful.Response) {
	res.WriteAsJson("hello i am sonal kumar")
}
func Hello(req *restful.Request, res *restful.Response) {
	res.WriteAsJson("hello world")
}
