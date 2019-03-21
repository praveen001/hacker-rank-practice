// https://www.hackerrank.com/challenges/staircase/submissions/code/67646789

package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, line int
	fmt.Scan(&n)
	line = n
	for line > 0 {
		line--
		var spaces, hashes int = line, n - line
		fmt.Print(strings.Repeat(" ", spaces), strings.Repeat("#", hashes), "\n")

	}
}
