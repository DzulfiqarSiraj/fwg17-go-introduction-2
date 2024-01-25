package main

import "fmt"

type DeretBilangan struct {
	N int
}

func (D DeretBilangan) ganjil() {
	for i := 1; i <= D.N; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}

func (D DeretBilangan) genap() {
	for i := 1; i <= D.N; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func (D DeretBilangan) fibonacci() {
	f0 := 0
	f1 := 1
	fmt.Println(f0)
	fmt.Println(f1)
	for i := 2; i <= D.N; i++ {
		f2 := f0 + f1
		f0 = f1
		f1 = f2
		fmt.Println(f2)
	}
}
