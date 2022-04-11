package recursion

type Factorial struct {
	value map[int]int
}

func NewFactorial(n int) *Factorial {
	return &Factorial{
		value: make(map[int]int, n),
	}
}

func (f *Factorial) Fac(n int) int {
	if f.value[n] != 0 {
		return f.value[n]
	}

	if n <= 1 {
		f.value[n] = 1
		return 1
	} else {
		res := n * f.Fac(n-1)
		f.value[n] = res
		return res
	}

}

func (f *Factorial) Print(n int) {
	println(f.value[n])
}
