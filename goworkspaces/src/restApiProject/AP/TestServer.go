package main

import (
	"bytes"
	"errors"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
	"testing"
	"time"
)

func RunRestfulCurlyRouterServer() {
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	u := UserResource{map[string]User{}}
	//u.Register(wsContainer)
	log.Println(u)

	log.Print("start listening on localhost:8090")
	server := &http.Server{Addr: ":8090", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}

func waitForServerUp(serverURL string) error {
	for start := time.Now(); time.Since(start) < time.Minute; time.Sleep(5 * time.Second) {
		_, err := http.Get(serverURL + "/")
		if err == nil {
			return nil
		}
	}
	return errors.New("waiting for server timed out")
}

func TestServer(t *testing.T) {
	serverURL := "http://localhost:8090"
	go func() {
		RunRestfulCurlyRouterServer()
	}()
	if err := waitForServerUp(serverURL); err != nil {
		t.Errorf("%v", err)
	}

	// GET should give a 405
	resp, err := http.Get(serverURL + "/users/")
	if err != nil {
		t.Errorf("unexpected error in GET /users/: %v", err)
	}
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("unexpected response: %v, expected: %v", resp.StatusCode, http.StatusOK)
	}

	// Send a POST request.
	var jsonStr = []byte(`{"id":"1","name":"user1"}`)
	req, err := http.NewRequest("POST", serverURL+"/users/", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", restful.MIME_JSON)

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("unexpected error in sending req: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected response: %v, expected: %v", resp.StatusCode, http.StatusOK)
	}

	// Test that GET works.
	resp, err = http.Get(serverURL + "/users/1")
	if err != nil {
		t.Errorf("unexpected error in GET /users/1: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected response: %v, expected: %v", resp.StatusCode, http.StatusOK)
	}
}
