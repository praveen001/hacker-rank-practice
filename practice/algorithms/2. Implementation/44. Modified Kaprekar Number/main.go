// https://www.hackerrank.com/challenges/kaprekar-numbers/problem

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func isKaprekarNumber(num int64, ln int) bool {
	var left, right string = strconv.Itoa(int(num * num)), ""
	for i := 0; i < ln; i++ {
		right = fmt.Sprintf("%s%s", string(left[len(left)-1]), right)
		left = fmt.Sprintf("%s", left[:len(left)-1])
	}

	leftInt, _ := strconv.Atoi(left)
	rightInt, _ := strconv.Atoi(right)
	return leftInt+rightInt == int(num)
}

// Complete the kaprekarNumbers function below.
func kaprekarNumbers(p int32, q int32) {
	flag := false
	for i := p; i <= q; i++ {
		if isKaprekarNumber(int64(i), len(strconv.Itoa(int(i)))) {
			flag = true
			fmt.Printf("%v ", i)
		}
	}

	if !flag {
		fmt.Println("INVALID RANGE")
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	pTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	p := int32(pTemp)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	kaprekarNumbers(p, q)
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
