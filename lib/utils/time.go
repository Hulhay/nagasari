package utils

import (
	"time"
)

func LoadTimeLocation() (*time.Location, error) {
	loc, err := time.LoadLocation(TimeZone)
	if err != nil {
		return nil, err
	}

	return loc, nil
}

func Now() time.Time {
	loc, _ := LoadTimeLocation()

	return time.Now().In(loc)
}

func ToWIB(t time.Time, format string) string {
	loc, _ := LoadTimeLocation()

	return t.In(loc).Format(format)
}
