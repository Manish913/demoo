package main

import (
	"bytes"
	"github.com/emicklei/go-restful/v3"
	_ "github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

//
//func waitForUr(URl string) error {
//	for start := time.Now(); time.Since(start) < time.Minute; time.Sleep(5 * time.Second) {
//		_, er := http.Get(URl + "")
//		if er == nil {
//			return nil
//		}
//	}
//	return errors.New("wait for server time out")
//}

func TestRest(t *testing.T) {
	url := "http://localhost:8089/test"
	go func() {
		Web() //it is working as thread
	}()
	//rs:=RestApp{map[string]Test{}}
	//TEst()
	//if er := waitForUr(url); er != nil {
	//	t.Errorf("%v", er)
	//}
	res, er := http.Get(url)
	if er != nil {
		t.Errorf("unexepected error %v", er)
	}
	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("unexepected response %v exepected %v", res.StatusCode, http.StatusOK)
	}
	//Post a request
	var jsonS = []byte(`{"Id":"1","Name":"manish"}`)
	re, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonS))
	re.Header.Set("context", restful.MIME_JSON)
	client := &http.Client{}
	resp, e := client.Do(re)
	if e != nil {
		t.Errorf("unexepected %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexepected %v , exepected %v", resp.StatusCode, http.StatusOK)
	}
	//res, er = http.Get(url + "1")
	//if er != nil {
	//	t.Errorf("unexepected  error%v", er)
	//}
	//if res.StatusCode != http.StatusOK {
	//	t.Errorf("unexepected %v , exepected %v", res.StatusCode, http.StatusOK)
	//}
}
