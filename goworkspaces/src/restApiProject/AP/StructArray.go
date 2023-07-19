package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

type Data struct {
	Name  string `json:"name"`
	Id    int    `json:"id"`
	Slary int    `json:"slary"`
}

// creating array of structure
var STR = []Data{}

type SAl struct {
	Mp map[string]Data
}

//var S=SAl{}

func (s SAl) StructArray(req *restful.Request, res *restful.Response) {
	var dt = []Data{}
	req.ReadEntity(&dt)
	log.Println(dt)
	for _, v := range s.Mp {
		dt = append(dt, v)
	}
	res.WriteAsJson(dt)
	log.Println(dt)
}
func main() {
	s := SAl{map[string]Data{}}
	ws := new(restful.WebService)
	ws.Path("/array")
	ws.Route(ws.POST("/ps/{id}").To(s.StructArray))
	restful.Add(ws)
	log.Println(http.ListenAndServe(":8083", nil))

}
