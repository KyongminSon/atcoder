package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var num, cost int
	var height uint64
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
