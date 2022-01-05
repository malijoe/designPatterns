package prototype

import "fmt"

type Address struct {
    StreetAddress, City, Country string
}

type Person struct {
    Name string
    Address *Address
    Friends []string
}

func PersonPrototype() {
    john := Person{Name: "John", Address: &Address{"123 London Rd", "London", "UK"}}

    //jane := john // only copies Address pointer not value
    //jane.Name = "Jane" // ok
    //jane.Address.StreetAddress = "321 Baker St" // changes both john's and jane's address

    // deep copy
    jane := john
    jane.Address = &Address {
        StreetAddress: john.Address.StreetAddress,
        City: john.Address.City,
        Country: john.Address.Country,
    }

    jane.Name = "Jane"
    jane.Address.StreetAddress = "321 Baker St"

    fmt.Println(john, john.Address)
    fmt.Println(jane, jane.Address)
}