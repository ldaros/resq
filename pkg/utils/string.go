package utils

import (
	"fmt"
	"strings"
)

func LimitLines(data []byte, numberOfLines int) []byte {
	lines := strings.Split(string(data), "\n")
	if len(lines) <= numberOfLines {
		return data
	}

	limited := strings.Join(lines[:numberOfLines], "\n")
	extraLines := fmt.Sprintf("\n... and %d more lines", len(lines)-numberOfLines)

	return []byte(limited + extraLines)
}
