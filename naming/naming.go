package naming
// package main
import (
// 	"fmt"
	"math/rand"
    "time"
)

var babyNames = []string{"Mike", "Jay", "Tim", "Wendy", "Elisabeth", "Lacus"}

func CreateBabyName(lastName string) string {
       rand.Seed(time.Now().UTC().UnixNano())
       randomize := rand.Intn(len(babyNames))
       fullName := babyNames[randomize] + " " + lastName
       return fullName
}

// func main() {
//     var lastName string;
//     fmt.Println("***** This is a Baby Naming Generator *****")
//     fmt.Println("Please enter your Last Name: " )
//
//     fmt.Scanln(&lastName)
//     fmt.Println("The baby name is " + CreateBabyName(lastName))
// }
