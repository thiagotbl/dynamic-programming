package fibtab

func Calculate(n uint) uint {
	if n <= 1 {
		return n
	}

	arr := make([]uint, n+1)
	arr[0] = 0
	arr[1] = 1

	for i := uint(0); i < n; i++ {
		incrementAt(arr, arr[i], i+1)
		incrementAt(arr, arr[i], i+2)
	}

	return arr[len(arr)-1]
}

func incrementAt(arr []uint, value uint, index uint) {
	lastIndex := len(arr) - 1

	if (lastIndex < 0) || index > uint(lastIndex) {
		return
	}

	arr[index] += value
}
