package helpers

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
)

func ReadInput(year int, day int) ([]byte, error) {
	gitRoot, err := gitProjectRoot()
	if err != nil {
		return nil, err
	}
	file := fmt.Sprintf("%s/input/%d/day%s.txt", gitRoot, year, strings.Trim(strconv.Itoa(day), "0"))
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func gitProjectRoot() (string, error) {
	path, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(path)), nil
}
