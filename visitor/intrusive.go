package visitor

import (
    "fmt"
    "strings"
)

// intrusive visitor implies modifying any object implementations to accommodate the visit

// Intrusive implementation
//type Expression interface {
//    Print(sb *strings.Builder)
//}
//
//type DoubleExpression struct {
//    value float64
//}
//
//func (d *DoubleExpression) Print(sb *strings.Builder) {
//    sb.WriteString(fmt.Sprintf("%g", d.value))
//}
//
//type AdditionExpression struct {
//    left, right Expression
//}
//
//func (a *AdditionExpression) Print(sb *strings.Builder) {
//    sb.WriteRune('(')
//    a.left.Print(sb)
//    sb.WriteRune('+')
//    a.right.Print(sb)
//    sb.WriteRune(')')
//}

// reflective implementation
//type Expression interface {
//}
//
//type DoubleExpression struct {
//   value float64
//}
//
//type AdditionExpression struct {
//   left, right Expression
//}
//
//func Print(e Expression, sb *strings.Builder) {
//    if de, ok := e.(*DoubleExpression); ok {
//        sb.WriteString(fmt.Sprintf("%g", de.value))
//    } else if ae, ok := e.(*AdditionExpression); ok {
//        sb.WriteRune('(')
//        Print(ae.left, sb)
//        sb.WriteRune('+')
//        Print(ae.right, sb)
//        sb.WriteRune(')')
//    }
//}


type ExpressionVisitor interface {
    VisitDoubleExpression(e *DoubleExpression)
    VisitAdditionExpression(e *AdditionExpression)
}

type Expression interface {
    Accept(ev ExpressionVisitor)
}

type DoubleExpression struct {
    value float64
}

func (d *DoubleExpression) Accept(ev ExpressionVisitor) {
    ev.VisitDoubleExpression(d)
}

type AdditionExpression struct {
    left, right Expression
}

func (a *AdditionExpression) Accept(ev ExpressionVisitor) {
    ev.VisitAdditionExpression(a)
}

type ExpressionPrinter struct {
    sb strings.Builder
}

func NewExpressionPrinter() *ExpressionPrinter {
    return &ExpressionPrinter{strings.Builder{}}
}

func (ep *ExpressionPrinter) VisitDoubleExpression(e *DoubleExpression) {
    ep.sb.WriteString(fmt.Sprintf("%g", e.value))
}

func (ep *ExpressionPrinter) VisitAdditionExpression(e *AdditionExpression) {
    ep.sb.WriteRune('(')
    e.left.Accept(ep)
    ep.sb.WriteRune('+')
    e.right.Accept(ep)
    ep.sb.WriteRune(')')
}

func (ep *ExpressionPrinter) String() string {
    return ep.sb.String()
}

type ExpressionEvaluator struct {
    result float64
}

func (ee *ExpressionEvaluator) VisitDoubleExpression(e *DoubleExpression) {
    ee.result = e.value
}

func (ee *ExpressionEvaluator) VisitAdditionExpression(e *AdditionExpression) {
    e.left.Accept(ee)
    x := ee.result
    e.right.Accept(ee)
    x += ee.result
    ee.result = x
}

func (ee *ExpressionEvaluator) String() string {
    return fmt.Sprintf("%g", ee.result)
}

func Visit() {
    // 1 + (2+3)
    e := &AdditionExpression{
        left: &DoubleExpression{value: 1.0},
        right: &AdditionExpression{
            left: &DoubleExpression{value: 2.0},
            right: &DoubleExpression{value: 3.0},
        },
    }

    ep := NewExpressionPrinter()
    e.Accept(ep)
    //fmt.Println(ep.String())

    ee := &ExpressionEvaluator{}
    e.Accept(ee)
    fmt.Println(ep,"=", ee)
}
