package main

// split array into two part, sort one by one, and merge together
func mergesort(a []int, begin, end int) {
	if begin+1 == end {
		return
	}
	if begin+2 == end {
		if a[begin] > a[end-1] {
			a[begin], a[end-1] = a[end-1], a[begin]
		}
		return
	}
	// split into two part, and sort
	h := (end + begin - 1) / 2
	mergesort(a, begin, h)
	mergesort(a, h, end)
	// merge two part together
	i, j := begin, h
	newarr := []int{}
	for i < h && j < end {
		if a[i] < a[j] {
			newarr = append(newarr, a[i])
			i++
		} else {
			newarr = append(newarr, a[j])
			j++
		}
	}
	newarr = append(newarr, a[i:h]...)
	newarr = append(newarr, a[j:end]...)
	for index, v := range newarr {
		a[index+begin] = v
	}
}
