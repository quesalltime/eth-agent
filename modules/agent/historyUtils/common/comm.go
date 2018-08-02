package common

import (
	"errors"
	"fmt"
	"strconv"
)

// ParseHexToInt64 convet heximal number into decimal number with type int64
func ParseHexToInt64(heximal string) (int64, error) {
	var decimalBytes32 int64
	var err error

	decimalBytes32, StrcovErr := strconv.ParseInt(heximal[2:], 16, 64)
	if StrcovErr != nil {
		message := fmt.Sprintf("Please check your parameters whether it is heximal format or not. Wrong Parameter:%s", heximal)
		err = errors.New(message)
	}
	return decimalBytes32, err
}

// ParseInt64ToHex convet decimal number into heximal number with type string
func ParseInt64ToHex(decimal int) string {
	var i64 int64
	i64 = int64(decimal)

	heximal := strconv.FormatInt(i64, 16)
	heximal = "0x" + heximal
	return heximal
}
