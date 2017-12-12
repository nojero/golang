package computer

import "fmt"

type instruction struct {
    ins string
    arg int
}

type computer struct {
    pc int
    sp int
    size int
    stack []instruction
}

func New(size int) computer {
    return computer{0,0,size,make([]instruction,size)}
}

func (c *computer) SetAddress(addr int) {
    if c.pc < c.size {
        c.pc = addr
    }
}

func (c *computer) Insert(ins string, arg int) {
    instr := instruction{ins, arg}
    c.stack[c.pc] = instr
    c.pc += 1
    if c.sp < c.pc && c.pc < c.size {
        c.sp = c.pc
    }
}

func (c *computer) Execute() {
    for {
        i := c.stack[c.pc]
        switch i.ins {
        case "MULT":
            a := c.stack[c.sp - 1]
            b := c.stack[c.sp - 2]
            c.stack[c.sp - 2] = instruction{"", a.arg * b.arg}
            c.sp -= 1
            c.pc += 1
        case "PUSH":
            c.stack[c.sp] = instruction{"", i.arg}
            c.sp += 1
            c.pc += 1
        case "PRINT":
            fmt.Println(c.stack[c.sp - 1].arg)
            c.sp -= 1
            c.pc += 1
        case "STOP":
            return
        case "RET":
            c.pc = c.stack[c.sp - 1].arg
            c.sp -= 1
        case "CALL":
            c.pc = i.arg
        default:
            fmt.Println("Instruction not supported: ", i.ins)
            return
        }

    }
}

