package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)
	//time struct by providing the year,month,dat,etc.Times are associated with a location.
	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	//Extract the various components of the time value.
	p(then.Year())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())

	//Compare two times,before or after.
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	//a Duration representing the interval between two times.
	diff := now.Sub(then)
	p(diff)
	diff = then.Sub(now)
	p(diff)

	//compute the length of the duration in varous units.
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	//Add to advance a time by a given duration ,or with a - to move backway by a duration.
	p(then.Add(diff))
	p(then.Add(-diff))
}
