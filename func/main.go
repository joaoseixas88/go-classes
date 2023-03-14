package main

import "fmt"

type Person struct {
	name    string
	surname string
	age     int
	job     string
}

type Developer struct {
	Person
	stack  string
	salary float64
}

type PhysicalTherapist struct {
	Person
	specialty string
	salary    float64
}

func (x Developer) getName() {
	fmt.Println("Meu nome é", x.name, " e eu trabalho como desenvolvedor na stack", x.stack, " e ganho: ", x.salary)
}
func (x PhysicalTherapist) getName() {
	fmt.Println("Meu nome é", x.name, " e eu trabalho como fisioterapeuta e sou especialista em", x.specialty, " e ganho: ", x.salary)
}

type Persons interface {
	getName()
}

func main() {

	person_1 := Developer{
		Person: Person{
			name:    "Joao",
			surname: "Seixas",
			age:     35,
			job:     "Developer",
		},
		stack:  "React",
		salary: 3580,
	}
	Persons(person_1).getName()
	person_1.getName()

}
