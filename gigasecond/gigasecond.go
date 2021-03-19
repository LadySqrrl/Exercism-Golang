//Add a gigasecond to the time
package gigasecond

import "time"

// Add a gigasecond to the time
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * 1000000000)
	return t
}
