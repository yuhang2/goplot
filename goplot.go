package main

import (
  "flag";
  "os";
  "io";
  "fmt";
  "bytes";
  "strings";
  "strconv";
  "container/vector";
  "math";
)

type Point struct {
  x float;
  y float;
}

func main() {
  
  const MAXLINES = 1000000;
  
  sourceFileName := flag.String("i","source.txt","Source data file name");
  destFileName := flag.String("o","vizdata.js","Output file name");

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
  series := vector.New(0);

  // need to test for error before saving the value
  stmp := Point{x:0.0, y:0.0};
  for ix:=0; ix < lineCount; ix++ {
    stmp , err = parseLine(srcLines[ix]);
    if err == nil {
      //fmt.Println(ix, stmp);
      series.Push(stmp);
    }
  }
  fmt.Println(series);
  jsonStr:="series=[";
  for ix:=0; ix < series.Len(); ix++ {
    fmt.Println(series.At(ix));
    jsonStr += "{x:" + strconv.Ftoa(series.At(ix).(Point).x,'f',3) + ",y:" + strconv.Ftoa(series.At(ix).(Point).y,'f',3) + "},";
  }
  jsonStr+="];\n";
  
  slope, intercept, stdError, correlation := linearRegression(series);
  jsonStr += fmt.Sprintf("regressionLine={slope:%f,intercept:%f,stdError:%f,correlation:%f};",slope, intercept, stdError, correlation);
  
  _=io.WriteFile(*destFileName, strings.Bytes(jsonStr), 777);
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

// perform linear regression on the data series
// based on Numerical Methods for Engineers, 2nd ed. by Chapra & Canal
func linearRegression(series *vector.Vector) (slope float, intercept float, stdError float, correlation float) {
  len := series.Len();
  flen := float(len); // convenience
  sumx := 0.0;
  sumy := 0.0;
  sumxy := 0.0;
  sumx2 := 0.0;
  for ix:=0; ix < len; ix++ {
    x := series.At(ix).(Point).x;
    y := series.At(ix).(Point).y;
    sumx += x;
    sumy += y;
    sumxy += x*y;
    sumx2 += x*x;
  }
  xmean := sumx / flen;
  ymean := sumy / flen;
  slope = (flen*sumxy - sumx*sumy) / (flen*sumx2 - sumx*sumx);
  intercept = ymean - slope * xmean;
  
  st := 0.0;
  sr := 0.0;
  for ix:=0; ix < len; ix++ {
    x := series.At(ix).(Point).x;
    y := series.At(ix).(Point).y;
    st += (y-ymean)*(y-ymean);
    // guessing the compiler sees this is constant & does sth faster than exponentiation
    sr += (y - (slope*x - intercept)) * (y - (slope*x - intercept));    
  }
  fmt.Println(st,sr);
  stdError = (float)(math.Sqrt((float64)(sr/(flen-2.0)))); // todo: must check that min 2 points are supplied
  correlation = (float)(math.Sqrt((float64)((st-sr)/st)));
  return slope, intercept, stdError, correlation;
}