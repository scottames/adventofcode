package credentials

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	// invalid
	invalidPassport1 = `eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926`
	invalidPassport2 = `iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946`
	invalidPassport3 = `hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277`
	invalidPassport4 = `hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007`

	// valid
	validPassport1 = `pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f`
	validPassport2 = `eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm`
	validPassport3 = `hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022`
	validPassport4 = "iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719"
)

func TestPassport_IsValidTrue(t *testing.T) {
	expectedValid := true
	expectedPP := &Passport{
		byr: 1980,
		iyr: 2012,
		eyr: 2030,
		hgt: "74in",
		hcl: "#623a2f",
		ecl: "grn",
		pid: "087499704",
		cid: "",
	}
	actualPP, actualValid := NewPassportFromString(validPassport1)
	msg := fmt.Sprintf("Expected %v. Got %v.", expectedValid, actualValid)
	assert.Equal(t, expectedValid, actualValid, msg)
	msg = fmt.Sprintf("Expected %#v. Got %#v.", expectedPP, actualPP)
	assert.Equal(t, expectedPP, actualPP, msg)
}

func TestPassport_IsValidFalse(t *testing.T) {
	expectedValid := false
	expectedPP := &Passport{
		byr: 1926,
		iyr: 2018,
		eyr: 1972,
		hgt: "170",
		hcl: "#18171d",
		ecl: "amb",
		pid: "186cm",
		cid: "100",
	}
	actualPP, actualValid := NewPassportFromString(invalidPassport1)
	msg := fmt.Sprintf("Expected %v. Got %v.", expectedValid, actualValid)
	assert.Equal(t, expectedValid, actualValid, msg)
	msg = fmt.Sprintf("Expected %#v. Got %#v.", expectedPP, actualPP)
	assert.Equal(t, expectedPP, actualPP, msg)
}

func TestPassport_isECLValidTrue(t *testing.T) {
	pp, _ := NewPassportFromString(validPassport2)
	expected := true
	actual := pp.isECLValid()
	msg := fmt.Sprintf("Expected %v. Got %v. (ecl: %s)", expected, actual, pp.ecl)
	assert.Equal(t, expected, actual, msg)
}

func TestPassport_isECLValidFalse(t *testing.T) {
	pp, _ := NewPassportFromString(invalidPassport4)
	expected := false
	actual := pp.isECLValid()
	msg := fmt.Sprintf("Expected %v. Got %v. (ecl: %s)", expected, actual, pp.ecl)
	assert.Equal(t, expected, actual, msg)
}

func TestPassport_isBYRValidTrue(t *testing.T) {
	pp, _ := NewPassportFromString(validPassport3)
	expected := true
	actual := pp.isBYRValid()
	msg := fmt.Sprintf("Expected %v. Got %v. (byr: %d)", expected, actual, pp.byr)
	assert.Equal(t, expected, actual, msg)
}

func TestPassport_isBYRValidFalse(t *testing.T) {
	pp, _ := NewPassportFromString(invalidPassport4)
	expected := false
	actual := pp.isBYRValid()
	msg := fmt.Sprintf("Expected %v. Got %v. (byr: %d)", expected, actual, pp.byr)
	assert.Equal(t, expected, actual, msg)
}

func TestPassport_isIYRValidTrue(t *testing.T) {
	pp, _ := NewPassportFromString(validPassport4)
	expected := true
	actual := pp.isIYRValid()
	msg := fmt.Sprintf("Expected %v. Got %v. (iyr: %d)", expected, actual, pp.iyr)
	assert.Equal(t, expected, actual, msg)
}

func TestPassport_isIYRValidFalse(t *testing.T) {
	pp, _ := NewPassportFromString(invalidPassport4)
	expected := false
	actual := pp.isIYRValid()
	msg := fmt.Sprintf("Expected %v. Got %v. (iyr: %d)", expected, actual, pp.iyr)
	assert.Equal(t, expected, actual, msg)
}

func TestPassport_isEYRValidTrue(t *testing.T) {
	pp, _ := NewPassportFromString(validPassport1)
	expected := true
	actual := pp.isEYRValid()
	msg := fmt.Sprintf("Expected %v. Got %v. (eyr: %d)", expected, actual, pp.eyr)
	assert.Equal(t, expected, actual, msg)
}

func TestPassport_isEYRValidFalse(t *testing.T) {
	pp, _ := NewPassportFromString(invalidPassport1)
	expected := false
	actual := pp.isEYRValid()
	msg := fmt.Sprintf("Expected %v. Got %v. (eyr: %d)", expected, actual, pp.eyr)
	assert.Equal(t, expected, actual, msg)
}

func TestPassport_isHCLValidTrue(t *testing.T) {
	pp, _ := NewPassportFromString(validPassport1)
	expected := true
	actual := pp.isHCLValid()
	msg := fmt.Sprintf("Expected %v. Got %v. (hcl: %s)", expected, actual, pp.hcl)
	assert.Equal(t, expected, actual, msg)
}

func TestPassport_isHCLValidFalse(t *testing.T) {
	pp, _ := NewPassportFromString(invalidPassport3)
	expected := false
	actual := pp.isHCLValid()
	msg := fmt.Sprintf("Expected %v. Got %v. (hcl: %s)", expected, actual, pp.hcl)
	assert.Equal(t, expected, actual, msg)
}

func TestPassport_isHGTValidTrue(t *testing.T) {
	pp, _ := NewPassportFromString(validPassport1)
	expected := true
	actual := pp.isHGTValid()
	msg := fmt.Sprintf("Expected %v. Got %v. (hgt: %s)", expected, actual, pp.hgt)
	assert.Equal(t, expected, actual, msg)
}

func TestPassport_isHGTValidFalse(t *testing.T) {
	pp, _ := NewPassportFromString(invalidPassport1)
	expected := false
	actual := pp.isHGTValid()
	msg := fmt.Sprintf("Expected %v. Got %v. (hgt: %s)", expected, actual, pp.hgt)
	assert.Equal(t, expected, actual, msg)
}

func TestPassport_isPIDValidTrue(t *testing.T) {
	pp, _ := NewPassportFromString(validPassport1)
	expected := true
	actual := pp.isPIDValid()
	msg := fmt.Sprintf("Expected %v. Got %v. (pid: %s)", expected, actual, pp.pid)
	assert.Equal(t, expected, actual, msg)
}

func TestPassport_isPIDValidFalse(t *testing.T) {
	pp, _ := NewPassportFromString(invalidPassport4)
	expected := false
	actual := pp.isPIDValid()
	msg := fmt.Sprintf("Expected %v. Got %v. (pid: %s)", expected, actual, pp.pid)
	assert.Equal(t, expected, actual, msg)
}
