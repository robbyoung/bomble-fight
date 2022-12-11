package players

import (
	"testing"
)

func TestNewApi(t *testing.T) {
	clearStorage()
	api := Api()

	if api == nil {
		t.Fatalf("Api() returned nil")
	}

	if api.application == nil {
		t.Fatalf("Api object has no application")
	}
}
