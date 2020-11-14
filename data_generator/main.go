package main

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
)

type RandomStruct struct {
	FirstName   string  `csv:"first_name"`
	MiddleName  string  `csv:"middle_name"`
	LastName    string  `csv:"last_name"`
	DateOfBirth string  `csv:"date_of_birth"`
	Gender      string  `csv:"gender"`
	Address     string  `csv:"address"`
	Salary      float32 `csv:"salary"`
}

func main() {
	random := RandomStruct{
		FirstName:   faker.FirstName(),
		MiddleName:  faker.FirstNameMale(),
		LastName:    faker.LastName(),
		DateOfBirth: faker.Date(),
		Gender:      faker.Gender(),
		Address:     "Random String Random String",
		Salary:      10000,
	}

	printStruct(random)
}

func printStruct(random RandomStruct) {
	fmt.Printf("\n***** DATA *****\n")
	fmt.Printf("Name: %s %s %s\n", random.FirstName, random.MiddleName, random.LastName)
	fmt.Printf("Date of birth: %s\n", random.DateOfBirth)
	fmt.Printf("Gender: %s\n", random.Gender)
	fmt.Printf("Address: %s\n", random.Address)
	fmt.Printf("Salary: %0.2f\n", random.Salary)
	fmt.Printf("***** END *****\n")
}
