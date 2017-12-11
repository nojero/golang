package computer

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
    c.pc = addr
}

func (c *computer) Insert(ins string, arg int) {
    instr := instruction{ins, arg}
    c.stack[c.pc] = instr
}

func (c *computer) Pepe() int {
    return c.pc
}

