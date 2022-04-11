package recursion

import "testing"

func TestFactorial_Fac(t *testing.T) {
	fac := NewFactorial(10)
	for i := 1; i < 15; i++ {
		fac.Fac(i)
		fac.Print(i)
	}
}
