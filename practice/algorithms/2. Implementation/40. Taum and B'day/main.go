// https://www.hackerrank.com/challenges/taum-and-bday/problem

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func min(x, y, z int64) int64 {
	if x < y && x < z {
		return x
	}

	if y < z {
		return y
	}

	return z
}

// Complete the taumBday function below.
func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64 {
	bb := int64(b)
	ww := int64(w)
	bcc := int64(bc)
	wcc := int64(wc)
	zz := int64(z)

	blackCost := bb * bcc
	whiteCost := ww * wcc

	var costBToW, costWToB, cost int64

	costWToB += whiteCost
	costWToB += bb * (wcc + zz)

	costBToW += blackCost
	costBToW += ww * (bcc + zz)

	cost = blackCost + whiteCost

	return min(costWToB, costBToW, cost)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		bw := strings.Split(readLine(reader), " ")

		bTemp, err := strconv.ParseInt(bw[0], 10, 64)
		checkError(err)
		b := int32(bTemp)

		wTemp, err := strconv.ParseInt(bw[1], 10, 64)
		checkError(err)
		w := int32(wTemp)

		bcWcz := strings.Split(readLine(reader), " ")

		bcTemp, err := strconv.ParseInt(bcWcz[0], 10, 64)
		checkError(err)
		bc := int32(bcTemp)

		wcTemp, err := strconv.ParseInt(bcWcz[1], 10, 64)
		checkError(err)
		wc := int32(wcTemp)

		zTemp, err := strconv.ParseInt(bcWcz[2], 10, 64)
		checkError(err)
		z := int32(zTemp)

		result := taumBday(b, w, bc, wc, z)

		fmt.Fprintf(writer, "%d\n", result)
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
