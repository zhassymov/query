package query

// Single returns the single element of the slice.
// If the slice does not contain exactly one element, Single returns
// the zero value for E and false.
//
// Example:
//
//	s := []int{10}
//	val, ok := Single(s) // val is 10, ok is true
//
//	s2 := []int{10, 20}
//	val2, ok2 := Single(s2) // val2 is 0, ok2 is false
//
//	s3 := []int{}
//	val3, ok3 := Single(s3) // val3 is 0, ok3 is false
func Single[S ~[]E, E any](s S) (E, bool) {
	if len(s) == 1 {
		return s[0], true
	}
	var zero E
	return zero, false
}

// SingleFunc returns the single element of the slice that satisfies the predicate f.
// If exactly one element satisfies the predicate, that element is returned along with true.
// Otherwise (if zero or more than one element satisfy the predicate),
// SingleFunc returns the zero value for E and false.
//
// Example:
//
//	s := []int{1, 2, 3, 4, 5}
//	val, ok := SingleFunc(s, func(x int) bool { return x == 3 }) // val is 3, ok is true
//
//	val2, ok2 := SingleFunc(s, func(x int) bool { return x > 5 }) // val2 is 0, ok2 is false
//
//	val3, ok3 := SingleFunc(s, func(x int) bool { return x%2 == 1 }) // val3 is 0, ok3 is false (multiple odd numbers)
func SingleFunc[S ~[]E, E any](s S, f func(E) bool) (E, bool) {
	var found E
	count := 0
	for i := range s {
		if f(s[i]) {
			found = s[i]
			count++
		}
	}
	if count == 1 {
		return found, true
	}
	var zero E
	return zero, false
}
