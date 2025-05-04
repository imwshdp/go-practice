package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func scanExample() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}
}

func scanWithMapCacheExample() {
	scanner := bufio.NewScanner(os.Stdin)
	uniqueSet := make(map[string]bool)

	for scanner.Scan() {
		text := scanner.Text()

		if _, ok := uniqueSet[text]; ok {
			continue
		}

		uniqueSet[text] = true
		fmt.Println(text)
	}
}

func ScanSortedExample(input io.Reader, output io.Writer) error {
	scanner := bufio.NewScanner(input)

	var prev string

	for scanner.Scan() {
		text := scanner.Text()

		if text < prev {
			return fmt.Errorf("File is not sorted")
		}

		if text == prev {
			continue
		}

		fmt.Fprintln(output, text)
		prev = text
	}

	return nil
}

func scanners() {
	// fmt.Println("SCANNER EXAMPLE\n===")
	// scanExample()
	// fmt.Println("===")

	// fmt.Println("SCANNER WITH CACHE EXAMPLE\n===")
	// scanWithMapCacheExample()
	// fmt.Println("===")

	fmt.Println("SORTED DATA SCANNER EXAMPLE\n===")
	if err := ScanSortedExample(os.Stdin, os.Stdout); err != nil {
		panic(err.Error())
	}
	fmt.Println("===")
}

func main() {
	scanners()
}
