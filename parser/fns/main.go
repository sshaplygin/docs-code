package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Example output:
// var SupportedTaxDepartments = map[ConstitutionRegionCode]TaxDepart{
// 	0: {
// 		Name: "ФНС России",
// 	},
// 	1: {
// 		Name: "УФНС России по Республике Адыгея",
// 		Branches: map[RegionTaxServiceNumber]string{
// 			1: "Межрайонная инспекция Федеральной налоговой службы №2 по Республике Адыгея",
// 			5: "Межрайонная инспекция Федеральной налоговой службы №1 по Республике Адыгея",
// 			7: "Межрайонная инспекция Федеральной налоговой службы №3 по Республике Адыгея",
// 		},
// 	},
// }

func main() {
	file, err := os.Open("input.txt")
	checkErr(err)

	defer func() {
		err = file.Close()
		checkErr(err)
	}()

	fmt.Println("package main", "\n")
	fmt.Println("var SupportedTaxDepartments = map[ConstitutionRegionCode]TaxDepart{")

	scanner := bufio.NewScanner(file)
	var code int

	var nextRegion bool
	for scanner.Scan() {
		row := scanner.Text()
		code, err = strconv.Atoi(row)
		if err != nil {
			var output string
			if nextRegion {
				//fmt.Println("},")
				output += "Name: "
			}
			fmt.Println(output + "`" + row + "`,")
			if nextRegion {
				fmt.Println("Branches: map[RegionTaxServiceNumber]string{")
			}

			nextRegion = false
			continue
		}

		if code%100 == 0 {
			fmt.Println("},")
			fmt.Println("},")
			fmt.Println(code/100, ": {")
			nextRegion = true
		} else {
			fmt.Print(code%100, ":")
		}
	}

	fmt.Println("},") // close branch map
	fmt.Println("},") // close global map
	fmt.Println("}")

	err = scanner.Err()
	checkErr(err)

}

func checkErr(e error) {
	if e != nil {
		log.Panic(e)
	}
}
