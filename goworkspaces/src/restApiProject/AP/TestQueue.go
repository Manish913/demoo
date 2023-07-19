package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

// encomplete
type Test struct {
	Ele int `json:"ele"`
}
type Que struct {
	Arr  []int
	Size int
}

var q = []Que{}

func (q *Que) En(req *restful.Request, res *restful.Response) {
	t := Test{}
	size := len(q.Arr)
	req.ReadEntity(t.Ele)
	if q.Size == size {
		log.Println("Queue is Full")
	}
	q.Arr = append(q.Arr, t.Ele)
	log.Println(t.Ele)
	res.WriteAsJson(t.Ele)

}
func (q *Que) IsEmp() bool {
	return len(q.Arr) == 0
}
func main() {
	q := Que{Size: 3}
	ws := new(restful.WebService)
	ws.Path("/queue")
	ws.Route(ws.POST("/").To(q.En))
	restful.Add(ws)
	log.Println(http.ListenAndServe(":8001", nil))
}
