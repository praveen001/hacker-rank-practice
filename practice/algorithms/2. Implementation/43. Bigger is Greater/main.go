// https://www.hackerrank.com/challenges/bigger-is-greater/problem

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func sortChars(ar []rune) {
	for i := 0; i < len(ar)-1; i++ {
		for j := i + 1; j < len(ar); j++ {
			if ar[i] > ar[j] {
				t := ar[i]
				ar[i] = ar[j]
				ar[j] = t
			}
		}
	}
}

// Complete the biggerIsGreater function below.
func biggerIsGreater(w string) string {
	wr := []rune(w)

	target := -1
	for i := len(wr) - 1; i > 0; i-- {

		// If next letter is greater than current letter
		// we can swap at the position to get a greater word
		if wr[i] > wr[i-1] {
			target = i - 1
			break
		}
	}

	if target == -1 {
		return "no answer"
	}

	// Find the smallest letter that is greater than target letter
	minIdx := target + 1
	for i := target; i < len(wr); i++ {
		if wr[i] < wr[minIdx] && wr[i] > wr[target] {
			minIdx = i
		}
	}

	// Swap target letter with smallest letter that is greater than target letter
	t := wr[target]
	wr[target] = wr[minIdx]
	wr[minIdx] = t

	// Sort remaining slice, to get the smallest word that is greater than given word
	sortChars(wr[target+1:])

	return string(wr)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	TTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	T := int32(TTemp)

	for TItr := 0; TItr < int(T); TItr++ {
		w := readLine(reader)

		result := biggerIsGreater(w)

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
