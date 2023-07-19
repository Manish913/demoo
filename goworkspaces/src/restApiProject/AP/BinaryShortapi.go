package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"net/http"
	"strconv"
)

type BinarySort struct {
	Arr    []int `json:"arr"`
	Target int   `json:"target"`
}

func Sort(req *restful.Request, res *restful.Response) {

	var ar = BinarySort{}
	req.ReadEntity(&ar)
	find := ar.Target

	start := 0
	end := len(ar.Arr) - 1

	for start <= end {
		mid := (start + end) / 2
		if find == ar.Arr[mid] {
			res.WriteAsJson("Element found at index  := " + strconv.Itoa(mid))
			break
		} else if find > ar.Arr[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}

	}
	if start > end {
		res.WriteAsJson("Element not found")
	}

}
func main() {
	//	bn := BinarySort{}
	ws := new(restful.WebService)
	ws.Path("/binary")
	ws.Route(ws.POST("/").To(Sort))
	restful.Add(ws)
	fmt.Println(http.ListenAndServe(":8088", nil))
}
