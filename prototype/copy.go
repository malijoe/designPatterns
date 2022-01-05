package prototype

import (
    "bytes"
    "encoding/gob"
    "fmt"
)

//func (a *Address) DeepCopy() *Address {
//    return &Address{
//        StreetAddress: a.StreetAddress,
//        City: a.City,
//        Country: a.Country,
//    }
//}
//
//func (p *Person) DeepCopy() *Person {
//    q := *p
//    q.Address = p.Address.DeepCopy()
//    copy(q.Friends, p.Friends)
//    return &q
//}

func (p *Person) DeepCopy() *Person {
    b := bytes.Buffer{}
    e := gob.NewEncoder(&b)
    _ = e.Encode(p)

    fmt.Println(string(b.Bytes()))

    d := gob.NewDecoder(&b)
    result := Person{}
    _ = d.Decode(&result)
    return &result
}

func Copy() {
    john := Person{
        Name: "John",
        Address: &Address{"123 London Rd", "London", "UK"},
        Friends: []string{"Chris", "Matt"},
    }

    jane := john.DeepCopy()
    jane.Name = "Jane"
    jane.Address.StreetAddress = "321 Baker St"
    jane.Friends = append(jane.Friends, "Angela")

    fmt.Println(john, john.Address)
    fmt.Println(jane, jane.Address)
}