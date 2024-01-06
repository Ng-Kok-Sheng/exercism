package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("dna have different lengths")
	}

	differences := 0

	for i, _ := range a {
		if a[i] != b[i] {
			differences++
		}
	}

	return differences, nil
}
