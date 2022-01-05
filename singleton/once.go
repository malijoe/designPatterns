package singleton

import (
    "fmt"
    "sync"
)

type Database interface {
    GetPopulation(string) int
}

type singletonDatabase struct {
    capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
    return db.capitals[name]
}

// thread safe (sync.Once or init())
// laziness (only done once asked) this ideal leans toward sync.Once approach

var once sync.Once
var instance *singletonDatabase

func GetDatabase() *singletonDatabase {
    once.Do(func() {
        // init instance make db connection etc...
        instance = &singletonDatabase{capitals: make(map[string]int)}
        instance.capitals["Seoul"] = 17500000
        instance.capitals["Mexico City"] = 17400000
    })
    return instance
}

func GetTotalPopulation(cities []string) int {
    result := 0
    for _, city := range cities {
        result += GetDatabase().GetPopulation(city)
    }
    // Depending directly on singletonDatabase (violates Dependency Inversion Principle DIP)
    return result
}

func GetTotalPopulationEx(db Database, cities []string) int {
    result := 0
    for _, city := range cities {
        result += db.GetPopulation(city)
    }
    // providing a Database interface provides abstraction and removes dependency on a real database object
    return result
}

type DummyDatabase struct {
    dummyData map[string]int
}

func (d *DummyDatabase) GetPopulation(s string) int {
    if len(d.dummyData) == 0 {
        d.dummyData = map[string]int {
            "alpha": 1,
            "beta": 2,
            "gamma": 3,
        }
    }
    return d.dummyData[s]
}

func Once() {
    //db := GetDatabase()
    //cities := []string{"Seoul", "Mexico City"}
    //tp := GetTotalPopulation(cities)
    //fmt.Println(tp)

    names := []string{"alpha", "gamma"}
    tp := GetTotalPopulationEx(&DummyDatabase{}, names)
    fmt.Println(tp == 4)
}
