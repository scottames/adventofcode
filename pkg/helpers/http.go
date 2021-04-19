package helpers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Get retrieves the specified AOC input from https://adventofcode.com/ and returns it as a slice of bytes
func Get(year int, day int) ([]byte, error) {
	sessionID := os.Getenv("AOC_SESSION_ID")
	if len(sessionID) == 0 {
		return nil, fmt.Errorf("set AOC_SESSION_ID in order to retrieve unique AOC input")
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	cookie := http.Cookie{
		Name:   "session",
		Value:  sessionID,
		Domain: ".adventofcode.com",
		Path:   "/",
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating http request: %v", err)
	}

	req.AddCookie(&cookie)
	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("reading response: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading body: %v", err)
	}

	return body, nil
}
