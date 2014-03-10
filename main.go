package main

import (
  "text/template"
  "bytes"
  "os"
)

type CfdgParams struct {
  Rotation1 float32
}


func main() {
  // Read template
  tmpl,err := template.ParseFiles("./tmpl/circle-branch.cfdg")
  if err != nil { panic(err) }

  // Create instance of struct and buffer to write template out to
  params := CfdgParams{1.0}
  var tmplBuf bytes.Buffer;

  // Execute template and put contents in a string
  err = tmpl.Execute(&tmplBuf, params)
  if err != nil { panic(err) }
  tmplString := tmplBuf.String() 

  outputName := "test0.cfdg"
  // Open file, close on program exit and write complete template out
  fo, err := os.Create("output/" + outputName) 
  if err != nil { panic(err) }
  defer func() {
      if err := fo.Close(); err != nil {
          panic(err)
      }
  }()

  fo.WriteString(tmplString)
}
