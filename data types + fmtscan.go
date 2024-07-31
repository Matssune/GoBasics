package main

import (
	"fmt"
)

func main() {
	// Оголошення змінних з різними типами даних
	var int1 int = 727
	var int2 int // значення за замовчуванням 0
	var int3 = 727

	var float1 float32 = 72.7
	var float2 float32 // значення за замовчуванням 0.0
	var float3 = 7.27

	var str1 string = "Foreground Eclipse"
	var str2 string // значення за замовчуванням порожній рядок ""
	var str3 = "Is truly amazing band"

	var bool1 bool = true
	var bool2 bool // значення за замовчуванням false
	var bool3 = false

	// Запрошення вибрати тип даних
	fmt.Println("Оберіть тип даних:")
	fmt.Println("1 для int")
	fmt.Println("2 для float32")
	fmt.Println("3 для string")
	fmt.Println("4 для bool")

	var choice int
	fmt.Scan(&choice)

	// Відображення значень змінних залежно від вибору користувача
	switch choice {
	case 1:
		fmt.Println("Змінні типу int:")
		fmt.Println("int1:", int1, "- задано значення 727")
		fmt.Println("int2:", int2, "- значення за замовчуванням 0")
		fmt.Println("int3:", int3, "- задано значення без указання типу змінної 727")
	case 2:
		fmt.Println("Змінні типу float32:")
		fmt.Println("float1:", float1, "- задано значення 72.7")
		fmt.Println("float2:", float2, "- значення за замовчуванням 0.0")
		fmt.Println("float3:", float3, "- задано значення без указання типу змінної  7.27")
	case 3:
		fmt.Println("Змінні типу string:")
		fmt.Println("str1:", str1, "- задано значення 'Foreground Eclipse'")
		fmt.Println("str2:", str2, "- значення за замовчуванням порожній рядок ''")
		fmt.Println("str3:", str3, "- задано значення без указання типу змінної 'Is truly amazing band'")
	case 4:
		fmt.Println("Змінні типу bool:")
		fmt.Println("bool1:", bool1, "- задано значення true")
		fmt.Println("bool2:", bool2, "- значення за замовчуванням false")
		fmt.Println("bool3:", bool3, "- задано значення без указання типу змінної false")
	default:
		fmt.Println("Nope.")
	}
}
