package healthcheckergo_test

import (
	"context"
	hth "healthchecker"
	"net/http"
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func sortURLstatus(arr []hth.URLStatus) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].URL < arr[j].URL
	})
}

func mockHandler(timeout int) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(timeout)*time.Second)
		w.WriteHeader(http.StatusOK)
	}
}

func demoServer(timeout int, port string, t *testing.T) *http.Server {
	server := http.Server{Addr: port}
	go func() {
		err := http.ListenAndServe(port,server.Handler)
		if err != nil {
			t.Log(err)
		}
	}()
	return &server
}

func TestPinger(t *testing.T) {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	
	t.Run("default pinger", func(t *testing.T) {
		
		pinger1 := hth.NewPinger(false, 3,log)

		inp1 := []string{"https://google.com","https://duckduckgo.com","https://github.com","https://stackoverflow.com"}
		expected1 := []hth.URLStatus{
			{"https://google.com",hth.Active}, 
			{"https://duckduckgo.com",hth.Active},
			{"https://github.com",hth.Active},
			{"https://stackoverflow.com",hth.Active},
		}
		sortURLstatus(expected1)
		res1, err := pinger1.PingEm(inp1)
		if err != nil {
			t.Fatalf("pingem error: %s",err.Error())
		}
		sortURLstatus(res1)
		if !reflect.DeepEqual(res1, expected1) {
			t.Fatalf("unexpected result")
		}

		// these might change
		inp2 := []string{"http://localhost:8079","duckduckgo.com","https://github.com","https://stackoverflow.com"}
		expected2 := []hth.URLStatus{
			{"http://localhost:8079",hth.Inactive}, // just inactive
			{"duckduckgo.com",hth.Inactive}, // expected protocol
			{"https://github.com",hth.Active},
			{"https://stackoverflow.com",hth.Active},
		}
		sortURLstatus(expected2)
		res2, err := pinger1.PingEm(inp2)
		if err != nil {
			t.Fatalf("pingem error: %s",err.Error())
		}
		sortURLstatus(res2)
		if !reflect.DeepEqual(res2, expected2) {
			t.Fatalf("unexpected result")
		}
	})

	http.HandleFunc("/",mockHandler(2))

	t.Run("timeout OK", func(t *testing.T) {
		server1 := demoServer(2,"localhost:8080", t)
		
		pinger1 := hth.NewPinger(false, 3,log)
		
		inp1 := []string{"http://localhost:8080/"}
		expected1 := []hth.URLStatus{
			{"http://localhost:8080/",hth.Active},
		}
		res1, err := pinger1.PingEm(inp1)
		if err != nil {
			t.Fatalf("pingem error: %s",err.Error())
		}
		if !reflect.DeepEqual(res1, expected1) {
			t.Fatalf("unexpected result")
		}

		
		server1.Shutdown(context.Background())
	})

	t.Run("timeout overrun", func(t *testing.T) {
		pinger2 := hth.NewPinger(false, 1,log)
		server1 := demoServer(2,"localhost:8081", t)

		inp2 := []string{"http://localhost:8081/"}
		expected2 := []hth.URLStatus{
			{"http://localhost:8081/",hth.Inactive},
		}
		res2, err := pinger2.PingEm(inp2)
		if err != nil {
			t.Fatalf("pingem error: %s",err.Error())
		}
		if !reflect.DeepEqual(res2, expected2) {
			t.Fatalf("unexpected result")
		}
		server1.Shutdown(context.Background())
	})

}