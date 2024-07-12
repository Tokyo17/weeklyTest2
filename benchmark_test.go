package main

import (
	"math/rand"
	"sync"
	"testing"
	"weeklytest2/model"
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

func BenchmarkCalculateTotalSalaryWithChannels(b *testing.B) {
	numEmployees := 100
	for i := 0; i < b.N; i++ {
		employeeCh := make(chan model.Employee)
		resultCh := make(chan []model.Employee, 1)

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			model.GenerateEmployees(numEmployees, employeeCh)
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			model.CalculateTotalSalary(employeeCh, resultCh)
		}()

		wg.Wait()
		<-resultCh
	}
}

// ================== Tanpa Channel================

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

func CalculateTotalSalaryNonChannel(employees []Employee) []Employee {
	var updatedEmployees []Employee
	for _, employee := range employees {
		employee.totalSalary = employee.salary + employee.insurance + employee.overtime + employee.allowance
		updatedEmployees = append(updatedEmployees, employee)
	}
	return updatedEmployees
}

func BenchmarkCalculateTotalSalaryWithoutChannels(b *testing.B) {
	numEmployees := 100
	for i := 0; i < b.N; i++ {
		var employees []Employee
		employeeCh := make(chan Employee, numEmployees)

		go GenerateEmployees(numEmployees, employeeCh)

		for emp := range employeeCh {
			employees = append(employees, emp)
		}

		_ = CalculateTotalSalaryNonChannel(employees)
	}
}
