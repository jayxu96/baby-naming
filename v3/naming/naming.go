package naming
// package main
import (
// 	"fmt"
	"math/rand"
    "time"
)

var maleNames = []string{"Mike", "Jay", "Tim", "Jack", "Andrew"}
var femaleNames = []string{"Lacus", "Wendy", "Elisabeth", "Christine", "Monica"}

func RandomString (length int) string { //https://golangdocs.com/generate-random-string-in-golang
    var letters = []rune("abcdefghijklmnopqrstuvwxyz")
    s := make([]rune, length)
    rand.Seed(time.Now().UTC().UnixNano())
    for i := range s {
        s[i] = letters[rand.Intn(len(letters))]
    }
    return string(s)
}

func GetNameList (male bool) []string {
 	var nameList []string
 	if male {
 		nameList = maleNames
 	} else {
 		nameList = femaleNames
 	}
 	return nameList
}

func CreateBabyName(male bool, minLen int, lastName string) string {
    nameList := GetNameList(male)
    rand.Seed(time.Now().UTC().UnixNano())
    randomize := rand.Intn(len(nameList))
    nameLen := int(len(nameList[randomize]))
    if nameLen < minLen {
       fullName := nameList[randomize] + RandomString(minLen-nameLen) + " " + lastName
       return fullName
    } else {
       fullName := nameList[randomize] + " " + lastName
       return fullName
    }
}

// func main() {
//     var lastName string
//     var gender string
//     var male bool
//     var minLen int
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
//     fmt.Println("Name length: ")
//     fmt.Scanln(&minLen)
//     fmt.Println("Congrats! The baby name is " + CreateBabyName(male, minLen, lastName))
// }
