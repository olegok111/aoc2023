package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	file, err := os.Open("aoc2023-02.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	valid_cube_cnts := map[string]int {
		"red": 12,
		"green": 13,
		"blue": 14,
	}
	res := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		left, right, found := strings.Cut(line, ":")
		if !found {
			break
		}

		_, gid_s, _ := strings.Cut(left, " ")
		gid, err := strconv.Atoi(gid_s)
		if err != nil {
			log.Fatal("gid_s:", err)
		}
		
		possible := true
		cubesets := strings.Split(strings.TrimLeft(right, " "), "; ")
		for _, cubeset := range cubesets {
			cubes := strings.Split(cubeset, ", ")
			for _, cube := range cubes {
				count_s, color, found := strings.Cut(cube, " ")
				if !found {
					log.Fatal("cube cut: \" \" not found")
				}
				count, err := strconv.Atoi(count_s)
				if err != nil {
					log.Fatal("count_s:", err)
				}

				if valid_cnt, ok := valid_cube_cnts[color]; !ok || count > valid_cnt {
					possible = false
					break
				}
			}
		}

		if possible {
			res += gid
		}
	}

	fmt.Println(res)
}
