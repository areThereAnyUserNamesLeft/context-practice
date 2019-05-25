package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

func showMeData() (string, error) {
	req, err := http.NewRequest("GET", "http://localhost:8080/showMeData", nil)
	if err != nil {
		return "", err
	}
	ctx, cancel := context.WithCancel(req.Context())
	defer cancel()
	req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(byt), nil
}

func main() {
	str, err := showMeData()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}
