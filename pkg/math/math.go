package math

// GCD Greatest Common Divison via euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM Less Common Multiple via GCD
func LCM(first int, integers []int) int {
	result := first * integers[0] / GCD(first, integers[0])
	for i := 1; i < len(integers); i++ {
		result = LCM(result, []int{integers[i]})
	}
	return result
}

/*
	2x2 System Linear Equations
	a * X + b * Y = c
	d * X + e * Y = f
*/
// SystemLinearEq2x2
func SystemLinearEq2x2(matrix [2][3]int) (float64, float64) {
	c := matrix[0][2]
	f := matrix[1][2]
	a := matrix[0][0]
	d := matrix[1][0]
	b := matrix[0][1]
	e := matrix[1][1]
	x := float64(c*e-f*b) / float64(a*e-d*b)
	y := float64(f*a-c*d) / float64(a*e-d*b)
	return x, y
}
