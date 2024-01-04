package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("base.xml")
	checkErr(err)

	var biks Biks
	err = xml.Unmarshal(file, &biks)
	checkErr(err)

	for _, bik := range biks.BikRows {
		fmt.Println(bik)
	}

	// printAvailablesBiks(&biks)
}

func printAvailablesBiks(biks *Biks) {
	fmt.Println("var existsBIKs = map[string]string{")
	for _, bik := range biks.BikRows {
		fmt.Println(`"` + bik.Bik + `": ` + "`" + bik.Name + "`,")
	}
	fmt.Println("}")
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
		panic(e)
	}
}
