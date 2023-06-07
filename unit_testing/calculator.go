package unit_testing

type Calculator struct {
}

//type calculator interface {
//	Multiply(a, b int) int
//}

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func (c Calculator) Multiply(a, b int) int {
	return a * b
}
