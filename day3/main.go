package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	var res = map[string]string{}
	urls := []string{
		"https://kitae0522.kro.kr",
		"https://www.google.co.kr",
		"https://www.naver.com",
	}
	for _, url := range urls {
		status := "OK"
		err := CheckURL(url)
		if err != nil {
			status = "Falied"
		}
		res[url] = status
	}
	for url, status := range res {
		fmt.Println(url, status)
	}
}

// CheckURL is Checking URL
func CheckURL(url string) error {
	fmt.Println("Checking URL: ", url)
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		fmt.Println(err, res.StatusCode)
		return errRequestFailed
	}
	return nil
}
