package main

import (
	"net/http"
	"fmt"
)


func main() {
	links := []string {
		"https://google.com.br",
		"https://uol.com.br",
		"https://terra.com.br",
		"https://amazon.com",
		"https://golang.org",
		"https://apple.com.br",
	}

	for _, link := range links {
		checkLink(link)
	}

}

func checkLink(link string)  {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
		return
	}

	fmt.Println(link, " is UP")
}