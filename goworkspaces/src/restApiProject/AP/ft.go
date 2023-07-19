package main

import (
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserResource struct {
	users map[string]User
}

func main() {
	ws := UserResource{map[string]User{}}
	restful.DefaultContainer.Add(ws.Register())
	log.Println(http.ListenAndServe(":8070", nil))

}
func (u UserResource) Register() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/users").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("/{user-id}").To(u.findUser))
	ws.Route(ws.POST("").To(u.updateUser))
	ws.Route(ws.PUT("/{user-id}").To(u.createUser))
	ws.Route(ws.DELETE("/{user-id}").To(u.removeUser))

	return ws
}

// GET http://localhost:8090/users/1
func (u UserResource) findUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	usr := u.users[id]
	if len(usr.Id) == 0 {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "User could not be found.")
	} else {
		response.WriteEntity(usr)
	}
}

// POST http://localhost:8090/users
// <User><Id>1</Id><Name>Melissa Raspberry</Name></User>
func (u *UserResource) updateUser(request *restful.Request, response *restful.Response) {
	usr := new(User)
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
func (u *UserResource) createUser(request *restful.Request, response *restful.Response) {
	usr := User{Id: request.PathParameter("user-id")}
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
func (u *UserResource) removeUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	delete(u.users, id)
}

// GET should give a 405
//	resp, err := http.Get(serverURL + "/users/")
//	if err != nil {
//		t.Errorf("unexpected error in GET /users/: %v", err)
//	}
//	if resp.StatusCode != http.StatusMethodNotAllowed {
//		t.Errorf("unexpected response: %v, expected: %v", resp.StatusCode, http.StatusOK)
//	}
//
//	// Send a POST request.
//	var jsonStr = []byte(`{"id":"1","name":"user1"}`)
//	req, err := http.NewRequest("POST", serverURL+"/users/", bytes.NewBuffer(jsonStr))
//	req.Header.Set("Content-Type", restful.MIME_JSON)
//
//	client := &http.Client{}
//	resp, err = client.Do(req)
//	if err != nil {
//		t.Errorf("unexpected error in sending req: %v", err)
//	}
//	if resp.StatusCode != http.StatusOK {
//		t.Errorf("unexpected response: %v, expected: %v", resp.StatusCode, http.StatusOK)
//	}
//
//	// Test that GET works.
//	resp, err = http.Get(serverURL + "/users/1")
//	if err != nil {
//		t.Errorf("unexpected error in GET /users/1: %v", err)
//	}
//	if resp.StatusCode != http.StatusOK {
//		t.Errorf("unexpected response: %v, expected: %v", resp.StatusCode, http.StatusOK)
//	}
//}
