package dicebag

import (
  "math/rand"
  "strconv"
  "time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func RollTest(testType string) string{
  result := ""
  if testType == "test" {
    result = rollDice(2,6)
  } else if testType == "disadvantage" {
    result = rollDice(1,6)
  } else if testType == "advantage" {
    result = rollDice(3,6)
  }
  return result
}

func rollDice(numDice int, dieType int) string{
  numRolls := 0
  numSuccess := 0
  result := "("
  for numRolls < numDice {
      numRolls++
      roll := 1 + r.Intn(6)
      result = result + strconv.Itoa(roll)
      if roll > 4 {
        numSuccess++
      }
      if numRolls != numDice {
        result = result + ", "
      } else {
        result = result + ")"
      }
  }
  if numSuccess > 0 {
    return result + " Test Succeeded!"
  }
  return result + " Test Failed!"
}
