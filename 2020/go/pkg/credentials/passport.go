package credentials

import (
	"regexp"
	"strconv"
	"strings"
)

// Passport struct representing a Passport
type Passport struct {
	byr int    // Birth Year
	iyr int    // Issue Year
	eyr int    // Expiration Year
	hgt string // Height
	hcl string // Hair Color
	ecl string // Eye Color
	pid string // Passport ID
	cid string // Country ID
}

// NewPassportFromString returns a new Passport
// and a bool whether the Passport is valid from a given string
func NewPassportFromString(s string) (*Passport, bool) {
	ss := strings.Fields(s)
	return NewPassportFromStringSlice(ss)
}

// NewPassportFromStringSlice returns a new Passport
// and a bool whether the Passport is valid from a given string slice
func NewPassportFromStringSlice(ss []string) (*Passport, bool) {
	p := Passport{}
	for _, val := range ss {
		kv := strings.Split(val, ":")
		if len(kv) != 2 {
			return nil, false
		}
		k := kv[0]
		v := kv[1]
		switch k {
		case "byr":
			i, err := strconv.Atoi(v)
			if err != nil {
				return nil, false
			}
			p.byr = i
		case "iyr":
			i, err := strconv.Atoi(v)
			if err != nil {
				return nil, false
			}
			p.iyr = i
		case "eyr":
			i, err := strconv.Atoi(v)
			if err != nil {
				return nil, false
			}
			p.eyr = i
		case "hgt":
			p.hgt = v
		case "hcl":
			p.hcl = v
		case "ecl":
			p.ecl = v
		case "pid":
			p.pid = v
		case "cid":
			p.cid = v
		}
	}
	return &p, p.IsValid()
}

// NewBasicPassportFromString returns a new "basic" (part 1) Passport
// and a bool whether the Passport is valid from a given string
func NewBasicPassportFromString(s string) (*Passport, bool) {
	ss := strings.Fields(s)
	return NewBasicPassportFromStringSlice(ss)
}

// NewBasicPassportFromStringSlice returns a new "basic" (part 1) Passport
// and a bool whether the Passport is valid from a given string slice
func NewBasicPassportFromStringSlice(ss []string) (*Passport, bool) {
	p, _ := NewPassportFromStringSlice(ss)
	if p.byr == 0 {
		return p, false
	}
	if p.ecl == "" {
		return p, false
	}
	if p.eyr == 0 {
		return p, false
	}
	if p.hgt == "" {
		return p, false
	}
	if p.iyr == 0 {
		return p, false
	}
	if p.pid == "" {
		return p, false
	}
	if p.hcl == "" {
		return p, false
	}
	return p, true
}

// IsValid returns whether the Passport is valid or not
func (p *Passport) IsValid() bool {
	return p.isBYRValid() &&
		p.isIYRValid() &&
		p.isECLValid() &&
		p.isEYRValid() &&
		p.isHGTValid() &&
		p.isHCLValid() &&
		p.isPIDValid()
}

// isECLValid returns whether the Passport's ecl value is valid
func (p Passport) isECLValid() bool {
	switch p.ecl {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	default:
		return false
	}
}

// isBYRValid returns whether the Passport's byr value is valid
// byr (Birth Year) - four digits; at least 1920 and at most 2002.
func (p Passport) isBYRValid() bool {
	return p.byr >= 1920 && p.byr <= 2002
}

// isIYRValid returns whether the Passport's iyr value is valid
// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
func (p Passport) isIYRValid() bool {
	return p.iyr >= 2010 && p.iyr <= 2020
}

// isEYRValid returns whether the Passport's eyr value is valid
// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
func (p Passport) isEYRValid() bool {
	return p.eyr >= 2020 && p.eyr <= 2030
}

// isHCLValid returns whether the Passport's hcl value is valid
// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func (p Passport) isHCLValid() bool {
	regex, err := regexp.Compile("^#(?:[0-9a-fA-F]{3}){1,2}$")
	if err != nil {
		return false
	}
	return regex.MatchString(p.hcl)
}

// isHGTValid returns whether the Passport's hgt value is valid
// hgt (Height) - a number followed by either cm or in:
//  If cm, the number must be at least 150 and at most 193.
//  If in, the number must be at least 59 and at most 76.
func (p Passport) isHGTValid() bool {
	if p.hgt == "" {
		return false
	}

	var measurement string
	if strings.Contains(p.hgt, "cm") {
		measurement = "cm"
	} else if strings.Contains(p.hgt, "in") {
		measurement = "in"
	} else {
		return false
	}

	height, err := strconv.Atoi(strings.Split(p.hgt, measurement)[0])
	if err != nil {
		return false
	}
	switch measurement {
	case "cm":
		return height >= 150 && height <= 193
	case "in":
		return height >= 59 && height <= 76
	}
	return false
}

// isPIDValid returns whether the Passport's pid value is valid
// pid (Passport ID) - a nine-digit number, including leading zeroes.
func (p Passport) isPIDValid() bool {
	if len(p.pid) != 9 {
		return false
	}
	for _, d := range p.pid {
		if _, err := strconv.Atoi(string(d)); err != nil {
			return false
		}
	}
	return true
}
