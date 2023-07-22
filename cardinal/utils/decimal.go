package utils

import (
	"math/big"

	"github.com/ericlagergren/decimal"
)

var (
	DecCtx = decimal.Context64
)

func IntToDec(i int) *decimal.Big {
	return decimal.New(int64(i), 0)
}

func DecToInt(dec *decimal.Big) int {
	bigInt := new(big.Int)
	dec.Int(bigInt)
	return int(bigInt.Int64())
}

func StrToDec(str string) *decimal.Big {
	dec := new(decimal.Big)
	DecCtx.SetString(dec, str)
	return dec
}

func Uint16ToDec(i uint16) *decimal.Big {
	return decimal.New(int64(i), 0)
}

func DecToStr(dec *decimal.Big) string {
	return dec.String()
}

func LessThan(a, b *decimal.Big) bool {
	return a.Cmp(b) == -1
}

func LessThanOrEqual(a, b *decimal.Big) bool {
	return a.Cmp(b) <= 0
}

func GreaterThan(a, b *decimal.Big) bool {
	return a.Cmp(b) == 1
}

func GreaterThanOrEqual(a, b *decimal.Big) bool {
	return a.Cmp(b) >= 0
}

func Equal(a, b *decimal.Big) bool {
	return a.Cmp(b) == 0
}

func DecMax(a, b *decimal.Big) *decimal.Big {
	if GreaterThan(a, b) {
		return a
	}
	return b
}

func DecMin(a, b *decimal.Big) *decimal.Big {
	if LessThan(a, b) {
		return a
	}
	return b
}
