package main

import "fmt"

func main() {
    var (
      nounOne string
      nounTwo string
      occupationJobONE string
      verbOne string
      placeLocation string
      nounThree string
      occupationJobTWO string
      nounFour string
      minutesTime string
      nounFive string
)
    fmt.Println("Choose a noun.")
  
    fmt.Scanln(&nounOne)

    fmt.Println("Choose another noun.")
    
    fmt.Scanln(&nounTwo)

    fmt.Println("Choose a occupation/job.")
    fmt.Scanln(&occupationJobONE)

    fmt.Println("Choose a verb.")
    
    fmt.Scanln(&verbOne)

    fmt.Println("Choose a place or location.")
    
    fmt.Scanln(&placeLocation)

    fmt.Println("Choose another noun.")
    
    fmt.Scanln(&nounThree)

    fmt.Println("Choose another occupation/job.")
    
    fmt.Scanln(&occupationJobTWO)

    fmt.Println("Choose another noun.")
    
    fmt.Scanln(&nounFour)

    fmt.Println("Choose a number.")
    
    fmt.Scanln(&minutesTime)

    fmt.Println("Choose another noun.")
    
    fmt.Scanln(&nounFive)

  fmt.Println("It was during the battle of", nounOne,". A",nounTwo,"went off right next to my platoon. Our", occupationJobONE," yelled us to ", verbOne," to the nearest",placeLocation," we could find. When we got to the",placeLocation, "we waited for our next orders more", nounOne, "continued to explode arround us. Over the radio, our commanding", occupationJobTWO," told us that", nounFour, "support was inbound. we hunkered down and waited for", minutesTime, "minutes until that glorious", nounFour, "support arrived, turning the tide of battle in our favor. We lost many good", nounFive, "that day, but we were victorious in the end.")
}