package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
	"strings"
)

type Emp struct {
	Name   string `json:"name"`
	Id     string `json:"ide"`
	Salary int64  `json:"salary"`
}
type Every struct {
	user map[string]Emp
}

//	func (e Every) Al() *restful.WebService {
//		ee := new(restful.WebService)
//		ee.Path("/users")
//		ee.Route(ee.GET("/all").To(e.all))
//		ee.Route(ee.GET("/{user-id}").To(e.getOne))
//		ee.Route(ee.POST("/{user-id}").To(e.PostOne))
//		ee.Route(ee.PUT("/{user-id}").To(e.Create))
//		ee.Route(ee.DELETE("/{user-id}").To(e.remove))
//		//restful.Add(ee)
//		log.Println(http.ListenAndServe(":8060", nil))
//		log.Println("listing on port")
//		return ee
//	}
func (e Every) all(req *restful.Request, res *restful.Response) {
	list := []Emp{} //here iterating structure
	for _, each := range e.user {
		list = append(list, each)
	}
	res.WriteEntity(list)
}
func (e Every) getOne(req *restful.Request, res *restful.Response) {
	id := req.PathParameter("user-id")
	er := e.user[id] //here searching in map
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
func (e *Every) PostOne(req *restful.Request, res *restful.Response) {
	a := new(Emp)
	er := req.ReadEntity(&a)
	if er == nil {
		e.user[a.Id] = *a
		res.WriteEntity(a)
	} else {
		log.Println("getting essue while posting")
		res.WriteErrorString(http.StatusInternalServerError, er.Error())
	}
}

// put
func (e *Every) Create(req *restful.Request, res *restful.Response) {
	es := Emp{Id: req.PathParameter("user-id")}
	er := req.ReadEntity(&es)
	if er == nil {
		e.user[es.Id] = es
		res.WriteEntity(es)
		res.WriteHeaderAndEntity(http.StatusCreated, es)

	}
}
func (e *Every) remove(req *restful.Request, res *restful.Response) {
	d := req.PathParameter("user-id")
	delete(e.user, d) //it deletes any particular record from Emp map
}
func getRoutee(path, doc, method string, handler func(*restful.Request, *restful.Response)) *restful.RouteBuilder {
	reout := &restful.RouteBuilder{}
	reout.Path(path)
	reout.To(handler)
	reout.Doc(doc)
	reout.Produces(restful.MIME_JSON)
	reout.Method(strings.ToUpper(method))
	return reout
}
func main() {

	u := Every{map[string]Emp{}}
	ws := &restful.WebService{}
	gt := getRoutee("/mk", "Testing api", "POST", u.Create)
	//restful.DefaultContainer.Add(gt)

	ws.Route(gt)
	//log.Println(t)
}
