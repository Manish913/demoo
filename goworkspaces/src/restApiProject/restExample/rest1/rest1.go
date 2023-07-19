package main

import (
	"github.com/emicklei/go-restful/v3"
	"io"
	"log"
	"net/http"
)

func main() {

	w := new(restful.WebService)
	w.Route(w.GET("/test").Filter(basicAuthenticates).To(test))
	//	w.Route(w.POST("/post").Filter(basicAuthenticates).To(test))
	//	restful.Add(w)
	restful.Add(w)
	log.Fatal(http.ListenAndServe("m2.plus:8082", nil))
}
func basicAuthenticates(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	u, p, ok := req.Request.BasicAuth()
	if !ok || u != "manish" || p != "7488" {
		resp.AddHeader("www-Authenticate", "Basic realm=Protected Area")
		resp.WriteErrorString(401, "401 unauthorising")
		return
	}
	chain.ProcessFilter(req, resp)

}
func test(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "hii")
	resp.Write([]byte("\nhey mani how r u man"))
	resp.Write([]byte("\n how is study going"))
}

//func post(req *restful.Request, res *restful.Response) {
//	ok, o := res.Write([]byte(""))
//	fmt.Println("heloo", ok, o)
//	//io.ReadAll()
//}
