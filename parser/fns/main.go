package main

import "os"

func main() {
	file, err := os.Open("input.txts")
	checkErr(err)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = file.Close()
		checkErr(err)
	}()

}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
