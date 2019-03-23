// https://www.hackerrank.com/challenges/repeated-string/problem

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	count := make([]int64, len(s)+1)

	var counter int64

	// Count number of a's at each index
	for i, l := range s {
		if l == 97 {
			counter++
		}
		count[i+1] = counter
	}

	var ans, ln int64
	ln = int64(len(s))

	// n / ln gives number of times the string can  be fully repeated
	// We multiple that with number of a's present in the whole string
	ans = counter * (n / ln)

	// Get then add the number of a's in remaining characters
	ans += count[n-(ln*(n/ln))]

	return ans
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

	fmt.Fprintf(writer, "%d\n", result)

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
