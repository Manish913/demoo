package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
	"strings"
)

type STD struct {
	Name string `json:"name"`
	Id   string `json:"id"`
	Age  int    `json:"age"`
}
type Emplo struct {
	emp map[string]STD
}

func route(path, doc, method string, handler func(req *restful.Request, res *restful.Response)) *restful.RouteBuilder {
	rou := &restful.RouteBuilder{}
	rou.Path(path) //it is rout path
	rou.Produces(restful.MIME_JSON)
	rou.Consumes(restful.MIME_JSON)
	rou.Doc(doc)
	rou.To(handler)
	rou.Method(strings.ToUpper(method))
	return rou
}
func (e Emplo) all1(req *restful.Request, res *restful.Response) {
	list := []STD{} //here iterating structure
	for _, eac := range e.emp {
		list = append(list, eac)
	}
	res.WriteEntity(list)
}
func (e Emplo) GetOne1(req *restful.Request, res *restful.Response) {
	id := req.PathParameter("user")
	er := e.emp[id] //here searching in map
	//len(er.Id)
	if len(er.Id) == 0 {
		res.WriteErrorString(http.StatusNotFound, "User not found")
	} else {
		err := res.WriteEntity(er)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
func (e *Emplo) PostOne1(req *restful.Request, res *restful.Response) {
	a := new(STD)
	er := req.ReadEntity(&a)
	if er == nil {
		e.emp[a.Id] = *a
		res.WriteEntity(a)
	} else {
		log.Println("getting essue while posting")
		res.WriteErrorString(http.StatusInternalServerError, er.Error())
	}
}

// put
func (e *Emplo) Create1(req *restful.Request, res *restful.Response) {
	es := STD{Id: req.PathParameter("user")}
	er := req.ReadEntity(&es)
	if er == nil {
		e.emp[es.Id] = es
		res.WriteEntity(es)
		res.WriteHeaderAndEntity(http.StatusCreated, es)

	}
}
func (e *Emplo) Remove1(req *restful.Request, res *restful.Response) {
	d := req.PathParameter("user")
	delete(e.emp, d) //it deletes any particular record from Emp map
}
func main() {
	u := new(Emplo)
	ws := &restful.WebService{}
	pt := route("/gt", "getting one by id", "GET", u.GetOne1)
	cs := restful.NewContainer()
	restful.Add(ws)
	add := "localhost:8010"
	log.Println("listing on", add)
	http.ListenAndServe(add, cs)
	ws.Route(pt)
	pt1 := route("/post", "posting one by id", "GET", u.Create1)
	ws.Route(pt1)
	pt3 := route("/all", "Get All  value", "UPDATE", u.all1)
	ws.Route(pt3)
	pt2 := route("/del", "removing one by id", "DELETE", u.Remove1)
	ws.Route(pt2)
}
