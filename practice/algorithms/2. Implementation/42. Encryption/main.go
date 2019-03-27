// https://www.hackerrank.com/challenges/encryption/problem

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

// Complete the encryption function below.
func encryption(s string) string {
	// Remove all spaces
	strArr := strings.Split(s, " ")
	s = strings.Join(strArr, "")

	l := len(s)
	lsqrt := math.Sqrt(float64(l))
	cols := int(math.Ceil(lsqrt))

	ans := ""
	for i := 0; i < cols; i++ {
		for j := i; j < l; j += cols {
			ans += string(s[j])
		}
		ans += " "
	}

	return ans
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	result := encryption(s)

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
