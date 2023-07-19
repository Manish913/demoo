package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

//type Arr []int

type Mp struct {
	M map[string][]int `json:"m"`
}

var sum1 = 0

func AddMap(req *restful.Request, res *restful.Response) {
	m := Mp{}          //storing structure in variable to access the map that created inside structure
	req.ReadEntity(&m) //it will get called when we hit http url
	log.Println(m)
	for _, value := range m.M {
		value = value
		sum1++

	}
	log.Println("sum of array", sum1)
	//it
	res.WriteAsJson(&sum1)

}

func main() {
	ws := new(restful.WebService)
	ws.Path("/adds")
	ws.Route(ws.POST("").To(AddMap))
	ws.Consumes(restful.MIME_JSON)
	restful.Add(ws)
	add := ":9070"
	log.Println("server is listing on 9070")
	http.ListenAndServe(add, nil)

}
