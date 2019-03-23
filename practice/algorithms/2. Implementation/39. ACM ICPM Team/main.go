package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the acmTeam function below.
func acmTeam(topic []string) []int32 {
	var maxSubjects, ways int32

	for i := 0; i < len(topic); i++ {
		for j := i + 1; j < len(topic); j++ {
			var subjects int32
			for k := 0; k < len(topic[i]); k++ {
				if topic[i][k] == 49 || topic[j][k] == 49 {
					subjects++
				}
			}

			if subjects > maxSubjects {
				maxSubjects = subjects
				ways = 1
			} else if subjects == maxSubjects {
				ways++
			}

		}
	}

	return []int32{maxSubjects, ways}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)
	fmt.Println(m)

	var topic []string

	for i := 0; i < int(n); i++ {
		topicItem := readLine(reader)
		topic = append(topic, topicItem)
	}

	result := acmTeam(topic)

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
