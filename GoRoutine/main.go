package main

import (
	"fmt"
	"sync"
)

type Student struct {
	ID    int
	Name  string
	Grade int
}

func main() {
	students := []Student{
		{ID: 1, Name: "John Doe", Grade: 85},
		{ID: 2, Name: "Jane Smith", Grade: 92},
		{ID: 3, Name: "Bob Johnson", Grade: 70},
		{ID: 4, Name: "Sarah Lee", Grade: 88},
		{ID: 5, Name: "Tom Wilson", Grade: 78},
	}

	var wg sync.WaitGroup
	wg.Add(len(students))

	for _, student := range students {
		go func(s Student) {
			processStudent(&s)
			wg.Done()
		}(student)
	}

	wg.Wait()
	fmt.Println("All students processed!")
}

func processStudent(student *Student) {
	// Perform some processing on student data
	if student.Grade >= 85 {
		fmt.Printf("%s (%d): Anda mendapatkan A\n", student.Name, student.ID)
	} else if student.Grade >= 70 {
		fmt.Printf("%s (%d): Anda Mendapatkan B\n", student.Name, student.ID)
	} else {
		fmt.Printf("%s (%d): Anda Mendapatkan C \n", student.Name, student.ID)
	}
}
