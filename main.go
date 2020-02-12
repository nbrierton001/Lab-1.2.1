package main

import "fmt"

func main() {
    var (
      distance int
      gallonsFuel int
      
)
    fmt.Println("How many miles did you travel?")
  
    fmt.Scanln(&distance)

    fmt.Println("How many gallons of fuel did you use?")
    
    fmt.Scanln(&gallonsFuel)

    
  fmt.Println("Your MPG was", distance/gallonsFuel)
}
