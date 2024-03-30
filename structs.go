package main

import ("fmt")

type Person struct {
  name string
  age int
  profession string
}

type Car struct {
  make string
  model string
  year int
}

type Book struct {
  title string
  author string
  year int
}

type Course struct {
  name string
  instructor string
  duration int
}

type Movie struct {
  title string
  director string
  year int
}

type Song struct {
  title string
  artist string
  year int
}

type Product struct {
  name string
  price float64
  category string
}

func main() {
  // Structs

  // Create a new person
  fmt.Println("------------------Person------------------")
  person := Person{"Google Jr", 25, "Software Developer"}
  fmt.Println("Name: ", person.name)
  fmt.Println("Age: ", person.age)
  fmt.Println("Profession: ", person.profession)


  // Create a new car
  fmt.Println("------------------Car------------------")
  car := Car{"Tesla", "Cybertruck", 2023}
  fmt.Println("Make: ", car.make)
  fmt.Println("Model: ", car.model)
  fmt.Println("Year: ", car.year)

  // Create a new book
  fmt.Println("------------------Book------------------")
  book := Book{"The Alchemist", "Paulo Coelho", 1988}
  fmt.Println("Title: ", book.title)
  fmt.Println("Author: ", book.author)
  fmt.Println("Year: ", book.year)

  // Create a new course
  fmt.Println("------------------Course------------------")
  course := Course{"Go Programming", "Andile Jaden", 30}
  fmt.Println("Name: ", course.name)
  fmt.Println("Instructor: ", course.instructor)
  fmt.Println("Duration: ", course.duration)

  // Create a new movie
  fmt.Println("------------------Movie------------------")
  movie := Movie{"Oppenheimer", "Christopher Nolan", 2023}
  fmt.Println("Title: ", movie.title)
  fmt.Println("Director: ", movie.director)
  fmt.Println("Year: ", movie.year)

  // Create a new song
  fmt.Println("------------------Song------------------")
  song := Song{"Keenan Te", "Mine", 2020}
  fmt.Println("Title: ", song.title)
  fmt.Println("Artist: ", song.artist)
  fmt.Println("Year: ", song.year)

  // Create a new product
  fmt.Println("------------------Product------------------")
  product := Product{"Laptop", 999.99, "Electronics"}
  fmt.Println("Name: ", product.name)
  fmt.Println("Price: ", product.price)
  fmt.Println("Category: ", product.category)

}
