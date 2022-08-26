package main

import (
	"fmt"
	"os"
)

func main() {
	const (
		sublink         = "https://"
		length          = len(sublink)
		mask            = '*'
		lengthOfWebsite = len("cnn.com")
	)
	link := os.Args[1:]
	website := link[0]
	lenOfLink := len(website)
	if lenOfLink <= 0 {
		fmt.Println("Please give me a link to work with")
		return
	}
	var buf []byte
	// buf = "https://"
	for i := 0; i < lenOfLink; i++ {
		if i >= length && i < lengthOfWebsite+length {
			buf = append(buf, mask)
		} else {
			buf = append(buf, website[i])
		}
		// fmt.Println(string(buf))
	}

	fmt.Println(string(buf[0:]))
}
