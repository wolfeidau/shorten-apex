package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/apex/go-apex"
	"github.com/wolfeidau/shorten"
)

const (
	domain = "https://s.wolfe.id.au/"

	// ~83733937890625 should be enough random values
	// this assumes 55^8
	length = 8
)

type message struct {
	ShortURL  string `json:"shortUrl"`
	URL       string `json:"url"`
	Timestamp int64  `json:"timestamp"`
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var m message

		if err := json.Unmarshal(event, &m); err != nil {
			return nil, err
		}

		if err := validateURL(m.URL); err != nil {
			return nil, err
		}

		m.ShortURL = domain + shorten.RandSeq(length)
		m.Timestamp = time.Now().UnixNano()

		return m, nil
	})
}

func validateURL(u string) (err error) {

	var url *url.URL
	if url, err = url.Parse(u); err != nil {
		return err
	}

	switch url.Scheme {
	case "http":
		return nil
	case "https":
		return nil
	default:
		return fmt.Errorf("Unsupported URL scheme: %s", url.Scheme)
	}
}
