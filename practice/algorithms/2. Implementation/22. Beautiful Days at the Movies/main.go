// https://www.hackerrank.com/challenges/beautiful-days-at-the-movies/problem

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func reverse(num int32) int32 {
	var ans int32

	for num != 0 {
		ans = (ans * 10) + (num % 10)
		num /= 10
	}

	return ans
}

func abs(num int32) int32 {
	if num < 0 {
		return num * -1
	}

	return num
}

// Complete the beautifulDays function below.
func beautifulDays(i int32, j int32, k int32) int32 {
	var days int32
	for i <= j {
		if abs(i-reverse(i))%k == 0 {
			days++
		}

		i++
	}

	return days
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	ijk := strings.Split(readLine(reader), " ")

	iTemp, err := strconv.ParseInt(ijk[0], 10, 64)
	checkError(err)
	i := int32(iTemp)

	jTemp, err := strconv.ParseInt(ijk[1], 10, 64)
	checkError(err)
	j := int32(jTemp)

	temp, err := strconv.ParseInt(ijk[2], 10, 64)
	checkError(err)
	k := int32(temp)

	result := beautifulDays(i, j, k)

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
