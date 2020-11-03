package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errErrorHandler = errors.New("Request Error")

func main() {
	// 맵을 정의 해줄땐 선언할때 무조건 할당을 해줘야됨

	// 방법1
	// result := make(map[string]string)
	// 방법2 // 이런식으로 지정해주지 않으면 map이 nil로 변
	results := map[string]string{}


	urls := []string {
		"https://www.airbnb.com",
		"https://www.facebook.com",
		"https://www.google.com",
	}

	for _,url := range urls {
		result := "OK"

		err := hitURL(url)
		if err != nil {
			result = "Fail"
		}
		results[url] = result
	}

	for url, result := range results {
		// 키와 벨류 분
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)

	if err != nil || resp.StatusCode >= 400 {
		return errErrorHandler
	}
	return nil
}