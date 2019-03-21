// https://www.hackerrank.com/challenges/mini-max-sum/problem

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
	var sum int64
	max := int64(arr[0])
	min := int64(arr[0])

	for _, x := range arr {
		y := int64(x)
		sum += y

		if y > max {
			max = y
		}

		if y < min {
			min = y
		}
	}

	fmt.Printf("%v %v\n", sum-max, sum-min)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
