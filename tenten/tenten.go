package main

import (
    "fmt";
    "github.com/nojero/computer"
)

func main () {
    comp := computer.New(100)
    comp.SetAddress(90)
    comp.Insert("AAAA",860)
    fmt.Println(comp)
}
