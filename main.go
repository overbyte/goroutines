package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://overbyte.co.uk",
		"https://google.com",
		"https://facebook.com",
		"https://twitter.com",
		"https://stackoverflow.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	fmt.Println("checking", link)

	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Error opening link", link, ":", err)
		c <- link
		return
	}

	fmt.Println(">", link, "is all good")
	c <- link

}
