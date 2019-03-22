// https://www.hackerrank.com/challenges/cut-the-sticks/problem

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the cutTheSticks function below.
func cutTheSticks(arr []int32) []int32 {
	var ans []int32

	for {
		var ar []int32

		// If arr is empty, end it
		if len(arr) == 0 {
			break
		}

		// Store the length before cutting
		ans = append(ans, int32(len(arr)))

		// Find Min
		var min int32
		for _, val := range arr {
			if min > val || min == 0 {
				min = val
			}
		}

		// Cut, remove 0s
		ar = []int32{}
		for i := 0; i < len(arr); i++ {
			if arr[i]-min != 0 {
				ar = append(ar, arr[i]-min)
			}
		}
		arr = ar
	}

	return ans
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := cutTheSticks(arr)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
