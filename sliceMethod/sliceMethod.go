package slicemethod

func IsInList[T comparable](list []T, item T) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == item {
			return true
		}
	}
	return false
}

func SliceFilter[T comparable](slice []T, compareFunction func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if compareFunction(v) {
			result = append(result, v)
		}
	}

	return result
}
