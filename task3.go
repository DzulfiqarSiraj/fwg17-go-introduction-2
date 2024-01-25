package main

type Hitung2d interface {
	Luas() float64
	Keliling() float64
}

type Hitung3d interface {
	Volume() float64
}

type Hitung interface {
	Hitung2d
	Hitung3d
}

type Persegi struct {
	Sisi float64
}

type Kubus struct {
	Sisi float64
}

type Lingkaran struct {
	Jari float64
}

func (P *Persegi) Luas() float64 {
	return P.Sisi * P.Sisi
}

func (P *Persegi) Keliling() float64 {
	return 4 * P.Sisi
}

func (K *Kubus) Volume() float64 {
	return K.Sisi * K.Sisi * K.Sisi
}

func (L *Lingkaran) Luas() float64 {
	return 3.14 * (L.Jari * L.Jari)
}
