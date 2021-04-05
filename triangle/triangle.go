//determine the type of triangle
package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = 1 // not a triangle
	Equ Kind = 2 // equilateral
	Iso Kind = 3 // isosceles
	Sca Kind = 4 // scalene
)

// return the type of traingle as type Kind
func KindFromSides(a, b, c float64) Kind {

	if a <= 0.0 || b <= 0.0 || c <= 0.0 || (a+b) < c || (a+c) < b || (b+c) < a {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || a == c || b == c {
		return Iso
	}

	return Sca
}
