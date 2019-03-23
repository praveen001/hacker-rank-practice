// https://www.hackerrank.com/challenges/non-divisible-subset/problem

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the nonDivisibleSubset function below.
func nonDivisibleSubset(k int32, S []int32) int32 {
	remCount := make([]int32, k)

	// Find count for each remainder 0 to k - 1
	for _, s := range S {
		remCount[s%k]++
	}

	// We can pick only 1 number with remainder zero
	// If we pick more than 1, their add will be divisible by K
	if remCount[0] > 1 {
		remCount[0] = 1
	}

	// If k is even, then we can pick only 1 number with remained k / 2
	// cuz their sum will be divisible by K
	if k%2 == 0 {
		if remCount[k/2] > 1 {
			remCount[k/2] = 1
		}
	}

	var ans, i int32

	// We can pick either set with remained i or k - i, they are mutually exclusive cuz their sum will be divisible by k
	// Hence we pick the max of those two sets.
	for i = 1; i <= k/2; i++ {
		if remCount[i] < remCount[k-i] {
			ans += remCount[k-i]
		} else {
			ans += remCount[i]
		}
	}

	// Pick one number with 0 remainder.
	ans += remCount[0]

	return ans
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

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	STemp := strings.Split(readLine(reader), " ")

	var S []int32

	for i := 0; i < int(n); i++ {
		SItemTemp, err := strconv.ParseInt(STemp[i], 10, 64)
		checkError(err)
		SItem := int32(SItemTemp)
		S = append(S, SItem)
	}

	result := nonDivisibleSubset(k, S)

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
