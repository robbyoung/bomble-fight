package spec

import "testing"

func ExpectEqualInts(t *testing.T, expected int, actual int, message string) {
	if expected != actual {
		t.Fatalf("ERROR: %s - Expected %d but received %d", message, expected, actual)
	}
}

func ExpectEqualStrings(t *testing.T, expected string, actual string, message string) {
	if expected != actual {
		t.Fatalf("ERROR: %s - Expected '%s' but received '%s'", message, expected, actual)
	}
}
