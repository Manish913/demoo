package main

import (
	"github.com/emicklei/go-restful/v3"
)

type Cus struct {
	Name  string
	Roll  int
	Phone int64
}

func main() {
	ws := new(restful.WebService)
	con := restful.NewContainer()
	ws.Route(ws.GET("/geloo").To(PostFirst))
	con.Add(ws)

}
func PostFirst(req *restful.Request, res *restful.Response) {

}
