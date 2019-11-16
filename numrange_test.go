package numrange

import (
	"errors"
	"testing"
)

func getMinMax(is IntSet) (int, int) {
	// Find the minimum/maximum values in the set
	var min, max int

	firstIteration := true

	for k := range is {
		if firstIteration {
			min = k
			max = k
			firstIteration = false
		} else {
			if k < min {
				min = k
			}
			if k > max {
				max = k
			}
		}
	}

	return min, max

}

func TestIntSet(t *testing.T) {

	tests := []struct {
		description    string
		rangeInput     string
		expectedErr    error
		expectedMin    int
		expectedMax    int
		expectedString string
		InRange        []int
		NotInRange     []int

		addRangeInput       string
		expectedErrAdd      error
		InRangeAfterAdd     []int
		NotInRangeAfterAdd  []int
		expectedMinAfterAdd int
		expectedMaxAfterAdd int

		delRangeInput       string
		expectedErrDel      error
		InRangeAfterDel     []int
		NotInRangeAfterDel  []int
		expectedMinAfterDel int
		expectedMaxAfterDel int
	}{
		{
			description:    "test most range parseing, adding, deleting",
			rangeInput:     "1..3,5..7,10,12..15",
			expectedErr:    nil,
			expectedMin:    1,
			expectedMax:    15,
			expectedString: "1..3,5..7,10,12..15",
			InRange:        []int{5},
			NotInRange:     []int{20},

			addRangeInput:       "20..25",
			InRangeAfterAdd:     []int{22},
			NotInRangeAfterAdd:  []int{26},
			expectedMinAfterAdd: 1,
			expectedMaxAfterAdd: 25,

			delRangeInput:       "20..25",
			InRangeAfterDel:     []int{5},
			NotInRangeAfterDel:  []int{25},
			expectedMinAfterDel: 1,
			expectedMaxAfterDel: 15,
		},
		{
			description:    "test negative range",
			rangeInput:     "-10..10",
			expectedErr:    nil,
			expectedMin:    -10,
			expectedMax:    10,
			expectedString: "-10..10",
			InRange:        []int{-5},
			NotInRange:     []int{-11},
		},
		{
			description:    "test negative range with commas present",
			rangeInput:     "-10..10,15",
			expectedErr:    nil,
			expectedMin:    -10,
			expectedMax:    15,
			expectedString: "-10..10,15",
			InRange:        []int{-5},
			NotInRange:     []int{16},
		},
		{
			description:    "test that String() returns the most efficient string",
			rangeInput:     "1,2,3,4,5,7,9,10",
			expectedErr:    nil,
			expectedMin:    1,
			expectedMax:    10,
			expectedString: "1..5,7,9..10",
			InRange:        []int{1},
			NotInRange:     []int{11},
		},
		{
			description: "test that missing endpoint in range fails",
			rangeInput:  "1..",
			expectedErr: errors.New("unable to parse \"..\" ending point \"\": strconv.Atoi: parsing \"\": invalid syntax"),
		},
		{
			description: "test that missing starting point in range fails",
			rangeInput:  "..1",
			expectedErr: errors.New("unable to parse \"..\" starting point \"\": strconv.Atoi: parsing \"\": invalid syntax"),
		},
		{
			description: "test that more than one .. range fails",
			rangeInput:  "1..1..1",
			expectedErr: errors.New("invalid \"..\" range: 1..1..1"),
		},
		{
			description: "test that non-numeric input fails",
			rangeInput:  "a",
			expectedErr: errors.New("unable to parse range part a: strconv.Atoi: parsing \"a\": invalid syntax"),
		},
		{
			description:    "test that an empty input succeeds",
			rangeInput:     "",
			expectedString: "",
			expectedErr:    nil,
			NotInRange:     []int{16},
		},
		{
			description:    "test that Add() and Del() with broken input fails",
			rangeInput:     "",
			expectedString: "",
			expectedErr:    nil,
			NotInRange:     []int{16},

			addRangeInput:  "..",
			expectedErrAdd: errors.New("unable to parse \"..\" starting point \"\": strconv.Atoi: parsing \"\": invalid syntax"),

			delRangeInput:  "..",
			expectedErrDel: errors.New("unable to parse \"..\" starting point \"\": strconv.Atoi: parsing \"\": invalid syntax"),
		},
	}

	for _, test := range tests {
		is, err := ParseIntSet(test.rangeInput)

		if test.expectedErr == nil {
			if err != test.expectedErr {
				t.Errorf("%s: expected error: %s, got: %s", test.description, test.expectedErr, err)
			}
		} else {
			if err != nil {
				if err.Error() != test.expectedErr.Error() {
					t.Errorf("%s: expected error: %s, got: %s", test.description, test.expectedErr, err)
				}
			} else {
				t.Errorf("%s: expected error: %s, got: %s", test.description, test.expectedErrDel, err)
			}

		}

		if err == nil {

			if test.expectedString != is.String() {
				t.Errorf("%s: expected String(): %s, got: %s", test.description, test.expectedString, is.String())
			}

			if len(is) > 0 {

				min, max := getMinMax(is)

				if test.expectedMin != min {
					t.Errorf("%s: expected min: %d, got: %d", test.description, test.expectedMin, min)
				}

				if test.expectedMax != max {
					t.Errorf("%s: expected max: %d, got: %d", test.description, test.expectedMax, max)
				}

				for _, sc := range test.InRange {
					if !is.InRange(sc) {
						t.Errorf("%s: set should contain %d, but it does not", test.description, sc)
					}
				}
			}

			// We can check for non-existence even when the set is empty
			for _, snc := range test.NotInRange {
				if is.InRange(snc) {
					t.Errorf("%s: set should not contain %d, but it does", test.description, snc)
				}
			}

			if test.addRangeInput != "" {
				err := is.Add(test.addRangeInput)

				if test.expectedErrAdd == nil {
					if err != test.expectedErrAdd {
						t.Errorf("%s: expected Add() error: %s, got: %s", test.description, test.expectedErrAdd, err)
					}
				} else {
					if err != nil {
						if err.Error() != test.expectedErrAdd.Error() {
							t.Errorf("%s: expected Add() error: %s, got: %s", test.description, test.expectedErrAdd, err)
						}
					} else {
						t.Errorf("%s: expected Add() error: %s, got: %s", test.description, test.expectedErrDel, err)
					}

				}

				min, max := getMinMax(is)

				if test.expectedMinAfterAdd != min {
					t.Errorf("%s: expected min after Add(): %d, got: %d", test.description, test.expectedMinAfterAdd, min)
				}

				if test.expectedMaxAfterAdd != max {
					t.Errorf("%s: expected max after Add(): %d, got: %d", test.description, test.expectedMaxAfterAdd, max)
				}

				for _, sc := range test.InRangeAfterAdd {
					if !is.InRange(sc) {
						t.Errorf("%s: set should contain %d after Add(), but it does not", test.description, sc)
					}
				}

				for _, snc := range test.NotInRangeAfterAdd {
					if is.InRange(snc) {
						t.Errorf("%s: set should not contain %d after Add(), but it does", test.description, snc)
					}
				}
			}

			if test.delRangeInput != "" {
				err := is.Del(test.addRangeInput)

				if test.expectedErrDel == nil {
					if err != test.expectedErrDel {
						t.Errorf("%s: expected Del() error: %s, got: %s", test.description, test.expectedErrAdd, err)
					}
				} else {
					if err != nil {
						if err.Error() != test.expectedErrDel.Error() {
							t.Errorf("%s: expected Del() error: %s, got: %s", test.description, test.expectedErrDel, err)
						}
					} else {
						t.Errorf("%s: expected Del() error: %s, got: %s", test.description, test.expectedErrDel, err)
					}
				}

				min, max := getMinMax(is)

				if test.expectedMinAfterDel != min {
					t.Errorf("%s: expected min after Del(): %d, got: %d", test.description, test.expectedMinAfterDel, min)
				}

				if test.expectedMaxAfterDel != max {
					t.Errorf("%s: expected max after Del(): %d, got: %d", test.description, test.expectedMaxAfterDel, max)
				}
				for _, sc := range test.InRangeAfterDel {
					if !is.InRange(sc) {
						t.Errorf("%s: set should contain %d after Del(), but it does not", test.description, sc)
					}
				}

				for _, snc := range test.NotInRangeAfterDel {
					if is.InRange(snc) {
						t.Errorf("%s: set should not contain %d after Del(), but it does", test.description, snc)
					}
				}
			}
		}
	}
}
