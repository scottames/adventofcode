package aoc2020

import (
	"fmt"
	"github.com/scottames/adventofcode/pkg/helpers"
	"strconv"
	"strings"
)

// Day02 - Day 1 Part 1 & 2
func Day02() error {
	get, err := helpers.ReadInput(2020, 2)
	if err != nil {
		return err
	}
	validators := []func(*password) bool{
		validatorPart1,
		validatorPart2,
	}
	nums, err := validatePassword(
		helpers.CovertBytesToStrings(get, "\n"),
		validators...,
	)
	if err != nil {
		return err
	}
	for i, num := range nums {
		fmt.Printf("--- Part %d ---\n\n", i+1)
		fmt.Printf("Valid passwords: %d\n\n", num)
	}
	return nil
}

// validatePassword returns a slice of ints based on the number of validators given and a possible error
// from a password entry as a string slice in the format: []string{"1-3","a","abcdefg"}
func validatePassword(input []string, validators ...func(*password) bool) ([]int, error) {
	var valid = make([]int, len(validators))
	for _, pps := range input {
		pp, err := newPassword(pps)
		if err != nil {
			return nil, err
		}
		for i, fn := range validators {
			if fn(pp) {
				valid[i]++
			}
		}
	}
	return valid, nil
}

// validatorPart1 validates the given password based on the formula given in "part 1"
func validatorPart1(pp *password) bool {
	if pp == nil {
		return false
	}
	var count int
	for _, c := range pp.password {
		if string(c) == pp.char {
			count++
		}
	}
	if count < pp.intRange.min || count > pp.intRange.max {
		return false
	}
	return true
}

// validatorPart2 validates the given password based on the formula given in "part 2"
func validatorPart2(pp *password) bool {
	if len(pp.password) <= pp.intRange.max-1 {
		return false
	}
	passwordAsRunes := []rune(pp.password)
	firstChar := string(passwordAsRunes[pp.intRange.min-1])
	lastChar := string(passwordAsRunes[pp.intRange.max-1])
	if (firstChar == pp.char && lastChar != pp.char) ||
		(lastChar == pp.char && firstChar != pp.char) {
		return true
	}
	return false
}

// intRange represents two ints that form a given range (min - max)
type intRange struct {
	max int
	min int
}

// newIntRangeFromString returns a pointer to an intRange and a possible error
// from a string of the format: "1-3"
func newIntRangeFromString(r string) (*intRange, error) {
	ss := strings.Split(r, "-")
	if lenSS := len(ss); lenSS != 2 {
		return nil, fmt.Errorf("expected numbered range (e.g. 1-3), got %s", r)
	}
	minString := ss[0]
	min, err := strconv.Atoi(minString)
	if err != nil {
		return nil, fmt.Errorf("invalid minimum found in range: %s", minString)
	}
	maxString := ss[1]
	max, err := strconv.Atoi(maxString)
	if err != nil {
		return nil, fmt.Errorf("invalid maximum found in range: %s", maxString)
	}
	return &intRange{
		max: max,
		min: min,
	}, nil
}

// password represents a given password and it's requirements
type password struct {
	char     string
	intRange intRange
	password string
}

// newPassword returns a pointer to a password and a possible error
// from a string containing the password & it's requirements as outlined in AOC 2020 Day 2
func newPassword(s string) (*password, error) {
	passStrings, err := passwordStrings(s)
	if err != nil {
		return nil, err
	}
	ir, err := newIntRangeFromString(passStrings[0])
	if err != nil {
		return nil, err
	} else if ir == nil {
		return nil, fmt.Errorf("invalid min/max values found")
	}
	return &password{
		char:     strings.Trim(passStrings[1], ":"),
		intRange: *ir,
		password: passStrings[2],
	}, nil
}

// passwordStrings returns a slice of strings from the given password string
// by splitting on the line's whitespace
func passwordStrings(p string) ([]string, error) {
	ss := strings.Split(p, " ")
	if lenSS := len(ss); lenSS != 3 {
		return nil, fmt.Errorf("expected 3 elements in password string, got %d", lenSS)
	}
	return ss, nil
}
