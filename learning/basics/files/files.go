package files

import (
	"fmt"
	"io"
	"log"
	"os"
)

type numsReader struct {
	nums string
}

func (reader numsReader) Read(buf []byte) (counter int, err error) {
	var count int

	for inx := 0; inx < len(reader.nums); inx++ {
		if reader.nums[inx] >= '0' && reader.nums[inx] <= '9' {
			buf[count] = reader.nums[inx]
			count++
		}
	}

	return count, io.EOF
}

func simpleReaderEx() {
	buf := make([]byte, 10)
	reader := numsReader{"1,2,3,4,5,6,7,8,=g9"}

	count, error := reader.Read(buf)
	if error != nil && error != io.EOF {
		log.Fatal(error)
	}

	fmt.Println(string(buf), count)

}

type rowsReader struct {
	text string
}

func (reader *rowsReader) Read(buf []byte) (counter int, err error) {
	var inx int

	for inx = 0; inx < len(reader.text); inx++ {
		if reader.text[inx] == '\n' {
			reader.text = reader.text[inx+1:]
			break
		}

		buf[inx] = reader.text[inx]

		if inx == len(reader.text)-1 {
			reader.text = ""
			return inx + 1, io.EOF
		}
	}

	return inx + 1, nil
}

func rowsReaderEx() {
	rowsReader := rowsReader{text: "1,2,3\n4,5,6\n7,8,9"}

	var (
		err   error
		count int
	)

	rowBuff := make([]byte, 100)
	for err != io.EOF {
		count, err = rowsReader.Read(rowBuff)
		fmt.Println(string(rowBuff), count)
	}
}

type numsWriter struct {
	stored []byte
}

func (writer numsWriter) Write(buf []byte) (counter int, err error) {
	var count int

	for inx := 0; inx < len(buf); inx++ {
		if buf[inx] >= '0' && buf[inx] <= '9' {
			writer.stored[count] = buf[inx]
			count++
		}
	}

	return count, nil
}

func simpleWriterEx() {
	nums := []byte{'1', ',', '2', '=', '3', '.'}
	writer := numsWriter{stored: make([]byte, 10)}

	count, err := writer.Write(nums)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(writer.stored), count)
}

func osFileEx() {
	newFile, err := os.Create("new.txt")
	if err != nil {
		log.Fatal(err)
	}

	descriptor, err := newFile.WriteString("Hello, world!")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(descriptor)

	err = newFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("new.txt")
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 100)
	descriptor, err = file.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Println(string(buf), descriptor)

	file, err = os.OpenFile("new.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if descriptor, err = file.WriteString("\nHello, world! 2"); err != nil {
		log.Fatal(err)
	}
}

func Files() {
	fmt.Println("Simple reader: ")
	simpleReaderEx()

	fmt.Println("\nRows reader: ")
	rowsReaderEx()

	fmt.Println("\nSimple writer: ")
	simpleWriterEx()

	fmt.Println("\nOS: ")
	osFileEx()
}
