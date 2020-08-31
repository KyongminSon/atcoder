package main

import (
	"bufio"
	"fmt"
	"os"

	"strconv"
	"strings"
)

func main() {
	var temp_list []int
	var sum int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		temp := strings.Split(scanner.Text(), " ")
		for _, str := range temp {
			i, _ := strconv.Atoi(str)
			temp_list = append(temp_list, i)
		}
		max_temp := max(temp_list)
		min_temp := min(temp_list)
		length := len(temp_list)
		for i, _ := range temp_list {
			if temp_list[i] != max_temp || temp_list[i] != min_temp {
				sum = temp_list[i]
			}
		}
		fmt.Println(sum / length)
	}
}

func max(a []int) int {
	max := a[0]
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}

func min(a []int) int {
	min := a[0]
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}
