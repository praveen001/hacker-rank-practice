// https://www.hackerrank.com/challenges/queens-attack-2/problem

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func abs(num int32) int32 {
	if num < 0 {
		return num * -1
	}

	return num
}

func min(a, b int32) int32 {
	if a > b {
		return b
	}

	return a
}

func isInPath(r_q, c_q, r, c int32) bool {
	// Same row?
	if r_q == r {
		return true
	}

	// Same col?
	if c_q == c {
		return true
	}

	if abs(r_q-r) == abs(c_q-c) {
		return true
	}

	return false
}

func getDirection(r_q, c_q, r, c int32) int32 {
	// Left
	if r_q == r && c_q > c {
		return 0
	}

	// Top-Left
	if r_q < r && c_q > c {
		return 1
	}

	// Top
	if r_q < r && c_q == c {
		return 2
	}

	// Top-Right
	if r_q < r && c_q < c {
		return 3
	}

	// Right
	if r_q == r && c_q < c {
		return 4
	}

	// Bottom-Right
	if r_q > r && c_q < c {
		return 5
	}

	// Bottom
	if r_q > r && c_q == c {
		return 6
	}

	// Bottom-Left
	if r_q > r && c_q > c {
		return 7
	}

	return 0
}

func getDistance(r_q, c_q, r, c, dir int32) int32 {
	switch dir {
	case 0, 1, 3, 4, 5, 7:
		return abs(c-c_q) - 1

	case 2, 6:
		return abs(r-r_q) - 1
	}

	return 0
}

func getDistanceToEdge(dim, r, c int32, dir int) int32 {
	var ans int32
	switch dir {
	case 0:
		ans = c - 1

	case 1:
		ans = min(c-1, dim-r)

	case 2:
		ans = dim - r

	case 3:
		ans = min(dim-r, dim-c)

	case 4:
		ans = dim - c

	case 5:
		ans = min(dim-c, r-1)

	case 6:
		ans = r - 1

	case 7:
		ans = min(r-1, c-1)

	}

	return ans
}

// Complete the queensAttack function below.
func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	minDistances := make([]int32, 8)
	for i := 0; i < 8; i++ {
		minDistances[i] = getDistanceToEdge(n, r_q, c_q, i)
	}

	for _, obstacle := range obstacles {
		if !isInPath(r_q, c_q, obstacle[0], obstacle[1]) {
			continue
		}

		dir := getDirection(r_q, c_q, obstacle[0], obstacle[1])
		dis := getDistance(r_q, c_q, obstacle[0], obstacle[1], dir)
		minDistances[dir] = min(minDistances[dir], dis)
	}

	var ans int32
	for _, dist := range minDistances {
		ans += dist
	}

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

	r_qC_q := strings.Split(readLine(reader), " ")

	r_qTemp, err := strconv.ParseInt(r_qC_q[0], 10, 64)
	checkError(err)
	r_q := int32(r_qTemp)

	c_qTemp, err := strconv.ParseInt(r_qC_q[1], 10, 64)
	checkError(err)
	c_q := int32(c_qTemp)

	var obstacles [][]int32
	for i := 0; i < int(k); i++ {
		obstaclesRowTemp := strings.Split(readLine(reader), " ")

		var obstaclesRow []int32
		for _, obstaclesRowItem := range obstaclesRowTemp {
			obstaclesItemTemp, err := strconv.ParseInt(obstaclesRowItem, 10, 64)
			checkError(err)
			obstaclesItem := int32(obstaclesItemTemp)
			obstaclesRow = append(obstaclesRow, obstaclesItem)
		}

		if len(obstaclesRow) != int(2) {
			panic("Bad input")
		}

		obstacles = append(obstacles, obstaclesRow)
	}

	result := queensAttack(n, k, r_q, c_q, obstacles)

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
