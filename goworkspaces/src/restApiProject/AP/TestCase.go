package main

//func waitForServerUpp(serveUrl string) error {
//	for st := time.Now(); time.Since(st) < time.Minute; time.Sleep(5 * time.Second) {
//		_, er := http.Get(serveUrl + "/")
//		if er == nil {
//			return nil
//		}
//	}
//	return errors.New("waiting for server time out")
//}
//func TestAdd(t *testing.T) {
//	serUrl := "http://localhost:9090"
//	go func() {
//		Regi()
//	}()
//	if err := waitForServerUpp(serUrl); err != nil {
//		t.Errorf("%v", err)
//	}
//	res, err := http.Get(serUrl + "/adds/")
//	if err != nil {
//		t.Errorf("unexpected error in get /adds/: %v", err)
//	}
//	if res.StatusCode != http.StatusMethodNotAllowed {
//		t.Errorf("unexepected response %v, exepected: %v", res.StatusCode, http.StatusOK)
//	}
//	//Send a Post request
//	var jsonStr=[]byte({"nu[]:})
//
