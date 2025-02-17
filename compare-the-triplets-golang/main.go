// https://github.com/runtimeninja
// https://www.hackerrank.com/profile/toufiq_py
// https://www.hackerrank.com/challenges/compare-the-triplets/problem?isFullScreen=true

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//DETAILS PROVIDED BY HACKERRANK WHILE SOLVING THE PROBLEM
/*
* Complete the 'compareTriplets' function below.
*
* The function is expected to return an INTEGER_ARRAY.
* The function accepts following parameters:
*  1. INTEGER_ARRAY a
*  2. INTEGER_ARRAY b
 */

// I HAVE STARTED WRITING CODE FROM HERE
// compareTriplets function compares the ratings of Alice and Bob and calculates their scores.
// param:
// - a: a slice of 3 integers representing Alice's ratings.
// - b: a slice of 3 integers representing Bob's ratings.
// ret:
// - a slice of 2 int: Alice's score and Bob's score.
func compareTriplets(a []int, b []int) []int {
	aliceScore := 0
	bobScore := 0

	// comparing the ratings for each category
	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			aliceScore++
		} else if a[i] < b[i] {
			bobScore++
		}
	}

	// returning the scores as a slice
	return []int{aliceScore, bobScore}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// reading Alice's ratings
	aTemp, _ := reader.ReadString('\n')
	aStr := strings.Fields(aTemp)
	var a []int
	for _, val := range aStr {
		num, _ := strconv.Atoi(val)
		a = append(a, num)
	}

	// reading Bob's ratings
	bTemp, _ := reader.ReadString('\n')
	bStr := strings.Fields(bTemp)
	var b []int
	for _, val := range bStr {
		num, _ := strconv.Atoi(val)
		b = append(b, num)
	}

	// calling compareTriplets and get the scores
	result := compareTriplets(a, b)

	// printing the result
	fmt.Println(result[0], result[1])
}
