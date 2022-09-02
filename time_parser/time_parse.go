package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hour, minute, second int
}

type TimeParserError struct {
	msg   string
	input string
}

func (t *TimeParserError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

func ParserTime(input string) (Time, error) {
	component := strings.Split(input, ":")
	if len(component) != 3 {
		return Time{}, &TimeParserError{"Invalid number of time components", input}
	}
	hour, err := strconv.Atoi(component[0])
	if err != nil {
		return Time{}, &TimeParserError{fmt.Sprintf("Error parsing hour: %v", err), input}
	}

	minute, err := strconv.Atoi(component[1])
	if err != nil {
		return Time{}, &TimeParserError{fmt.Sprintf("Error parsing minute: %v", err), input}
	}

	second, err := strconv.Atoi(component[2])
	if err != nil {
		return Time{}, &TimeParserError{fmt.Sprintf("Error parsing second: %v", err), input}
	}

	if hour > 23 || hour < 0 {
		return Time{}, &TimeParserError{"Hour out of range : 0 <= hour <= 23", fmt.Sprintf("%v", hour)}
	}

	if minute > 59 || minute < 0 {
		return Time{}, &TimeParserError{"Minute out of range : 0 <= minute <= 59", fmt.Sprintf("%v", minute)}
	}

	if second > 59 || second < 0 {
		return Time{}, &TimeParserError{"Second out of range : 0 <= second <= 59", fmt.Sprintf("%v", second)}
	}

	return Time{hour, minute, second}, nil

}
