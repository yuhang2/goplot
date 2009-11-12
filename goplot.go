package main

import (
  "flag";
  "os";
  "io";
  "fmt";
  "bytes";
  "strings";
)

// Plan:
// - read a data file
// - create path data
// - write out the path in an svg file

func main() {
  
  const MAXLINES = 1000000;
  
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
  
  // split the buffer into an array of strings, one per source line
  src := sourceBuf.String();
  srcLines := strings.Split(src,"\n",MAXLINES);

  lineCount := len(srcLines);
  plotPath := "";
  if lineCount > 0 {
    plotPath = "M"+srcLines[0];
    for ix:=1 ; ix < lineCount ; ix++ {
      plotPath += "L" + srcLines[ix];
    }
  }
  svgStr := "<?xml version=\"1.0\"?>\n";
  svgStr += "<svg xmlns:xlink=\"http://www.w3.org/1999/xlink\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 100 100\">\n";
  svgStr += "<path stroke-width=\"2\" stroke=\"#BB5511\" fill=\"none\" d=\"" + plotPath + "\"/>\n";
  svgStr += "</svg>\n";

  fmt.Println(svgStr);
  
  _=io.WriteFile(*destFileName, strings.Bytes(svgStr), 777);
}
