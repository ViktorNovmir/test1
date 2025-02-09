package main

// Importing fmt
import (
	"fmt"
)

// Calling main
func main() {
	b := 1
	for ; b < 5; b++ {
		fmt.Println(b)
	}

	var n [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(n)

	var n1 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(n1)

	n2 := [3]string{0: "van", 1: "tu", 2: "fri"}
	fmt.Println(n2[2])

	switch b {
	case 3:
		otlik()
	case 4:
		add(5, 5)
	case 5:
		get(2, 4)
	default:
		fmt.Println("Значение небыло введено")
	}

}
func otlik() {
	fmt.Println("Отклик функции.")
}
func add(x int, y int) {
	var z = x + y
	fmt.Println("x + y =", z)
	fmt.Println("x = ", x, "\n ", "y =", y)

}

func get(a int, s int) {
	if a < s {
		fmt.Println("первое число меньше второго")
	} else {
		fmt.Println("первое число больше второго")
	}

}
