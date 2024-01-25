package main

import "fmt"

func main() {
	// Pembulatan persepuluh
	var angka1 float32 = 4.37
	var angka2 float32 = 4.32
	var angka3 float32 = 4.324
	fmt.Println(Pembulatan(angka1))
	fmt.Println(Pembulatan(angka2))
	fmt.Println(Pembulatan(angka3))
	fmt.Println("============")

	// Deret bilangan
	D := DeretBilangan{40}
	D.ganjil()
	D.genap()
	D.fibonacci()
}
