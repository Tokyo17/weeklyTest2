package main

import (
	"sync"
	"weeklytest2/model"
)

func main() {

	numEmployees := 100

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

	model.PrintResult(<-resultCh)

}
