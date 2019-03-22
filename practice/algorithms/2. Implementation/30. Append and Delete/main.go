// https://www.hackerrank.com/challenges/append-and-delete/problem

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the appendAndDelete function below.
func appendAndDelete(s string, t string, k int32) string {
	sLen := int32(len(s))
	tLen := int32(len(t))

	var commonPrefix int32
	ln := sLen
	if sLen > tLen {
		ln = tLen
	}
	for i := 0; i < int(ln); i++ {
		if s[i] != t[i] {
			break
		}
		commonPrefix++
	}

	sLen -= commonPrefix
	tLen -= commonPrefix

	// Delete the remaining characters from S
	k -= sLen

	// Enuf steps to do additions?
	k -= tLen

	// Not enuf steps?
	if k < 0 {
		return "No"
	}

	// Exactly enuf?
	if k == 0 {
		return "Yes"
	}

	// Extra?
	if k > 0 {
		// Can be remove all common prefix, and do empty deletion to waste steps?
		if commonPrefix*2 <= k {
			return "Yes"
		}

		// We can delete some characters and put it back? to do that, extra should be even
		if k%2 == 0 {
			return "Yes"
		}

	}

	return "No"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	t := readLine(reader)

	temp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	k := int32(temp)

	result := appendAndDelete(s, t, k)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
