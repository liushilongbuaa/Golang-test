package main

// if root is smaller than child, then swap.
func shiftdown(a []int, root, end int) {
	l, r := 2*root+1, 2*root+2
	if r < end {
		if a[r] > a[l] && a[r] > a[root] {
			a[r], a[root] = a[root], a[r]
			shiftdown(a, r, end)
		} else if a[l] > a[root] {
			a[l], a[root] = a[root], a[l]
			shiftdown(a, l, end)
		}
	} else if l < end {
		if a[l] > a[root] {
			a[l], a[root] = a[root], a[l]
			shiftdown(a, l, end)
		}
	}
	return
}

func heaploop(a []int, end int) {
	if end == 1 {
		return
	}
	a[0], a[end-1] = a[end-1], a[0]
	shiftdown(a, 0, end-1)
	heaploop(a, end-1)
}

// sort a[0:end), example: heapsort(a,len(a))
func heapsort(a []int, end int) {
	// construct MaxHeap
	for j := 0; 2*j+1 < end; j += j + 1 {
		for i := (len(a) - 2) / 2; i >= j; i-- {
			shiftdown(a, i, end)
		}
	}
	// sort
	heaploop(a, end)
}
