package expenses

import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(records []Record, predicate func(Record) bool) []Record {
	var filteredRecords []Record
	for _, v := range records {
		if predicate(v) {
			filteredRecords = append(filteredRecords, v)
		}
	}
	return filteredRecords
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(record Record) bool {
		if record.Day >= p.From && record.Day <= p.To {
			return true
		}

		return false
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(record Record) bool {
		if c == record.Category {
			return true
		}
		return false
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(records []Record, p DaysPeriod) float64 {
	amount := 0.0

	for _, record := range records {
		if record.Day >= p.From && record.Day <= p.To {
			amount += record.Amount
		}
	}

	return amount
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(records []Record, p DaysPeriod, c string) (float64, error) {
	filteredRecords := Filter(records, ByCategory(c))

	if len(filteredRecords) == 0 {
		return 0.0, errors.New("unknown category entertainment")
	}

	filteredRecords = Filter(filteredRecords, ByDaysPeriod(p))

	if len(filteredRecords) == 0 {
		return 0.0, nil
	}

	return TotalByPeriod(filteredRecords, p), nil
}
