package dt

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Dtols returns a long date format i.e. Sunday August 18, 2019
func Dtols(dt time.Time) string {
	return fmt.Sprint(dt.Weekday().String(), " ", dt.Month().String(), " ", dt.Day(), ", ", dt.Year())
}

// Today returns todays date with no time component.
func Today() time.Time {
	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)
}

// Stod returns a date time from a string in yyyy-MM-dd, yyyy/mm/dd, mm-dd-yyyy, or mm/dd/yyyy
func Stod(value string) (time.Time, error) {
	sep := "-"
	if strings.ContainsRune(value, rune('/')) {
		sep = "/"
	}
	parts := strings.Split(value, sep)
	a, e := strconv.Atoi(parts[0])
	if e != nil {
		return time.Now(), e
	}
	b, e := strconv.Atoi(parts[1])
	if e != nil {
		return time.Now(), e
	}
	c, e := strconv.Atoi(parts[2])
	if e != nil {
		return time.Now(), e
	}
	var y, m, d int
	if a > 31 {
		y = a
		m = b
		d = c
	} else {
		m = a
		d = b
		y = c
	}
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local), nil
}

// Dtos returns a short date formatted as yyyy-MM-dd
func Dtos(dt time.Time) string {
	return fmt.Sprintf("%4d-%2d-%2d", dt.Year(), dt.Month(), dt.Day())
}
