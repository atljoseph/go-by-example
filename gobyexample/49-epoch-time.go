package main

import (
	"fmt"
	"time"
)

func main() {

	p := fmt.Println

	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	p(now)

	millisecs := nanos / 1000000
	p(secs)
	p(millisecs)
	p(nanos)

	// convert secs or nanos
	// into time object
	p(time.Unix(secs, 0))
	p(time.Unix(0, nanos))

}
