package util

func StringsContains(s []string, target string) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}

func IntsContains(s []int, target int) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}

func Int64sContains(s []int64, target int64) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}
