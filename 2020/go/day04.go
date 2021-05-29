package aoc2020

import (
	"fmt"
	"github.com/scottames/adventofcode/2020/go/pkg/credentials"
	"github.com/scottames/adventofcode/pkg/helpers"
)

const (
	day04 = 4
)

// Day04 - Day 4 Part 1 & 2
func Day04() error {
	input, err := helpers.ReadInput(year, day04)
	if err != nil {
		return err
	}

	passportStrings := helpers.CovertBytesToStrings(input, "\n\n")

	// Part 1
	validBasicPassports := validBasicPassportsFromStringSlice(passportStrings)

	helpers.PrintPart1()
	fmt.Printf("Number of valid valid passports: %d\n\n", len(validBasicPassports))

	// Part 2
	validPassports := validPassportsFromStringSlice(passportStrings)

	helpers.PrintPart2()
	fmt.Printf("Number of valid valid passports: %d\n\n", len(validPassports))

	return nil
}

// validPassportsFromStringSlice returns a slice of valid Passports
//  from a given string slice
func validPassportsFromStringSlice(ss []string) []credentials.Passport {
	validPassports := []credentials.Passport{}
	for _, p := range ss {
		pp, valid := credentials.NewPassportFromString(p)
		if valid && pp != nil {
			validPassports = append(validPassports, *pp)
		}
	}
	return validPassports
}

// validBasicPassportsFromStringSlice returns a slice of valid
// "basic" (part 1) Passports from a given string slice
func validBasicPassportsFromStringSlice(ss []string) []credentials.Passport {
	validPassports := []credentials.Passport{}
	for _, p := range ss {
		pp, valid := credentials.NewBasicPassportFromString(p)
		if valid && pp != nil {
			validPassports = append(validPassports, *pp)
		}
	}
	return validPassports
}
