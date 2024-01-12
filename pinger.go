package healthcheckergo

import (
	"fmt"
	"net/http"
	"sync/atomic"
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

type pingTask struct {
	active atomic.Bool
	resc chan URLStatus
}

func NewPinger(stopOnError bool, timeout int, log *logrus.Logger) *Pinger {
	return &Pinger {
		cli: http.Client{Timeout: time.Duration(timeout)*time.Second},
		stopOnError: stopOnError,
		log: log,
	}
}

func (p *Pinger) PingEm(urls []string) ([]URLStatus, error) {
	
	//var mx sync.Mutex
	result := make([]URLStatus,0,len(urls))
	task := pingTask{
		active: atomic.Bool{},
		resc: make(chan URLStatus),
	}
	task.active.Store(true)
	
	for _, u := range urls {
		go p.pingSingle(u, &task)
	}

	var err error
	for i := 0; i < len(urls); i++ {
		select {
		case res := <- task.resc:
			if (res.Status == Inactive || res.Status == Error) && p.stopOnError {
				task.active.Store(false)
				close(task.resc)
				return nil, fmt.Errorf("failed to ping: %s",res.URL)
			}
			result = append(result, res)
		}
	}
	return result, err


}

func (p *Pinger) pingSingle(url string, pt *pingTask) {
	res := URLStatus{URL: url}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("failed to fetch %s. Cause: %s", url, err.Error())
		res.Status = Error
		if pt.active.Load() {
			pt.resc <- res
		}
		return
	}
	resp,err := p.cli.Do(req)
	if err != nil {
		res.Status = Inactive
		if pt.active.Load() {
			pt.resc <- res
		}
		return 
	}
	if resp.StatusCode != http.StatusOK || err != nil {
		res.Status = Inactive
	} else {
		res.Status = Active
	}
	
	if pt.active.Load() {
		pt.resc <- res
	}
}


// func PingEm(urls []string, ) ([]HealthStatus, error) {
// 	cli := http.Client{Timeout: time.Duration(timeout)*time.Second}
// 	return nil, nil
//}