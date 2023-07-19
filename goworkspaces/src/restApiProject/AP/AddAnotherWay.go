package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

type Sum struct {
	Array []int `json:"array"`
}

var sum = 0

func AddSum(req *restful.Request, res *restful.Response) {
	ar := Sum{} //storing structure using shorthand operator soo it must be declared inside function only
	req.ReadEntity(&ar)
	for _, value := range ar.Array {
		sum += value
	}
	res.WriteAsJson(sum)
	log.Println("Sum of Arrays", sum)
}
func main() {
	ws := new(restful.WebService)
	ws.Path("/array")
	//ws.Consumes(restful.MIME_JSON)
	//	ws.Produces(restful.MIME_JSON)
	ws.Route(ws.POST("").To(AddSum))
	restful.DefaultContainer.Add(ws)
	log.Println("port listing on 8008")
	log.Println(http.ListenAndServe(":8008", nil))
}
