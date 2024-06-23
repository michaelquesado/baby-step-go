package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	payload := bytes.NewBuffer([]byte(`{ "foo": "bar" }`))
	res, err := c.Post("https://www.google.com", "application/json", payload)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	io.CopyBuffer(os.Stdout, res.Body, nil)
}
