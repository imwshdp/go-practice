package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const (
	DIR_FOUND        = "├"
	DIR_LAST_FOUND   = "└"
	DIR_BRANCH       = "───"
	DIR_PARENT_LEVEL = "│\t"
	SKIP_LEVEL       = "\t"
)

const (
	EMPTY_POSTFIX     = " (empty)"
	NOT_EMPTY_PREFIX  = " ("
	NOT_EMPTY_POSTFIX = "b)"
)

func getFileInfo(path string) (os.FileInfo, error) {
	return os.Stat(path)
}

func addPostfixToName(name string, fileinfo os.FileInfo) (string, error) {
	isDir := fileinfo.Mode().IsDir()

	if !isDir {
		fiSize := fileinfo.Size()
		if fiSize == 0 {
			return name + EMPTY_POSTFIX, nil
		}
		return name + NOT_EMPTY_PREFIX + strconv.Itoa(int(fiSize)) + NOT_EMPTY_POSTFIX, nil
	}

	return name, nil
}

func getFilteredInsides(dirPath string, printFiles bool) ([]string, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	filteredEntries := make([]string, 0, len(entries))

	for _, entry := range entries {
		isDir := entry.IsDir()
		nameToPrint := entry.Name()

		if !printFiles && !isDir {
			continue
		}

		filteredEntries = append(filteredEntries, nameToPrint)
	}

	sort.Strings(filteredEntries)
	return filteredEntries, nil
}

func directoryWalk(out io.Writer, dirPath string, printFiles bool, prefix string) (err error) {
	folderInsides, err := getFilteredInsides(dirPath, printFiles)
	if err != nil {
		return err
	}

	for inx, file := range folderInsides {
		var (
			outputResult string
			nameToPrint  string = file
			fullFilePath string = strings.Join([]string{dirPath, file}, string(filepath.Separator))
		)

		fileinfo, err := getFileInfo(fullFilePath)
		if err != nil {
			return err
		}

		if printFiles {
			nameToPrint, err = addPostfixToName(nameToPrint, fileinfo)
			if err != nil {
				return err
			}
		}

		isLastFileInDir := inx == len(folderInsides)-1

		if isLastFileInDir {
			outputResult += prefix + DIR_LAST_FOUND + DIR_BRANCH + nameToPrint
		} else {
			outputResult += prefix + DIR_FOUND + DIR_BRANCH + nameToPrint
		}

		fmt.Fprintln(out, outputResult)

		if fileinfo.Mode().IsDir() {
			var newPrefix string
			if isLastFileInDir {
				newPrefix = SKIP_LEVEL
			} else {
				newPrefix = DIR_PARENT_LEVEL
			}

			directoryWalk(out, fullFilePath, printFiles, prefix+newPrefix)
		}
	}

	return nil
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	err := directoryWalk(out, path, printFiles, "")
	if err != nil {
		return err
	}

	return nil
}

func main() {
	out := os.Stdout

	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}

	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"

	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
