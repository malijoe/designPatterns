package builder

import "fmt"

type Person struct {
    // address
    StreetAddress, Postcode, City string
    // job
    CompanyName, Position string
    AnnualIncome int
}

type PersonBuilder struct {
    person *Person
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
    return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
    return &PersonJobBuilder{*b}
}

func NewPersonBuilder() *PersonBuilder {
    return &PersonBuilder{&Person{}}
}

type PersonAddressBuilder struct {
    PersonBuilder
}

func (b *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
    b.person.StreetAddress = streetAddress
    return b
}

func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
    b.person.City = city
    return b
}

func (b *PersonAddressBuilder) WithPostCode(code string) *PersonAddressBuilder {
    b.person.Postcode = code
    return b
}

type PersonJobBuilder struct {
    PersonBuilder
}

func (b *PersonJobBuilder) At(company string) *PersonJobBuilder {
    b.person.CompanyName = company
    return b
}

func (b *PersonJobBuilder) AsA(title string) *PersonJobBuilder {
    b.person.Position = title
    return b
}

func (b *PersonJobBuilder) Earning(income int) *PersonJobBuilder {
    b.person.AnnualIncome = income
    return b
}

func (b *PersonBuilder) Build() *Person {
    return b.person
}

func BuildPerson() {
    pb := NewPersonBuilder()
    pb.Lives().At("123 London Road").In("London").WithPostCode("SW12BC").
        Works().At("Fabrikam").AsA("Programmer").Earning(123000)
    person := pb.Build()
    fmt.Println(person)
}