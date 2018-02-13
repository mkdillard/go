package main

import (
  "fmt"
  "io"
  //"log"
  "net/http"

  "github.com/mkdillard/pmToolkit/pmAdventureGenerator"
  //"github.com/mkdillard/pmToolkit/pmCompanyGenerator"
)

func main() {

/*
  mux := http.NewServeMux()
  mux.HandleFunc("/generateAdventure", generateAdventure)
  mux.HandleFunc("/generateCompany", generateCompany)



  http.ListenAndServe(":8000", mux)
  */
  myAdv := pmAdventureGenerator.NewAdventure()
  myAdv.PrintAdventure()
  fmt.Println("")
}

func generateAdventure(w http.ResponseWriter, r *http.Request) {
  newAdv := &pmAdventureGenerator.Adventure{}
  io.WriteString(w, "ADVENTURE:\n")
  io.WriteString(w, "  EMPLOYER: " + newAdv.Employer + "\n")
  io.WriteString(w, "  JOB: " + newAdv.Job + "\n")
  io.WriteString(w, "  TARGET: " + newAdv.Target + "\n")
  io.WriteString(w, "  TWIST: " + newAdv.Twist + "\n")
  io.WriteString(w, "\n")
}

/*
func generateCompany(w http.ResponseWriter, r *http.Request) {
  var newCom adventureCompany.Company
  newAdv = adventureGenerator.GenerateAdventure()
  io.WriteString(w, "ADVENTURE:\n")
  io.WriteString(w, "  EMPLOYER: " + newAdv.Employer + "\n")
  io.WriteString(w, "  JOB: " + newAdv.Job + "\n")
  io.WriteString(w, "  TARGET: " + newAdv.Target + "\n")
  io.WriteString(w, "  TWIST: " + newAdv.Twist + "\n")
  io.WriteString(w, "\n")
}
*/
