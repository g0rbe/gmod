package slices

import (
	"fmt"
	"testing"
)

type Test struct {
	i int
	f float64
}

// Implement comparable interface on Test
func (a Test) Compare(b Test) int {
	switch {
	case a.i < b.i:
		return -1
	case a.i > b.i:
		return 1
	case a.f < b.f:
		return -1
	case a.f > b.f:
		return 1
	default:
		return 0
	}
}

func (t Test) String() string {
	return fmt.Sprintf("%d|%.0f", t.i, t.f)
}

// Test with a user defined type
func TestCopyUser(t *testing.T) {

	a := []Test{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}}

	b := Copy(a)

	if b == nil {
		t.Fatalf("FAILED!")
	}
}

// Test with a builtin type
func TestCopyBuiltin(t *testing.T) {

	a := []int{0, 1, 2, 3, 4}

	b := Copy(a)

	if b == nil {
		t.Fatalf("FAILED!")
	}
}

func TestRemoveElemUser(t *testing.T) {

	a := []Test{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}}

	b := RemoveElem(a, Test{0, 0})

	c := Test{1, 1}
	if b[0] != c {
		t.Fatalf("FAILED!")
	}
}

func TestRemoveElemBuiltin(t *testing.T) {

	a := []int{0, 1, 2, 3, 4}

	b := RemoveElem(a, 0)

	if b[0] != 1 {
		t.Fatalf("FAILED!")
	}

}

func TestRemoveIndexUser(t *testing.T) {

	a := []Test{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}}

	b := RemoveIndex(a, 2)

	c := Test{3, 3}
	if b[2] != c {
		t.Fatalf("FAILED!")
	}
}

func TestRemoveIndexBuiltin(t *testing.T) {

	a := []int{0, 1, 2, 3, 4}

	a = RemoveIndex(a, 2)

	if a == nil || a[0] != 0 || a[1] != 1 || a[2] != 3 || a[3] != 4 {
		t.Fatalf("FAILED!")
	}

	a = RemoveIndex(a, 2)

	if a == nil || a[0] != 0 || a[1] != 1 || a[2] != 4 {
		t.Fatalf("FAILED: %#v", a)
	}

	a = RemoveIndex(a, 1)

	if a == nil || a[0] != 0 || a[1] != 4 {
		t.Fatalf("FAILED: %#v", a)
	}

	a = RemoveIndex(a, 1)

	if a == nil || a[0] != 0 {
		t.Fatalf("FAILED: %#v", a)
	}

	a = RemoveIndex(a, 0)

	if a == nil || len(a) != 0 {
		t.Fatalf("FAILED: %#v", a)
	}
}

func BenchmarkRemoveIndexBuiltin(b *testing.B) {

	a := []int{0, 1, 2, 3, 4}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RemoveIndex(a, 2)
	}

}

func TestJoinUser(t *testing.T) {

	a := []Test{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}}

	b := Join(a, " ")

	if b != "0|0 1|1 2|2 3|3 4|4" {
		t.Fatalf("FAILED!")
	}
}

func TestContainsUser(t *testing.T) {

	a := []Test{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}}

	if !Contains(a, Test{0, 0}) {
		t.Fatalf("FAILED!")
	}
}

func TestContainsBuiltin(t *testing.T) {

	a := []int{0, 1, 2, 3, 4}

	if !Contains(a, 0) {
		t.Fatalf("FAILED!")
	}
}

func TestContainsNUser(t *testing.T) {

	a := []Test{{0, 0}, {0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}}

	if ContainsN(a, Test{0, 0}) != 2 {
		t.Fatalf("FAILED!")
	}
}

func TestContainsNBuiltin(t *testing.T) {

	a := []int{0, 0, 1, 2, 3, 4}

	if ContainsN(a, 0) != 2 {
		t.Fatalf("FAILED!")
	}
}

func TestStringsUser(t *testing.T) {

	a := []Test{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}}

	b := Strings(a)

	if b[0] != "0|0" {
		t.Fatalf("FAILED!")
	}
}

func TestAppendUniqueUser0(t *testing.T) {

	a := []Test{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}}
	b := AppendUnique(a, Test{i: 0, f: 0})

	if len(b) != len(a) {
		t.Fatalf("FAILED!")
	}
}

// Test when s is nil
func TestAppendUniqueUser1(t *testing.T) {

	var c []Test
	d := AppendUnique(c, Test{i: 0, f: 0})

	if len(c) != 0 || len(d) != 1 {
		t.Fatalf("FAILED!")
	}
}

func TestAppendUniqueBuiltin0(t *testing.T) {

	a := []int{0, 1, 2, 3, 4}
	b := AppendUnique(a, 2)

	if len(b) != len(a) {
		t.Fatalf("FAILED!")
	}
}

// Test when s is nil with a builtin type.
func TestAppendUniqueBuiltin1(t *testing.T) {

	var c []int
	d := AppendUnique(c, 1)

	if len(c) != 0 || len(d) != 1 {
		t.Fatalf("FAILED!")
	}
}

func TestAppendUniquesBuiltin0(t *testing.T) {

	a := []int{0, 1, 2, 3, 4}
	b := AppendUniques(a, []int{0, 1}...)

	if len(b) != len(a) {
		t.Fatalf("FAILED!")
	}
}

// Test when s is nil with a builtin type.
func TestAppendUniquesBuiltin1(t *testing.T) {

	var a []int
	b := AppendUniques(a, []int{0, 1}...)

	if len(a) != 0 || len(b) != 2 {
		t.Fatalf("FAILED!")
	}
}

func TestAppendUniquesUser0(t *testing.T) {

	a := []Test{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}}
	b := AppendUniques(a, []Test{{0, 0}}...)

	if len(b) != len(a) {
		t.Fatalf("FAILED!")
	}
}

// Test when s is nil
func TestAppendUniquesUser1(t *testing.T) {

	var a []Test
	b := AppendUniques(a, Test{i: 0, f: 0})

	if len(a) != 0 || len(b) != 1 {
		t.Fatalf("FAILED!")
	}
}

func TestContainsDuplicateUser(t *testing.T) {

	a := []Test{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}}

	if ContainsDuplicate(a) {
		t.Fatalf("FAILED!")
	}

	a = []Test{{0, 0}, {0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}}

	if !ContainsDuplicate(a) {
		t.Fatalf("FAILED!")
	}
}

func TestContainsDuplicateBuiltin(t *testing.T) {

	a := []int{0, 1, 2, 3, 4}

	if ContainsDuplicate(a) {
		t.Fatalf("FAILED!")
	}

	a = []int{0, 0, 1, 2, 3, 4}

	if !ContainsDuplicate(a) {
		t.Fatalf("FAILED!")
	}
}

func TestRemoveDuplicatesUser(t *testing.T) {

	a := []Test{{0, 0}, {0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}}

	a = RemoveDuplicates(a)

	if ContainsDuplicate(a) {
		t.Fatalf("FAILED!")
	}
}

func TestRemoveDuplicatesBuiltin(t *testing.T) {

	a := []int{0, 0, 1, 2, 3, 4}

	a = RemoveDuplicates(a)

	if ContainsDuplicate(a) {
		t.Fatalf("FAILED!")
	}
}
