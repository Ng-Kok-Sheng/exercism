package raindrops

import "fmt"

const pling = "Pling"
const plang = "Plang"
const plong = "Plong"

func Convert(number int) string {
	sound := ""

	if number%3 == 0 {
		sound += pling
	}

	if number%5 == 0 {
		sound += plang
	}

	if number%7 == 0 {
		sound += plong
	}

	if sound == "" {
		return fmt.Sprint(number)
	}

	return sound
}
