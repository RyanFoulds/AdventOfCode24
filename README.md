# Advent of Code '24

Solutions are all in Go this year.

Each day is a standalone module, from the `day{day-num}/` directory:
- Compile with `go build -o bin/ day{day-num}.go`
- Execute with:
  - `bin/day{day-num} <input file path>` (Linux, MacOS)
  - `bin\day{day-num}.exe <input file path>` (Windows)

 The solutions for both parts will be printed to the terminal. Except for day 24, which will print the solution to part1 to the terminal, but will write a graph in the [DOT format](https://graphviz.org/doc/info/lang.html) to `./bin/adder-graph.gv`, use this graph with a Visualiser to manually solve part 2.
