package robot

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
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

type coord struct {
  i, j int
}

type KeyPadRobot struct {
  c coord
  r rune
}

func GetAllShortestSequences(code string) []string {
  robot := KeyPadRobot{coordOfVal['A'], 'A'}
  retVal := []string{""}

  for _, target := range code {
    nextRetVal := make([]string, 0)
    extras := robot.getShortestSequences(target)
    for _, existing := range retVal {
      for _, extra := range extras {
        nextRetVal = append(nextRetVal, existing+extra)
      }
    }
    retVal = nextRetVal
  }

  return retVal
}

func (robot *KeyPadRobot) getShortestSequences(target rune) []string {
  retVal := make([]string, 0)
  targetCoord := coordOfVal[target]
  right, down := targetCoord.j - robot.c.j, targetCoord.i - robot.c.i

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

  // All L/R first
  if isLeftFirstValid(robot.c, targetCoord) {
    var rlBuilder strings.Builder
    for i:=0; i<abs(right); i++ {
      rlBuilder.WriteRune(rightLeft)
    }
    for i:=0; i<abs(down); i++ {
      rlBuilder.WriteRune(upDown)
    }
    rlBuilder.WriteRune('A')
    retVal = append(retVal, rlBuilder.String())
  }

  // All U/D first
  if isDownFirstValid(robot.c, targetCoord) {
    var udBuilder strings.Builder
    for i := 0; i<abs(down); i++ {
      udBuilder.WriteRune(upDown)
    }
    for i:=0; i<abs(right); i++ {
      udBuilder.WriteRune(rightLeft)
    }
    udBuilder.WriteRune('A')
    ud := udBuilder.String()
    if len(retVal) < 1 || retVal[0] != ud {
      retVal = append(retVal, ud)
    }
  }
  robot.c = targetCoord
  robot.r = target
  return retVal
}

func isLeftFirstValid(start, end coord) bool {
  return !(start.i == 3 && end.j == 0)
}

func isDownFirstValid(start, end coord) bool {
  return !(start.j == 0 && end.i == 3)
}

func abs(n int) int {
  if n < 0 {
    return -1 * n
  } else {
    return n
  }
}

func getShortestDpadInputLengthForCode(code string, intermediateDpads int) int {
  inputs := GetAllShortestSequences(code)

  for i:=0; i<intermediateDpads; i++ {
    fmt.Println(len(inputs[0]))
    inputs = getAllShortestDpadSequences(inputs[0:1])
  }
  return len(inputs[0])
}

func getAllShortestDpadSequences(shortestOutputs []string) []string {
  var retVal []string
  shortestLength := math.MaxInt
  for _, output := range shortestOutputs {
    inputs := getAllShortestDistancesFirstDirPad(output)
    
    if len(inputs[0]) < shortestLength {
      retVal = inputs
      shortestLength = len(inputs[0])
    } else if len(inputs[0]) == shortestLength {
      retVal = append(retVal, inputs...)
    }
  }
  return retVal
}

func getAllShortestDistancesFirstDirPad(buttonSequence string) []string {
  robot := KeyPadRobot{coordOfArrow['A'], 'A'}
  retVal := []string{""}

  for _, target := range buttonSequence {
    nextRetVal := make([]string, 0)
    extras := robot.getShortestDirPadSequences(target)
    for _, existing := range retVal {
      for _, extra := range extras {
        nextRetVal = append(nextRetVal, existing+extra)
      }
    }
    retVal = nextRetVal
  }
  
  return retVal
}

func (robot *KeyPadRobot) getShortestDirPadSequences(target rune) []string {
  retVal := make([]string, 0)
  targetCoord := coordOfArrow[target]
  upDown, leftRight := getHorizontalAndVerticalButtons(robot.c, targetCoord)

  if isUpFirstValidDirPad(robot.c, targetCoord) {
    var udBuilder strings.Builder
    if robot.c.i != targetCoord.i {
      udBuilder.WriteRune(upDown)
    }
    for i := 0; i<abs(targetCoord.j - robot.c.j); i++ {
      udBuilder.WriteRune(leftRight)
    }
    udBuilder.WriteRune('A')
    retVal = append(retVal, udBuilder.String())
  }
  if isLeftFirstValidDirPad(robot.c ,targetCoord) {
    var lrBuilder strings.Builder
    for i:=0; i<abs(targetCoord.j - robot.c.j); i++ {
      lrBuilder.WriteRune(leftRight)
    }
    if robot.c.i != targetCoord.i {
      lrBuilder.WriteRune(upDown)
    }
    lrBuilder.WriteRune('A')
    lr := lrBuilder.String()
    if len(retVal) < 1 || retVal[0] != lr {
      retVal = append(retVal, lr)
    }
  }
  robot.c = targetCoord
  robot.r = target
  return retVal
}

func getHorizontalAndVerticalButtons(start, end coord) (upDown, leftRight rune) {
  if end.j > start.j {
    leftRight = '>'
  } else {
    leftRight = '<'
  }
  if end.i > start.i {
    upDown = 'v'
  } else {
    upDown = '^'
  }
  return
}

func isUpFirstValidDirPad(start, end coord) bool {
  return !(start.j == 0 && end.i == 0)
}

func isLeftFirstValidDirPad(start, end coord) bool {
  return !(start.i == 0 && end.j == 0)
}

func Solve(codes []string, intermediateDpads int) (sum int) {
  for _, code := range codes {
    num, err := strconv.ParseInt(code[:len(code)-1], 10, 0)
    if err != nil {
      log.Fatal("Badly formatted code:", code)
    }
    inputLength := getShortestDpadInputLengthForCode(code, intermediateDpads)
    sum += inputLength*int(num)
  }

  return
}
