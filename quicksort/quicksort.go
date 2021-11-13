package quicksort

func QuickSort(a []int) {
	qSort(a, 0, len(a)-1)
}

func qSort(a []int, l, r int) {
	i, j := l, r
	p := a[(l+r)/2]

	for i < j {
		for a[i] < p {
			i++
		}
		for a[j] > p {
			j--
		}
		if i <= j {
			a[i], a[j] = a[j], a[i]
			i++
			j--
		}
	}

	if l < j {
		qSort(a, l, j)
	}
	if i < r {
		qSort(a, i, r)
	}
}
