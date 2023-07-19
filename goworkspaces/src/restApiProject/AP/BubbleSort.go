package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

type Bubble struct {
	Arr []int `json:"ar"`
}

// var Arr[]int
// var ar= []Bubble{}
var temp = 0

func BubbleSort(req *restful.Request, res *restful.Response) {
	ar := Bubble{}
	//a := ar.Arr
	req.ReadEntity(&ar)
	log.Println("Value", ar.Arr)
	for v := range ar.Arr {
		fmt.Println("Before Sort", v)
	}

	n := len(ar.Arr) - 1
	for i := 0; i < n; i++ {
		for j := 1; j < n-i; j++ {
			if ar.Arr[j-i] > ar.Arr[j] {
				temp, ar.Arr[j-i], ar.Arr[j] = ar.Arr[j-i], ar.Arr[j], temp
			}
		}
		res.WriteAsJson(ar.Arr)
	}
	var i = 0
	for i = 0; i < len(ar.Arr); i++ {
		fmt.Println("After shorting", ar.Arr[i])

	}
	res.WriteAsJson(i)
}
func main() {
	ws := new(restful.WebService)
	ws.Path("/sort")
	ws.Route(ws.POST("/").To(BubbleSort))
	restful.Add(ws)
	log.Println("Port listing on :8000")
	log.Println(http.ListenAndServe(":8000", nil))
}
