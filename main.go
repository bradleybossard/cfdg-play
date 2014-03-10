package main

import (
  "text/template"
  "os"
  //"fmt"
)

type CfdgParams struct {
  Rotation1 float32
}


func main() {
  params := CfdgParams{1.0}
  tmpl,err := template.ParseFiles("./tmpl/circle-branch.cfdg")
  err = tmpl.Execute(os.Stdout, params)
  if err != nil { panic(err) }
}
