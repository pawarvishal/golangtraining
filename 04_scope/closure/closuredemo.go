package closure

func Wrapper() func() int {
	x := 0 //value of x is remembered by closure

	return func() int {
		x++
		return x
	}
}
