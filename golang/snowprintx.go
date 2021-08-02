package snowprintx

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// max string length of fingerprint < 128 chars
// int64 chars unix second + 32 chars description + 32 random hex

const (
	colonDelimiter       = ":"
	randomness           = 32
	descriptionMaxLength = 32
)

var (
	errDescriptionMaxLength = errors.New(
		fmt.Sprint(
			"description is over max length: ",
			descriptionMaxLength,
		),
	)
)

func generateRandomByteArray(n int) (*[]byte, error) {
	token := make([]byte, n)
	length, errRandom := rand.Read(token)
	if errRandom != nil || length != n {
		return nil, errRandom
	}

	return &token, nil
}

func CreateSnowprint(description string) (*string, error) {
	if len(description) > descriptionMaxLength {
		return nil, errDescriptionMaxLength
	}

	randomBytes, errRandomBytes := generateRandomByteArray(randomness)
	if errRandomBytes != nil {
		return nil, errRandomBytes
	}

	now := time.Now().Unix()
	bytesAsStr := hex.EncodeToString(*randomBytes)
	preparedID := fmt.Sprint(
		now,
		colonDelimiter,
		description,
		colonDelimiter,
		bytesAsStr,
	)

	return &preparedID, nil
}
