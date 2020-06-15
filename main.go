package main

import (
	"fmt"
	"net/http"
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

	for i := 0; i < len(links); i++ {
		fmt.Println("received", <-c)
	}
}

func checkLink(link string, c chan string) {
	fmt.Println("checking", link)

	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Error opening link", link, ":", err)
		c <- link
		return
	} else {
		fmt.Println(">", link, "is all good")
		c <- link
	}

}
