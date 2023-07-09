package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	input_file, f_err := os.Open("input")

	if f_err != nil {
		log.Fatal(f_err)
	}
	defer input_file.Close()
	// read the file line by line using scanner
	scanner := bufio.NewScanner(input_file)

	var line_counter int = 1
	var elves int = 1
	var current_calories int = 0
	var max_calories = [3]int{0, 0, 0}

	for scanner.Scan() {
		line_counter++
		// do something with a line
		// fmt.Printf("line: %s\n", scanner.Text())

		line_text := scanner.Text()
		cal, str_err := strconv.Atoi(line_text)
		if str_err != nil {
			if current_calories > max_calories[0] {
				max_calories[0] = current_calories
				sort.Ints(max_calories[:])
			}
			current_calories = 0
			elves += 1
		} else {
			current_calories += cal
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum_calories_top_3_elves := 0

	for _, val := range max_calories {
		sum_calories_top_3_elves += val
	}

	fmt.Printf("Total lines: %d\n", line_counter)
	fmt.Printf("Total elves: %d\n", elves)
	fmt.Printf("Max Calories of an elf is : %d\n", max_calories[2])
	fmt.Printf("Sum of top 3 max Calorie elves is: %d\n", sum_calories_top_3_elves)

	// TODO Convert input to int
}
