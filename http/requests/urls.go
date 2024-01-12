package requests

import (
	"fmt"
	"net/url"
)

type UrlsRequest struct {
	Urls []string 
}

func (ur *UrlsRequest) VerifyURLs() error {
	for _,u := range ur.Urls {
		url, err := url.Parse(u)
		if err != nil {
			return fmt.Errorf("failed to parse url: %s", u)
		}
		if url.Scheme != "http" && url.Scheme != "https" {
			return fmt.Errorf("unknown scheme: %s", url.Scheme)
		}
	}
	return nil
}