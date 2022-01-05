package factories

import "fmt"

const (
    Developer = iota
    Manager
)

func NewEmployee(role int) *Employee {
    switch role {
    case Developer:
        return &Employee{"", "developer", 60000}
    case Manager:
        return &Employee{Name: "", Position: "manager", AnnualIncome: 80000}
    default:
        panic("unsupported role")
    }
}

func PrototypeFactory() {
    m := NewEmployee(Manager)
    m.Name = "Sam"
    fmt.Println(m)
}
