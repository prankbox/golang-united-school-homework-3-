package homework

func reverse(input []int64) (result []int64) {
	inputLen := len(input)
	output := make([]int, inputLen)

	for i, n := range input {
		j := inputLen - i - 1

		output[j] = n
	}

	return output
}
