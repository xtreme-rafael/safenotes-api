package crypto

import (
	"errors"
)

var InvalidNumberOfBytesError = errors.New("invalid number of bytes requested")
var UnableToGenerateBytesError = errors.New("unable to generate random bytes")

func GenerateRandomBytes(numBytes int) ([]byte, error) {
	if numBytes < 0 {
		return nil, InvalidNumberOfBytesError
	}

	byteArray := make([]byte, numBytes)
	n, err := RandRead(byteArray)
	if err != nil || n != numBytes {
		return nil, UnableToGenerateBytesError
	}

	return byteArray, nil
}

func GenerateRandomBase64String(numBytes int) (string, error) {
	randomBytes, err := GenerateRandomBytes(numBytes)
	if err != nil {
		return "", err
	}

	return Base64EncodeToString(randomBytes), nil
}
