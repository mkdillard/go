package pmAdventureGenerator

import (
  "fmt"
  "math/rand"
  "time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// Adventure is a struct that contains the information about the adventure.
type Adventure struct {
  Employer string
  Job string
  Resources string
  Target string
  Twist string
}

//NewAdventure returns a new random adventure.
func NewAdventure() Adventure{
  employer := determineEmployer()
  job, resources := determineJob()
  target := determineTarget()
  twist := determineTwist()
  return Adventure{
    Employer: employer,
    Job: job,
    Resources: resources,
    Target: target,
    Twist: twist,
  }
}

//NewAdventurePointer returns a pointer to a new random adventure.
func NewAdventurePointer() *Adventure{
  adv := NewAdventure()
  return &adv
}

var employerMap = map[int]string{
  1: "Business",
  2: "Government",
  3: "Charity",
  4: "Individual",
  5: "Religious Institution",
}

var jobMap = map[int]string{
  2: "Blackmail",
  3: "Protection",
  4: "Heist",
  5: "Extortion",
  6: "Military Action",
  7: "Inciting Revolution",
  8: "Transporting Goods",
  9: "Elimination",
  10: "Retrieval",
  11: "Policing",
}

var resourceMap = map[int]string{
  2: "2-3",
  3: "4-6",
  4: "2-6",
  5: "1-5",
  6: "4-8",
  7: "3-6",
  8: "1-3",
  9: "2-6",
  10: "4-6",
  11: "1-2",
}

var targetMap = map[int]string{
  2: "Interplanetary Church",
  3: "Rival Business",
  4: "Politician",
  5: "Street Gang",
  6: "Rogue AI",
  7: "Crime Syndicate",
  8: "Rival Mercenary Company",
  9: "\"Empty\" area of space",
  10: "Planetary Government",
  11: "System Government",
}

var twistMap = map[int]string{
  1: "Traitor in the mercenaries' midst.",
  2: "Wrong target.",
  3: "Employer double cross.",
  4: "Misinformation.",
  5: "Rivals swoop in and take target.",
  6: "Target has a redundancy plan.",
}

// GenerateNewAdventure populates and adventure struct and returns it.
// The struct is populated following the rules lain out in the
// Planet Mercenary core rule book.
func (a *Adventure) GenerateNewAdventure() {
  a.Employer = determineEmployer()
  a.Job, a.Resources = determineJob()
  a.Target = determineTarget()
  a.Twist = determineTwist()
}

// GenerateNewEmployer generates a new random employer, but does not change
// any of the other properties of the current adventure.
func (a *Adventure) GenerateNewEmployer() {
  a.Employer = determineEmployer()
}

// GenerateNewJob generates a new random job, but does not change
// any of the other properties of the current adventure.
func (a *Adventure) GenerateNewJob() {
  a.Job, a.Resources = determineJob()
}

// GenerateNewTarget generates a new random target, but does not change
// any of the other properties of the current adventure.
func (a *Adventure) GenerateNewTarget() {
  a.Target = determineTarget()
}

// GenerateNewTwist generates a new random twist, but does not change
// any of the other properties of the current adventure.
func (a *Adventure) GenerateNewTwist() {
  a.Twist = determineTwist()
}


//PrintAdventure prints the fields of an adventure object.
func (a *Adventure) PrintAdventure() {
  fmt.Println("  EMPLOYER: " + a.Employer)
  fmt.Println("  JOB: " + a.Job)
  fmt.Println("  Resources: " + a.Resources)
  fmt.Println("  TARGET: " + a.Target)
  fmt.Println("  TWIST: " + a.Twist)
  fmt.Println()
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

func determineJob() (string, string) {
  job := ""
  resources := ""
  roll := 2 + r.Intn(6) + rand.Intn(6)
  if (roll == 12) {
    newroll1 := 2 + r.Intn(6) + rand.Intn(6)
    newroll2 := 2 + r.Intn(6) + rand.Intn(6)
    if (newroll1 != 12 && newroll2 != 12) {
      job = jobMap[newroll1] + ", " + jobMap[newroll2]
      resources = resourceMap[newroll1] + ", " + resourceMap[newroll2]
    } else if (newroll1 != 12) {
      job = jobMap[newroll1]
      resources = resourceMap[newroll1]
    } else if (newroll2 != 12) {
      job = jobMap[newroll2]
      resources = resourceMap[newroll2]
    } else {
      //Call function recursively.
      //It's possible (unlikely, but possible) we rolled all 6's
      job, resources = determineJob()
    }
  } else {
    job = jobMap[roll]
    resources = resourceMap[roll]
  }
  return job, resources
}

func determineTarget() string {
  target := ""
  roll := 2 + r.Intn(6) + rand.Intn(6)
  if (roll == 12) {
    newroll1 := 2 + r.Intn(6) + rand.Intn(6)
    newroll2 := 2 + r.Intn(6) + rand.Intn(6)
    if (newroll1 != 12 && newroll2 != 12) {
      target = targetMap[newroll1] + ", " + targetMap[newroll2]
    } else if (newroll1 != 12) {
      target = targetMap[newroll1]
    } else if (newroll2 != 12) {
      target = targetMap[newroll2]
    } else {
      //Call function recursively.
      //It's possible (unlikely, but possible) we rolled all 6's
      target = determineTarget()
    }
  } else {
    target = targetMap[roll]
  }
  return target
}

func determineTwist() string {
  twist := ""
  roll := 1 + r.Intn(6)
  twist = twistMap[roll]
  return twist
}
