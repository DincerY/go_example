package main

import "fmt"

func main() {

	fmt.Println("Program başlatıldı")

	// Hatanın yakalanıp yakalanmadığını kontrol etmek için flag
	success := true

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("❌ Panik yakalandı:", r)
			success = false // Hata olduysa başarıyla bitmedi
		}

		if success {
			fmt.Println("✅ Program başarıyla tamamlandı")
		} else {
			fmt.Println("⚠️ Program hata ile sonlandırıldı")
		}
	}()

	numbers := []int{10, 5, 0, 2}
	for _, num := range numbers {
		result := divide(100, num)
		fmt.Println("Sonuç:", result)
	}
}

func divide(a, b int) int {
	if b == 0 {
		panic("Sıfıra bölme hatası!") // Kritik hata
	}
	return a / b
}
