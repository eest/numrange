package numrange

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// IntSet contains a set of ints that can be added to/deleted from and allows
// for checking if a value fits inside range.
type IntSet map[int]bool

// ParseIntSet takes a string such as "-10..10,12,15,20..25" and returns an
// IntSet
func ParseIntSet(s string) (IntSet, error) {

	is := IntSet{}

	if s == "" {
		return is, nil
	}

	parts := strings.Split(s, ",")

	for _, part := range parts {
		if strings.Contains(part, "..") {
			endpoints := strings.Split(part, "..")

			if len(endpoints) != 2 {
				return nil, fmt.Errorf("invalid \"..\" range: %s", part)
			}

			start, err := strconv.Atoi(endpoints[0])
			if err != nil {
				return nil, fmt.Errorf("unable to parse \"..\" starting point \"%s\": %s", endpoints[0], err)
			}

			end, err := strconv.Atoi(endpoints[1])
			if err != nil {
				return nil, fmt.Errorf("unable to parse \"..\" ending point \"%s\": %s", endpoints[1], err)
			}

			for i := start; i <= end; i++ {
				if _, ok := is[i]; !ok {
					is[i] = true
				}
			}

		} else {
			i, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("unable to parse range part %s: %s", part, err)
			}
			if _, ok := is[i]; !ok {
				is[i] = true
			}
		}
	}

	return is, nil
}

// Add adds a range to the set
func (is IntSet) Add(s string) error {

	addSet, err := ParseIntSet(s)
	if err != nil {
		return err
	}

	for k := range addSet {
		if _, ok := is[k]; !ok {
			is[k] = true
		}
	}

	return nil

}

// Del deletes a range from the set
func (is IntSet) Del(s string) error {

	delSet, err := ParseIntSet(s)
	if err != nil {
		return err
	}

	for k := range delSet {
		if _, ok := is[k]; ok {
			delete(is, k)
		}
	}

	return nil

}

// InRange checks if the supplied value fits inside the range in the set
// (including start and end values)
func (is IntSet) InRange(num int) bool {

	if len(is) == 0 {
		return false
	}

	var min, max int

	// Find out the minimum and maximum values in the set
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

	if min <= num && num <= max {
		return true
	}

	return false
}

// String returns a string representation of the range in the set
func (is IntSet) String() string {

	if len(is) == 0 {
		return ""
	}

	var keys []int

	for k := range is {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	var b strings.Builder
	insideRange := false

	for i, k := range keys {
		if i == 0 {
			// The first number will always be printed.
			fmt.Fprintf(&b, "%d", k)
		} else {
			// If the current value is just one more than the
			// previous value we are in a ".." range.
			if k == keys[i-1]+1 {
				// This is part of an ongoing ".." range, keep going
				insideRange = true
			} else {
				// If we were in a range up until this number,
				// end the range on the previous number
				if insideRange {
					fmt.Fprintf(&b, "..%d", keys[i-1])
					insideRange = false
				}
				// Since we are starting a new part delimit with comma
				fmt.Fprintf(&b, ",%d", k)
			}
		}
	}

	// If we were in a range when we ended the loop, add the end of the range
	if insideRange {
		fmt.Fprintf(&b, "..%d", keys[len(keys)-1])
		insideRange = false
	}

	return b.String()
}
