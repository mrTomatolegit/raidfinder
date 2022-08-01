package util

func Contains[T string | int | struct{}](source *[]T, element T) bool {
	for _, v := range *source {
		if v == element {
			return true
		}
	}
	return false
}

func Filter[T any](source *[]T, fn func(T, int) bool) (removed []T) {
	newarr := []T{}
	for i, v := range *source {
		if fn(v, i) {
			newarr = append(newarr, v)
		}
	}
	return newarr
}

func Splice[T any](source *[]T, start int, delete int, item ...T) (removed []T) {
	if start > len(*source) {
		start = len(*source)
	}
	if start < 0 {
		start = len(*source) + start
	}
	if start < 0 {
		start = 0
	}
	if delete < 0 {
		delete = 0
	}
	if delete > 0 {
		for i := 0; i < delete; i++ {
			if i+start < len(*source) {
				removed = append(removed, (*source)[i+start])
			}
		}
	}
	delete = len(removed) // Adjust to actual delete count
	grow := len(item) - delete
	switch {
	case grow > 0: // So we grow
		*source = append(*source, make([]T, grow)...)
		copy((*source)[start+delete+grow:], (*source)[start+delete:])
	case grow < 0: // So we shrink
		from := start + len(item)
		to := start + delete
		copy((*source)[from:], (*source)[to:])
		*source = (*source)[:len(*source)+grow]
	}
	copy((*source)[start:], item)
	return
}
