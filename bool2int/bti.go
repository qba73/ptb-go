package bti

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func BoolToInt2(b bool) int {
	r := 0
	if b {
		r = 1
	}
	return r
}

func BoolToInt3(b bool) int {
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return i
}
