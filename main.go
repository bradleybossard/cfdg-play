package main

import (
  "text/template"
  "bytes"
  "os"
  "fmt"
)

type CfdgParams struct {
  Rotation1 float32
}

// Open file, close on program exit and write complete template out
func writeTmplString(filename string, output string) {

  fo, err := os.Create("output/" + filename) 
  if err != nil { panic(err) }

  fo.WriteString(output)

  err = fo.Close();
  if err != nil { panic(err) }
}

func main() {
  // Read template
  tmpl,err := template.ParseFiles("./tmpl/circle-branch.cfdg")
  if err != nil { panic(err) }

  // Create instance of struct and buffer to write template out to
  params := CfdgParams{1.0}

  for i := 0; i <= 10; i++ {
    var tmplBuf bytes.Buffer;
    params.Rotation1 = float32(i) * 0.2

    // Execute template and put contents in a string
    err = tmpl.Execute(&tmplBuf, params)
    if err != nil { panic(err) }
    tmplString := tmplBuf.String() 

    outputName := fmt.Sprintf("test%05d.cfdg", i)
    writeTmplString(outputName, tmplString) 
  }


}
