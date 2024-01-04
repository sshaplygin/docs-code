package main

import (
	"fmt"
	"os"
	"strings"

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

	// fmt.Println("var fisrtLevel = map[int]string{")
	doc.Find(".doc-table").Each(func(_ int, s *goquery.Selection) {
		// fmt.Println(s.Text())
		s.Find("td").Each(func(_ int, s *goquery.Selection) {
			// s.
			fmt.Println("->", strings.TrimSpace(s.Text()))
			// href, exists := s.Attr("href")
			// if !exists {
			// 	return
			// }

			// parts := strings.Split(href, "/")
			// name := strings.Split(s.Text(), " - ")

			// fmt.Println(parts[len(parts)-1]+":", `"`+strings.TrimSpace(excludeDigits(name[1]))+`",`)
		})
	})
	// fmt.Println("}")
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
