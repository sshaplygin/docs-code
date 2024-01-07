package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/sshaplygin/docs-code/models"
)

// RandomDigits generate random digits required length. Required len > 0.
func RandomDigits(len int) int64 {
	if len <= 0 {
		len = 1
	}

	max, _ := strconv.Atoi(strings.Repeat("9", len))
	min, _ := strconv.Atoi("1" + strings.Repeat("0", len-1))

	num, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	return num.Int64() + int64(min)
}

// Random generate random digit in range [min, max]. Required max > 0.
func Random(min, max int) int {
	if max == 0 || min == max {
		max += 1
	}

	randomNumber, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	if err != nil {
		panic(fmt.Errorf("generate random number: %w", err))
	}

	return int(randomNumber.Int64()) + min
}

func StrToArr(str string) ([]int, error) {
	numbers := strings.Split(str, "")
	arr := make([]int, 0, len(numbers))
	for _, number := range numbers {
		number, err := strconv.Atoi(number)
		if err != nil {
			return nil, models.ErrInvalidValue
		}
		arr = append(arr, number)
	}
	return arr, nil
}

func SliceToInt(data []int) int {
	var res int
	for _, num := range data {
		res = res*10 + num
	}
	return res
}

func CodeToInts(code int) []int {
	var res []int

	for code > 0 {
		digit := code % 10
		res = append([]int{digit}, res...)
		code /= 10
	}

	return res
}

// StrCode method could throw two panics.
func StrCode(val, length int) string {
	if length < 1 {
		panic("invalid required code length")
	}

	var str strings.Builder
	code := strconv.Itoa(int(val))

	n := length
	if len(code) > length {
		panic(fmt.Sprintf("invalid int code length: %d, %d", len(code), length))
	}

	str.Grow(n)

	for i := 0; i+len(code) < length; i++ {
		str.WriteString("0")
	}
	str.WriteString(code)

	return str.String()
}

func FillSlice(from, to []int, fromIdx int) {
	idx := fromIdx
	for i := len(from) - 1; i >= 0; i-- {
		to[idx] = from[i]
		idx--
	}
}
