package main

import (
	"fmt"
	"log"
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

	var i int
	fmt.Println("var SupportedRegionsCodes = map[ConstitutionRegionCode]string{")
	doc.Find(".doc-table").Each(func(_ int, s *goquery.Selection) {
		s.Find("td").Each(func(_ int, s *goquery.Selection) {
			if i%2 == 0 {
				fmt.Println(strings.TrimSpace(s.Text()) + `: `)
			} else {
				fmt.Println(`"` + strings.TrimSpace(s.Text()) + `",`)
			}
			i++
		})
	})
	fmt.Println("}")
}

func checkErr(e error) {
	if e != nil {
		log.Panic(e)
	}
}
