package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	subDirs    = 3
	minDisplay = 4
	pathSep    = string(os.PathSeparator)
)

func main() {
	dir, _ := os.Getwd()
	parts := strings.Split(dir, pathSep)

	switch l := len(parts); {
	case l <= subDirs || l == minDisplay:
		fmt.Print(dir)
	default:
		start := l - subDirs
		fmt.Print("../" + strings.Join(parts[start:], pathSep))
	}

	os.Exit(0)
}
