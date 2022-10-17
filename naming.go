package naming
// package main
import (
// 	"fmt"
	"math/rand"
    "time"
)

var maleNames = []string{"Mike", "Jay", "Tim", "Jack", "Andrew"}
var femaleNames = []string{"Lacus", "Wendy", "Elisabeth", "Christine", "Monica"}

func CreateBabyName(male bool, lastName string) string {
       rand.Seed(time.Now().UTC().UnixNano())
       if male {
            randomize := rand.Intn(len(maleNames))
            fullName := maleNames[randomize] + " " + lastName
            return fullName
       } else {
           randomize := rand.Intn(len(femaleNames))
           fullName := femaleNames[randomize] + " " + lastName
           return fullName
       }
}

// func main() {
//     var lastName string
//     var gender string
//     var male bool
//     fmt.Println("***** This is a Baby Naming Generator *****")
//     fmt.Println("1. Please enter your Last Name: " )
//     fmt.Scanln(&lastName)
//     fmt.Println("2. Is your baby a boy or girl?")
//     fmt.Scanln(&gender)
//     fmt.Println("Your baby is a " + gender)
//     if gender == "boy" {
//        male = true
//     } else if gender == "girl" {
//       male = false
//     } else {
//       fmt.Println ("Sorry, this generator does not support generate names for other genders.")
//       return
//     }
//     fmt.Println("Congrats! The baby name is " + CreateBabyName(male, lastName))
// }
