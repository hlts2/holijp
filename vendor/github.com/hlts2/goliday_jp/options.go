package goliday

import "time"

// Option configures holiday search.
type Option func(*evaluateOption)

type evaluateOption struct {
	year  int
	month int
	day   int
}

func newEvaluateOption() *evaluateOption {
	return &evaluateOption{
		year:  -1,
		month: -1,
		day:   -1,
	}
}

// WithYear returns an option that sets the year.
func WithYear(year int) Option {
	return func(e *evaluateOption) {
		e.year = year
	}
}

// WithMonth returns an option that sets the month.
func WithMonth(month int) Option {
	return func(e *evaluateOption) {
		e.month = month
	}
}

// WithDay returns an option that sets the day.
func WithDay(day int) Option {
	return func(e *evaluateOption) {
		e.day = day
	}
}

// WithTime returns an option that sets the year and month and day from time.Time instance.
func WithTime(t *time.Time) Option {
	return func(e *evaluateOption) {
		e.year = t.Year()
		e.month = int(t.Month())
		e.day = t.Day()
	}
}
