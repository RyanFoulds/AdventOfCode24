package robot

import (
	"log"
	"strconv"
)

var coordOfVal map[rune]coord = map[rune]coord{
  'A' : {3,2},
  '0' : {3,1},
  '1' : {2,0},
  '2' : {2,1},
  '3' : {2,2},
  '4' : {1,0},
  '5' : {1,1},
  '6' : {1,2},
  '7' : {0,0},
  '8' : {0,1},
  '9' : {0,2},
}

var keyPad map[coord]rune = map[coord]rune {
  {3,2} : 'A',
  {3,1} : '0',
  {2,0} : '1',
  {2,1} : '2',
  {2,0} : '3',
  {1,0} : '4',
  {1,1} : '5',
  {1,2} : '6',
  {0,0} : '7',
  {0,1} : '8',
  {0,2} : '9',
}

var coordOfArrow map[rune]coord = map[rune]coord{
  'A' : {0,2},
  '^' : {0,1},
  '<' : {1,0},
  'v' : {1,1},
  '>' : {1,2},
}

var dirPad map[coord]rune = map[coord] rune {
  {0,2} : 'A',
  {0,1} : '^',
  {1,0} : '<',
  {1,1} : 'v',
  {1,2} : '>',
}

type key struct {
  seq string
  depth int
}

var cache map[key]int = make(map[key]int)

type coord struct {
  i, j int
}

type KeyPadRobot struct {
  c coord
  r rune
}

func getShortestSequenceTo(start rune, target rune, isNumpad bool) []rune {
  var startCoord, targetCoord coord
  if isNumpad {
    startCoord, targetCoord = coordOfVal[start], coordOfVal[target] 
  } else {
    startCoord, targetCoord = coordOfArrow[start], coordOfArrow[target]
  }
  right, down := targetCoord.j - startCoord.j, targetCoord.i - startCoord.i

  var upDown rune
  var rightLeft rune
  if right < 0 {
    rightLeft = '<'
  } else {
    rightLeft = '>'
  }
  if down < 0 {
    upDown = '^'
  } else {
    upDown = 'v'
  }

  path := make([]rune, 0)
  // All L/R first
  if isLeftFirst(startCoord, targetCoord, isNumpad) {
    for i:=0; i<abs(right); i++ {
      path = append(path, rightLeft)
    }
    for i:=0; i<abs(down); i++ {
      path = append(path, upDown)
    }
    path = append(path, 'A')
    return path
  } else {
    for i := 0; i<abs(down); i++ {
      path = append(path, upDown)
    }
    for i:=0; i<abs(right); i++ {
      path = append(path, rightLeft)
    }
    path = append(path, 'A')
    return path
  }
}

func dfs(k key) (sum int) {
  if v, ok := cache[k]; ok {
    return v
  }
  if k.depth == 0 {
    return len(k.seq)
  }
  prev := 'A'

  paths := make([][]rune, 0)
  for _, c := range k.seq {
    paths = append(paths, getShortestSequenceTo(prev, c, false))
    prev = c
  }
  for _, path := range paths {
    sum += dfs(key{string(path), k.depth-1})
  }

  cache[k] = sum
  return

}

func isLeftFirst(start coord, end coord, isNumpad bool) bool {
  if isNumpad {
    crossGap := (start.i == 3 && end.j == 0) || (start.j == 0 && end.i == 3)
    left := end.j < start.j
    return left != crossGap
  } else {
    crossGap := (start.i == 0 && end.j == 0) || (start.j == 0 && end.i == 0)
    left := end.j < start.j
    return left != crossGap
  }
}

func abs(n int) int {
  if n < 0 {
    return -1 * n
  } else {
    return n
  }
}

func isUpFirstValidDirPad(start, end coord) bool {
  return !(start.j == 0 && end.i == 0)
}

func isLeftFirstValidDirPad(start, end coord) bool {
  return !(start.i == 0 && end.j == 0)
}

func Solve(codes []string, depth int) (retval int) {
  for _, code := range codes {
    num, err := strconv.ParseInt(code[:len(code)-1], 10, 0)
    if err != nil {
      log.Fatal("Couldn't parse number in code", code)
    }

    path := make([][]rune, 0)
    prev := 'A'
    for _, c := range code {
      path = append(path, getShortestSequenceTo(prev, c, true))
      prev = c
    }

    length := 0
    for _, p := range path {
      length += dfs(key{string(p), depth})
    }

    retval += int(num) * length
  }
  return
}

