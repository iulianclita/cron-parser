package cron_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/iulianclita/cron-parser/cron"
)

func TestExtractValuesInInterval(t *testing.T) {
	testCases := map[string]struct {
		input     string
		min, max  int
		values    []int
		wantError bool
	}{
		"invalid input": {
			input:     "invalid input",
			min:       0,
			max:       59,
			values:    nil,
			wantError: true,
		},
		"input is all values (*)": {
			input:     "*",
			min:       1,
			max:       12,
			values:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			wantError: false,
		},
		"input is of type */x but x is not a number:": {
			input:     "*/not-a-number",
			min:       1,
			max:       31,
			values:    nil,
			wantError: true,
		},
		"input is of type */x but x is not between min and max:": {
			input:     "*/10",
			min:       1,
			max:       7,
			values:    nil,
			wantError: true,
		},
		"input is of type */x and x is valid": {
			input:     "*/15",
			min:       0,
			max:       59,
			values:    []int{0, 15, 30, 45},
			wantError: false,
		},
		"input is of type x,y,z but z is not a number": {
			input:     "10,20,not-a-number",
			min:       0,
			max:       59,
			values:    nil,
			wantError: true,
		},
		"input is of type x,y,z but z is not between min and max": {
			input:     "10,20,99",
			min:       0,
			max:       59,
			values:    nil,
			wantError: true,
		},
		"input is of type x,y,z and all values are valid": {
			input:     "10,20,30",
			min:       0,
			max:       59,
			values:    []int{10, 20, 30},
			wantError: false,
		},
		"input is of type x-y-z but this is invaliud format": {
			input:     "10-20-30",
			min:       0,
			max:       59,
			values:    nil,
			wantError: true,
		},
		"input is of type x-y but x is not a number": {
			input:     "notANumber-10",
			min:       0,
			max:       59,
			values:    nil,
			wantError: true,
		},
		"input is of type x-y but y is not a number": {
			input:     "10-notANumber",
			min:       0,
			max:       59,
			values:    nil,
			wantError: true,
		},
		"input is of type x-y but z is not between min and max": {
			input:     "10-99",
			min:       0,
			max:       59,
			values:    nil,
			wantError: true,
		},
		"input is of type x-y but y is smaller than x": {
			input:     "15-10",
			min:       0,
			max:       59,
			values:    nil,
			wantError: true,
		},
		"input is of type x-y and all values are valid": {
			input:     "10-20",
			min:       0,
			max:       59,
			values:    []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			wantError: false,
		},
		"input is not a valid number": {
			input:     "xxx",
			min:       0,
			max:       59,
			values:    nil,
			wantError: true,
		},
		"input is not a valid number but not between min and max": {
			input:     "99",
			min:       0,
			max:       59,
			values:    nil,
			wantError: true,
		},
		"input is a valid number": {
			input:     "30",
			min:       0,
			max:       59,
			values:    []int{30},
			wantError: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			gotValues, err := cron.ExtractValuesInInterval("minute", tc.input, tc.min, tc.max)
			if tc.wantError && err == nil {
				t.Fatalf("expected error but got no error")
			}
			if !tc.wantError && err != nil {
				t.Fatalf("expected no error but got an error")
			}

			wantValues := tc.values
			if !cmp.Equal(gotValues, wantValues) {
				t.Errorf("wanted %v but got %v. diff: %v", wantValues, gotValues, cmp.Diff(gotValues, wantValues))
			}
		})
	}
}
