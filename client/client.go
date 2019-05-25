package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

func knockKnock() (string, error) {
	req, err := http.NewRequest("GET", "http://localhost:8080/knockKnock", nil)
	if err != nil {
		return "", err
	}
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

func whoIsThere() (string, error) {
	req, err := http.NewRequest("GET", "http://localhost:8080/whoIsThere", nil)
	if err != nil {
		return "", err
	}
	req.WithContext(context.WithValue(req.Context(), "jokes", "Yes")) // Key should not be a string but I am ignoring that on purpose
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
	str, err := knockKnock()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
	str, err = whoIsThere()
	if err != nil {
		fmt.Println()
	}
	fmt.Println(str)
}
