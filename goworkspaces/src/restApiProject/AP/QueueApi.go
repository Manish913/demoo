package main

import (
	"errors"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

type Element struct {
	Ar int
}
type Queue struct {
	Arr  []int
	Size int
}

func (qu *Queue) Enqueue(req *restful.Request, res *restful.Response) {
	e := Element{}
	req.ReadEntity(&e)

	qu.Arr = append(qu.Arr, e.Ar, 10) //appending elements in slice array
	res.WriteAsJson(qu.Arr)
	log.Println(qu.Arr)
}
func (qu *Queue) Dequeue(req *restful.Request, res *restful.Response) {
	if qu.IsEmPty() {
		res.WriteAsJson("Queue Is Empty")
	}
	e := qu.Arr[0]
	if qu.GetSiz() == 1 {
		qu.Arr = nil
		log.Println(e)
	}
	qu.Arr = qu.Arr[1:]
	res.WriteAsJson(qu.Arr)
	res.WriteAsJson("Dequeue Element")

}
func (qu *Queue) GetSiz() int {
	return len(qu.Arr)
}
func (qu *Queue) Peek() (int, error) {
	if qu.IsEmPty() {
		return 0, errors.New("Empty")
	}
	return qu.Arr[0], nil
}
func (qu *Queue) IsEmPty() bool {
	return len(qu.Arr) == 0
}
func main() {
	q := Queue{}
	ws := new(restful.WebService)
	ws.Path("/qu")
	ws.Route(ws.POST("").To(q.Enqueue))
	ws.Route(ws.POST("/d").To(q.Dequeue))
	restful.Add(ws)
	log.Println(http.ListenAndServe(":9000", nil))
}
