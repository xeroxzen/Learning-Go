package main
import "fmt"

func main() {
    // Call the function
    fmt.Println(add(3, 2))
    fmt.Println(sub(4, 2))
    fmt.Println(mul(7, 2))
    fmt.Println(div(4, 2))

}

func add(x int, y int) int {
    return x + y
}

func sub(x int, y int) int {
    return x - y
}

func mul(x int, y int) int {
    return x * y
}

func div(x int, y int) int {
    return x / y
}
