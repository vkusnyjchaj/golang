package main

// Модуль числа вычисляется условием принципиально, чтобы не подключать пакет math

import (
	"fmt"
)

func GetDiff(x float64, z float64) float64 {
	diff := (z*z - x) / (2 * z)
	return diff
}

func Sqrt(x float64) float64 {
	// корень числа
	z := 1.0
	// Допустимая погрешность
	a := 0.000000001
	// Разница между предыдущим и текущим значениями корня
	diff := GetDiff(x, z)

	if diff < 0 {
		diff = -diff
	}

	for diff > a {
		diff = GetDiff(x, z)
		z -= diff

		if diff < 0 {
			diff = -diff
		}
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
