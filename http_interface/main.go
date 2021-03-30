package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("cannot get response from google: ", err)
		os.Exit(1)
	}

	// // make a byte slice for the read function
	// bs := make([]byte, 99999)
	// // read function translates incoming data into bytes and stores it in the bs
	// resp.Body.Read(bs)
	// // printing out bs as string type
	// fmt.Println(string(bs))

	lw := logWriter{}

	// The equivalent of doing the above
	// Copy fn expects a dst Writer fn, and src Reader fn, which are satisfied by os.Stdout.Writer, and resp.Body.Reader
	// copyBuffer then is called and using src dst to output values from http call

	// we pass in our os.stdout (logWriter) method in to copy
	// it works because Copy expects the Write fn and implicitly gives us a new type interface
	// any fn that has the Write fn satisfies the interface guideline
	// this is similar to doing io.Copy(os.stdout, resp.Body)

	io.Copy(lw, resp.Body)

}

// this our implimentation of the os.stdout method in go
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes")
	return len(bs), nil
}
