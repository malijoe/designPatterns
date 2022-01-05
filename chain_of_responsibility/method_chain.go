package chain_of_responsibility

//type Creature struct {
//    Name string
//    Attack, Defense int
//}
//
//func (c *Creature) String() string {
//    return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack, c.Defense)
//}
//
//func NewCreature(name string, attack int, defense int) *Creature {
//    return &Creature{Name: name, Attack: attack, Defense: defense}
//}

type Modifier interface {
    Add(m Modifier)
    Handle()
}

//type CreatureModifier struct {
//    creature *Creature
//    next Modifier
//}
//
//func (c *CreatureModifier) Add(m Modifier) {
//    if c.next != nil {
//        c.next.Add(m)
//    } else {
//        c.next = m
//    }
//}
//
//func (c *CreatureModifier) Handle() {
//    if c.next != nil {
//        c.next.Handle()
//    }
//}
//
//func NewCreatureModifier(creature *Creature) *CreatureModifier {
//    return &CreatureModifier{creature: creature}
//}
//
//type DoubleAttackModifier struct {
//    CreatureModifier
//}
//
//func (d *DoubleAttackModifier) Handle() {
//    fmt.Println("Doubling", d.creature.Name, "\b's attack")
//    d.creature.attack *= 2
//    d.CreatureModifier.Handle()
//}
//
//func NewDoubleAttackModifier(c *Creature) *DoubleAttackModifier {
//    return &DoubleAttackModifier{CreatureModifier{creature: c}}
//}
//
//type NoBonusesModifier struct {
//    CreatureModifier
//}
//
//func (n *NoBonusesModifier) Handle() {
//    //empty
//}
//
//func NewNoBonusesModifier(c *Creature) *NoBonusesModifier {
//    return &NoBonusesModifier{CreatureModifier{creature: c}}
//}

func MethodChain() {
    //goblin := NewCreature("Goblin", 1, 1)
    //fmt.Println(goblin.String())

    //root := NewCreatureModifier(goblin)
    //
    //root.Add(NewNoBonusesModifier(goblin))
    //
    //root.Add(NewDoubleAttackModifier(goblin))
    //
    //root.Add(NewDoubleAttackModifier(goblin))
    //root.Handle()
    //fmt.Println(goblin.String())
}