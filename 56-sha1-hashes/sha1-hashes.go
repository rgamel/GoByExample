package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"
	h := sha1.New()

	h.Write([]byte(s)) // write expects bytes. s is a string and so needs to be coerced

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
