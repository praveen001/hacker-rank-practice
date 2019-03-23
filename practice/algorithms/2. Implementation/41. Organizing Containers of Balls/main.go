// https://www.hackerrank.com/challenges/organizing-containers-of-balls/problem

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the organizingContainers function below.
func organizingContainers(container [][]int32) string {
	sizeToContainerCount := make(map[int]int)
	categoryToCount := make(map[int]int)

	for i := 0; i < len(container); i++ {
		var size int
		for j := 0; j < len(container[i]); j++ {
			categoryToCount[j] += int(container[i][j])
			size += int(container[i][j])
		}
		sizeToContainerCount[size]++
	}

	var match int
	for _, count := range categoryToCount {
		if containerCount, ok := sizeToContainerCount[count]; !ok || containerCount == 0 {
			break
		}
		sizeToContainerCount[count]--
		match++
	}

	if match == len(container) {
		return "Possible"
	}

	return "Impossible"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		var container [][]int32
		for i := 0; i < int(n); i++ {
			containerRowTemp := strings.Split(readLine(reader), " ")

			var containerRow []int32
			for _, containerRowItem := range containerRowTemp {
				containerItemTemp, err := strconv.ParseInt(containerRowItem, 10, 64)
				checkError(err)
				containerItem := int32(containerItemTemp)
				containerRow = append(containerRow, containerItem)
			}

			if len(containerRow) != int(int(n)) {
				panic("Bad input")
			}

			container = append(container, containerRow)
		}

		result := organizingContainers(container)

		fmt.Fprintf(writer, "%s\n", result)
	}

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
