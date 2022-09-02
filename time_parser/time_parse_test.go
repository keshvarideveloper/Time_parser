package timeparse

import "testing"

func TestParseTime(t *testing.T) {
	table := []struct {
		time string
		ok   bool
	}{
		{"22:05:59", true},
		{"4:5:9", true},
		{"12:00:00", true},
		{"00:00:00", true},
		{"-1:05:00", false},
		{"23:05:60", false},
		{"15:65:01", false},
		{"s:h:59", false},
	}

	for _, data := range table {
		_, err := ParserTime(data.time)
		if data.ok && err != nil {
			t.Errorf("%v: %v, error should be nil", data.time, err)
		}
	}

}
