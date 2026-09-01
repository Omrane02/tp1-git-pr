package main
import "fmt"

func Add(a, b int) int { return a + b }

func Multiply(a, b int) int { return a * b }
func Divide(a, b int) int { return a / b }
func main() {
	fmt.Println("Add(2, 3) =", Add(2, 3))
	fmt.Println("Divide(10,0) =", Divide(10,0))
}
