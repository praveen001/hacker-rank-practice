// https://www.hackerrank.com/challenges/jumping-on-the-clouds-revisited/problem

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32, k int32) int32 {
	// Start with energy 100.
	var e int32 = 100

	var cur int32
	for e == 100 || cur != 0 {
		// Jump
		e--
		cur = (cur + k) % int32(len(c))

		// Check the jump
		if c[cur] == 1 {
			e -= 2
		}
	}

	return e
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	temp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(temp)

	cTemp := strings.Split(readLine(reader), " ")

	var c []int32

	for i := 0; i < int(n); i++ {
		cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
		checkError(err)
		cItem := int32(cItemTemp)
		c = append(c, cItem)
	}

	result := jumpingOnClouds(c, k)

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
