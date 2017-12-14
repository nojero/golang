package computer

import "fmt"

type instruction struct {
	ins string
	arg int
}

type Computer struct {
	pc    int
	sp    int
	size  int
	stack []instruction
}

func New(size int) Computer {
	return Computer{0, 0, size, make([]instruction, size)}
}

func (c *Computer) SetAddress(addr int) {
	if c.pc < c.size {
		c.pc = addr
	}
}

func (c *Computer) Insert(ins string, arg int) {
	instr := instruction{ins, arg}
	c.stack[c.pc] = instr
	c.pc += 1
	if c.sp < c.pc && c.pc < c.size {
		c.sp = c.pc
	}
}

func (c *Computer) Print() {
	for i := 0; i < c.size; i++ {
		fmt.Println(i, c.stack[i].ins, c.stack[i].arg)
	}
}

func (c *Computer) Execute() (string, error) {
	ret := ""
	for {
		i := c.stack[c.pc]
		switch i.ins {
		case "MULT":
			a := c.stack[c.sp-1]
			b := c.stack[c.sp-2]
			c.stack[c.sp-2] = instruction{"", a.arg * b.arg}
			c.sp -= 1
			c.pc += 1
		case "PUSH":
			c.stack[c.sp] = instruction{"", i.arg}
			c.sp += 1
			c.pc += 1
		case "PRINT":
			ret = fmt.Sprintf("%s\n%s", ret, c.stack[c.sp-1].arg)
			c.sp -= 1
			c.pc += 1
		case "STOP":
			return ret, nil
		case "RET":
			c.pc = c.stack[c.sp-1].arg
			c.sp -= 1
		case "CALL":
			c.pc = i.arg
		default:
			return "", fmt.Errorf("Instruction not supported: %s", i.ins)
		}

	}
}
