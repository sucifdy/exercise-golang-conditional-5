package main

import "fmt"

func TicketPlayground(height, age int) int {
	// Memeriksa jika umur di bawah 5 tahun
	if age < 5 {
		return -1 // Tidak dapat membeli tiket
	}

	// Menentukan tarif berdasarkan umur dan tinggi
	if age > 12 {
		return 100000 // Tiket khusus Remaja
	} else if age >= 12 || height > 160 {
		return 60000
	} else if age >= 10 || height > 150 {
		return 40000
	} else if age >= 8 || height > 135 {
		return 25000
	} else if age >= 5 || height > 120 {
		return 15000
	}

	return -1 // Jika tidak ada kondisi yang cocok
}

// Debug
func main() {
	fmt.Println(TicketPlayground(160, 11)) // Output: 40000
	fmt.Println(TicketPlayground(165, 10)) // Output: 60000
	fmt.Println(TicketPlayground(140, 13)) // Output: 100000
	fmt.Println(TicketPlayground(85, 4))   // Output: -1
}
