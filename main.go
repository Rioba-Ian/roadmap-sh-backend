package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	scanner   = bufio.NewScanner(os.Stdin)
	low       = flag.Int("low", 1, "Lowest number to start range")
	high      = flag.Int("high", 100, "Highest number to end range")
	timeLimit = flag.Int64("time-limit", 10, "Set the total time the game will take")
)

type Game struct {
	answer int
	guess  string
}

func loadGame() *Game {
	var game Game

	game.answer = randRange(*low, *high)

	fmt.Println("The answer is", game.answer)

	return &game
}

func (game *Game) runGame() {
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	answerCh := make(chan string)
	fmt.Println("Make your guess:")
	go func() {
		scanner.Scan()
		userAnswer := scanner.Text()
		answerCh <- userAnswer
	}()

	select {
	case <-timer.C:
		fmt.Println("Time Ran out. Please try again.")

	case guessed := <-answerCh:
		guessedNum, err := strconv.Atoi(guessed)
		fatalError("Error in parsing number", err)

		if guessedNum == game.answer {
			fmt.Println("You got it correct! Kudos!!")
			close(answerCh)
		} else {
			fmt.Println("Got wrong answer")
		}

	}

	return
}

func main() {
	fmt.Println(`
		Number guessing game cli!
		I'm thinking of a number between 1 and 100.
		You have 5 chances to guess the correct number.
		`)
	game := loadGame()
	game.runGame()
}

func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func fatalError(message string, err error) {
	if err != nil {
		log.Fatalln(message, ":", err)
	}
}
