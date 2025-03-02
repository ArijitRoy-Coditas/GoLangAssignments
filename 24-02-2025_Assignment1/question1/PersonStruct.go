// Refer to the Questions.md for more detail

package question1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Person struct {
	name string
	age  uint64
}

func (P *Person) intro() {
	fmt.Printf("Hi my name is %v and ", P.name)
	fmt.Printf("I'm %v year old.\n", P.age)
}

func getUserInput(message string) string {
	fmt.Println(message)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func isValidageInput() uint64 {
	var newAge uint64
	var input string
	//Infinite loop which only exits when a valid input is given which is integer in this case.
	for {
		input = getUserInput("Please enter your new age: ")
		if input == "" {
			log.Println("Age cannot be empty")
			continue
		}

		newAge, _ = strconv.ParseUint(input, 10, 64)
		break
	}

	return newAge
}

func ageInput() uint64 {
	var ageInput uint64
	//Assign value and validate if the user input value is an integer type and more than zero
	if ageInput = isValidageInput(); ageInput > 0 {
		return ageInput
	}

	return ageInput
}

func (P *Person) updateAge() {
	var newAge uint64 = ageInput()

	if newAge == P.age {
		log.Println("No changes made to the age")
		return
	}

	if newAge < P.age {
		log.Printf("Age must be greater than current age: %v\n", P.age)
		return
	}

	P.age = newAge
	time.Sleep(time.Millisecond * 800)
	fmt.Printf("Your new age is: %v\n", P.age)
}

func isEligible(age uint64) (uint64, bool) {
	return age, age >= 18
}

func (P *Person) vote() {
	if _, eligible := isEligible(P.age); eligible {
		fmt.Println("You're eligible for vote")
	} else {
		fmt.Println("You're not eligible for vote.")
		age := 18 - P.age
		years := ""
		if age > 1 {
			years = "years"
		} else {
			years = "year"
		}
		fmt.Printf("After %v %v you will be eligible for vote.\n", age, years)
	}
}

func test(intro, Update, vote func()) {
	fmt.Println("========================================")
	intro()
	Update()
	vote()
	fmt.Println("========================================")
	time.Sleep(time.Millisecond * 1000)
}

func Run() {
	person1 := Person{
		name: "Arijit Roy",
		age:  22,
	}
	person2 := Person{
		name: "Anirban Paul",
		age:  23,
	}
	person3 := Person{
		name: "HrishiRaj Modi",
		age:  17,
	}

	test(
		person1.intro,
		person1.updateAge,
		person1.vote,
	)

	test(
		person2.intro,
		person2.updateAge,
		person2.vote,
	)

	test(
		person3.intro,
		person3.updateAge,
		person3.vote,
	)
}
