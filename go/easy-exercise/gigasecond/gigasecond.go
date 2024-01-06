// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package gigasecond uses seconds to express time intervals
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond adds 10 billion seconds to a provided time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
