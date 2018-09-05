package dicebag

import (
  "fmt"
  "math/rand"
  "regexp"
  "strconv"
  "strings"
  "time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func RollTest(testType string) string{
  result := ""
  if testType == "test" {
    result = rollTestDice(2,6)
  } else if testType == "disadvantage" {
    result = rollTestDice(1,6)
  } else if testType == "advantage" {
    result = rollTestDice(3,6)
  } else if testType == "focusTest" {
    result = rollFocusDice(2,6)
  } else if testType == "focusDisadvantage" {
    result = rollFocusDice(1,6)
  } else if testType == "focusAdvantage" {
    result = rollFocusDice(3,6)
  } else if testType == "focusMarksman" {
    result = rollMarksmanDice(2,6)
  } else if testType == "focusMarksmanDisadvantage" {
    result = rollMarksmanDice(1,6)
  } else if testType == "focusMarksmanAdvantage" {
    result = rollMarksmanDice(3,6)
  }
  return result
}

func HurthaansWill() string{
  result := singleRoll(6)
  if result < 4 {
    return "Hurthaan Approves!"
  }
  return "Hurthaan Disapproves!"
}

func GenericRoll(roll string) string{
  inputs := strings.Split(roll, "d")
  numDice, err := strconv.ParseInt(inputs[0], 10, 0)
  if err != nil {
    fmt.Printf("Error getting numDice: %s, %d\n", err, numDice)
  }
  dieType, err := strconv.ParseInt(inputs[1], 10, 0)
  if err != nil {
    fmt.Printf("Error getting dieType: %s, %d\n", err, dieType)
  }
  result := multiRollString(int(numDice), int(dieType))
  total := totalMultiRollString(result)

  result = result + " = " + strconv.Itoa(total)
  return result
}

func rollTestDice(numDice int, dieType int) string{
  numSuccess := 0
  result := multiRollString(numDice, dieType)
  numSuccess += strings.Count(result, "5")
  numSuccess += strings.Count(result, "6")
  if numSuccess > 0 {
    return result + " Test Succeeded!"
  }
  return result + " Test Failed!"
}

func rollFocusDice(numDice int, dieType int) string{
  numSuccess := 0
  result := multiRollString(numDice, dieType)
  numSuccess += strings.Count(result, "4")
  numSuccess += strings.Count(result, "5")
  numSuccess += strings.Count(result, "6")
  if numSuccess > 0 {
    return result + " Test Succeeded!"
  }
  return result + " Test Failed!"
}

func rollMarksmanDice(numDice int, dieType int) string{
  numSuccess := 0
  result := multiRollString(numDice, dieType)
  numSuccess += strings.Count(result, "3")
  numSuccess += strings.Count(result, "4")
  numSuccess += strings.Count(result, "5")
  numSuccess += strings.Count(result, "6")
  if numSuccess > 0 {
    return result + " Test Succeeded!"
  }
  return result + " Test Failed!"
}

func totalMultiRollString(resultString string) int{
  total := 0
  re := regexp.MustCompile("[0-9]+")
  dieResults := re.FindAllString(resultString, -1)

  for _, element := range dieResults {
    value, err := strconv.ParseInt(element, 10, 0)
    if err != nil {
      fmt.Printf("Error, value: %s, %d\n", err, value)
    }
    total += int(value)
  }
  return total
}

func multiRollString(numDice int, dieType int) string{
  numRolls := 0
  result := "("
  for numRolls < numDice {
    numRolls++
    roll := singleRoll(dieType)
    result = result + strconv.Itoa(roll)
    if numRolls != numDice {
      result = result + ", "
    } else {
      result = result + ")"
    }
  }
  return result
}

func singleRoll(dieType int) int{
  return 1 + r.Intn(dieType)
}
