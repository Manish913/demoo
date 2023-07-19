package main

import (
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
)

type Users struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserResources struct {
	users map[string]Users
}

func (u UserResources) Register() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/users").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("/{user-id}").To(u.findUsers))
	ws.Route(ws.POST("").To(u.updateUsers))
	ws.Route(ws.PUT("/{user-id}").To(u.createUsers))
	ws.Route(ws.DELETE("/{user-id}").To(u.removeUsers))

	return ws
}

// GET http://localhost:8090/users/1
func (u UserResources) findUsers(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	usr := u.users[id]
	if len(usr.Id) == 0 {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "User could not be found.")
	} else {
		response.WriteEntity(usr)
	}
}
func main() {
	rs := UserResources{map[string]Users{}}
	restful.DefaultContainer.Add(rs.Register())
	log.Println(http.ListenAndServe(":8010", nil))
	log.Println("serving on 8010 port")

}

// POST http://localhost:8090/users
// <User><Id>1</Id><Name>Melissa Raspberry</Name></User>
func (u *UserResources) updateUsers(request *restful.Request, response *restful.Response) {
	usr := new(Users)
	err := request.ReadEntity(&usr)
	if err == nil {
		u.users[usr.Id] = *usr
		response.WriteEntity(usr)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

// PUT http://localhost:8090/users/1
// <User><Id>1</Id><Name>Melissa</Name></User>
func (u *UserResources) createUsers(request *restful.Request, response *restful.Response) {
	usr := Users{Id: request.PathParameter("user-id")}
	err := request.ReadEntity(&usr)
	if err == nil {
		u.users[usr.Id] = usr
		response.WriteHeader(http.StatusCreated)
		response.WriteEntity(usr)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

// DELETE http://localhost:8090/users/1
func (u *UserResources) removeUsers(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	delete(u.users, id)
}

func RunRestfulCurlyRouterServer1() {
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	u := UserResources{map[string]Users{}}
	u.Register()

	log.Print("start listening on localhost:8090")
	server := &http.Server{Addr: ":8090", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}

//func waitForServerUp1(serverURL string) error {
//	for start := time.Now(); time.Since(start) < time.Minute; time.Sleep(5 * time.Second) {
//		_, err := http.Get(serverURL + "/")
//		if err == nil {
//			return nil
//		}
//	}
//	return errors.New("waiting for server timed out")
//}
//
//func TestServer1(t *testing.T) {
//	serverURL := "http://localhost:8090"
//	go func() {
//		RunRestfulCurlyRouterServer()
//	}()
//	if err := waitForServerUp(serverURL); err != nil {
//		t.Errorf("%v", err)
//	}

// GET should give a 405
//resp, err := http.Get(serverURL + "/users/")
//if err != nil {
//	t.Errorf("unexpected error in GET /users/: %v", err)
//}
//if resp.StatusCode != http.StatusMethodNotAllowed {
//	t.Errorf("unexpected response: %v, expected: %v", resp.StatusCode, http.StatusOK)
//}
//
//// Send a POST request.
//var jsonStr = []byte(`{"id":"1","name":"user1"}`)
//req, err := http.NewRequest("POST", serverURL+"/users/", bytes.NewBuffer(jsonStr))
//req.Header.Set("Content-Type", restful.MIME_JSON)
//
//client := &http.Client{}
//resp, err = client.Do(req)
//if err != nil {
//	t.Errorf("unexpected error in sending req: %v", err)
//}
//if resp.StatusCode != http.StatusOK {
//	t.Errorf("unexpected response: %v, expected: %v", resp.StatusCode, http.StatusOK)
//}
//
//// Test that GET works.
//resp, err = http.Get(serverURL + "/users/1")
//if err != nil {
//	t.Errorf("unexpected error in GET /users/1: %v", err)
//}
//if resp.StatusCode != http.StatusOK {
//	t.Errorf("unexpected response: %v, expected: %v", resp.StatusCode, http.StatusOK)
//}
