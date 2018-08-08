package dicebag

import (
  "math/rand"
  "time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func rollTest(testType string) string{
  result := ""
  if testType == "test" {
    result = rollTest(2,6)
  } else if testType == "disadvantage" {
    result = rollTest(1,6)
  } else if testType == "advantage" {
    result = rollTest(3,6)
  }
  return result
}

func rollDice(numDice int, dieType int) string{
  numRolls := 0
  numSuccess := 0
  for numRolls < numDice {
      roll := 1 + r.Intn(6)
      if roll > 4 {
        numSuccess++
      }
      numRolls++
  }
  if numSuccess > 0 {
    return "Test Succeeded!"
  }
  return "Test Failed!"
}
