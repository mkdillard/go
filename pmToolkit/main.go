package main

import (
  "flag"
  "io"
  "net/http"

  "github.com/mkdillard/pmAdventureGenerator/adventureGenerator"
)

func main() {
  var numAdventures int
  flag.IntVar(&numAdventures,
              "numAdventures",
              1,
              "The number of Planet Mercenary adventures to generate.")
  flag.Parse()

  mux := http.NewServeMux()
  mux.HandleFunc("/generateAdventure", generateAdventure)
  http.ListenAndServe(":8000", mux)

  /*var newAdv adventureGenerator.Adventure

  fmt.Println("Generating " + strconv.Itoa(numAdventures) + " random adventures.\n")
  for i := 1; i <= numAdventures; i++ {
    newAdv = adventureGenerator.GenerateAdventure()
    fmt.Println("ADVENTURE " + strconv.Itoa(i) + ":")
    adventureGenerator.PrintAdventure(newAdv)
  }
  */
}

func generateAdventure(w http.ResponseWriter, r *http.Request) {
  newAdv := new(adventureGenerator.Adventure)
  io.WriteString(w, "ADVENTURE:\n")
  io.WriteString(w, "  EMPLOYER: " + newAdv.Employer + "\n")
  io.WriteString(w, "  JOB: " + newAdv.Job + "\n")
  io.WriteString(w, "  TARGET: " + newAdv.Target + "\n")
  io.WriteString(w, "  TWIST: " + newAdv.Twist + "\n")
  io.WriteString(w, "\n")
}
