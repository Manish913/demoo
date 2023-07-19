package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

type St struct {
	Ar int
}
type Arrr struct {
	S     []int //it is slice array its size is not define
	Sizee int
}

func (ar *Arrr) Encod(req *restful.Request, res *restful.Response) {
	s := St{}
	req.ReadEntity(&s) //reading value from st structure and appending that value in Arrr struct inside S array
	//if ar.Sizee == ar.IsFul() {
	//	log.Println("Log Is Full")
	//}
	ar.S = append(ar.S, s.Ar)
	log.Println(ar.S)
	res.WriteAsJson(s)
}
func (a *Arrr) DeQue(req *restful.Request, res *restful.Response) {
	if a.IsEmp() {
		r := recover()
		log.Println("Queue is Empty", r)
		res.WriteAsJson("Queue is Empty")
	}
	ele := a.S[0]
	if a.Sizee == 1 {
		a.S = nil
		e := ele //skip this e variable
		log.Println(e)
	}
	a.S = a.S[1:] //it will skip the 0th position elements

	res.WriteAsJson(a.S)
	log.Println("dequeue:= ", a.S)
	//return ele
}
func (a *Arrr) IsEmp() bool {
	return len(a.S) == 0
}
func main() {
	q := Arrr{}
	ws := new(restful.WebService)
	ws.Path("/q")
	ws.Route(ws.POST("").To(q.Encod))
	ws.Route(ws.POST("/ps").To(q.DeQue))
	restful.Add(ws)
	log.Println(http.ListenAndServe(":8002", nil))
}
