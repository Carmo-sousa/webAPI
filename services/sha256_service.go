package services

import (
	"crypto/sha256"
	"fmt"
)

func SHA265Encoder(s string) string {
	str := sha256.Sum256([]byte(s))

	return fmt.Sprintf("%x", str)
}
