package gofetch

import (
	"io"
	"net/http"
)

func Get(URL string) io.Reader {
	res, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	return res.Body
}

func Post(URL string, ContentType string, Body io.Reader) io.Reader {
	res, err := http.Post(URL, ContentType, Body)
	if err != nil {
		panic(err)
	}
	return res.Body
}
