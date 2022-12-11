package players

import (
	"bomble-fight/internal/spec"
	"testing"
)

func TestNewLocalPlayerStorage(t *testing.T) {
	cs := newLocalStorage()

	if cs == nil {
		t.Fatalf("NewLocalPlayerStorage() returned nil")
	}

	if len(cs.players) != 0 {
		t.Fatalf("NewLocalPlayerStorage() didn't return an empty store")
	}
}

func TestPlayerStorageSaveAndLoad(t *testing.T) {
	cs := newLocalStorage()

	original := newAggregate("aname")
	original.Money = 200

	cs.SavePlayer(original)

	loaded := cs.LoadPlayer(original.Id)

	if original == loaded {
		t.Fatalf("Loaded player uses same pointer as original")
	}

	spec.ExpectEqualStrings(t, original.Id, loaded.Id, "Id mismatch")
	spec.ExpectEqualStrings(t, original.Name, loaded.Name, "Name mismatch")
	spec.ExpectEqualInts(t, original.Money, loaded.Money, "MaxHealth mismatch")

}
