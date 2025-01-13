// https://github.com/runtimeninja
// https://www.hackerrank.com/profile/toufiq_py
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
 * Complete the 'simpleArraySum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY ar as parameter.
 */

// I HAVE STARTED WRITING CODE FROM HERE
// simpleArraySum calculates the sum of an array of integers.
// param:
// - ar: a slice of integers
// returns:
// - the summation of the integers in the array.
func simpleArraySum(ar []int) int {
	sum := 0
	for _, value := range ar {
		sum += value
	}
	return sum
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// reading the size of the array
	nTemp, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nTemp))

	// reading the array elements
	arrTemp, _ := reader.ReadString('\n')
	arrStr := strings.Fields(arrTemp)

	// converting the array elements to integers
	var ar []int
	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(arrStr[i])
		ar = append(ar, num)
	}

	// calculating the sum of the array elements
	result := simpleArraySum(ar)

	// printing the result
	fmt.Println(result)
}
