package requests

import "testing"

func TestURl(t *testing.T) {
	t.Run("test urls",func(t *testing.T) {
		inp1 := []string{"http://valid-url.com","http://localhost:8000","http://drive.google.com/21512512"}

		urls := UrlsRequest{Urls: inp1}
		err := urls.VerifyURLs()
		if err != nil {
			t.Fatal("unexpected result: no errors expected")
		}

		inp2 := []string{"wrong-address"}
		urls2 := UrlsRequest{Urls: inp2} 
		err = urls2.VerifyURLs()
		if err == nil {
			t.Fatal("unexpected result: wanted error")
		}
	})
}