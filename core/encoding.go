package core

import (
	"encoding/base64"
)

func EncodeString(rawStr string) string {

	// base64.StdEncoding: Standard encoding with padding
	// It requires a byte slice so we cast the string to []byte
	encodedStr := base64.StdEncoding.EncodeToString([]byte(rawStr))

	return encodedStr
}

func DecodeString(encodedStr string) string {

	decodedStr, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		panic("malformed input" + encodedStr)
	}

	return string(decodedStr)
}
