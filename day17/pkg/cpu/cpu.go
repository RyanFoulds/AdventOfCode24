package cpu

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

type cpu struct {
	A, B, C uint
}

type program []uint

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
		log.Fatal("Couldn't parse an uint!")
	}

	program := strings.Split(parts[1], ": ")[1]
	prog := make([]uint, (len(program)+1)/2)
	for i, r := range program {
		if i%2 == 1 {
			continue
		}
		prog[i/2] = uint(r - '0')
	}

	return Computer{cpu{uint(A), uint(B), uint(C)}, prog}
}

func (com *Computer) Run() string {
	i, maxI := 0, len(com.p)-1
	out := make([]string, 0)

	for i < maxI {
		opCode := com.p[i]
		operand := uint(com.p[i+1])

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

func (c *cpu) adv(operand uint) {
	operand = c.resolveCombo(operand)
	den := pow2(operand)
	c.A = c.A / den
}

func (c *cpu) bxl(operand uint) {
	c.B = c.B ^ operand
}

func (c *cpu) bst(operand uint) {
	operand = c.resolveCombo(operand)
	c.B = operand % 8
}

func (c *cpu) jnz(operand uint, currentInstruction int) int {
	if c.A == 0 {
		return 0
	}
	return int(operand) - currentInstruction - 2
}

func (c *cpu) bxc(operand uint) {
	c.B = c.B ^ c.C
}

func (c *cpu) out(operand uint) string {
	return strconv.Itoa(int(c.resolveCombo(operand) % 8))
}

func (c *cpu) bdv(operand uint) {
	operand = c.resolveCombo(operand)
	den := pow2(operand)
	c.B = c.A / den
}

func (c *cpu) cdv(operand uint) {
	operand = c.resolveCombo(operand)
	den := pow2(operand)
	c.C = c.A / den
}

func (c cpu) resolveCombo(operand uint) uint {
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

func pow2(n uint) uint {
	if n == 0 {
		return 1
	} else {
		return 2 << (n - 1)
	}
}

func withA(com Computer, a uint) Computer {
	return Computer{cpu{a, com.c.B, com.c.C}, com.p}
}

func (com Computer) Search() uint {
	cache := map[uint]struct{}{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}, 6: struct{}{}, 7: struct{}{}}
	solutions := make([]uint, 0)
	strs := make([]string, 0)
	for _, n := range com.p {
		strs = append(strs, strconv.Itoa(int(n)))
	}
	target := strings.Join(strs, ",")

	queue := []uint{1, 2, 3, 4, 5, 6, 7}
	var a uint

	for len(queue) > 0 {
		a, queue = queue[0], queue[1:]

		for i := 0; i < 8; i++ {
			nextA := a<<3 + uint(i)
			_, alreadySeen := cache[nextA]
			if alreadySeen {
				continue
			}

			c := withA(com, nextA)
			output := c.Run()

			if target == output {
				solutions = append(solutions, nextA)
			} else if strings.HasSuffix(target, output) {
				queue = append(queue, nextA)
			}
		}
	}
	return slices.Min(solutions)
}
