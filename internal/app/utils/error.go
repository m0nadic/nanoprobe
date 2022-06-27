package utils

import (
	"fmt"
	"os"
)

func ErrorExit(exitCode int, err error) {
	fmt.Fprintln(os.Stderr, "ERROR:", err)
	os.Exit(exitCode)
}
