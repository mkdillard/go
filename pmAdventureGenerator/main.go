package main

import (
  "flag"
  "fmt"
  "strconv"

  "github.com/mkdillard/pmAdventureGenerator/adventureGenerator"
)

func main() {
  var numAdventures int
  flag.IntVar(&numAdventures,
              "numAdventures",
              1,
              "The number of Planet Mercenary adventures to generate.")
  flag.Parse()

  var newAdv adventureGenerator.Adventure

  fmt.Println("Generating " + strconv.Itoa(numAdventures) + " random adventures.\n")
  for i := 1; i <= numAdventures; i++ {
    newAdv = adventureGenerator.GenerateAdventure()
    fmt.Println("ADVENTURE " + strconv.Itoa(i) + ":")
    fmt.Println("  EMPLOYER: " + newAdv.Employer)
    fmt.Println("  JOB: " + newAdv.Job)
    fmt.Println("  TARGET: " + newAdv.Target)
    fmt.Println("  TWIST: " + newAdv.Twist)
    fmt.Println()
  }
}
