package healthcheckergo

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/sirupsen/logrus"
)

type HealthStatus string

const (
	Active HealthStatus = "active"
	Inactive = "inactive"
	Error = "error"
)

type URLStatus struct {
	URL    string       `json:"url"`
	Status HealthStatus `json:"status"`
}

type Pinger struct {
	cli http.Client
	stopOnError bool
	log *logrus.Logger
}

func NewPinger(stopOnError bool, timeout int, log *logrus.Logger) *Pinger {
	return &Pinger {
		cli: http.Client{Timeout: time.Duration(timeout)*time.Second},
		stopOnError: stopOnError,
		log: log,
	}
}

func (p *Pinger) PingEm(urls []string) ([]URLStatus, error) {
	
	var mx sync.Mutex

	result := make([]URLStatus,0,len(urls))
	resc := make(chan URLStatus)
	defer close(resc)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	
	for _, u := range urls {
		go p.pingSingle(u, ctx, resc)
	}

	for {
		select {
		case res := <- resc:
			if (res.Status == Inactive || res.Status == Error) && p.stopOnError {
				cancel()
				return nil, fmt.Errorf("failed to ping: %s",res.URL)
			}
			mx.Lock()
			result = append(result, res)
			mx.Unlock()
			if len(result) == len(urls){
				return result, nil
			}
		}
	}

}

func (p *Pinger) pingSingle(url string, ctx context.Context, resc chan URLStatus) {
	res := URLStatus{URL: url}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("failed to fetch %s. Cause: %s", url, err.Error())
		res.Status = Error
		resc <- res
		return
	}
	resp,err := p.cli.Do(req)
	if err != nil {
		res.Status = Inactive
		resc <- res
		return 
	}
	if resp.StatusCode != http.StatusOK || err != nil {
		res.Status = Inactive
	} else {
		res.Status = Active
	}
	resc <- res
}


// func PingEm(urls []string, ) ([]HealthStatus, error) {
// 	cli := http.Client{Timeout: time.Duration(timeout)*time.Second}
// 	return nil, nil
//}