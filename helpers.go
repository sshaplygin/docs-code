package ru_doc_code

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"strings"
)

func RandomDigits(len int) int64 {
	if len <= 0 {
		len = 1
	}
	max, _ := strconv.Atoi(strings.Repeat("9", len))
	min, _ := strconv.Atoi("1" + strings.Repeat("0", len-1))

	num, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	return num.Int64() + int64(min)
}

func StrToArr(str string) ([]int, error) {
	numbers := strings.Split(str, "")
	arr := make([]int, 0, len(numbers))
	for _, number := range numbers {
		number, err := strconv.Atoi(number)
		if err != nil {
			return []int{}, ErrInvalidValue
		}
		arr = append(arr, number)
	}
	return arr, nil
}
