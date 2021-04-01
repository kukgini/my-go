package main

import (
    "log"
)

type A struct {
    a string
    B
}

type B interface {
    b() error
}

type C interface {
    c() error
}

func main() {
    a := NewB()
    log.Println("i got new A:", a)

    a.b() // it can call b() anyway.

    b, ok := a.(B)
    if ok {
        log.Println("a has B interface", b)
        b.b()
    } else {
        log.Println("a has no B interface", b)
    }

    c, ok := a.(C)
    if ok {
        log.Println("a has C interface", c)
        c.c()
    } else {
        log.Println("a has no C interface", c)
        c.c() // it will panic
    }
}

func NewA() *A {
    return &A{
        a: "a",
    }
}

func NewB() B {
    return &A{
        a: "a",
    }
}

func (a *A) b() error {
    log.Println("b() called !!!")
    return nil
}

func (a *A) c() { // this will fail type assertion because return parameter is missing
    log.Println("c() called !!!")
    return nil
}
