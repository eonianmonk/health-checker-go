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
		_, err := url.Parse(u)
		if err != nil {
			return fmt.Errorf("failed to parse url: %s", u)
		}
	}
	return nil
}