package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

func main() {
	e := Employe{map[string]Emp{}}
	ws := new(restful.WebService)
	ws.Path("/all")
	ws.Produces(restful.MIME_JSON)
	ws.Consumes(restful.MIME_JSON)
	ws.Route(ws.GET("/").To(e.GetAll).Writes(Emp{}))
	ws.Route(ws.GET("/gt/{user-id}").To(e.GetById).Writes(Emp{}))
	ws.Route(ws.POST("/pst/{id}").To(e.Pst).Reads(Emp{}))
	ws.Route(ws.PUT("/pt/{id}").To(e.Update))
	ws.Route(ws.DELETE("/delete/{id}").To(e.Remove))
	restful.Add(ws)
	log.Println("port listing on 8081")
	log.Println(http.ListenAndServe(":8081", nil))

}

type Emp struct {
	Ename   string `json:"ename"`
	Id      string `json:"id"`
	Salary  int64  `json:"salary"`
	Address string `json:"address"`
	//Ss      struct{
	//}
}
type Employe struct {
	Em map[string]Emp
}

func (e Employe) GetAll(req *restful.Request, response *restful.Response) {
	lis := []Emp{}
	for _, v := range e.Em {
		lis = append(lis, v)
	}
	response.WriteEntity(lis)

}
func (e Employe) GetById(req *restful.Request, res *restful.Response) {
	id := req.PathParameter("user-id")
	er := e.Em[id]
	if len(er.Id) == 0 {
		res.WriteErrorString(http.StatusInternalServerError, "User not found")
	} else {
		res.WriteEntity(id)
	}
}
func (e *Employe) Pst(req *restful.Request, res *restful.Response) {
	usr := Emp{Id: req.PathParameter("user-id")}
	er := req.ReadEntity(&usr)
	log.Println("Data is nil", er)
	if er == nil {
		e.Em[usr.Id] = usr
		res.WriteHeaderAndEntity(http.StatusCreated, usr)
	} else {
		log.Println("An Error is ")
		res.WriteError(http.StatusInternalServerError, nil)
	}

}
func (e *Employe) Update(req *restful.Request, res *restful.Response) {
	us := new(Emp)
	er := req.ReadEntity(&us)
	if er == nil {
		e.Em[us.Id] = *us
		res.WriteEntity(us)
	} else {
		res.WriteHeaderAndEntity(http.StatusInternalServerError, er)
	}
}
func (e *Employe) Remove(req *restful.Request, res *restful.Response) {
	id := req.PathParameter("user-id")
	log.Println("Record deleted Sucessfuly")
	res.WriteAsJson("Record Deleted")
	delete(e.Em, id)
}
