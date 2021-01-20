package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://ebuinvest.com",
	}
	// Make channel to communicate between Main routine
	// and Child routines.
	c := make(chan string)

	for _, url := range links {
		// launch additional go routines
		go checkURL(url, c)
	}

	// Receiving messages from all of the
	// child routines.
	// i := 0; i < len(links); i++

	// looping for infinity to check websites
	// for {
	// 	go checkURL(<-c, c)
	// }

	for l := range c {
		go func(url string) {
			time.Sleep(time.Second * 5)
			checkURL(url, c)
		}(l)
	}
	// Receiving message from channel.
	// fmt.Println(<-c)
	// os.Exit(0)
}

func checkURL(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down!")
		c <- url
		return
	}
	fmt.Println(url, "seems working!")
	c <- url
	return
}
