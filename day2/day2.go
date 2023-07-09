package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func HandVal(hand byte) int {
	return int(hand - 'X' + 1)
}

func WinnerVal(opp_hand, your_hand byte) int {
	/*
		The first column is what your opponent is going to play: A for Rock, B for Paper, and C for Scissors
		Second coloumn X for Rock, Y for Paper, and Z for Scissors
		0 if you lost, 3 if the round was a draw, and 6 if you won
	*/
	switch string([]byte{opp_hand, your_hand}) {
	case "AX": // opp: rock - you: rock - draw
		return 3
	case "AY": // opp: rock - you: paper - win
		return 6
	case "AZ": // opp: rock - you: scissors - lose
		return 0
	case "BX": // opp: paper - you: rock - lose
		return 0
	case "BY": // opp: paper - you: paper - draw
		return 3
	case "BZ": // opp: paper - you: scissors - win
		return 6
	case "CX": // opp: scissors - you: rock - win
		return 6
	case "CY": // opp: scissors - you: paper - lose
		return 0
	case "CZ": // opp: scissors - you: scissors - draw
		return 3
	default:
		log.Fatal("Invalid winner_val value: " + string(opp_hand) + string(your_hand))
		return -999999999
	}
}

func main() {

	input_file, f_err := os.Open("input")

	if f_err != nil {
		log.Fatal(f_err)
	}
	defer input_file.Close()
	// read the file line by line using scanner
	scanner := bufio.NewScanner(input_file)

	var games int = 0
	var won_games int = 0
	var lost_games int = 0
	var score int = 0
	var b_score = 0

	for scanner.Scan() {
		games++

		line_text := scanner.Text()
		if len(line_text) < 3 {
			log.Fatal(`Invalid line ${games}: ${line_text}`)
		}
		opp_hand := line_text[0]
		your_hand := line_text[2]

		winner := WinnerVal(opp_hand, your_hand)
		hand_val := HandVal(your_hand)

		if winner == 6 {
			won_games++
		}
		if winner == 0 {
			lost_games++
		}
		score += hand_val
		score += winner

		a := int(line_text[0] - 'A')
		b := int(line_text[2] - 'X')

		b_score += (b+2+a)%3 + 1 /* shape score (part B) */
		b_score += b * 3         /* outcome score (part B) */

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total games: %d\n", games)
	fmt.Printf("Won games: %d\n", won_games)
	fmt.Printf("Lost games: %d\n", lost_games)
	fmt.Printf("Score: %d\n", score)
	fmt.Printf("B Score: %d\n", b_score)

}
