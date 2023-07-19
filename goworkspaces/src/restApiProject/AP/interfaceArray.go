package main

import (
	"encoding/json"
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
	"reflect"
)

/*
algorithm
13,40,6,34
divide the array by 2
*/
func MergeSort(req *restful.Request, res *restful.Response) {
	var sl []int
	mp := make(map[string]interface{})
	req.ReadEntity(&mp)
	if mp != nil {
		log.Println(mp)
		for _, v := range mp {
			fmt.Println(reflect.TypeOf(v)) //checking type of array
			switch a := v.(type) {
			case []interface{}: //if array is interface type then execute this block
				for _, value := range a { //iterating interface array
					fmt.Println(value)
					if value != nil {
						v, _ := value.(json.Number).Int64() //typecasting interface array to
						sl = append(sl, int(v))             //appending interface array to integer array after typecasting
					}
				}
			}

			res.WriteAsJson(sl)
		}
	}

}
func main() {
	ws := new(restful.WebService)
	ws.Path("/inter")
	ws.Route(ws.POST("").To(MergeSort))
	restful.Add(ws)
	log.Println(http.ListenAndServe(":8090", nil))

}
