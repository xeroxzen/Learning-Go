package main
import "fmt"

func main() {
    // Declare an array of 5 integers that is initialized with some values
    
    var a [5]int = [5]int{1, 2, 3, 4, 5}
    fmt.Println("emp:", a)

    // Set a value at an index
    a[4] = 100
    fmt.Println("set:", a)

    // Get the value at an index
    fmt.Println("get:", a[3])

    // Get the length of the array
    fmt.Println("len:", len(a))

    // Declare a dynamic array
    b := []int{4, 11, 3, 7, 24}
    fmt.Println("dcl:", b)

    // Append to an array
    b = append(b, 6)
    fmt.Println("append:", b)

    // Copy an array
    c := make([]int, len(b))
    copy(c, b)
    fmt.Println("cpy:", c)

    // Slice an array
    d := b[2:5]
    fmt.Println("sl1:", d)

    // Declare a 2D array
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
    
    // Reverse an array
    for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
        a[i], a[j] = a[j], a[i]
    }
    fmt.Println("Reverse array:", a)

    // Add all elements of an array
    sum := 0
    for i := 0; i < len(a); i++ {
        sum += a[i]
    }
    fmt.Println("Sum of array elements:", sum)

    // Add all elements of an array using range
    sum = 0
    for _, num := range a {
        sum += num
    }
    fmt.Println("Sum of array elements:", sum)

    // Sort an array in ascending order
    for i := 0; i < len(b); i++ {
        for j := i + 1; j < len(b); j++ {
            if b[i] > b[j] {
                b[i], b[j] = b[j], b[i]
            }
        }
    }

    fmt.Println("Sorted array:", b)

    // bubbleSort function
    bubbleSort(b)
    fmt.Println("Sorted array:", b)
}

func bubbleSort(arr []int) {
  n := len(arr)
  swapped := true
  for i := 0; i < n-1 && swapped; i++ {
    swapped = false
    for j := 0; j < n-i-1; j++ {
      if arr[j] > arr[j+1] {
        arr[j], arr[j+1] = arr[j+1], arr[j]
        swapped = true
      }
    }
  }
}


// Output
// emp: [1 2 3 4 5]
// set: [1 2 3 4 100]
// get: 4
// len: 5
// dcl: [4 11 3 7 24]
// append: [4 11 3 7 24 6]
// cpy: [4 11 3 7 24 6]
// sl1: [3 7 24]
// 2d:  [[0] [1 2] [2 3 4]]
// Reverse array: [100 4 3 2 1]
// Sum of array elements: 110
// Sum of array elements: 110
// Sorted array: [3 4 6 7 11 24]
// Bubble sort: [3 4 6 7 11 24]


