package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Game struct {
	targetWord  string
	guesses     []string
	maxAttempts int
	attempts    int
}

func getRandomWord() string {
	words := []string{"apple", "banana", "orange", "grape", "melon"}
	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))]
}

func startGame() *Game {
	word := getRandomWord()
	game := &Game{
		targetWord:  word,
		guesses:     make([]string, 0),
		maxAttempts: 15,
		attempts:    0,
	}
	return game
}

func displayWord(game *Game) string {
	display := ""
	for _, letter := range game.targetWord {
		if contains(game.guesses, string(letter)) {
			display += string(letter)
		} else {
			display += "_"
		}
		display += " "
	}
	return display
}

func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func makeGuess(game *Game, guess string) bool {
	for _, previousGuess := range game.guesses {
		if previousGuess == guess {
			fmt.Println("You've already guessed that letter.")
			return false
		}
	}

	game.attempts++
	game.guesses = append(game.guesses, guess)

	if strings.Contains(game.targetWord, guess) {
		fmt.Println("Correct guess!")
		return true
	} else {
		fmt.Println("Incorrect guess.")
		return false
	}
}

func isWinner(game *Game) bool {
	for _, letter := range game.targetWord {
		if !contains(game.guesses, string(letter)) {
			return false
		}
	}
	return true
}

func gameStatus(game *Game) string {
	if game.attempts >= game.maxAttempts {
		return "Game over!. The word was: " + game.targetWord
	} else if isWinner(game) {
		return "Congratulations! You've guessed the word: " + game.targetWord
	} else {
		return "Keep guessing! Attempts remaining: " + fmt.Sprint(game.maxAttempts-game.attempts)
	}
}

func main() {
	game := startGame()

	for {
		fmt.Println("Current word:", displayWord(game))
		fmt.Println("Previous guesses:", game.guesses)

		var guess string
		fmt.Print("Guess the word: ")
		fmt.Scan(&guess)

		if len(guess) != 1 {
			fmt.Println("Please enter only one letter.")
			continue
		}

		if makeGuess(game, guess) {
			if isWinner(game) {
				fmt.Println(gameStatus(game))
				break
			}
		}

		fmt.Println(gameStatus(game))

		if game.attempts >= game.maxAttempts {
			break
		}
	}
}