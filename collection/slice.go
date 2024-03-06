package collection

import "fmt"

//-----------------------------------------------------------------------------
// Commonly used sliced tricks

// Remove the element at the specified index while preserving the element order.
// NOTE: The underlying array is modified.
// This function does not check that the index is within bounds.
func SliceRemoveAt[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}

// Remove the element at the specified index without preserving the element order.
// The last element in the slice is copied to s[index].
// NOTE: The underlying array is modified.
// This function does not check that the index is within bounds.
func SliceRemoveAtFast[T any](s []T, index int) []T {
	s[index] = s[len(s)-1]
	return s[:len(s)-1]
}

// Remove the element at the specified index while preserving the element order.
// NOTE: The underlying array is modified.
// An error will be returned if the index is out of bounds.
func SliceSafeRemoveAt[T any](s []T, index int) ([]T, error) {
	if (index < 0) || (index >= len(s)) {
		return nil, fmt.Errorf("invalid index %d", index)
	}
	return SliceRemoveAt(s, index), nil
}

// Remove the element at the specified index without preserving the element order.
// NOTE: The underlying array is modified.
// An error will be returned if the index is out of bounds.
func SliceSafeRemoveAtFast[T any](s []T, index int) ([]T, error) {
	if (index < 0) || (index >= len(s)) {
		return nil, fmt.Errorf("invalid index %d", index)
	}
	return SliceRemoveAtFast(s, index), nil
}

//-----------------------------------------------------------------------------
// Common slice/array operations

// The following functions are now being provided by the std lib `slices` package.
// slices.Contains, slices.Index

// // Return true if the slice contains the needle.
// // This does a linear search for the needle.
// func SliceContains[T comparable](s []T, needle T) bool {
// 	found := SliceFind[T](s, needle)
// 	return found != -1
// }

// // Look for the needle in the slice and return the index at which it is found.
// // This does a linear search for the needle.
// // Returns -1 if the needle could not be found.
// func SliceFind[T comparable](s []T, needle T) int {
// 	for i, val := range s {
// 		if val == needle {
// 			return i
// 		}
// 	}
// 	return -1
// }
