package main
import "fmt"

func main() {
    // If-else statement
    x := 10
    y := 20

    if x < y {
        fmt.Println("x is less than y")
    } else {
        fmt.Println("y is less than x")
    }

    // If-else if-else statement
    age := 18

    if age < 18 {
        fmt.Println("You are a minor")
    }
    if age >= 18 && age < 60 {
        fmt.Println("You are an adult")
    } else {
        fmt.Println("You are a senior citizen")
    }
}

// Output
// x is less than y
// You are an adult
