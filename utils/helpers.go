package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/sshaplygin/docs-code/models"
)

// RandomDigits generates a random positive integer with exactly length decimal
// digits, i.e. in the inclusive range [10^(length-1), 10^length-1]. A length < 1
// is clamped to 1. It uses big.Int arithmetic internally so it never overflows
// for the length while producing the value.
func RandomDigits(length int) int64 {
	if length < 1 {
		length = 1
	}

	// min = 10^(length-1), the smallest number with exactly length digits.
	min := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(length-1)), nil)
	// span = 9 * 10^(length-1), the count of numbers with exactly length digits.
	span := new(big.Int).Mul(min, big.NewInt(9))

	num, err := rand.Int(rand.Reader, span)
	if err != nil {
		panic(fmt.Errorf("generate random digits: %w", err))
	}

	return num.Add(num, min).Int64()
}

// Random generates a uniformly distributed random integer in the inclusive
// range [min, max]. It requires min <= max, otherwise it panics.
func Random(min, max int) int {
	if min > max {
		panic(fmt.Errorf("invalid random range: min %d > max %d", min, max))
	}

	span := int64(max-min) + 1

	randomNumber, err := rand.Int(rand.Reader, big.NewInt(span))
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
		panic(fmt.Sprintf("invalid int code '%s' length: %d, %d", code, len(code), length))
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
