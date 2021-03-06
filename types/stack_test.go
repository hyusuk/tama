package types

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack(100)
	if s.Len() != 100 {
		t.Fatalf("expected %d, but got %d", 100, s.Len())
	}
	if s.Top() != nil {
		t.Fatalf("expected nil")
	}
}

func TestStackPushAndPop(t *testing.T) {
	s := NewStack(100)
	s.Push(Number(1))
	num := s.Pop()
	if num == nil {
		t.Fatalf("unexpected value")
	}
	if num.String() != "1" {
		t.Fatalf("expected %s, but got %s", "1", num.String())
	}
	if s.Pop() != nil {
		t.Fatalf("expected nil")
	}
	if s.Top() != nil {
		t.Fatalf("expected nil")
	}
}
