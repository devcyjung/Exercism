package complexnumbers

import "math/cmplx"

type Number struct {
    r	float64
    i	float64
}

func ToNumber(c complex128) Number {
    return Number{
        r:	real(c),
        i:	imag(c),
    }
}

func (n Number) ToCmplx() complex128 {
    return complex(n.r, n.i)
}

func (n Number) Real() float64 {
	return real(n.ToCmplx())
}

func (n Number) Imaginary() float64 {
	return imag(n.ToCmplx())
}

func (n1 Number) Add(n2 Number) Number {
	return ToNumber(n1.ToCmplx() + n2.ToCmplx())
}

func (n1 Number) Subtract(n2 Number) Number {
	return ToNumber(n1.ToCmplx() - n2.ToCmplx())
}

func (n1 Number) Multiply(n2 Number) Number {
	return ToNumber(n1.ToCmplx() * n2.ToCmplx())
}

func (n Number) Times(factor float64) Number {
	return ToNumber(n.ToCmplx() * complex(factor, 0))
}

func (n1 Number) Divide(n2 Number) Number {
	return ToNumber(n1.ToCmplx() / n2.ToCmplx())
}

func (n Number) Conjugate() Number {
	return ToNumber(cmplx.Conj(n.ToCmplx()))
}

func (n Number) Abs() float64 {
	return cmplx.Abs(n.ToCmplx())
}

func (n Number) Exp() Number {
	return ToNumber(cmplx.Exp(n.ToCmplx()))
}
