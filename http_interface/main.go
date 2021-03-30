package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("cannot get response from google: ", err)
		os.Exit(1)
	}

	// make a byte slice for the read function
	bs := make([]byte, 99999)
	// read function translates incoming data into bytes and stores it in the bs
	resp.Body.Read(bs)
	// printing out bs as string type
	fmt.Println(string(bs))

	// The equivalent of doing the above
	io.Copy(os.Stdout, resp.Body)

}
