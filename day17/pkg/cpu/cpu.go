package cpu

import (
	"log"
	"strconv"
	"strings"
)

type cpu struct {
	A, B, C int
}

type program []int

type Computer struct {
	c cpu
	p program
}

func NewComputer(s string) Computer {
	parts := strings.Split(s, "\n\n")
	registers := strings.Split(parts[0], "\n")
	A, errA := strconv.ParseInt(strings.Split(registers[0], ": ")[1], 10, 0)
	B, errB := strconv.ParseInt(strings.Split(registers[1], ": ")[1], 10, 0)
	C, errC := strconv.ParseInt(strings.Split(registers[2], ": ")[1], 10, 0)

	if errA != nil || errB != nil || errC != nil {
		log.Fatal("Couldn't parse an int!")
	}

	program := strings.Split(parts[1], ": ")[1]
	prog := make([]int, (len(program)+1)/2)
	for i, r := range program {
		if i%2 == 1 {
			continue
		}
		prog[i/2] = int(r - '0')
	}

	return Computer{cpu{int(A), int(B), int(C)}, prog}
}

func (com *Computer) Run() string {
	i, maxI := 0, len(com.p)-1
	out := make([]string, 0)

	for i < maxI {
		opCode := com.p[i]
		operand := com.p[i+1]

		switch opCode {
		case 0:
			com.c.adv(operand)
		case 1:
			com.c.bxl(operand)
		case 2:
			com.c.bst(operand)
		case 3:
			i += com.c.jnz(operand, i)
		case 4:
			com.c.bxc(operand)
		case 5:
			out = append(out, com.c.out(operand))
		case 6:
			com.c.bdv(operand)
		case 7:
			com.c.cdv(operand)
		}
		i += 2
	}
	return strings.Join(out, ",")
}

func (c *cpu) adv(operand int) {
	operand = c.resolveCombo(operand)
	den := pow2(operand)
	c.A = c.A / den
}

func (c *cpu) bxl(operand int) {
	c.B = c.B ^ operand
}

func (c *cpu) bst(operand int) {
	operand = c.resolveCombo(operand)
	c.B = operand % 8
}

func (c *cpu) jnz(operand int, currentInstruction int) int {
	if c.A == 0 {
		return 0
	}
	return operand - currentInstruction - 2
}

func (c *cpu) bxc(operand int) {
	c.B = c.B ^ c.C
}

func (c *cpu) out(operand int) string {
	return strconv.Itoa(c.resolveCombo(operand) % 8)
}

func (c *cpu) bdv(operand int) {
	operand = c.resolveCombo(operand)
	den := pow2(operand)
	c.B = c.A / den
}

func (c *cpu) cdv(operand int) {
	operand = c.resolveCombo(operand)
	den := pow2(operand)
	c.C = c.A / den
}

func (c cpu) resolveCombo(operand int) int {
	switch operand {
	case 4:
		return c.A
	case 5:
		return c.B
	case 6:
		return c.C
	}
	return operand
}

func pow2(n int) int {
	if n == 0 {
		return 1
	} else {
		return 2 << (n - 1)
	}
}

//TODO implement a reverse machine, that computes the initial value of the registers given an output ??
