package main
import "fmt"

func main() {
    // variable declaration
    var x int = 10
    var y int = 20
    var sum int = x + y
    fmt.Println("Sum of x and y is: ", sum)

    // variable declaration with type inference
    var a = 10
    var b = 20
    var add = a + b
    fmt.Println("Addition of a and b is: ", add)

    // variable declaration with shorthand
    c := 10
    d := 20
    e := c + d
    fmt.Println("Sum of c and d is: ", e)

    // multiple variable declaration
    var (
        f = 10
        g = 20
        h = f + g
    )
    fmt.Println("Sum of f and g is: ", h)
}
