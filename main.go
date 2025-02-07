package main

// Importing fmt
import (
	"fmt"
)

// Calling main
func main() {
	otlik()
	add(5, 5)

}
func otlik() {
	fmt.Println("Отклик функции.")
}
func add(x int, y int) {
	var z = x + y
	fmt.Println("x + y =", z)
	fmt.Println("x = ", x, "\n ", "y =", y)
}
