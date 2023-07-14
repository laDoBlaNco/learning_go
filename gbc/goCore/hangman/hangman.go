package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

var pl = fmt.Println

/*
 +--+
    |
 0  |
/|\ |
/ \ |
   ===

Secret Word: M_N__
Incorrect Guesses; A

Guess a Letter; Y

Sorry Your Dead! The is MONKEY
Yes the Secret Word is MONKEY

Please enter only one letter.
Please enter only letters.
Please enter a letter you haven't guessed.


*/

var hmArr = [7]string{
	" +---+\n" +
		"     |\n" +
		"     |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"     |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		" |   |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|   |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\ |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\ |\n" +
		"/    |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\ |\n" +
		"/ \\ |\n" +
		"    ===\n",
}

var wordArr = [7]string{
	"JAZZ", "ZIGZAG", "ZILCH", "ZIPPER", "ZODIAC", "ZOMBIE", "FLUFF",
}
var randWord string
var guessedLetters string
var correctLetters []string
var wrongGuesses []string

func getRandomWord() string {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randWord = wordArr[rand.Intn(7)]
	correctLetters = make([]string, len(randWord))
	return randWord
}

func showBoard() {
	fmt.Println(hmArr[len(wrongGuesses)])
	fmt.Print("Secret Word: ")
	for _, v := range correctLetters {
		if v == "" {
			fmt.Print("_")
		}
		fmt.Print(v)
	}
	fmt.Print("\nIncorrect Guesses: ")
	if len(wrongGuesses) > 0 {
		for _, v := range wrongGuesses {
			fmt.Print(v + " ")
		}
	}
	fmt.Println()
}

func getUserLetter() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\nGuess a letter: ")
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.ToUpper(guess)
		guess = strings.TrimSpace(guess)
		var isLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
		// this is similar to our 'var pl = fmt.Println', we are basically putting
		// a regexp MatchString function into our  var isLetter
		if len(guess) != 1 {
			fmt.Println("Please enter only 1 letter!")
		} else if !isLetter(guess) {
			fmt.Println("Please enter only letters!")
		} else if strings.Contains(guessedLetters, guess) {
			fmt.Println("Please enter a letter you haven't already guessed!")
		} else {
			return guess
		}
	}
}

func getAllIndices(theStr, subStr string) (indices []int) {
	if len(subStr) == 0 || len(theStr) == 0 {
		return indices
	}
	offset := 0
	for {
		i := strings.Index(theStr[offset:], subStr)
		if i == -1 {
			return indices
		}
		offset += i
		indices = append(indices, offset)
		offset += len(subStr)
	}
}

func updateCorrectLetters(letter string) {
	indexMatches := getAllIndices(randWord, letter)
	for _, v := range indexMatches {
		correctLetters[v] = letter
	}
}

func sliceHasEmptys(theSlice []string) bool {
	for _, v := range theSlice {
		if len(v) == 0 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(getRandomWord())
	for {
		// show gameboard
		showBoard()

		// Get a letter from user
		guess := getUserLetter()
		// A. If they guessed letter in word add to correctLetter
		if strings.Contains(randWord, guess) {
			updateCorrectLetters(guess)
			// 1. Are there more letters to guess?
			if sliceHasEmptys(correctLetters) {
				fmt.Println("More letters to Guess")
			} else {
				// 2. If no more letters to guess (You Win!)
				fmt.Println("Yes the secret word is", randWord)
				break
			}
		} else {
			// B. If they guessed letter not in word.
			// 1. Add new letter to guessedLetters, wrongGuesses
			guessedLetters += guess
			wrongGuesses = append(wrongGuesses, guess)
			// 2. Check if they died.
			if len(wrongGuesses) >= 6 { // 6 turns will get to a full hangman
				showBoard()
				fmt.Println("Sorry your dead! The word is", randWord)
				break
			}
		}

	}
}
