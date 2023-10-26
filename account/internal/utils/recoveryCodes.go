package utils

import (
	"math/rand"
	"time"
)

// Function to generate recovery codes
func RecoveryCodes(numCodes int, codeLength int) []string {
	recoveryCodes := make([]string, numCodes)
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	for i := 0; i < numCodes; i++ {
		code := ""
		for j := 0; j < codeLength; j++ {
			code += string(byte(65 + random.Intn(26))) // Generate random uppercase letters (A-Z)
		}
		recoveryCodes[i] = code
	}

	return recoveryCodes
}
