package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

// Single Responsibility Principal

var entryCount = 0
type Journal struct {
    entries []string
}

func (j *Journal) AddEntry(text string) int {
    entryCount++
    entry := fmt.Sprintf("%d: %s", entryCount, text)
    j.entries = append(j.entries, entry)
    return entryCount
}

func (j *Journal) RemoveEntry(index int) {
    // ...
}

func (j *Journal) String() string {
    return strings.Join(j.entries, "\n")
}

// Separation of Concerns
// Different problems of a system should resolve in different constructs (package, object, etc...)
// Avoid God Objects (anti-pattern) where everything that needs to be done for an object resides within that object

// These functions break the Single Responsibility Principle. The Journal is responsible for managing entries not persistence.
// func(j *Journal) Save(filename string)
// func(j *Journal) Load(filename string)
// func(*Journal) LoadFromWeb(url *url.URL)

// Persistence should be handled by a separate construct (package, type, etc...)

// Package approach

var LineSeparator = "\n"
func SaveToFile(j *Journal, filename string) {
    _ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, LineSeparator)), 0644)
}

// Struct approach

type Persistence struct {
    lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
    _ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func srp() {
    j := Journal{}
    j.AddEntry("I cried today")
    j.AddEntry("I ate a bug")
    fmt.Println(j.String())
    // SaveToFile(&j, "journal.txt")

    // p := Persistence{"/r/n"}
    //p.SaveToFile(&j, "journal.txt")
}