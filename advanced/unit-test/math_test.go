package math

import "testing"

const defaultError = "Expected value %v, but the result found is %v."

// click run test
func TestMedia(t *testing.T) {
	expectedValue := 7.28
	value := Media(7.2, 9.9, 6.1, 5.9)

	if value != expectedValue {
		t.Error(defaultError, expectedValue)
	}
}
