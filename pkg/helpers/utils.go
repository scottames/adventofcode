package helpers

import (
	"fmt"
	"os"
)

func ExitOnError(err error) {
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}
}
