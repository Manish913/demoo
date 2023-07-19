package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

type CustomeArr map[string]interface{}

// creating array of map[string]interface{}
var Mp1 = []CustomeArr{}

func JsonMAp(req *restful.Request, response *restful.Response) {
	//Mp1:=[]CustomeArr{}
	//	Mp1=m
	req.ReadEntity(&Mp1)
	//var DB []interface{}
	log.Println("name:", Mp1, "#")
	for _, value := range Mp1 {
		//	log.Printf("%T", reflect.TypeOf(Mp1))
		//log.Println("key :", key, "value: ", value)
		//DB = append(DB, value)
		log.Println("name :", "#", value)
		for k, v := range value {
			log.Println("Map of", v, k)
		}
		//Mp1 = append(Mp1, value)

	}
	response.WriteAsJson(Mp1)

	log.Println()
}
func main() {
	//c := CustomeArr{}
	ws := new(restful.WebService)
	ws.Path("/app")
	ws.Route(ws.POST("/ps").To(JsonMAp))
	restful.Add(ws)
	log.Println(http.ListenAndServe(":9090", nil))
}
