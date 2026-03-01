package helper

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateUniqueNumber generates a unique number in the format: {prefix}-{ddmmyyyy}-{9-digit random number}
func GenerateUniqueNumber(prefix string) string {
	now := time.Now()
	datePart := now.Format("02012006") // Use 8 digits for year (ddmmyyyy)

	// Create a new random source
	r := rand.New(rand.NewSource(now.UnixNano()))

	// Generate a random number between 100,000,000 and 999,999,999 (9 digits)
	randomPart := r.Intn(900000000) + 100000000

	return fmt.Sprintf("%s-%s-%d", prefix, datePart, randomPart)
}
