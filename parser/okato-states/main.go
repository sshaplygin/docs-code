package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	file, err := os.Open("data.html")
	checkErr(err)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = file.Close()
		checkErr(err)
	}()

	doc, err := goquery.NewDocumentFromReader(file)
	checkErr(err)

	fmt.Println("var fisrtLevel = map[int]string{")
	doc.Find(".list-group").Each(func(_ int, s *goquery.Selection) {
		s.Find("a").Each(func(_ int, s *goquery.Selection) {
			href, exists := s.Attr("href")
			if !exists {
				return
			}

			parts := strings.Split(href, "/")
			name := strings.Split(s.Text(), " - ")

			fmt.Println(parts[len(parts)-1]+":", `"`+strings.TrimSpace(excludeDigits(name[1]))+`",`)
		})
	})
	fmt.Println("}")
}

func excludeDigits(input string) string {
	var result strings.Builder
	result.Grow(len(input))
	for _, char := range input {
		if !unicode.IsDigit(char) {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
