// https://www.hackerrank.com/challenges/time-conversion/problem

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
	isPostMeridian := strings.HasSuffix(s, "PM")
	hour, _ := strconv.Atoi(s[:2])

	if (isPostMeridian && hour != 12) || (!isPostMeridian && hour == 12) {
		hour += 12
	}
	hour %= 24

	hourStr := strconv.Itoa(hour)
	if hour < 10 {
		hourStr = fmt.Sprintf("0%v", hour)
	}

	return fmt.Sprintf("%v%v", hourStr, s[2:])
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

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
