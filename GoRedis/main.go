// package main
// Masalah path Go Redis saja
// import (
// 	"fmt"
// 	"sync"

// 	"github.com/go-redis/redis"
// )

// type Student struct {
// 	ID    int
// 	Name  string
// 	Grade int
// }

// func main() {
// 	students := []Student{
// 		{ID: 1, Name: "John Doe", Grade: 85},
// 		{ID: 2, Name: "Jane Smith", Grade: 92},
// 		{ID: 3, Name: "Bob Johnson", Grade: 70},
// 		{ID: 4, Name: "Sarah Lee", Grade: 88},
// 		{ID: 5, Name: "Tom Wilson", Grade: 78},
// 	}

// 	// Buat koneksi ke Redis
// 	client := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Password: "",
// 		DB:       0,
// 	})

// 	var wg sync.WaitGroup
// 	wg.Add(len(students))

// 	for _, student := range students {
// 		go func(s Student) {
// 			processStudent(&s, client)
// 			wg.Done()
// 		}(student)
// 	}

// 	wg.Wait()
// 	fmt.Println("\nAll students processed!")
// }

// func processStudent(student *Student, client *redis.Client) {
// 	// Perform some processing on student data
// 	if student.Grade >= 85 {
// 		fmt.Printf("%s (%d): Anda mendapatkan A\n", student.Name, student.ID)

// 		// Simpan hasil nilai student ke Redis
// 		err := client.Set(fmt.Sprintf("student:%d", student.ID), "A", 0).Err()
// 		if err != nil {
// 			panic(err)
// 		}
// 	} else if student.Grade >= 70 {
// 		fmt.Printf("%s (%d): Anda Mendapatkan B\n", student.Name, student.ID)

// 		// Simpan hasil nilai student ke Redis
// 		err := client.Set(fmt.Sprintf("student:%d", student.ID), "B", 0).Err()
// 		if err != nil {
// 			panic(err)
// 		}
// 	} else {
// 		fmt.Printf("%s (%d): Anda Mendapatkan C\n", student.Name, student.ID)

// 		// Simpan hasil nilai student ke Redis
// 		err := client.Set(fmt.Sprintf("student:%d", student.ID), "C", 0).Err()
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// }
