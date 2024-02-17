package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}

	charactersRegExp := regexp.MustCompile(`^[0-9]+$`)
	matches := charactersRegExp.MatchString(id)
	if !matches {
		return false
	}

	var digRev []int
	ss := []rune(id)
	for i := len(ss) - 1; i >= 0; i-- {
		num, err := strconv.Atoi(string(ss[i]))
		if err != nil {
			return false
		}

		digRev = append(digRev, num)
	}

	totalSum := 0
	for i, n := range digRev {
		if i%2 != 0 {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}

		totalSum += n
	}

	if totalSum%10 == 0 {
		return true
	}
	return false
}
