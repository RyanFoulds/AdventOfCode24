package logic

import (
	"log"
	"strconv"
	"strings"
)

type Puzzle struct{
  values map[string]int
  gates map[string]gate
  allWires map[string]struct{}
}

type gate struct {
  arg0, arg1 string
  operator func(a, b int) int
}

var functions map[string]func(a,b int) int = map[string]func(a int, b int) int {
  "AND" : and,
  "XOR": xor,
  "OR" : or}

func and(a, b int) int {
  return a & b
}

func xor(a,b int) int {
  return a ^ b
}

func or(a, b int) int {
  return a | b
}

func NewPuzzle(s string) Puzzle {
  parts := strings.Split(s, "\n\n")
  inputs := make(map[string]int)
  allWires := make(map[string]struct{})
  allGates := make(map[string]gate)

  inputStrs := strings.Split(parts[0], "\n")
  for _, inputStr := range inputStrs {
    split := strings.Split(inputStr, ": ")
    val, err := strconv.ParseInt(split[1], 10, 0)
    if err != nil {
      log.Fatal("Couldn't parse input", inputStr)
    }
    inputs[split[0]] = int(val)
    allWires[split[0]] = struct{}{}
  }

  gateStrs := strings.Split(parts[1], "\n")
  for _, gateStr := range gateStrs {
    split := strings.Split(gateStr, " -> ")
    operation := strings.Split(split[0], " ")
    g := gate{operation[0], operation[2], functions[operation[1]]}
    allGates[split[1]] = g

    allWires[split[1]] = struct{}{}
    allWires[operation[0]] = struct{}{}
    allWires[operation[2]] = struct{}{}
  }

  return Puzzle{inputs, allGates, allWires}
}

func (p *Puzzle) resolveValue(address string) int {
  if val, ok := p.values[address]; ok {
    return val
  }
  g, gOk := p.gates[address]
  if !gOk {
    log.Fatal("Could not resolve address", address)
  }
  a, b := p.resolveValue(g.arg0), p.resolveValue(g.arg1)
  result := g.operator(a, b)
  p.values[address] = result
  return result
}

func (p *Puzzle) SolvePartOne() (sum int) {
  for wire := range p.allWires {
    if !strings.HasPrefix(wire, "z") {
      continue
    }
    shift, err := strconv.ParseInt(wire[1:], 10, 0)
    if err != nil {
      log.Fatal("Could not parse Z wire", wire)
    }
    value := p.resolveValue(wire)
    sum += value << int(shift)
  }
  return
}

