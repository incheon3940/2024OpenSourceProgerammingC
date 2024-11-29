package main

import "fmt"

type student struct {
	id   int
	name string
	gpa  float32
}

func main() {
	var student1 student
	student1.id = 20241234
	student1.name = "인하"
	student1.gpa = 4.5
	fmt.Println(student1.gpa)

	var student2 student
	student2.id = 20245678
	student2.name = "공전"
	student2.gpa = 4.4
	fmt.Println(student2.id)
}
