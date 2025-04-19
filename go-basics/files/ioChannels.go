package files

import (
	"fmt"
	"io"
	"log"
	"os"
)

func output() {
	count, err := fmt.Println("Some text here")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)

	_, err = fmt.Fprintln(os.Stdout, "Writed to stdout")
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintln(os.Stderr, "Writed to stderr")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("io_channels.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	_, err = fmt.Fprintln(file, "Written to file")
	_, err = fmt.Fprintln(file, "Hello from fmt.Fprintln!")
	if err != nil {
		log.Fatal(err)
	}
}

func input() {
	var (
		text  string
		text2 string
	)

	count, err := fmt.Scan(&text, &text2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Read: %s %s (%d items)\n", text, text2, count)

	count, err = fmt.Fscan(os.Stdin, &text, &text2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Read: %s %s (%d items)\n", text, text2, count)

	file, err := os.Open("io_channels.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	count, err = fmt.Fscan(file, &text)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	fmt.Printf("Read: %s (%d items)\n", text, count)

}

func args() {
	for _, arg := range os.Args {
		fmt.Println(arg)
	}
}

func IoChannels() {
	output()
	input()
	args()
}
