package main

import (
  "flag";
  "os";
  "io";
  "fmt";
  "bytes";
  "strings";
  "strconv";
  //"containers/vector";
)

type Point struct {
  x float;
  y float;
}

func main() {
  
  const MAXLINES = 1000000;
  
  sourceFileName := flag.String("i","source.txt","Source data file name");
  destFileName := flag.String("o","vizdata.js","Output file name");

  //fmt.Println(*sourceFileName);
  //fmt.Println(*destFileName);

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
  //fmt.Println(sourceBuf.String());
  
  // split the buffer into an array of strings, one per source line
  src := sourceBuf.String();
  srcLines := strings.Split(src,"\n",MAXLINES);

  lineCount := len(srcLines);
  // make a slice to hold the point data we read in
  series := make([]Point, lineCount, lineCount);
  
  // if an error is returned we still save the value in the series array
  //for ix:=0; ix < lineCount; ix++ {
  //  series[ix] , err = parseLine(srcLines[ix]);
  //}

  // need to test for error before saving the value
  stmp := Point{x:0.0, y:0.0};
  for ix:=0; ix < lineCount; ix++ {
    stmp , err = parseLine(srcLines[ix]);
    if err == nil {
      //fmt.Println(ix, stmp);
      series[ix] = stmp;
    }
  }


  jsonStr:="series=[";
  for ix:=0; ix < lineCount; ix++ {
    // fmt.Printf("{x:%f,y:%f}",series[ix].x, series[ix].y);
    jsonStr += "{x:" + strconv.Ftoa(series[ix].x,'f',3) + ",y:" + strconv.Ftoa(series[ix].y,'f',3) + "},";
  }
  jsonStr+="];";

  //fmt.Println(jsonStr);
  
  _=io.WriteFile(*destFileName, strings.Bytes(jsonStr), 777);
}

func foo(p Point) {
  
}

func parseLine(coords string) (p Point, err os.Error) {
  if len(coords) > 0 {
    coordsAr := strings.Split(strings.TrimSpace(coords), ",", 3);
    if len(coordsAr) > 1 {
      // ignore conversion errors
      p.x, err = strconv.Atof(coordsAr[0]);
      if err == nil {
        p.y, err = strconv.Atof(coordsAr[1]);
      }
    }
  } else {
    err = os.NewError("parseLine: No data");
  }
  return p, err;
}