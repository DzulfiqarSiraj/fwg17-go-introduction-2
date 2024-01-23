package main

func Pembulatan(n float32) float32 {
	bilBulat := int(n)
	bilDesimal := n - float32(bilBulat)
	persepuluh := int(bilDesimal * 10)
	perseratus := int(bilDesimal*100) - (persepuluh * 10)

	// fmt.Printf("Bilangan inputan = %v %v\n", n, reflect.TypeOf(n))
	// fmt.Printf("Bilangan desimal = %v %v\n", bilDesimal, reflect.TypeOf(bilDesimal))
	// fmt.Printf("Bilangan bulat = %v %v\n", bilBulat, reflect.TypeOf(bilBulat))
	// fmt.Printf("Bilangan desimal persepuluhan = %v %v\n", persepuluh, reflect.TypeOf(persepuluh))
	// fmt.Printf("Bilangan desimal perseratus = %v %v\n", perseratus, reflect.TypeOf(perseratus))

	bilBulatResult := float32(bilBulat)
	persepuluhResult := float32(persepuluh) * 0.1
	perseratusResult := float32(perseratus) * 0.01

	if perseratus > 5 && persepuluh == 9 {
		bilBulatResult = float32(bilBulat + 1)
		persepuluhResult = 0.0
		perseratusResult = 0.00
	} else if perseratus > 5 && persepuluh < 9 {
		persepuluhResult = float32(persepuluh+1) * 0.1
		perseratusResult = 0.00
	} else if perseratus < 5 {
		perseratusResult = 0.00
	}

	result := bilBulatResult + persepuluhResult + perseratusResult

	// fmt.Printf("Bilangan bulat result = %v %v\n", bilBulatResult, reflect.TypeOf(bilBulatResult))
	// fmt.Printf("Bilangan desimal persepuluhan result = %v %v\n", persepuluhResult, reflect.TypeOf(persepuluhResult))
	// fmt.Printf("Bilangan desimal perseratus result = %v %v\n", perseratusResult, reflect.TypeOf(perseratusResult))

	// fmt.Printf("hasil = %v %v\n", result, reflect.TypeOf(result))
	return result
}
