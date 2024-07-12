package model

import (
	"fmt"
	"math/rand"
)

type Employee struct {
	id          int
	fullName    string
	salary      float64
	status      string
	insurance   float64
	overtime    float64
	allowance   float64
	totalSalary float64
}

func GenerateEmployees(num int, ch chan Employee) {
	statusOptions := []string{"Permanent", "Contract", "Trainee"}
	names := []string{"a1", "b1", "a2", "c4", "a5"}

	for i := 1; i <= num; i++ {
		employee := Employee{
			id:          100 + i,
			fullName:    names[rand.Intn(len(names))],
			salary:      float64(5000 + rand.Intn(10_000)),
			status:      statusOptions[rand.Intn(len(statusOptions))],
			insurance:   500_000,
			overtime:    55_000,
			allowance:   100_000,
			totalSalary: 0.0,
		}
		ch <- employee
	}
	close(ch)
}

func CalculateTotalSalary(inCh chan Employee, outCh chan []Employee) {
	var updatedEmployees []Employee
	for employee := range inCh {
		employee.totalSalary = employee.salary + employee.insurance + employee.overtime + employee.allowance
		updatedEmployees = append(updatedEmployees, employee)
	}
	outCh <- updatedEmployees
	close(outCh)
}

func PrintResult(updatedEmployees []Employee) {
	for _, employee := range updatedEmployees {
		fmt.Printf("ID: %d, Name: %s, Total Salary: %.2f\n", employee.id, employee.fullName, employee.totalSalary)
	}
}
