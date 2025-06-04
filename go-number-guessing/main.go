package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	low       = flag.Int("low", 1, "Lowest number to start range")
	high      = flag.Int("high", 100, "Highest number to end range")
	timeLimit = flag.Int64("time-limit", 15, "Set the total time the game will take")
)

type Game struct {
	answer     int
	difficulty int
	chances    int
}

func loadGame() *Game {
	var game Game

	game.answer = randRange(*low, *high)

	return &game
}

func (game *Game) runGame(wg *sync.WaitGroup) {
	timer := time.NewTicker(time.Second * time.Duration(*timeLimit))
	done := make(chan bool)
	defer wg.Done()
	wg.Add(2)

	scanner := bufio.NewScanner(os.Stdin)

	var userAnswer string
	numTries := 0

	go func() {
		game.Difficulty()

		fmt.Println("Make your guess:")
	gameLoop:
		for scanner.Scan() {

			userAnswer = scanner.Text()
			userNumber, err := strconv.Atoi(userAnswer)
			fatalError("Failed to parse Number", err)

			if userNumber == game.answer {
				fmt.Println("You got it right. Kudos!")
				done <- true
				break gameLoop
			}

			if userNumber != game.answer {
				fmt.Print("Incorrect! ")
				fmt.Println("Enter your guess: ")

				numTries++

				if numTries > 5 {
					fmt.Println("You have exceeded all chances. Try again.")
					done <- true
					break gameLoop
				}

				continue gameLoop
			}

		}
	}()

	select {
	case <-done:
	case <-timer.C:
		fmt.Println("time's up, too bad you didn't get it. The answer was ", game.answer)
	}

}

func (game Game) Difficulty() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(`Please select the difficulty level:
			1. Easy (10 chances)
			2. Medium (5 chances)
			3. Hard (3 chances)`)
	fmt.Printf("Enter your choice: ")
	scanner.Scan()
	userDifficultyAnswer := scanner.Text()
	difficultyInt, err := strconv.Atoi(userDifficultyAnswer)
	fatalError("Failed to parse difficulty level: ", err)

	game.difficulty = difficultyInt

	switch game.difficulty {
	case 1:
		game.chances = 10
		fmt.Println("Great! You have selected the Easy level")
	case 2:
		game.chances = 5
		fmt.Println("Great! You have selected the Medium level")
	case 3:
		game.chances = 3
		fmt.Println("Great! You have selected the Hard level")
	}

}

func main() {
	var wg sync.WaitGroup
	fmt.Println(`
		Welcome to the Number Guessing Game!
		I'm thinking of a number between 1 and 100.
		You have 5 chances to guess the correct number.
		`)
	game := loadGame()
	game.runGame(&wg)

}

func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func fatalError(message string, err error) {
	if err != nil {
		log.Fatalln(message, ":", err)
	}
}
