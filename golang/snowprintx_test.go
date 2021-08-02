package snowprintx

import (
	"fmt"
	"testing"
)

func TestGenerateRandomByteArray(t *testing.T) {
	testLength := 128

	randomBytes, errRandomBytes := generateRandomByteArray(testLength)
	if errRandomBytes != nil {
		t.Fail()
		t.Logf(errRandomBytes.Error())
	}

	if randomBytes == nil {
		t.Fail()
		t.Logf("randomBytes should not be nil")
	}

	randomByteLength := len(*randomBytes)

	if randomByteLength != testLength {
		t.Fail()
		t.Logf(
			fmt.Sprint(
				"randomBytes should be:",
				testLength,
				", instead found:",
				randomByteLength,
			),
		)
	}
}

func TestCreateSnowprint(t *testing.T) {
	snowprintTest, errSnowprintTest := CreateSnowprint("thirty_two_character_test_string")
	if errSnowprintTest != nil {
		t.Fail()
		t.Logf(errSnowprintTest.Error())
	}

	if snowprintTest == nil {
		t.Fail()
		t.Logf("snowprintTest should not be nil")
	}

	snowprintTestLength := len(*snowprintTest)
	if snowprintTestLength > 128 {
		t.Fail()
		t.Logf("snowprintTest is over 128 char length")
	}
}
