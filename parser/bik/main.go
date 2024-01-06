package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

var devNull *os.File

func main() {
	file, err := os.ReadFile("base.xml")
	checkErr(err)

	var biks Biks
	err = xml.Unmarshal(file, &biks)
	checkErr(err)

	writer := devNull
	fmt.Fprint(writer, "package main", "\n\n")

	printAvailablesBiks(writer, &biks)
}

func printAvailablesBiks(writer io.Writer, biks *Biks) {
	fmt.Fprint(writer, "var existsBIKs = map[string]string{", "\n")
	for _, bik := range biks.BikRows {
		fmt.Fprint(writer, "\t", `"`+bik.Bik+`": `+"`"+bik.Name+"`,", "\n")
	}
	fmt.Fprint(writer, "}", "\n")
}

type Biks struct {
	XMLName xml.Name `xml:"biks"`
	BikRows []Bik    `xml:"bik"`
}

type Bik struct {
	XMLName    xml.Name `xml:"bik"`
	Version    string   `xml:"version,attr"`
	Bik        string   `xml:"bik,attr"`
	Ks         string   `xml:"ks,attr"`
	Name       string   `xml:"name,attr"`
	NameMini   string   `xml:"namemini,attr"`
	Index      string   `xml:"index,attr"`
	City       string   `xml:"city,attr"`
	Address    string   `xml:"address,attr"`
	Phone      string   `xml:"phone,attr"`
	Okato      string   `xml:"okato,attr"`
	Okpo       string   `xml:"okpo,attr"`
	RegNum     string   `xml:"regnum,attr"`
	Srok       string   `xml:"srok,attr"`
	DateAdd    string   `xml:"dateadd,attr"`
	DateChange string   `xml:"datechange,attr"`
}

func checkErr(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func init() {
	var err error
	devNull, err = os.OpenFile(os.DevNull, os.O_WRONLY, 0600)
	checkErr(err)
}
