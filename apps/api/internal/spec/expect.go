package spec

import "testing"

func ExpectEqualInts(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Fatalf("ERROR: expected %d but received %d", expected, actual)
	}
}

func ExpectEqualStrings(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Fatalf("ERROR: expected '%s' but received '%s'", expected, actual)
	}
}
