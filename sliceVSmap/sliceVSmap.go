package slicevsmap

// search returns an index of element `e` in slice `s`.
// s must be sorted slice to execute binary search.
func search(e int, s []int) int {
	idx := binarySearch(e, 0, len(s)-1, s)
	if idx < 0 {
		panic("not found")
	}
	return idx
}

func binarySearch(e, l, r int, s []int) int {
	i := (l + r) / 2

	if i == l {
		switch {
		case s[l] == e:
			return l
		case s[r] == e:
			return r
		default:
			return -1
		}
	}

	switch {
	case s[i] < e:
		return binarySearch(e, i+1, r, s)
	case s[i] > e:
		return binarySearch(e, l, i-1, s)
	case s[i] == i:
		return i
	}

	return -1
}
