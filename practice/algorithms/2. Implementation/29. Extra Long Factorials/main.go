// https://www.hackerrank.com/challenges/extra-long-factorials/problem

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func extraLongFactorials(n int32) {
	a := make([]int32, 1000)

	if n == 1 {
		fmt.Println(1)
		return
	}

	cpy := n

	i := 0
	for i = 0; cpy != 0; i++ {
		a[i] = cpy % 10
		cpy /= 10
	}

	for n != 1 {
		n--

		ln := i
		i = 0
		var carry int32
		for i = 0; i < ln || carry != 0; i++ {
			fc := n*a[i] + carry
			a[i] = fc % 10
			carry = fc / 10
		}
	}

	i--
	for i >= 0 {
		fmt.Printf("%v", a[i])
		i--
	}
	fmt.Println()
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	extraLongFactorials(n)
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
