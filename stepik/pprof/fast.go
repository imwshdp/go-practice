package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	// "log"
)

type Set struct {
	items map[string]struct{}
}

func NewSet() *Set {
	return &Set{items: make(map[string]struct{})}
}

func (s *Set) Add(item string) {
	s.items[item] = struct{}{}
}

func (s *Set) Has(item string) bool {
	_, exists := s.items[item]
	return exists
}

func (s *Set) Clear() {
	s.items = make(map[string]struct{})
}

func (s *Set) ToSlice() []string {
	result := make([]string, 0, len(s.items))
	for item := range s.items {
		result = append(result, item)
	}
	return result
}

type User struct {
	Browsers []string `json:"browsers"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
}

const ANDROID, MSIE = "Android", "MSIE"

func FastSearch(out io.Writer) {
	seenBrowsers := NewSet()
	var foundUsers strings.Builder
	userIndex := -1

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var user User

		if err := json.Unmarshal(scanner.Bytes(), &user); err != nil {
			panic(err)
		}

		userIndex++
		isAndroid, isMSIE := false, false

		for _, browser := range user.Browsers {
			if strings.Contains(browser, ANDROID) {
				isAndroid = true
				if !seenBrowsers.Has(browser + ANDROID) {
					seenBrowsers.Add(browser + ANDROID)
				}
			}

			if strings.Contains(browser, MSIE) {
				isMSIE = true
				if !seenBrowsers.Has(browser + MSIE) {
					seenBrowsers.Add(browser + MSIE)
				}
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		email := strings.Replace(user.Email, "@", " [at] ", -1)
		foundUsers.WriteString(fmt.Sprintf("[%d] %s <%s>\n", userIndex, user.Name, email))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Fprintf(out, "found users:\n%s\n", foundUsers.String())
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers.ToSlice()))
	seenBrowsers.Clear()
}
