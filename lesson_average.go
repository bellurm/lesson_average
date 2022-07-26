package main

import "fmt"

func main() {
	var vize float64
	var final float64

	fmt.Print("Vize notunuz: ")
	fmt.Scanln(&vize)

	if vize <= 100 && vize >= 0 {
		fmt.Print("Final notunuz: ")
		fmt.Scanln(&final)

		if final <= 100 && final >= 0 {
			ortalama := vize*0.40 + final*0.60

			if ortalama < 50 {
				fmt.Printf("Kaldınız. Ortalama: %v", ortalama)
			} else if ortalama > 50 && ortalama < 55 {
				fmt.Printf("Şartlı Geçtiniz. Ortalama: %v", ortalama)
			} else {
				fmt.Printf("Geçtiniz. Ortalama: %v", ortalama)
			}
		} else {
			fmt.Println("Lütfen geçerli bir final notu giriniz.")
		}

	} else {
		fmt.Println("Lütfen geçerli bir vize notu giriniz.")
	}

}
