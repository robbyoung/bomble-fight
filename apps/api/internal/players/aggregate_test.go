package players

import (
	"bomble-fight/internal/spec"
	"testing"
)

func TestAggregate_New(t *testing.T) {
	agg := newAggregate("aname")

	if agg == nil {
		t.Fatalf("newAggregate() returned nil")
	}

	if len(agg.Id) != 36 {
		t.Fatalf("Aggregate ID has unexpected value %s", agg.Id)
	}

	spec.ExpectEqualStrings(t, "aname", agg.Name, "Unexpected name value")
	spec.ExpectEqualInts(t, 100, agg.Money, "Unexpected money value")
}

func TestAggregate_Persistence(t *testing.T) {
	orig := newAggregate("aname")
	orig.Money = 250

	model := orig.toPersistence()
	loaded := fromPersistence(model)

	if orig == loaded {
		t.Fatalf("Original and loaded players use same reference")
	}

	spec.ExpectEqualStrings(t, orig.Id, loaded.Id, "Original and loaded IDs do not match")
	spec.ExpectEqualStrings(t, orig.Name, loaded.Name, "Original and loaded names do not match")
	spec.ExpectEqualInts(t, orig.Money, loaded.Money, "Original and loaded money do not match")
}

func TestAggregate_SpendMoney(t *testing.T) {
	agg := newAggregate("aname")

	err := agg.SpendMoney(30)
	spec.ExpectEqualInts(t, 70, agg.Money, "Unexpected money value")

	if err != nil {
		t.Error("Unexpected error from SpendMoney()")
	}
}

func TestAggregate_SpendMoney_TooMuch(t *testing.T) {
	agg := newAggregate("aname")

	err := agg.SpendMoney(110)
	spec.ExpectEqualInts(t, 100, agg.Money, "Unexpected money value")

	if err == nil {
		t.Error("Expected error from SpendMoney(), received none")
	}
}

func TestAggregate_SpendMoney_Negative(t *testing.T) {
	agg := newAggregate("aname")

	err := agg.SpendMoney(-30)
	spec.ExpectEqualInts(t, 100, agg.Money, "Unexpected money value")

	if err == nil {
		t.Error("Expected error from SpendMoney(), received none")
	}
}

func TestAggregate_EarnMoney(t *testing.T) {
	agg := newAggregate("aname")

	err := agg.EarnMoney(30)
	spec.ExpectEqualInts(t, 130, agg.Money, "Unexpected money value")

	if err != nil {
		t.Error("Unexpected error from EarnMoney()")
	}
}

func TestAggregate_EarnMoney_Negative(t *testing.T) {
	agg := newAggregate("aname")

	err := agg.EarnMoney(-30)
	spec.ExpectEqualInts(t, 100, agg.Money, "Unexpected money value")

	if err == nil {
		t.Error("Expected error from EarnMoney(), received none")
	}
}
