package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

type Test struct {
	Id, Name string
}
type RestApp struct {
	mp map[string]Test
}

func Web() restful.Route {
	w := new(restful.WebService)

	w.Path("/test")
	w.Route(w.POST("").To(TEst))
	//	w.Route(w.GET("").To(GEt))
	w.Consumes(restful.MIME_JSON)
	w.Produces(restful.MIME_JSON)
	restful.Add(w)
	log.Println(http.ListenAndServe(":8089", nil))
	return restful.Route{}
}

func TEst(req *restful.Request, res *restful.Response) {
	//re := RestApp{}
	//req.ReadEntity(&re)
	//res.WriteEntity(re)
}

//func GEt(req *restful.Request, res *restful.Response) {
//	r := RestApp{}
//	res.WriteAsJson(r)
//}

func main() {
	Web()
}
