package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"github.com/gorilla/schema"
	"io"
	"log"
	"net/http"
)

type Validation struct {
	UserName string
	Password int
}

var decoded *schema.Decoder

func main() {
	decoded = schema.NewDecoder()
	ws := new(restful.WebService)
	ws.Route(ws.GET("/st").To(Validations))
	ws.Route(ws.POST("/st").Consumes("application/x-www-form-urlencoded").To(validForm))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8099", nil))

}
func validForm(req *restful.Request, res *restful.Response) {
	er := req.Request.ParseForm()
	if er != nil {
		res.WriteErrorString(http.StatusBadRequest, er.Error())
		return
	}
	v := new(Validation)
	er = decoded.Decode(v, req.Request.PostForm)
	if er != nil {
		res.WriteErrorString(http.StatusBadRequest, er.Error())
		return
	}
	io.WriteString(res.ResponseWriter, fmt.Sprintf("<html><body>UserName=%s,Password=%d</body></html", v.UserName, v.Password))
}
func Validations(req *restful.Request, res *restful.Response) {
	io.WriteString(res.ResponseWriter, `
<html><div class="container">
  <form method="post">
    <label for="UserName">Username</label>
    <input type="text" id="UserName" name="UserName" required><br>

    <label for="pwd">Password</label>
  <input type="Password" name=Password required/>
    <input type="submit" value="Submit">
  </form>
</div>
</html>`)
}
