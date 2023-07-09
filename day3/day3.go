package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// fmt.Println(fmt.Sprintf("a int is %d, and A is %d", int('a'), int('A')))
	// os.Exit(0)

	input_file, f_err := os.Open("input")
	if f_err != nil {
		log.Fatal(f_err)
	}
	defer input_file.Close()
	scanner := bufio.NewScanner(input_file)

	var sac int = 0
	var sum = 0
	var b_sum int = 0
	last2_str := [2]string{"", ""}

	for scanner.Scan() {
		if sac%3 < 2 {
			last2_str[sac%3] = scanner.Text()
		}
		line_text := scanner.Text()
		mid := len(line_text) / 2
		firstHalf := line_text[:mid]
		secondHalf := line_text[mid:]
		var match byte = '-'
		for i := 0; i < mid; i++ {
			for j := 0; j < mid; j++ {
				if firstHalf[i] == secondHalf[j] {
					match = firstHalf[i]
					break
				}
			}
			if match != '-' {
				break
			}
		}
		pos := ((int(match) - 64) % 32)
		scale := ((int(match) - 64 + 32) % 64 / 32)
		calced := pos + 26*scale
		sum += calced
		if sac < 10 {
			fmt.Printf(
				"firstHalf: %s, secondHalf: %s, match: %c, int: %d, calced: %d, pos %d, scale %d\n",
				firstHalf, secondHalf, match, int(match), calced, pos, scale)
		}

		if (sac)%3 == 2 {
			var card_match byte = '-'
			for i := 0; i < len(line_text); i++ {
				for j := 0; j < len(last2_str[0]); j++ {
					if line_text[i] == last2_str[0][j] {
						for k := 0; k < len(last2_str[1]); k++ {
							if line_text[i] == last2_str[1][k] {
								card_match = line_text[i]
								break
							}
						}
						break
					}
				}
				if card_match != '-' {
					break
				}
			}
			b_sum += ((int(card_match) - 64) % 32) + 26*((int(card_match)-64+32)%64/32)
		}

		sac++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total lines: %d\n", sac)
	fmt.Printf("Sum of all matches: %d\n", sum)
	fmt.Printf("Sum of all b matches: %d\n", b_sum)

}
