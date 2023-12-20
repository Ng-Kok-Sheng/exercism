package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	var floatString = strconv.FormatFloat(f, 'f', 1, 64)
	return fmt.Sprintf("This is the number %s", floatString)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	number := float64(nb.Number())
	return fmt.Sprintf("This is a box containing the number %.1f", number)
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	isFancyNumber, ok := fnb.(FancyNumber)

	if !ok {
		return 0
	}

	number, _ := strconv.Atoi(isFancyNumber.n)
	return number
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	extractedFancyNumber := float64(ExtractFancyNumber(fnb))
	return fmt.Sprintf("This is a fancy box containing the number %.1f", extractedFancyNumber)
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	integerNumber, isInteger := i.(int)
	if isInteger {
		return DescribeNumber(float64(integerNumber))
	}

	floatNumber, isFloatNumber := i.(float64)
	if isFloatNumber {
		return DescribeNumber(floatNumber)
	}

	numberBox, isNumberBox := i.(NumberBox)
	if isNumberBox {
		return DescribeNumberBox(numberBox)
	}

	fancyNumberBox, isFancyNumberBox := i.(FancyNumberBox)
	if isFancyNumberBox {
		return DescribeFancyNumberBox(fancyNumberBox)
	}

	return "Return to sender"
}
