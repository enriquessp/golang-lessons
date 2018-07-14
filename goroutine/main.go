package main

import (
	"time"
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

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

//	for i := 0; i < len(links); i++ {
//		fmt.Println(<- c)
//	}
	
//	for {
//		go checkLink(<- c, c)
//	}

	for l := range c {
//		go checkLink(l, c)
		go func(m string) {
			time.Sleep(5 * time.Second)
			checkLink(m, c)
		}(l)
	}
	
}

func checkLink(link string, c chan string)  {
//	time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
//		c <- "Might be down"
		c <- link
		return
	}

	fmt.Println(link, " is UP")
//	c <- "Yes, it is UP"
	c <- link
}