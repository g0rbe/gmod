package slices

import (
	"fmt"
)

// Copy create a copy of s and returns it.
// If copy fail, returns nil.
func Copy[T any](s []T) []T {

	n := len(s)

	v := make([]T, n)

	if copy(v, s) != n {
		return nil
	}

	return v
}

// RemoveElem removes element e from s.
func RemoveElem[T comparable](s []T, e T) []T {

	var v []T

	for i := range s {

		if s[i] == e {
			continue
		}

		v = append(v, s[i])
	}

	return v
}

// RemoveIndex removes element with index i from s.
// Returns nil, s is nil or empty, or i is not in range.
func RemoveIndex[T any](s []T, i int) []T {

	ls := len(s)

	if s == nil || ls == 0 || i < 0 || i > ls-1 {
		return nil
	}

	// Remove the first elem
	if i == 0 {
		return s[i+1:]
	}

	// Remove the last elem
	if i == ls-1 {
		return s[:i]
	}

	//v := make([]T, 0, ls-1)

	v := s[:i]

	return append(v, s[i+1:]...)
}

// Join join elements of s with sep and return as a string.
// If s is nil or zero-length, returns an empty string ("").
func Join[T fmt.Stringer](s []T, sep string) string {

	n := len(s)
	var v string

	for i := range s {

		v += s[i].String()

		if i != n-1 {
			v += sep
		}
	}

	return v
}

// Contains returns true if s contains e.
func Contains[T comparable](s []T, e T) bool {

	for i := range s {
		if s[i] == e {
			return true
		}
	}

	return false
}

// ContainsN returns the number of same element e in s.
func ContainsN[T comparable](s []T, e T) int {

	n := 0

	for i := range s {
		if s[i] == e {
			n++
		}
	}

	return n
}

// Strings convert a slice of T to a string slice.
// If s is nil, returns nil.
func Strings[T fmt.Stringer](s []T) []string {

	var v []string

	for i := range s {
		v = append(v, s[i].String())
	}

	return v
}

// AppendUnique appends e to s if e is not in s.
// If e is in s, than the returned slice is equal to s, i.e. return the same slice.
func AppendUnique[T comparable](s []T, e T) []T {

	if Contains(s, e) {
		return s
	}

	return append(s, e)
}

// AppendUniques appends the elements of e to s if element of e is not in s.
// If e is nil or empty, returns the same slice.
func AppendUniques[T comparable](s []T, e ...T) []T {

	for i := range e {
		s = AppendUnique(s, e[i])
	}

	return s
}

// ContainsDuplicate returns whether slice s contains any duplicate element.
func ContainsDuplicate[T comparable](s []T) bool {

	for i := range s {
		if ContainsN(s, s[i]) != 1 {
			return true
		}
	}

	return false
}

// RemoveDuplicates removes the duplicate elements from s.
// The returned slice contains only unique elements.
func RemoveDuplicates[T comparable](s []T) []T {

	v := make([]T, 0)

	for i := range s {
		if !Contains(v, s[i]) {
			v = append(v, s[i])
		}
	}

	return v
}
