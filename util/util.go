package util

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ContainsLessThan(s []int, e int) bool {
	for _, a := range s {
		if a < e {
			return true
		}
	}
	return false
}
