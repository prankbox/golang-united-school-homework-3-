package homework

func reverse(input []int64) (result []int64) {
	values := []string{}
	keys := []int{}
	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		values = append(values, input[k])
	}
	return values
}
