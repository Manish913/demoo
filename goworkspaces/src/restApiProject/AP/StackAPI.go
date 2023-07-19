package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

type Stack1 struct {
	Stk int `json:"s"`
}
type STK struct {
	StkArray []int //it is slice array
	Size     int
}

func (st *STK) Push1(req *restful.Request, res *restful.Response) {
	s := Stack1{}
	req.ReadEntity(&s)
	log.Println(s.Stk)
	st.StkArray = append(st.StkArray, s.Stk)
	res.WriteAsJson(st.StkArray)
}
func (st *STK) Pop(req *restful.Request, res *restful.Response) {
	if st.IsEmptyy() {
		res.WriteAsJson("Stack is Empty")
	}
	lst := st.StkArray[len(st.StkArray)-1]
	if st.Size == len(st.StkArray)-1 {
		st.StkArray = nil
		log.Println(lst)
	}
	st.StkArray = st.StkArray[:len(st.StkArray)-1]
	log.Println("Pop", st.StkArray)
	res.WriteAsJson(st.StkArray)

}
func (st *STK) Top(req *restful.Request, res *restful.Response) {

	req.ReadEntity(&st)
	log.Println("Top element:= ", st.StkArray[len(st.StkArray)-1:])
	res.WriteAsJson(st.StkArray[len(st.StkArray)-1:])
}

// if stack is empty then it will return 0
func (st *STK) IsEmptyy() bool {
	return len(st.StkArray) == 0
}
func main() {
	s := STK{}
	ws := new(restful.WebService)
	ws.Path("/stack")
	ws.Route(ws.POST("").To(s.Push1))
	ws.Route(ws.POST("/p").To(s.Pop))
	ws.Route(ws.POST("/tp").To(s.Top))
	restful.Add(ws)
	log.Println("port is running on: ")
	log.Println(http.ListenAndServe(":8087", nil))
}
