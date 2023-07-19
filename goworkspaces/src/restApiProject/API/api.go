package main

import (
	"github.com/emicklei/go-restful/v3"
	"net/http"
)

type ST struct {
	Id int `json:"id"`
}
type DTail struct {
	dt map[string]ST
}

func (r DTail) WB() *restful.WebService {
	ws := new(restful.WebService)
	ws.Produces(restful.MIME_JSON)
	ws.Consumes(restful.MIME_JSON)
	return ws
}
func main() {

}
func Ts(req *restful.Request, Res *restful.Response) {
	Res.WriteAsJson("hey manish")
}
func (d DTail) Pst(req *restful.Request, res *restful.Response) {
	s := new(DTail)
	res.WriteAsJson("posting")
	res.WriteHeaderAndEntity(http.StatusCreated, s)
}
