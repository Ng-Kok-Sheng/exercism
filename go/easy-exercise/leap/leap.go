// Package leap is a small library for determining whether the passed in year is a leap year
package leap

// IsLeapYear accepts year as an integer.
// Return true if the year is a leap year, else false.
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
