package dicebag

import (
  "regexp"
  "strconv"
  "testing"
)

func TestRollTest(t *testing.T){
  disadvantageTest := "disadvantage"
  normalTest := "test"
  advantageTest := "advantage"

  ExpectedDisadvantagePass := regexp.MustCompile("\\([1-4]\\) Test Failed!")
  ExpectedDisadvantageFail := regexp.MustCompile("\\([5-6]\\) Test Succeeded!")
  ExpectedNormalPass := regexp.MustCompile("\\([1-4], [1-4]\\) Test Failed!")
  ExpectedNormalFail := regexp.MustCompile("\\([1-6], [1-6]\\) Test Succeeded!")
  ExpectedAdvantagePass := regexp.MustCompile("\\([1-4], [1-4], [1-4]\\) Test Failed!")
  ExpectedAdvantageFail := regexp.MustCompile("\\([1-6], [1-6], [1-6]\\) Test Succeeded!")

  disadvantageResult := RollTest(disadvantageTest)
  normalResult := RollTest(normalTest)
  advantageResult := RollTest(advantageTest)

  if !ExpectedDisadvantageFail.MatchString(disadvantageResult) {
    if !ExpectedDisadvantagePass.MatchString(disadvantageResult) {
      t.Errorf("Disadvantage Roll incorrect, got %s, want: %s.", disadvantageResult, "([1-4]) Test Failed! || ([5-6]) Test Succeeded!")
    }
  }

  if !ExpectedNormalFail.MatchString(normalResult) {
    if !ExpectedNormalPass.MatchString(normalResult) {
      t.Errorf("Normal Roll incorrect, got %s, want: %s.", normalResult, "([1-4], [1-4]) Test Failed! || (<at least one 5 or 6>) Test Succeeded!")
    }
  }

  if !ExpectedAdvantageFail.MatchString(advantageResult) {
    if !ExpectedAdvantagePass.MatchString(advantageResult) {
      t.Errorf("Advantage Roll incorrect, got %s, want: %s.", advantageResult, "([1-4], [1-4], [1-4]) Test Failed! || (<at least one 5 or 6>) Test Succeeded!")
    }
  }
}

func TestHurthaansWill(t *testing.T) {
    passResult := "Hurthaan Approves!"
    failResult := "Hurthaan Disapproves!"

    result := HurthaansWill()

    if passResult != result {
      if failResult != result {
        t.Errorf("Unexpected result from HurthaansWill\ngot %s\nwant %s or %s",result, passResult, failResult)
      }
    }
}

func TestGenericRoll(t *testing.T) {
  var tests = []struct {
    input string
    expected *regexp.Regexp
  }{
    {"1d2", regexp.MustCompile("\\([1-2]\\) = [1-2]")},
    {"1d3", regexp.MustCompile("\\([1-3]\\) = [1-3]")},
    {"1d4", regexp.MustCompile("\\([1-4]\\) = [1-4]")},
    {"1d6", regexp.MustCompile("\\([1-6]\\) = [1-6]")},
    {"1d8", regexp.MustCompile("\\([1-8]\\) = [1-8]")},
    {"1d10", regexp.MustCompile("\\([0-9]+\\) = [0-9]+")},
    {"1d12", regexp.MustCompile("\\([0-9]+\\) = [0-9]+")},
    {"1d20", regexp.MustCompile("\\([0-9]+\\) = [0-9]+")},
    {"1d100", regexp.MustCompile("\\([0-9]+\\) = [0-9]+")},
    {"2d10", regexp.MustCompile("\\([0-9]+, [0-9]+\\) = [0-9]+")},
    {"3d6", regexp.MustCompile("\\([1-6], [1-6], [1-6]\\) = [0-9]+")},
  }

  for _, test := range tests {
    result := GenericRoll(test.input)
    if !test.expected.MatchString(result) {
      t.Errorf("GenericRoll test failed, result: %s expected: %s\n", result, test.expected.String())
    }
  }
}

func TestRollTestDice(t *testing.T){
  var tests = []struct{
    numDice int
    dieType int
    expectedFail *regexp.Regexp
    expectedPass *regexp.Regexp
  }{
    {1,6, regexp.MustCompile("\\([1-4]\\) Test Failed!"), regexp.MustCompile("\\([5-6]\\) Test Succeeded!")},
    {2,6, regexp.MustCompile("\\([1-4], [1-4]\\) Test Failed!"), regexp.MustCompile("\\([1-6], [1-6]\\) Test Succeeded!")},
    {3,6, regexp.MustCompile("\\([1-4], [1-4], [1-4]\\) Test Failed!"), regexp.MustCompile("\\([1-6], [1-6], [1-6]\\) Test Succeeded!")},
  }
  for _, test := range tests{
    result := rollTestDice(test.numDice, test.dieType)
    if !test.expectedFail.MatchString(result) {
      if !test.expectedPass.MatchString(result) {
        t.Errorf("rollTestDice test failed, result: %s expected: %s or %s\n", result, test.expectedFail.String(), test.expectedPass.String())
      }
    }
  }
}

func TestTotalMultiRollString(t *testing.T){
  var tests = []struct{
    input string
    expected int
  }{
    {"(3)", 3},
    {"(1, 2, 3)", 6},
    {"(4, 2, 6)", 12},
  }

  for _, test := range tests{
    result := totalMultiRollString(test.input)
    if result != test.expected{
      t.Errorf("totalMultiRollString unexpected result, got: %d, want: %d\n", result, test.expected)
    }
  }
}

func TestMultiRollString(t *testing.T){
  var tests = []struct{
    numDice int
    dieType int
    expected *regexp.Regexp
  }{
    {1,6,regexp.MustCompile("\\([1-6]\\)")},
    {2,6,regexp.MustCompile("\\([1-6], [1-6]\\)")},
    {3,6,regexp.MustCompile("\\([1-6], [1-6], [1-6]\\)")},
    {1,20,regexp.MustCompile("\\([0-9]+\\)")},
    {4,10,regexp.MustCompile("\\([0-9]+, [0-9]+, [0-9]+, [0-9]+\\)")},
  }

  for _, test := range tests{
    result := multiRollString(test.numDice, test.dieType)
    if !test.expected.MatchString(result) {
      t.Errorf("multiRollString unexpected result, got: %s, want: %s\n", result, test.expected.String())
    }
  }
}

func TestSingleRoll(t *testing.T){
  var tests = []struct{
    input int
    expected *regexp.Regexp
  }{
    {2,regexp.MustCompile("[1-2]")},
    {3,regexp.MustCompile("[1-3]")},
    {6,regexp.MustCompile("[1-6]")},
    {8,regexp.MustCompile("[1-8]")},
    {20,regexp.MustCompile("[0-9]+")},
    {100,regexp.MustCompile("[0-9]+")},
  }

  for _, test := range tests{
    result := singleRoll(test.input)
    if !test.expected.MatchString(strconv.Itoa(result)) {
      t.Errorf("singleRoll unexpected result, got: %d, want: %s", result, test.expected.String())
    }
  }
}
