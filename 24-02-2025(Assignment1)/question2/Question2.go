// Refer to the Questions.md for more detail

package question2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
	"github.com/olekukonko/tablewriter"
)

type employee struct {
	name string
	age int
	salary float64
}

type department struct {
	name string
	employees []employee
}

func NewEmployeeDetails() []string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the name of the new employee")
	scanner.Scan()
	name := scanner.Text()
	fmt.Println("Enter the age of the new employee")
	scanner.Scan()
	age := scanner.Text()
	fmt.Println("Enter the salary of the new employee")
	scanner.Scan()
	salary := scanner.Text()
	return []string{name,age,salary}
}

func (d *department) AddEmployee() {
	empDetails := NewEmployeeDetails()
	age, _ := strconv.Atoi(empDetails[1])
	salary, _ := strconv.ParseFloat(empDetails[2],64)
	d.employees = append(d.employees, employee{empDetails[0], age, salary})
	fmt.Println("Hang on! we are adding new employee details...")
	time.Sleep(time.Millisecond * 1500)
	EmployeeDetails(d)
	
}

func (d *department) RemoveEmployee() {
	var empID int
	fmt.Printf("Enter the employee ID whose record you want to delete: ")
	_, err := fmt.Scanln(&empID)
	if err != nil {
		fmt.Println("Invalid response ",err)
	}
	var newEmpDetails []employee
	for ID, eachEmp := range d.employees {
		if ID+1 != empID {
			newEmpDetails = append(newEmpDetails, eachEmp)
		}
	}

	d.employees = newEmpDetails
	fmt.Printf("Hang on! we are removing the employee details with ID%v...\n",empID)
	time.Sleep(time.Millisecond * 1500)
	EmployeeDetails(d)
}

func EmployeeDetails(d *department){
	table := tablewriter.NewWriter(os.Stdout)
	fmt.Println("====================================")
	fmt.Println("\tDepartment -- ",d.name)
	fmt.Println("====================================")
	table.SetHeader([]string{"ID","Name","Age","Salary"})
	for ID, eachEmp := range d.employees {
		table.Append([]string{
			fmt.Sprintf("%d",ID+1),
			eachEmp.name,
			fmt.Sprintf("%d",eachEmp.age),
			fmt.Sprintf("%.2f",eachEmp.salary),
		})
	}	
	table.Render()
}


func (d *department) GiveRaise() {
	var empID int
	fmt.Printf("Enter the employee ID whom you want to give raise: ")
	_, err := fmt.Scanln(&empID)
	if err != nil {
		fmt.Println("Invalid response ",err)
	}


	newSalary := 0.0
	fmt.Printf("Enter the raise amount you want to give for the empID%v: ",empID)
	fmt.Scanln(&newSalary)
	var newEmpDetails []employee
	for ID, eachEmp := range d.employees {
		if ID+1 == empID {
			if newSalary < eachEmp.salary {
				fmt.Printf("Raise amount need to be higher than old amount %v\n",eachEmp.salary)
				return
			} else {
				eachEmp.salary = newSalary
			}
		}
		newEmpDetails = append(newEmpDetails, eachEmp)
	}

	d.employees = newEmpDetails
	fmt.Printf("Hang on! we are updating the salary of the EmpID%v...\n",empID)
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
	fmt.Printf("The average salary of %v department employees is: %.2f\n",d.name, average)
}


func Run(){
	
	department1 := department{
		name: "GoLang",
		employees: []employee{
			{"Arijit",22,11000.0},
			{"Anirban",22,35000.0},
			{"HrishiRaj",22,18000.0},
		},
	}

	EmployeeDetails(&department1)
	department1.AddEmployee()
	department1.RemoveEmployee()
	department1.GiveRaise()
	department1.AverageSalary()

	department2 := department{
		name: "DotNet",
		employees: []employee{
			{"Subham",22,30000.0},
			{"Tanmoy",22,35000.0},
			{"Mostafizur",22,18000.0},
		},
	}

	EmployeeDetails(&department2)
	department2.AddEmployee()
	department2.RemoveEmployee()
	department2.GiveRaise()
	department2.AverageSalary()
}