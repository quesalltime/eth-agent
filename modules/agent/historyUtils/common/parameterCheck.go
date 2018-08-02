package common

import (
	"errors"
	"fmt"
)

var (
	// AddressLength42 define the length of Address should be 42 with prefix 0x.
	AddressLength42 = 42
	// AddressLength66 define the length of Address should be 66 with prefix 0x.
	AddressLength66 = 66
)

// CheckBlockNumberIndexFormat is check whether the heximal format of block index is correct ot not
func CheckBlockNumberIndexFormat(blockNumber interface{}) (string, error) {
	var blockNumberIndex string
	var message string
	var err error

	switch blockNumber.(type) {
	case nil:
		blockNumberIndex = "latest"
	case string:
		blockNumberIndex = blockNumber.(string)

		if blockNumberIndex != "latest" && blockNumberIndex[0:2] != "0x" {
			message := fmt.Sprintf("Invalid argument %s: hex string should be 0x prefix", blockNumberIndex)
			err = errors.New(message)
		}

	default:
		message = fmt.Sprintf("Invalid argument")
		err = errors.New(message)
	}

	return blockNumberIndex, err
}

// CheckLeftIsLowerThanTheRight check the input a which is less then b or not
func CheckLeftIsLowerThanTheRight(a int64, b int64) bool {
	if a <= b {
		return true
	}
	return false
}

// Check40HeximalFormat will check an address whether it has 0x prefix and is 40 length or not.
func Check40HeximalFormat(address interface{}) (string, error) {
	var addressCheck string
	var message string
	var err error

	switch address.(type) {
	case nil:
		addressCheck = "0x0"
	case string:
		addressCheck = address.(string)

		if addressCheck[0:2] != "0x" {
			message = fmt.Sprintf("Invalid argument address: hex string without 0x prefix")
			err = errors.New(message)
		}
		if len(addressCheck) != AddressLength42 {
			message := fmt.Sprintf("Invalid argument address: hex string should be %d length (include 0x prefix)", AddressLength42)
			err = errors.New(message)
		}
	default:
		message := fmt.Sprintf("Error address format")
		err = errors.New(message)
	}

	return addressCheck, err
}

// Check64HeximalFormat check the addres is heximal format or not and its lenght should be 64 bit long.
func Check64HeximalFormat(address interface{}) (string, error) {
	var addressCheck string
	var message string
	var err error

	switch address.(type) {
	case nil:
		addressCheck = "0x0"
	case string:
		addressCheck = address.(string)

		if addressCheck[0:2] != "0x" {
			message = fmt.Sprintf("Invalid argument address: hex string should have 0x prefix")
			err = errors.New(message)
		}
		if len(addressCheck) != AddressLength66 {
			message = fmt.Sprintf("Invalid argument address: hex string should be %d length (include 0x prefix)", AddressLength66)
			err = errors.New(message)
		}
	default:
		message = "error address type"
		err = errors.New(message)
	}

	return addressCheck, err
}

// CheckBlockNumberFormat only check the value is heximal or not.
// if not, return error.
func CheckBlockNumberFormat(blockNumber string) error {
	var message string
	var err error
	if blockNumber[0:2] != "0x" {
		message = fmt.Sprintf("Invalid argument : hex string should be 0x prefix")
		err = errors.New(message)

		return err
	}

	return nil
}
