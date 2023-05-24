package xtime

import "time"

// String2timestamp ...
func String2timestamp(str string, layout string) (timestamp int64, err error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return
	}

	t, err := time.ParseInLocation(layout, str, loc)
	if err != nil {
		return
	}

	timestamp = t.UnixNano() / 1e6
	return
}

// GetTimestampInMilli ...
func GetTimestampInMilli() int64 {
	return int64(time.Now().UnixNano() / 1e6)
}

// Elapse ...
func Elapse(f func()) int64 {
	now := time.Now().UnixNano()
	f()
	return time.Now().UnixNano() - now
}
