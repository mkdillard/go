package pmCompanyGenerator

import (
  "math/rand"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

//Company is a struct that contains information relevant to a company
type Company struct {
  AIRating int
  CaptainContract int
  CompanyResources int
  NumberPCs int
  ReputationBonus int
  Resources int
  TenureBonus int
  Reputation string
  Tenure string
}

var tenureMap = map[int]string{
  1: "New",
  2: "Experienced",
  3: "Experienced",
  4: "Experienced",
  5: "Veteran",
  6: "Venerable",
}

var tenureReputationBonus = map[int]int{
  1: 0,
  2: 1,
  3: 1,
  4: 1,
  5: 2,
  6: 4,
}

var reputation = map[int]string{
  1: "Scum and Villany"
  2: "Ignoble Rakes"
  3: "Untrustworthy"
  4: "Neutral"
  5: "Reputed"
  6: "Always Delivers"
  7: "Discreet and Efficient"
  8: "Sector's Best"
  9: "Legends"
  10: "Always Drink for Free"
}

var reputationBonus = map[int]int{
  1: -3
  2: -2
  3: -1
  4: 0
  5: 1
  6: 2
  7: 3
  8: 4
  9: 5
  10: 6
}
// Roll 1d6+3 for each PC (1d6 for AI), total is company resources

func determineResources() int {
  return 0
}

var aiRating = map[int]int{
  1: 2
  2: 5
  3: 10
  4: 20
}

AIRating int
CaptainContract int
CompanyResources int
NumberPCs int
ReputationBonus int
Resources int
TenureBonus int
Reputation string
Tenure string

func (c *Company) setAIRating(airating int) {
  c.AIRating = airating
}

func (c *Company) setCaptainContract(cc int) {
  c.CaptainContract = cc
}

func (c *Company) determineResources() {
  //TODO fill in details
  resources := 0
  aiRatingCost = aiRatingCost[c.AIRating]
  for i := 0; i < c.NumberPCs; i++ {
    resources += 4 + r.Intn(6)
  }
  c.Resources = resources - aiRatingCost
}

func determineEmployer() string {
  employer := ""
  roll := 1 + r.Intn(6)
  if (roll == 6) {
    newroll1 := 1 + r.Intn(6)
    newroll2 := 1 + r.Intn(6)
    if (newroll1 != 6 && newroll2 != 6) {
      employer = employerMap[newroll1] + ", " + employerMap[newroll2]
    } else if (newroll1 != 6) {
      employer = employerMap[newroll1]
    } else if (newroll2 != 6) {
      employer = employerMap[newroll2]
    } else {
      //Call function recursively.
      //It's possible (unlikely, but possible) we rolled all 6's
      employer = determineEmployer()
    }
  } else {
    employer = employerMap[roll]
  }
  return employer
}
