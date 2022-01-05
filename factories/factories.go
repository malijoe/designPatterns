package factories

import "fmt"

// Factories
// Object creation logic becomes too convoluted
// Struct has too many fields, need to initialize all correctly
// Wholesale object creation (non-piecewise, unlike Builder) can be outsourced to
//      - A separate function (Factory Function a.k.a. Constructor)
//      - That may exist in a separate struct (Factory)

/*
    Factory - a component responsible for the wholesale (not piecewise) creation of objects.
 */

/*
    Summary:
        a factory function (constructor) is a helper function for making struct instances
        a factory is any entity that can take care of object creation
        can be a function or a dedicated struct
 */

type Person interface {
    SayHello()
}

type person struct {
    Name string
    Age int
    EyeCount int
}

func (p *person) SayHello() {
    fmt.Printf("Hi, my name is %s, I am %d years old\n", p.Name, p.Age)
}

type tiredPerson struct {
    Name string
    Age int
    EyeCount int
}

func (tp *tiredPerson) SayHello() {
    fmt.Println("Sorry, I'm too tired")
}
// NewPerson is a factory function that performs the initialization of person object. Can set default or check conditions on creation
//func NewPerson(name string, age int) *person {
//    if age < 16 {
//        //...
//    }
//    return &person{Name: name, Age: age, EyeCount: 2}
//}

// NewPerson returns interface to Person allowing functionality without modification
func NewPerson(name string, age int) Person {
    if age > 100 {
        return &tiredPerson{Name: name, Age: age}
    }
    return &person{Name: name, Age: age}
}

func Factory() {
    p := NewPerson("John", 33)
    p.SayHello()

    p = NewPerson("Agnes", 134)
    p.SayHello()
}