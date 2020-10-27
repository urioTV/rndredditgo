package rndredditgo

import (
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/imroc/req"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func RndImg(sub string) string {

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: tr,
	}

	url := fmt.Sprintf("https://www.reddit.com/r/%s.json?sort=hot", sub)
	header := make(http.Header)
	header.Set("User-Agent", "TestBot")
	r := req.New()
	r.SetClient(client)
	resp, err := r.Get(url, header)
	if err != nil {
		log.Fatal(err)
	}

	data, err := resp.ToBytes()

	if err != nil {
		log.Fatal(err)
	}
	aLen, err := jsonparser.GetInt(data, "data", "dist")
	if aLen == 0 {
		return "Invalid subreddit"
	}
	rand.Seed(time.Now().UnixNano())
	rndLen := rand.Intn(int(aLen) - 1)
	strRndImg := fmt.Sprintf("[%d]", rndLen)
	if err != nil {
		log.Fatal(err)
	}

	s, err := jsonparser.GetString(data, "data", "children", strRndImg, "data", "url")

	if err != nil {
		log.Fatal(err)
	}

	return s
}
