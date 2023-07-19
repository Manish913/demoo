package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
	"strconv"
)

type Details struct {
	Name   string `json:"name default:"manish"`
	Id     int    `json:"id default:"1"`
	Age    int    `json:"age default:"24"`
	Salary int64  `json:"salary default:"16000"`
}
type All struct {
	User map[string]Details
}

func (a All) Regis(res *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/user")
	ws.Consumes(restful.MIME_JSON, restful.MIME_JSON) //i want to consume json data
	ws.Produces(restful.MIME_JSON, restful.MIME_JSON) //and well produce json data
	ws.Route(ws.GET("{/user-id}").To(a.GetbyId))
	ws.Route(ws.POST("{/user-id}").To(a.updateUs))
	ws.Route(ws.DELETE("{/user-id}").To(a.delete))
	restful.Add(ws)
	log.Println(http.ListenAndServe((":8090"), nil))

}

func (a All) GetbyId(req *restful.Request, res *restful.Response) {
	id := req.PathParameter("{user-i}")
	us := a.User[id] //we are searching the id in map
	if us.Id == 0 {
		res.AddHeader("content-type", "text")
		res.WriteErrorString(http.StatusNotFound, "User could not found")
	} else {
		res.WriteEntity(&us)
	}
}
func (a *All) updateUs(req *restful.Request, res *restful.Response) {
	us := new(Details)
	err := req.ReadEntity(&us)
	if err == nil {
		a.User[strconv.Itoa(us.Id)] = *us
		res.WriteEntity(us)
	} else {
		res.AddHeader("content-type", "text/plain")
		res.WriteErrorString(http.StatusInternalServerError, err.Error())
	}

}

func (a *All) delete(red *restful.Request, res *restful.Response) {
	re := red.PathParameter("{user-id}")
	delete(a.User, re) //it will delete any particular data
}
func main() {
	//u := All{map[string]Details{}}
	//restful.DefaultContainer.Add(u)
	a := All{}.Regis
	log.Fatal(http.ListenAndServe(":8092", nil))
	//a.Regis()
	fmt.Println(a)

}
