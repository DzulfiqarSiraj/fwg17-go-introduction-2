package main

import "fmt"

func main() {
	// Pembulatan persepuluh
	var angka1 float32 = 4.37
	var angka2 float32 = 4.32
	var angka3 float32 = 4.324
	fmt.Println("Pembulatan angka")
	fmt.Printf("%v -> %v\n", angka1, Pembulatan(angka1))
	fmt.Printf("%v -> %v\n", angka2, Pembulatan(angka2))
	fmt.Printf("%v -> %v\n", angka3, Pembulatan(angka3))
	fmt.Println("============")

	// Deret bilangan
	D := DeretBilangan{40}
	fmt.Print()
	fmt.Printf("Bilangan Prima : %v\n", D.prima())
	fmt.Printf("Bilangan Ganjil : %v\n", D.ganjil())
	fmt.Printf("Bilangan Genap : %v\n", D.genap())
	fmt.Printf("Bilangan Fibonacci : %v\n", D.fibonacci())
	fmt.Println("============")

	// Hitung Persegi
	p := &Persegi{}
	p.Sisi = 5
	// Luas
	fmt.Printf("Luas persegi dengan panjang sisi %v = %v\n", p.Sisi, p.Luas())
	// Keliling
	fmt.Printf("Keliling persegi dengan panjang sisi %v = %v\n", p.Sisi, p.Keliling())

	// Hitung Kubus
	k := &Kubus{}
	k.Sisi = 6
	fmt.Printf("Volume kubus dengan panjang sisi %v = %v\n", k.Sisi, k.Volume())

	// Hitung Lingkaran
	l := &Lingkaran{}
	l.Jari = 5
	// Luas
	fmt.Printf("Luas lingkaran dengan jari-jari %v = %v\n", l.Jari, l.Luas())

}
