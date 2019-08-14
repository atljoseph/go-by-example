package main

import (
	"fmt"
	"time"
)

func main() {

	// alias the function
	p := fmt.Println

	// now
	now := time.Now()
	p("now", now)

	// then
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p("then", then)

	// various time segments
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())

	// location = time zone
	p(then.Location())

	p(then.Weekday())

	// comparisons
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// difference / Subtract
	diff := now.Sub(then)
	p(diff)

	// time span
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// Add
	p(then.Add(diff))
	p(then.Add(-diff))

}
