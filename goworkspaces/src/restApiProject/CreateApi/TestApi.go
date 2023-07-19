package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

type Api struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type Test struct {
	Usr map[string]Api
}

func main() {
	t := Test{make(map[string]Api)}
	ws := new(restful.WebService)
	ws.Path("usr")
	//ws.Filter("")
	ws.Route(ws.GET("/").To(t.Gt))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8070", nil))
	fmt.Println("listening on port :8070")

}
func (a Api) Regi(req *restful.Request, res *restful.Response) *restful.WebService {
	n := new(restful.WebService)

	return n
}
func (t Test) Gt(req *restful.Request, res *restful.Response) {
	t = Test{make(map[string]Api)}

	body := req.Request.Body
	er := res.WriteEntity(&t)
	if er == nil {
		fmt.Println("Say")
	}
	fmt.Println(body)
	fmt.Println("Getting")
}
