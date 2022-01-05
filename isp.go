package main

// Interface Segregation Principle
// Shouldn't put too much into an interface

type Document struct {

}

// Machine is too broad. Not all printers can support all methods
type Machine interface {
    Print(d Document)
    Fax(d Document)
    Scan(d Document)
}

type MultiFunctionPrinter struct {

}

func (m MultiFunctionPrinter) Print(d Document) {
}

func (m MultiFunctionPrinter) Fax(d Document) {
}

func (m MultiFunctionPrinter) Scan(d Document) {
}

type OldFashionedPrinter struct {

}

func (o OldFashionedPrinter) Print(d Document) {
    //ok
}

// Deprecated: ...
func (o OldFashionedPrinter) Fax(d Document) {
    panic("operation not supported")
}

func (o OldFashionedPrinter) Scan(d Document) {
    panic("implement me")
}

// ISP

type Printer interface {
    Print(d Document)
}

type Scanner interface {
    Scan(d Document)
}

type MyPrinter struct {

}

func (m MyPrinter) Print(d Document) {
    //ok
}

type Photocopier struct {

}

func (p Photocopier) Print(d Document) {
    //ok
}

func (p Photocopier) Scan(d Document) {
    //ok
}

type MultiFunctionDevice interface {
    Printer
    Scanner
    // Fax
}

// decorator

type MultiFunctionMachine struct {
    printer Printer
    scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
    m.printer.Print(d)
}


func isp() {
}