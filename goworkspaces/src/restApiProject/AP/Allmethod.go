package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

type APT struct {
	Name   string `json:"name"`
	Id     string `json:"id"`
	Salary int64  `json:"salary"`
}
type Manage struct {
	Detail map[string]APT
}

func (m Manage) All(req *restful.Request, res *restful.Response) {
	list := []APT{}
	for _, valu := range m.Detail {
		list = append(list, valu)
	}
	res.WriteEntity(&list)
}
func (m Manage) GetByOne(req *restful.Request, res *restful.Response) {
	id := req.PathParameter("user-id")
	l := m.Detail[id]
	if len(l.Id) == 0 {
		res.WriteErrorString(http.StatusInternalServerError, "User did not found")
	} else {
		res.WriteEntity(l)
	}

}
func (m *Manage) Posts(req *restful.Request, res *restful.Response) {
	use := APT{Id: req.PathParameter("user-id")}
	er := req.ReadEntity(&use)

	if er == nil {
		m.Detail[use.Id] = use
		res.WriteHeaderAndEntity(http.StatusCreated, use)
	} else {
		res.WriteError(http.StatusInternalServerError, er)
	}
}

func main() {
	e := Manage{map[string]APT{}}
	ws := new(restful.WebService)
	ws.Path("/every")
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)
	ws.Route(ws.GET("").To(e.All))
	ws.Route(ws.GET("/one/{id}").To(e.GetByOne).Writes(APT{}).Returns(200, "Ok", APT{}).
		Returns(404, "Not Present", nil))
	ws.Route(ws.POST("/ps/{id}").To(e.Posts).Reads(APT{}))
	//Reads(APT{}))

	restful.Add(ws)
	log.Println(http.ListenAndServe(":8002", nil))
	log.Println("Port is listing on 8002")
}
