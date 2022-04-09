package goset

import "testing"

func TestCreate(t *testing.T) {
	if s := NewSet(); s == nil {
		t.Error("NewSet() returned nil")
	}
}

func TestAddingDuplicates(t *testing.T) {
	s := NewSet()
	s.Add(1)
	s.Add(1)
	s.Add(1)
	if len(s.Get()) != 1 {
		t.Error("Adding duplicate items to a set should not increase the size of the set")
	}
}

func TestRemoveItemFromSet(t *testing.T) {
	s := NewSet()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Remove(2)
	if len(s.Get()) != 2 {
		t.Error("Removing an item from a set should decrease the size of the set")
	}
}
