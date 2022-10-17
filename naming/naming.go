package naming
// package main
import (
// 	"fmt"
	"math/rand"
        "time"
)

var maleNames = []string{"Mike", "Jay", "Tim", "Jack", "Andrew"}
var femaleNames = []string{"Lacus", "Wendy", "Elisabeth", "Christine", "Monica"}

func GetNameList (male bool) []string {
 	var nameList []string
 	if male {
 		nameList = maleNames
 	} else {
 		nameList = femaleNames
 	}
 	return nameList
}

func CreateBabyName(male bool, lastName string) string {
    nameList := GetNameList(male)
    rand.Seed(time.Now().UTC().UnixNano())
    randomize := rand.Intn(len(nameList))
    fullName := nameList[randomize] + " " + lastName
    return fullName
}

// func main() {
//     var lastName string
//     var gender string
//     var male bool
//     fmt.Println("***** This is a Baby Naming Generator *****")
//     fmt.Println(GetNameList(true))
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
