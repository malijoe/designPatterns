package main

import "fmt"

// OCP (Open Close Principle)
// open for extension, closed for modification

type Color int

const (
    red Color = iota
    green
    blue
)

type Size int

const (
    small Size = iota
    medium
    large
)

type Product struct {
    name string
    color Color
    size Size
}

// Filter Using this method every time a new filter behavior is desired a new method will have to be created. Meaning the Filter object will endlessly need to be modified. This violates the OCP principle.
type Filter struct {}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
    result := make([]*Product, 0)

    for i, v := range products {
        if v.color == color {
            result = append(result, &products[i])
        }
    }

    return result
}

// Specification Pattern
// New Specifications can be created and ambiguously used using BetterFilter.Filter

type Spec interface {
    IsSatisfied(p *Product) bool
}

type ColorSpec struct {
    color Color
}

func (s ColorSpec) IsSatisfied(p *Product) bool {
    return p.color == s.color
}

type SizeSpec struct {
    size Size
}

func (s SizeSpec) IsSatisfied(p *Product) bool {
    return p.size == s.size
}

type AndSpec struct {
    first, second Spec
}

func (s AndSpec) IsSatisfied (p *Product) bool {
    return s.first.IsSatisfied(p) && s.second.IsSatisfied(p)
}

type BetterFilter struct{}

func (bf BetterFilter) Filter(products []Product, spec Spec) []*Product {
    result := make([]*Product, 0)

    for i, v := range products {
        if spec.IsSatisfied(&v) {
            result = append(result, &products[i])
        }
    }

    return result
}

func ocp() {
    apple := Product{"Apple", green, small}
    tree := Product{"Tree", green, large}
    house := Product{"House", blue, large}

    products := []Product{apple, tree, house}
    fmt.Println("Green Products (old): ")
    f := Filter{}
    for _, v := range f.FilterByColor(products, green) {
        fmt.Printf(" - %s is green\n", v.name)
    }

    fmt.Println("Green Products (new): ")
    greenSpec := ColorSpec{green}
    bf := BetterFilter{}
    for _, v := range bf.Filter(products, greenSpec) {
        fmt.Printf(" - %s is green\n", v.name)
    }

    fmt.Println("Large green Products")
    largeSpec := SizeSpec{large}
    lgSpec := AndSpec{largeSpec, greenSpec}
    for _, v := range bf.Filter(products, lgSpec) {
        fmt.Printf(" - %s is large and green\n", v.name)
    }
}