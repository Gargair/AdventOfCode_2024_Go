package intmath

func Abs(x *int) int {
	y := 0
	return AbsDiff(x, &y)
}

func AbsDiff(x, y *int) int {
	if *x < *y {
		return *y - *x
	}
	return *x - *y
}
