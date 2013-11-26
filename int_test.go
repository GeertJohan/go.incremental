package incremental

import (
	"testing"
)

type user struct {
	i Int
}

func TestIntPtr(t *testing.T) {
	i := &Int{}
	num := i.Next()
	if num != 1 {
		t.Fatal("expected 1, got %d", num)
	}
	num = i.Next()
	if num != 2 {
		t.Fatal("expected 2, got %d", num)
	}
}

func TestAsField(t *testing.T) {
	u := user{}
	num := u.i.Next()
	if num != 1 {
		t.Fatal("expected 1, got %d", num)
	}
	num = u.i.Next()
	if num != 2 {
		t.Fatal("expected 2, got %d", num)
	}
}