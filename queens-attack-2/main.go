// https://github.com/runtimeninja
// https://www.hackerrank.com/profile/toufiq_py
// https://toufiq.info

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'queensAttack' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER k
 *  3. INTEGER r_q
 *  4. INTEGER c_q
 *  5. 2D_INTEGER_ARRAY obstacles
 */

// started writing code from here.
func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	directions := [8][2]int32{
		{1, 0},   // up
		{-1, 0},  // down
		{0, 1},   // right
		{0, -1},  // left
		{1, 1},   // up-right
		{1, -1},  // up-left
		{-1, 1},  // down-right
		{-1, -1}, // down-left
	}

	// created a map of obstacles for faster lookup
	obstacleMap := make(map[[2]int32]bool)
	for _, obs := range obstacles {
		obstacleMap[[2]int32{obs[0], obs[1]}] = true
	}

	totalMoves := int32(0)

	// calculated moves in each direction
	for _, dir := range directions {
		step := int32(1)
		for {
			newRow := r_q + step*dir[0]
			newCol := c_q + step*dir[1]

			// Check if the position is outside the board or blocked by an obstacle
			if newRow < 1 || newRow > n || newCol < 1 || newCol > n || obstacleMap[[2]int32{newRow, newCol}] {
				break
			}

			totalMoves++
			step++
		}
	}

	return totalMoves
}

// completed queensAttack function

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	secondMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	r_qTemp, err := strconv.ParseInt(secondMultipleInput[0], 10, 64)
	checkError(err)
	r_q := int32(r_qTemp)

	c_qTemp, err := strconv.ParseInt(secondMultipleInput[1], 10, 64)
	checkError(err)
	c_q := int32(c_qTemp)

	var obstacles [][]int32
	for i := 0; i < int(k); i++ {
		obstaclesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var obstaclesRow []int32
		for _, obstaclesRowItem := range obstaclesRowTemp {
			obstaclesItemTemp, err := strconv.ParseInt(obstaclesRowItem, 10, 64)
			checkError(err)
			obstaclesRow = append(obstaclesRow, int32(obstaclesItemTemp))
		}

		if len(obstaclesRow) != 2 {
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
