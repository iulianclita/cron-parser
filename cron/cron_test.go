package cron_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/iulianclita/cron-parser/cron"
)

func TestExtractMinutes(t *testing.T) {
	testCases := map[string]struct {
		input     string
		values    []int
		wantError bool
	}{
		"invalid input": {
			input:     "invalid input",
			values:    nil,
			wantError: true,
		},
		"input is all values (*)": {
			input: "*",
			values: []int{
				0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
				11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
				21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
				31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
				41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
				51, 52, 53, 54, 55, 56, 57, 58, 59,
			},
			wantError: false,
		},
		"input is of type */x but x is not a number:": {
			input:     "*/not-a-number",
			values:    nil,
			wantError: true,
		},
		"input is of type */x but x is not between 0-59:": {
			input:     "*/99",
			values:    nil,
			wantError: true,
		},
		"input is of type */x and x is valid": {
			input:     "*/15",
			values:    []int{0, 15, 30, 45},
			wantError: false,
		},
		"input is of type x,y,z but z is not a number": {
			input:     "10,20,not-a-number",
			values:    nil,
			wantError: true,
		},
		"input is of type x,y,z but z is not between 0-59": {
			input:     "10,20,99",
			values:    nil,
			wantError: true,
		},
		"input is of type x,y,z and all values are valid": {
			input:     "10,20,30",
			values:    []int{10, 20, 30},
			wantError: false,
		},
		"input is of type x-y but z is not a number": {
			input:     "10-notANumber",
			values:    nil,
			wantError: true,
		},
		"input is of type x-y but z is not between 0-59": {
			input:     "10-99",
			values:    nil,
			wantError: true,
		},
		"input is of type x-y and all values are valid": {
			input:     "10-20",
			values:    []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			wantError: false,
		},
		"input is not a valid number": {
			input:     "xxx",
			values:    nil,
			wantError: true,
		},
		"input is a valid number": {
			input:     "30",
			values:    []int{30},
			wantError: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			gotValues, err := cron.ExtractMinutes(tc.input)
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
