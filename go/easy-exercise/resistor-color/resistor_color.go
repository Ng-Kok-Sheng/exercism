package resistorcolor

// Colors returns the list of all colors.
func Colors() []string {
	return []string{
		"black",
		"brown",
		"red",
		"orange",
		"yellow",
		"green",
		"blue",
		"violet",
		"grey",
		"white",
	}
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	colorsList := Colors()

	for i, v := range colorsList {
		if v == color {
			return i
		}
	}

	panic("No value found.")
}
