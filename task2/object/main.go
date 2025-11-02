package main

import (
	"fmt"
)

// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
// 然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}

func (r *Rectangle) Area() {
	fmt.Println("Rectangle Area")
}

func (r *Rectangle) Perimeter() {
	fmt.Println("Rectangle Perimeter")
}

type Circle struct {
}

func (c *Circle) Area() {
	fmt.Println("Circle Area")
}

func (c *Circle) Perimeter() {
	fmt.Println("Circle Perimeter")
}

// 考察点 ：接口的定义与实现、面向对象编程风格。
// 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
// 再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
// 为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
// 考察点 ：组合的使用、方法接收者。

type PrintInfo interface {
	PrintInfo()
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person     Person
	EmployeeID string
}

func (employee *Employee) EmployeePrint() {
	fmt.Println(employee.EmployeeID, employee.Person.Age, employee.Person.Name)
}

func main() {
	circle := &Circle{}
	rectangle := &Rectangle{}
	circle.Area()
	circle.Perimeter()
	rectangle.Area()
	rectangle.Perimeter()

	Employee := &Employee{Person: Person{Name: "test1", Age: 18}, EmployeeID: "19203"}
	Employee.EmployeePrint()
}
