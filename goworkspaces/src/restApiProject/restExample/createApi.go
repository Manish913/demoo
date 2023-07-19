package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

type UserService struct {
	users map[string]User
}

func (u UserService) getAll(req *restful.Request, res *restful.Response) {
	lis := []User{}
	for _, l := range lis {
		lis = append(lis, l)
	}
	res.WriteEntity(lis)
}
func (u UserService) getById(req *restful.Request, res *restful.Response) {
	id := req.PathParameter("user-id")
	usr := u.users[id]
	if usr.Id == 0 {
		res.WriteErrorString(http.StatusNotFound, "User could not found")
	} else {
		res.WriteEntity(usr)
	}
}
func (u *UserService) post(req *restful.Request, res *restful.Response) {

}
func (u *UserService) put(req *restful.Request, res *restful.Response) {

}
func (u *UserService) removeUser(req *restful.Request, res *restful.Response) {

}
func (u UserService) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/users")
	ws.Consumes(restful.MIME_JSON, restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON, restful.MIME_JSON)
	//ws.Route(ws.GET(""))
	//tg := []string{"users"}
	ws.Route(ws.GET("/").To(u.getAll))
	ws.Doc("get all")

	ws.Route(ws.GET("{user-id}").To(u.getById)).
		Doc("get a user")

	return ws
}

func main() {
	mp := UserService{map[string]User{}}
	ws := new(restful.WebService)
	restful.DefaultContainer.Add(mp.WebService())
	ws.Route(ws.GET(""))
	container := new(restful.Container)
	container.Add(ws)
	log.Println(http.ListenAndServe(":8090", nil))

}

type User struct {
	Id   int    `json:"id"`
	name string `json:"name"`
}
