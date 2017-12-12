package main

import (
    "github.com/nojero/computer"
)

func main () {
    comp := computer.New(100)
    comp.SetAddress(50)
    comp.Insert("MULT", 0)
    comp.Insert("PRINT", 0)
    comp.Insert("RET", 0)
    comp.SetAddress(0)
    comp.Insert("PUSH", 1009)
    comp.Insert("PRINT", 0)
    comp.Insert("PUSH", 6)
    comp.Insert("PUSH", 101)
    comp.Insert("PUSH", 10)
    comp.Insert("CALL", 50)
    comp.Insert("STOP", 0)
    comp.SetAddress(0)
    comp.Execute()
}
