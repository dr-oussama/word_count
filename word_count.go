package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func lineCount(scanner *bufio.Scanner) int {
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	return lineCount
}

func wordCount(scanner *bufio.Scanner) int {
	wordCount := 0
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wordCount++
	}
	return wordCount
}

func charCount(scanner *bufio.Scanner) int {
	charCount := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		charCount += len(line)
	}
	return charCount
}

func resetScanner(file *os.File, scanner *bufio.Scanner) {
	_, err := file.Seek(0, 0) // Go to the beginning of the file
	check(err)
	*scanner = *bufio.NewScanner(file) // Reset the scanner
}

func main() {
	// Flags
	wordPtr := flag.Bool("w", false, "word count")
	linePtr := flag.Bool("l", false, "line count")
	charPtr := flag.Bool("c", false, "character count")

	flag.Parse()

	// Argument
	path := flag.Arg(0)

	// Open file
	dat, err := os.Open(path)
	check(err)
	defer dat.Close()

	scanner := bufio.NewScanner(dat)
	if *linePtr {
		fmt.Println("number of lines: ", lineCount(scanner))
	} else if *wordPtr {
		fmt.Println("number of words: ", wordCount(scanner))
	} else if *charPtr {
		fmt.Println("number of characters: ", charCount(scanner))
	} else {
		fmt.Print("line: ", lineCount(scanner))
		resetScanner(dat, scanner)
		fmt.Print(", words: ", wordCount(scanner))
		resetScanner(dat, scanner)
		fmt.Print(", char: ", charCount(scanner))
	}
}
