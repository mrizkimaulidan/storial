package time

import "time"

// Generate current time to unix timestamp.
func CurrentTimeToUnixTimestamp() uint64 {
	n := time.Now().UnixNano() / 1000000

	return uint64(n)
}

// Format unix timestamp to time format.
func UnixToTime(n uint64) time.Time {
	t := (int64(n) / 1000)

	return time.Unix(t, 0)
}
