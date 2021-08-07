package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/elliotchance/pie/pie"
)

func ParseLinesToInts(year int, day int) (pie.Ints, error) {
	gitRoot, err := gitProjectRoot()
	if err != nil {
		return nil, err
	}
	fileName := fmt.Sprintf("%s/input/%d/day%s.txt", gitRoot, year, strconv.Itoa(day))
	fmt.Println(fileName)
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ints := pie.Ints{}
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if line == "" {
			continue
		}
		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("Invalid input on line %d: %s", lineNum, line)
		}
		ints = append(ints, i)
	}
	return ints, nil
}
