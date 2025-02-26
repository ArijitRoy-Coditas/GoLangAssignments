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

type Person struct{
	name string
	age uint64
}

func (P *Person) Intro() {
	fmt.Printf("Hi my name is %v and ", P.name)
	fmt.Printf("I'm %v year old.\n",P.age)
}

func GetUserInput(message string) string {
	fmt.Println(message)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func IsValidAgeInput() uint64 {
	var newAge uint64
	var input string
	//Infinite loop which only exits when a valid input is given which is integer in this case.
	for {
		input = GetUserInput("Please enter your new age: ")
		if input == "" {
			log.Println("Age cannot be empty")
			continue
		}

		newAge, _ = strconv.ParseUint(input, 10, 64)
		break
	}

	return newAge
}

func AgeInput() uint64 {
	var ageInput uint64
	//Assign value and validate if the user input value is an integer type and more than zero
	if ageInput = IsValidAgeInput(); ageInput > 0 {
		return ageInput
	}

	return ageInput
}

func (P *Person) UpdateAge() {
	var newAge uint64 = AgeInput()

	if newAge == P.age {
		log.Println("No changes made to the age")
		return
	}

	if newAge < P.age {
		log.Printf("Age must be greater than current age: %v\n",P.age)
		return
	}

	P.age = newAge
	time.Sleep(time.Millisecond * 800)
	fmt.Printf("Your new age is: %v\n",P.age)
}

func IsEligible(age uint64) (uint64, bool) {
	return age, age >= 18
}

func (P *Person) Vote() {
	if _, eligible := IsEligible(P.age); eligible {
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
		fmt.Printf("After %v %v you will be eligible for vote.\n",age,years)
	}
}

func test(Intro, Update, Vote func() ){
	fmt.Println("========================================")
	Intro()
	Update()
	Vote()
	fmt.Println("========================================")
	time.Sleep(time.Millisecond * 1000)
}

func Run() {
	person1 := Person{
		name: "Arijit Roy",
		age: 22,
	}
	person2 := Person{
		name: "Anirban Paul",
		age: 23,
	}
	person3 := Person{
		name: "HrishiRaj Modi",
		age: 17,
	}

	test(
		person1.Intro,
		person1.UpdateAge,
		person1.Vote,
	)

	test(
		person2.Intro,
		person2.UpdateAge,
		person2.Vote,
	)

	test(
		person3.Intro,
		person3.UpdateAge,
		person3.Vote,
	)
}