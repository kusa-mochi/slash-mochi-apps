package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"time"
)

func main() {
	s256 := sha256.New()

	now := time.Now().Format("2000-01-01 11:11:11.000000000")
	io.WriteString(s256, now)

	h := s256.Sum(nil)
	hstr := fmt.Sprintf("%x", h)
	fmt.Println(hstr)
}
