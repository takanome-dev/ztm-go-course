//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//
//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"unicode"
)

type Count struct {
	count int
	sync.Mutex
}

func getWords(line string) []string {
	return strings.Split(line, " ")
}

func countLetters(word string) int {
	count := 0

	for _, char := range word {
		if unicode.IsLetter(char) {
			count += 1
		}
	}

	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	totalLetters := Count{}
	var wg sync.WaitGroup

	for {
		if scanner.Scan() {
			line := scanner.Text()
			words := getWords(line)

			for _, word := range words {
				wordCopy := word
				wg.Add(1)

				go func() {
					totalLetters.Lock()
					defer totalLetters.Unlock()
					defer wg.Done()
					count := countLetters(wordCopy)
					totalLetters.count += count
				}()
			}
		} else {
			break
		}
	}

	wg.Wait()

	totalLetters.Lock()
	sum := totalLetters.count
	totalLetters.Unlock()

	fmt.Println("total letters =", sum)
}
