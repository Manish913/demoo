package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

type Dt struct {
	One string `json:"one"`
	Two string `json:"two"`
}
type USS struct {
	us map[string]Dt
}

func (u *USS) PostAdd(req *restful.Request, res *restful.Response) {
	dt := Dt{One: req.PathParameter("id")}
	er := req.ReadEntity(&dt)
	if er == nil {
		u.us[dt.One] = dt
		u.us[dt.Two] = dt
		res.WriteHeaderAndEntity(http.StatusCreated, dt)
	} else {
		res.WriteEntity(http.StatusInternalServerError)
	}
}
func (u USS) GetAdd(req *restful.Request, res *restful.Response) {
	add := []USS{}
	for _, data := range u.us {
		ad := data.One + data.Two
		log.Println("Add of all", ad)
	}
	res.WriteEntity(add)
}

func (u USS) Web() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/add")
	ws.Route(ws.GET("/").To(u.GetAdd).
		Writes([]Dt{}).Returns(200, "OK", []Dt{}).
		Returns(404, "Not found", nil))
	ws.Route(ws.POST("/").To(u.PostAdd).
		Reads([]Dt{}).
		Returns(201, "Created", []Dt{}))
	ws.Consumes(restful.MIME_JSON, restful.MIME_JSON)
	ws.Consumes(restful.MIME_JSON, restful.MIME_JSON)

	return ws
}
func main() {
	ws := USS{map[string]Dt{}}
	restful.DefaultContainer.Add(ws.Web())
	log.Println(http.ListenAndServe(":8070", nil))
	log.Println("port Listing on 8070")
}
