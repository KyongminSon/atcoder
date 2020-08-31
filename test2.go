package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var save []int
	now := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		switch scanner.Text() {
		case "add":
			now++
		case "print":
			fmt.Println(now)
		case "rollback":
			{
				now = save[len(save)-1]
				save = save[:len(save)-1]
			}
		case "pin":
			save = append(save, now)
		default:
			continue
		}
	}
}
