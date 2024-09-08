package main

import (
	"fmt"
)

func main() {
	var x, y float64
	x = 10
	y = 5

	// Penjumlahan
	sum := x + y
	fmt.Printf("Penjumlahan: %.2f + %.2f = %.2f\n", x, y, sum)

	// Pengurangan
	sub := x - y
	fmt.Printf("Pengurangan: %.2f - %.2f = %.2f\n", x, y, sub)

	// Perkalian
	mul := x * y
	fmt.Printf("Perkalian: %.2f * %.2f = %.2f\n", x, y, mul)

	// Pembagian
	div := x / y
	fmt.Printf("Pembagian: %.2f / %.2f = %.2f\n", x, y, div)

	// Modulus (sisa hasil bagi)
	mod := int(x) % int(y)
	fmt.Printf("Modulus: %d %% %d = %d\n", int(x), int(y), mod)
}
