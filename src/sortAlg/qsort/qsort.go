package qsort

func qsort(values []int, left int, right int) {
	if left < right {
		l, r, mark := left, right, values[left]
		for l < r {
			for l < r && values[r] >= mark {
				r --
			}
			if l < r {
				values[l] = values[r]
				l++
			}
			for l < r && values[l] < mark {
				l ++
			}
			if l < r {
				values[r] = values[l]
				r--
			}
		}
		values[l] = mark
		qsort(values, 1, l-1)
		qsort(values, l+1, r)
	}
}

func Qsort(values []int)  {
	qsort(values, 0, len(values) - 1)
}
