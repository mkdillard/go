package adventureGenerator

import (
  "math/rand"
  "time"
)

// Adventure is a struct that contains the information about the adventure.
type Adventure struct {
  Employer string
  Job string
  Target string
  Twist string
}

var employerMap = map[int]string{
  1: "Business",
  2: "Government",
  3: "Charity",
  4: "Individual",
  5: "Religious Institution",
}

var jobMap = map[int]string{
  2: "Blackmail \n  RESOURCES: 2-3",
  3: "Protection \n  RESOURCES: 4-6",
  4: "Heist \n  RESOURCES: 2-6",
  5: "Extortion \n  RESOURCES: 1-5",
  6: "Military Action \n  RESOURCES: 4-8",
  7: "Inciting Revolution \n  RESOURCES: 3-6",
  8: "Transporting Goods \n  RESOURCES: 1-3",
  9: "Elimination \n  RESOURCES: 2-6",
  10: "Retrieval \n  RESOURCES: 4-6",
  11: "Policing \n  RESOURCES: 1-2",
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

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenerateAdventure populates and adventure struct and returns it.
// The struct is populated following the rules lain out in the
// Planet Mercenary core rule book.
func GenerateAdventure() Adventure {
  employer := determineEmployer()
  job := determineJob()
  target := determineTarget()
  twist := determineTwist()
  adv := Adventure{Employer: employer, Job: job, Target: target, Twist: twist}
  return adv
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

func determineJob() string {
  job := ""
  roll := 2 + r.Intn(6) + rand.Intn(6)
  if (roll == 12) {
    newroll1 := 2 + r.Intn(6) + rand.Intn(6)
    newroll2 := 2 + r.Intn(6) + rand.Intn(6)
    if (newroll1 != 12 && newroll2 != 12) {
      job = jobMap[newroll1] + ", " + jobMap[newroll2]
    } else if (newroll1 != 12) {
      job = jobMap[newroll1]
    } else if (newroll2 != 12) {
      job = jobMap[newroll2]
    } else {
      //Call function recursively.
      //It's possible (unlikely, but possible) we rolled all 6's
      job = determineJob()
    }
  } else {
    job = jobMap[roll]
  }
  return job
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
