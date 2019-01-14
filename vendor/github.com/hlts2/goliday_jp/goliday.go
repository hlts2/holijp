package goliday

// HasHoliday returns true if the option values exists within the holiday.
func HasHoliday(ops ...Option) bool {
	return table.contains(evaluatedOption(ops...))
}

// Holidays returns Holiday slice that matches options.
func Holidays(ops ...Option) []Holiday {
	return table.holidaysByEvaluateOption(evaluatedOption(ops...))
}

func evaluatedOption(ops ...Option) *evaluateOption {
	evalOp := newEvaluateOption()
	for _, op := range ops {
		op(evalOp)
	}
	return evalOp
}
