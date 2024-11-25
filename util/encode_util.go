package util

import "encoding/base64"

func EncodeCredentials(credentials string) string {
	var encodedCredential = base64.StdEncoding.EncodeToString([]byte(credentials))
	return encodedCredential
}
