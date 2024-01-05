package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(calculator FodderCalculator, amountOfCows int) (float64, error) {
	fodderAmount, fodderErr := calculator.FodderAmount(amountOfCows)
	flatteningFactor, flatteningErr := calculator.FatteningFactor()

	if fodderErr != nil {
		return 0, fodderErr
	}

	if flatteningErr != nil {
		return 0, flatteningErr
	}

	return fodderAmount * flatteningFactor / float64(amountOfCows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(calculator FodderCalculator, amountOfCows int) (float64, error) {
	if amountOfCows < 1 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(calculator, amountOfCows)
}

type InvalidCowsError struct {
	amountOfCows int
	message      string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.amountOfCows, e.message)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(amountOfCows int) error {
	if amountOfCows < 0 {
		return &InvalidCowsError{
			amountOfCows: amountOfCows,
			message:      "there are no negative cows",
		}
	}

	if amountOfCows == 0 {
		return &InvalidCowsError{
			amountOfCows: amountOfCows,
			message:      "no cows don't need food",
		}
	}

	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
