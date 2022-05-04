package slices

func Delete[T comparable](s []T, id T) {
	j := 0
	for _, val := range s {
		if val != id {
			s[j] = val
			j++
		} else {
			break
		}
	}
}

func Find[T comparable](s []T, id T) (i int) {
	var val T
	for i, val = range s {
		if val == id {
			return i
		}
	}
	return -1
}
