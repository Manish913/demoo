package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/emicklei/go-restful"
)

// This example show how to test one particular RouteFunction (getIt)
// It uses the httptest.ResponseRecorder to capture output

func getIt(req *restful.Request, resp *restful.Response) {
	resp.WriteHeader(204)
}

func TestCallFunction(t *testing.T) {
	httpReq, _ := http.NewRequest("GET", "/", nil)
	req := restful.NewRequest(httpReq)

	recorder := new(httptest.ResponseRecorder)
	resp := restful.NewResponse(recorder)

	getIt(req, resp)
	if recorder.Code != 204 {
		t.Fatalf("Missing or wrong status code:%d", recorder.Code)
	}
}
func main() {
	//getIt()
	ws := new(restful.WebService)
	ws.Path("/test")
	ws.Route(ws.GET("").To(getIt))
	//ws.Route(ws.GET("/").To())
	ws.Method(http.MethodGet)
	restful.Add(ws)
	fmt.Println(http.ListenAndServe(":8087", nil))
}
