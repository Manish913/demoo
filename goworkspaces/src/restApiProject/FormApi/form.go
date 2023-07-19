package main

import (
	"encoding/json"
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"github.com/gorilla/schema"
	"io"
	"log"
	"net/http"
	"os"
)

type Detail struct {
	Name string
	Id   int
	Age  int
}

var decode *schema.Decoder

func main() {
	decode = schema.NewDecoder() //schema used to fill the struct form value
	ws := new(restful.WebService)
	ws.Route(ws.GET("/").To(PostForm))
	ws.Route(ws.POST("/posts").Consumes("application/x-www-form-urlencoded").To(Post))
	//contai := restful.Container{}
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8094", nil))

}
func Post(req *restful.Request, res *restful.Response) {
	err := req.Request.ParseForm()
	if err != nil {
		res.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	d := new(Detail)
	err = decode.Decode(d, req.Request.PostForm)
	if err != nil {
		res.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	io.WriteString(res.ResponseWriter, fmt.Sprintf("<html><body>Name=%s,Id=%d,Age=%d</body></html>", d.Name, d.Id, d.Age))
	j, _ := json.Marshal(d)
	os.WriteFile("/home/manish/Desktop/goworkspaces/src/restApiProject/file1.txt", j, 0666)
	//bufio.NewWriter()
}
func PostForm(req *restful.Request, res *restful.Response) {
	io.WriteString(res.ResponseWriter,
		`<html>
		<body>
		<h1>Enter Profile</h1>
		<form method="post">
		    <label>Name:</label>
			<input type="text" name="Name"/>
          <label>ID:</label>
          <input type="text" name="Id">
			<label>Age:</label>
		    <input type="text" name="Age"/>
			<input type="Submit" />
		</form>
		</body>
		</html>`)
}
