// Refer to the Questions.md for more detail

package question2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

type employee struct {
	name   string
	age    uint
	salary float64
}

type department struct {
	name      string
	employees map[int]employee
}

func GetUserInput(message string) string {
	fmt.Println(message)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func GetUserID(d *department, message string) uint64 {
	var input string
	var empID uint64

	for {
		input = GetUserInput(message)
		if input == "" {
			log.Println("Employee ID cannot be empty")
			continue
		}

		if strings.HasPrefix(input, "-") {
			log.Println("Employee ID cannot be negative")
			continue
		}

		var err error
		empID, err = strconv.ParseUint(input, 10, 64)
		if err != nil {
			log.Println("Invalid employee ID:", err)
			continue
		}

		if empID < 1 || int(empID) > len(d.employees) {
			log.Printf("Employee ID %v does not exist. Try again.", empID)
			continue
		}
		break
	}

	return empID
}

func NewEmployeeDetails() []string {
	var name string
	for {
		name = GetUserInput("Enter the name of the new employee:")
		if name == "" {
			log.Println("Name cannot be empty")
			continue
		}
		if _, err := strconv.Atoi(name); err == nil {
			log.Println("Employee name cannot be a number.")
			continue
		} else {
			break
		}
	}

	// Get and validate the employee's age
	var age uint64
	for {
		input := GetUserInput("Enter the age of the new employee:")
		if strings.HasPrefix(input, "-") {
			log.Println("Age cannot be negative.")
			continue
		}

		if input == "0" {
			log.Println("Age cannot be zero")
			continue
		}

		var err error
		age, err = strconv.ParseUint(input, 10, 32)
		if err != nil {
			log.Println("Invalid age:", err)
			continue
		}
		break
	}

	// Get and validate the employee's salary
	var salary float64
	for {
		input := GetUserInput("Enter the salary of the new employee:")
		if strings.HasPrefix(input, "-") {
			log.Println("Salary cannot be negative.")
			continue
		}

		if input == "0" {
			log.Println("Salary cannot be zero")
			continue
		}
		var err error
		salary, err = strconv.ParseFloat(input, 64)
		if err != nil {
			log.Println("Invalid salary:", err)
			continue
		}
		break
	}

	return []string{name, fmt.Sprintf("%d", age), fmt.Sprintf("%.2f", salary)}
}

func (d *department) AddEmployee() {
	var empDetails []string
	empDetails = NewEmployeeDetails()
	if empDetails == nil {
		log.Println("Employee details were not valid. Employee not added...")
		empDetails = NewEmployeeDetails()
	}

	age, _ := strconv.ParseUint(empDetails[1], 10, 32)
	salary, _ := strconv.ParseFloat(empDetails[2], 64)
	//
	empId := len(d.employees) + 1
	d.employees[empId] = employee{
		name:   empDetails[0],
		age:    uint(age),
		salary: salary,
	}

	fmt.Println("Hang on! we are adding new employee details...")
	time.Sleep(time.Millisecond * 1500)
	EmployeeDetails(d)

}

func (d *department) RemoveEmployee() {
	empID := GetUserID(d, "Enter the employee ID whose record you want to delete: ")

	delete(d.employees, int(empID))

	fmt.Printf("Hang on! we are removing the employee details with ID%v...\n", empID)
	time.Sleep(time.Millisecond * 1500)
	EmployeeDetails(d)
}

func EmployeeDetails(d *department) {
	keys := make([]int, 0, len(d.employees))
	for key := range d.employees {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	table := tablewriter.NewWriter(os.Stdout)
	fmt.Println("====================================")
	fmt.Println("\tDepartment -- ", d.name)
	fmt.Println("====================================")
	table.SetHeader([]string{"ID", "Name", "Age", "Salary"})
	for _, ID := range keys {
		table.Append([]string{
			fmt.Sprintf("%d", ID),
			d.employees[ID].name,
			fmt.Sprintf("%d", d.employees[ID].age),
			fmt.Sprintf("%.2f", d.employees[ID].salary),
		})
	}
	table.Render()
}

func (d *department) GiveRaise() {

	empID := GetUserID(d, "Enter the employee ID whom you want to give raise: ")

	newSalary := 0.0
	for {
		fmt.Printf("Enter the raise amount you want to give for the empID%v: ", empID)
		fmt.Scanln(&newSalary)
		if emp, exists := d.employees[int(empID)]; exists {
			if newSalary < emp.salary {
				log.Printf("Raise amount must be higher than the current salary: %.2f\n", emp.salary)
				continue
			}
			emp.salary = newSalary
			d.employees[int(empID)] = emp
			break
		}
	}

	fmt.Printf("Hang on! we are updating the salary of the EmpID%v...\n", empID)
	time.Sleep(time.Millisecond * 1500)
	EmployeeDetails(d)

}

func (d *department) AverageSalary() {
	sum := 0.0
	count := 0.0
	average := 0.0
	for _, eachEmp := range d.employees {
		sum += eachEmp.salary
		count++
	}
	average = sum / count
	fmt.Printf("The average salary of %v department employees is: %.2f\n", d.name, average)
}

func Run() {

	department1 := department{
		name: "GoLang",
		employees: map[int]employee{
			1: {"Arijit", 22, 11000.0},
			2: {"Anirban", 22, 35000.0},
			3: {"HrishiRaj", 22, 18000.0},
		},
	}

	EmployeeDetails(&department1)
	department1.AddEmployee()
	department1.RemoveEmployee()
	department1.GiveRaise()
	department1.AverageSalary()

	department2 := department{
		name: "DotNet",
		employees: map[int]employee{
			1: {"Subham", 22, 30000.0},
			2: {"Tanmoy", 22, 35000.0},
			3: {"Mostafizur", 22, 18000.0},
		},
	}

	EmployeeDetails(&department2)
	department2.AddEmployee()
	department2.RemoveEmployee()
	department2.GiveRaise()
	department2.AverageSalary()
}
