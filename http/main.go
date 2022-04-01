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
		fmt.Println("Error :", err)
		os.Exit(1)
	}
	fmt.Println(resp.Body)

	// Print HTML
	// bs := make([]byte, 9999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// M2 - Print HTML
	io.Copy(os.Stdin, resp.Body)

	// M3 - Custom writer interface
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (n int, e error) {
	fmt.Println(bs)
	fmt.Println("Just wrote this many bytes : %v", len(bs))
	return len(bs), nil
}
