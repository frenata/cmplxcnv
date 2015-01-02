package cmplxcnv

import (
	"testing"
)

var cmplx = []complex128{0i, 011i, 0.i, 2.71828, 1E6i, 1.e+0i, 0, 3.0 + 2i, 2.7e-4 + 2i, 7.7e+4 + 0.9i, 3i + 4.0, 2.1i + 4.55e+10}
var str = []string{"0i", "011i", "0.i", "2.71828", "1E6i", "1.e+0i", "0", "3+2i", "2.7e-4+2i", "7.7e+4+0.9i", "3i+4.0", "2.1i+4.55e+10"}

func Test_ParseComplex(t *testing.T) {
	for i := range str {
		c, _ := ParseComplex(str[i], 64)
		if c != cmplx[i] {
			t.Fatalf("%v did not parse correctly to %v, instead: %v", str[i], cmplx[i], c)
		}
	}
}
