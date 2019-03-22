// https://www.hackerrank.com/challenges/day-of-the-programmer/problem

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the dayOfProgrammer function below.
func dayOfProgrammer(year int32) string {
	daysInMonth := []int32{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30}

	if year < 1918 && year%4 == 0 {
		daysInMonth[1] = 29
	} else if year > 1918 && (year%400 == 0 || (year%4 == 0 && year%100 != 0)) {
		daysInMonth[1] = 29
	} else if year == 1918 {
		daysInMonth[1] = 15
	}

	var total, month int32
	for i, days := range daysInMonth {
		if total+days > 256 {
			month = int32(i) + 1
			break
		}
		total += days
	}

	day := 256 - total

	return fmt.Sprintf("%02d.%02d.%02d", day, month, year)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	year := int32(yearTemp)

	result := dayOfProgrammer(year)

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
