package main 
import "fmt"

func main() {
    // for loop
    
    fmt.Println("-----------For loop-----------")
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // while loop
    j := 0
    fmt.Println("-----------While loop-----------")
    for j < 5 {
        fmt.Println(j)
        j++
    }
}
