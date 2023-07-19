package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

type LearSearch struct {
	Arr    []int
	Search int
}

func Search(req *restful.Request, res *restful.Response) {
	arr := LearSearch{}
	req.ReadEntity(&arr)
	for _, v := range arr.Arr {
		log.Println("value", v)
		if v == arr.Search {
			fmt.Println("found at index ", v, "value", arr.Search)
			res.WriteAsJson(v)
		}
	}
}
func main() {
	ws := new(restful.WebService)
	ws.Path("/search")
	ws.Route(ws.POST("").To(Search))
	restful.Add(ws)
	log.Println("Server is running ")
	log.Println(http.ListenAndServe(":8081", nil))
}
