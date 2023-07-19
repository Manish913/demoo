package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
)

func Regi() {
	ws := new(restful.WebService)
	//ws.Routes(restful.CurlyRouter{})
	ws.Path("/adds")
	ws.Route(ws.POST("/").To(Addd))
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)
	//ws.Produces()
	restful.Add(ws)
	log.Println("port listing on :9090")
	//9090server:=&http.Server{Addr: ":9090",Handler: ws}
	//	log.Println(http.ListenAndServe(":9090", nil))
}
func main() {

	Regi()
}

type ST struct {
	Nu []int `json:"nu"`
}

func Addd(request *restful.Request, response *restful.Response) {

	n := ST{}
	//fmt.Println(n)//it will get empty Array
	var sum = 0
	request.ReadEntity(&n)
	for _, valu := range n.Nu {
		sum = sum + valu
	}
	//fmt.Printf("%T", sum)
	//request.ReadEntity(&n)
	log.Println("Sum of Array", sum)
	response.WriteAsJson(sum)
}
