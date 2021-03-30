package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// declared a slice of strings
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"https://na.leagueoflegends.com/en-us/",
		"https://www.youtube.com/",
		"https://www.reddit.com/",
		"http://jesuslives.com/",
		// "http://poggerwoggerwidepeepohappy.com/",
	}

	// channel to establish comm with child routines
	c := make(chan string)

	for _, link := range links {

		go checkLink(link, c)

		// attempt 1 at using one round of channel calls and use index to initiate the 5 seconds sleep
		// if i == (len(links) - 1) {
		// 	time.Sleep(time.Second * 5)
		// 	go checkLink(link, c)
		// } else {
		// 	go checkLink(link, c)
		// }

	}

	// this loop creates as many receiving channels as the links we have
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(c)
	// }

	// // infinite loop
	// for {
	// 	go checkLink(<-c, c)
	// }

	// alternative syntax tellking the checkLink fn to run whenever a c chan receives a string val, take the val (l) and run checkLink again
	for l := range c {

		// fn literal or anonymous fn in js
		// we are doing this to prevent the main routine from sleeping, and we don't want the checkLink to sleep
		// this is purgatory
		go func(link string) {
			time.Sleep(5 * time.Second)
			// go knows this first channel (<-c is actually the result of our earlier call) is going to produce a string, this link ie: "google.com" is passed along with another channel into the checkLink fn
			checkLink(link, c)

			// (l)
			// to avoid loop variable l captured by func literalloopclosure
			// we want to pass l into the function call
			// the function is expecting a string of link
			// (l)
		}(l)
	}

}

//                          c channel expects strings to for comm
func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link+" might be down ", err)

		c <- link
		// return here so the fn stops and go next
		return
	}

	c <- link
	fmt.Println(link + " is currently up and running!")
}
