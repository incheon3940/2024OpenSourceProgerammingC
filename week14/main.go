package main

import "fmt"

func main() {
	var student1 struct {
		id   int
		name string
		gpa  float32
	}
	student1.id = 20241234
	student1.name = "인하"
	student1.gpa = 4.5
	fmt.Println(student1.gpa)

	var student2 struct {
		id   int
		name string
		gpa  float32
	}
	student2.id = 20245678
	student2.name = "공전"
	student2.gpa = 4.4
	fmt.Println(student2.id)
}
