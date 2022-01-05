package factories

import "fmt"

type Employee struct {
    Name, Position string
    AnnualIncome int
}

// functional

func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
    return func(name string) *Employee {
        return &Employee{name, position, annualIncome}
    }
}

// structural

type EmployeeFactory struct {
    Position string
    AnnualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
    return &Employee{name, f.Position, f.AnnualIncome}
}

func NewEmployeeFactoryStruct(position string, annualIncome int) *EmployeeFactory {
    return &EmployeeFactory{position, annualIncome}
}

func EmployeeFactories() {
    // cannot modify values after creation
    // more idiomatic
    developerFactory := NewEmployeeFactory("developer", 60000)
    managerFactory := NewEmployeeFactory("manager", 80000)

    developer := developerFactory("Adam")
    manager := managerFactory("Jane")

    fmt.Println(developer)
    fmt.Println(manager)

    // is modifiable
    bossFactory := NewEmployeeFactoryStruct("CEO", 100000)
    bossFactory.AnnualIncome = 110000
    boss := bossFactory.Create("Sam")
    fmt.Println(boss)
}
