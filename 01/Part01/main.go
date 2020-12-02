package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	amounts, err := readAmounts("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

OUTER:
	for _, x := range amounts {
		for _, y := range amounts {
			if x+y == 2020 {
				fmt.Println(x * y)
				break OUTER
			}
		}
	}
}

func readAmounts(filename string) ([]int, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}

	amounts := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		amount, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		amounts = append(amounts, amount)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return amounts, nil
}
