package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	// "strconv"
)

func main() {
	file, err := os.Open("aoc2023-01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	number_words := []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	answer := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		first_digit_set := false
		first_digit := -1
		last_digit := -1
		digit := -1
		line := scanner.Text()

		for beg := 0; beg < len(line); beg++ {
			for n, w := range number_words {
				if strings.HasPrefix(line[beg:], w) {
					digit = n
					break
				}
			}
			b0 := byte('0')
			b9 := byte('9')
			if b0 <= line[beg] && line[beg] <= b9 {
				//fmt.Println("line[beg] =", line[beg], " byte(line[beg]) =", byte(line[beg]))
				digit = int(line[beg] - b0)
			}

			if digit != -1 {
				if !first_digit_set {
					first_digit_set = true
					first_digit = digit
				} else {
					last_digit = digit
				}
			}
		}

		if last_digit == -1 {
			fmt.Printf("Only one digit (%v) on line\n", first_digit)
			last_digit = first_digit
		}

		fmt.Println(first_digit, last_digit)
		answer += 10*first_digit + last_digit
	}

	fmt.Println(answer)
}
