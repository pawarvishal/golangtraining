package greeter

func Makegreeter() func() string {
	return func() string {
		return "Hello World"
	}
}