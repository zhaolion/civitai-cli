package api

func containsInt(list []int, e int) bool {
	for _, a := range list {
		if a == e {
			return true
		}
	}

	return false
}
