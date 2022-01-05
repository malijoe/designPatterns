package builder

import "fmt"

type FPerson struct {
    name, position string
}


type fpersonMod func(*FPerson)
type FPersonBuilder struct {
    actions []fpersonMod
}

func (b *FPersonBuilder) Called(name string) *FPersonBuilder {
    b.actions = append(b.actions, func(p *FPerson) {
        p.name = name
    })

    return b
}

func (b *FPersonBuilder) WorksAsA(position string) *FPersonBuilder {
    b.actions = append(b.actions, func(p *FPerson) {
        p.position = position
    })

    return b
}

func (b *FPersonBuilder) Build() *FPerson {
    fp := FPerson{}
    for _, a := range b.actions {
        a(&fp)
    }
    return &fp
}

func FunctionalPerson() {
    b := FPersonBuilder{}
    p := b.Called("Dmitri").WorksAsA("Programmer").Build()
    fmt.Println(*p)
}