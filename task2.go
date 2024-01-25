package main

type DeretBilangan struct {
	N int
}

func (D DeretBilangan) prima() []int {
	results := []int{}

	for i := 2; i <= D.N; i++ {
		bilPrima := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				bilPrima = false
				break
			}
		}
		if bilPrima {
			results = append(results, i)
		}
	}
	return results
}

func (D DeretBilangan) ganjil() []int {
	results := []int{}
	for i := 1; i <= D.N; i++ {
		if i%2 != 0 {
			results = append(results, i)
		}
	}
	return results
}

func (D DeretBilangan) genap() []int {
	results := []int{}

	for i := 1; i <= D.N; i++ {
		if i%2 == 0 {
			results = append(results, i)
		}
	}

	return results
}

func (D DeretBilangan) fibonacci() []int {
	results := []int{0, 1}

	f0 := 0
	f1 := 1
	for i := 2; i <= D.N; i++ {
		f2 := f0 + f1
		f0 = f1
		f1 = f2
		if f2 > 40 {
			break
		} else {
			results = append(results, f2)
		}
	}

	return results
}
