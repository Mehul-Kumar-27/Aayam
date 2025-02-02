package utils

import (
	"fmt"
	"reflect"
)

// AssertEqual is a generic function that can compare any comparable types
func AssertEqual[T comparable](got, expected T) bool {
	var logEntry Log

	if got == expected {
		return true
	}

	logEntry = Log{
		Level: WarningLevel,
		Info: fmt.Sprintf("Assertion failed: expected %v (type %v), got %v (type %v)",
			expected,
			reflect.TypeOf(expected),
			got,
			reflect.TypeOf(got)),
	}
	logger := NewLogger()
	logger.LogMessage(logEntry)
	return false
}

// AssertNotEqual is a generic function that checks if values are not equal, the first argument is the value that we have
// and the second argument represent the value we do not expect the the first argument to be equal to.
func AssertNotEqual[T comparable](got, unexpected T) bool {
	var logEntry Log

	if got != unexpected {
		return true
	}

	logEntry = Log{
		Level: WarningLevel,
		Info: fmt.Sprintf("Assertion failed: value should not equal %v (type %v)",
			unexpected,
			reflect.TypeOf(unexpected)),
	}
	logger := NewLogger()
	logger.LogMessage(logEntry)
	return false
}

// AssertRange is a generic function that checks if the value that we have is in the range that we expect
func AssertRange[T comparable](got, min_val, max_val int) bool {
	var logEntry Log

	if got >= min_val && got <= max_val {
		return true
	}

	logEntry = Log{
		Level: WarningLevel,
		Info: fmt.Sprintf("Assertion failed: value %v (type %v) not in range [%v, %v]",
			got,
			reflect.TypeOf(got),
			min_val,
			max_val),
	}
	logger := NewLogger()
	logger.LogMessage(logEntry)
	return false
}

// AssertDeepEqual is a generic function that uses reflect.DeepEqual to compare two values
func AssertDeepEqual(got, expected interface{}) bool {
	var logEntry Log

	if reflect.DeepEqual(got, expected) {
		return true
	}

	logEntry = Log{
		Level: WarningLevel,
		Info: fmt.Sprintf("Assertion failed: expected %v (type %v), got %v (type %v)",
			expected,
			reflect.TypeOf(expected),
			got,
			reflect.TypeOf(got)),
	}
	logger := NewLogger()
	logger.LogMessage(logEntry)
	return false
}
