package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

//	type MAp struct{
//		//M struct{
//			Name string `json:"name"`
//			Age string `json:"age"`
//		//}
//	}
func AddMAp(req *restful.Request, res *restful.Response) {
	su := 0
	m := make(map[string]int)
	req.ReadEntity(&m)
	for _, val := range m {
		su = su + val
		//	log.Println(k)
	}
	log.Println("Sum of Array", su)
	res.WriteAsJson(su)

}

func main() {
	ws := new(restful.WebService)
	ws.Path("/root")
	ws.Route(ws.POST("/ps/{id}").To(AddMAp))

	restful.Add(ws)
	log.Println(http.ListenAndServe(":8082", nil))
}
