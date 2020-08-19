package compare

func Max2(lhs, rhs float64) float64 {
	if lhs > rhs {
		return lhs
	}

	return rhs
}

func Max3(a, b, c float64) float64 {
	rhs := Max2(b, c)

	return Max2(a, rhs)
}

func Max4(a, b, c, d float64) float64 {
	rhs := Max3(b, c, d)

	return Max2(a, rhs)
}
