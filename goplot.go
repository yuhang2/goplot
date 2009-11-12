package main

import (
  "flag";
  "os";
  "io";
  "fmt";
  "bytes";
)

// Plan:
// - read a data file
// - create path data
// - write out the path in an svg file

func main() {
  
  sourceFileName := flag.String("i","source.txt","Source data file name");
  destFileName := flag.String("o","viz.svg","Output file name");

  fmt.Println(*sourceFileName);
  fmt.Println(*destFileName);

   //read the data
   //f is a file.File*
  //var x float;
  //var y float;

  sourceData, err := io.ReadFile(*sourceFileName);
  if err != nil {
    errStr := err.String();
    fmt.Fprintf(os.Stderr, "failed to read %s: %s\n", *sourceFileName, errStr);
  }
  sourceBuf := bytes.NewBuffer(sourceData);
  fmt.Println(sourceBuf.String());

}



