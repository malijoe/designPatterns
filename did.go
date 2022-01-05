package main

import "fmt"

// Dependency Inversion Principle
// HLM should not depend on LLM
// Both should depend on abstractions

type Relationship int

const (
    Parent Relationship = iota
    Child
    Sibling
)

type Person struct {
    name string
    //
}

type Info struct {
    from *Person
    relationship Relationship
    to *Person
}

// low-level module

type RelationshipBrowser interface {
    FindAllChildrenOf(name string) []*Person
}

type Relationships struct {
    relations []Info
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
    result := make([]*Person, 0)
    for i, v := range r.relations {
        if v.from.name == name {
            result = append(result, r.relations[i].to)
        }
    }

    return result
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
    r.relations = append(r.relations,
        Info{parent, Parent, child})
    r.relations = append(r.relations,
        Info{child, Child, parent})
}

// high-level module

type Research struct {
    // break DIP
    // any change to how relationships stores data will break Investigate function
    //relationships Relationships
    browser RelationshipBrowser
}

// Breaks because of DIP
//func (r *Research) Investigate() {
//    relations := r.relationships.relations
//    for _, rel := range relations {
//        if rel.from.name == "John" && rel.relationship == Parent {
//            fmt.Println("John has a child called ", rel.to.name)
//        }
//    }
//}

func (r *Research) Investigate() {
    for _, p := range r.browser.FindAllChildrenOf("John") {
        fmt.Println("John has a child called ", p.name)
    }
}

func did() {
    parent := Person{"John"}
    child1 := Person{"Chris"}
    child2 := Person{"Matt"}

    relationships := Relationships{}
    relationships.AddParentAndChild(&parent, &child1)
    relationships.AddParentAndChild(&parent, &child2)
    r := Research{&relationships}
    r.Investigate()
}