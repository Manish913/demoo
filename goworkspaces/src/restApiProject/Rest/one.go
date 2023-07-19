package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"github.com/gorilla/schema"
	"io"
	"log"
	"net/http"
)

type Dt struct {
	Name       string
	Roll       int
	Mobile     int64
	TotalMarks int
	Pass       string
}

var decodes *schema.Decoder

func main() {
	decodes = schema.NewDecoder() //it is used to fill the struct value form
	ds := new(restful.WebService)
	ds.Route(ds.GET("/st").To(postAll))
	ds.Route(ds.POST("/st").Consumes("application/x-www-form-urlencoded").To(GettAll))
	restful.Add(ds)
	log.Fatal(http.ListenAndServe(":7080", nil))
}
func GettAll(req *restful.Request, res *restful.Response) {
	err := req.Request.ParseForm()
	if err != nil {
		res.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	p := new(Dt)
	err = decodes.Decode(p, req.Request.PostForm)
	if err != nil {
		res.WriteErrorString(http.StatusBadRequest, err.Error())
		//return
	}

	io.WriteString(res.ResponseWriter, fmt.Sprintf("<html><body>Name=%s, Roll=%d,Mobile=%d,TotalMarks=%d,Result=%s</body></html>", p.Name, p.Roll, p.Mobile, p.TotalMarks, p.Pass))
}
func postAll(req *restful.Request, res *restful.Response) {
	io.WriteString(res.ResponseWriter,
		`<html>
            <body>
 <h1>M2-Plush</h1>
        <form method="post">
      <lavel>Name:</lavel>
<input type="text" name="Name"/><br>
    <lavel>Roll:</lavel>
<input type="text" name="Roll"/><br>
    <lavel>Mobile:</lavel>
<input type=number name="Mobile"/><br>
    <lavel>Total:</lavel>
<input type=number name="TotalMarks"/><br>
    <lavel>Result:</lavel>
<input type="text" name="Pass"/>
<input type="Submit"/>
</form>
</body>
</html>`)
}
