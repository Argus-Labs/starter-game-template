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

// TODO: this is pretty illegal lol
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
