package main

import (
  "os";
  "io";
  "fmt";
  "strings";
  "strconv";
  "container/vector";
  "math";
  "http";
)

type Point struct {
  x float;
  y float;
}

func main() {
  http.Handle("/goplot/viz", http.HandlerFunc(dataSampleServer));
  // serve our own files instead of using http.FileServer for very tight access control
  http.Handle("/goplot/graph.js", http.HandlerFunc(fileServe));
  // in order
  err := http.ListenAndServe("192.168.3.235:6060", nil); // todo: clearly this needs to be detected/configured
  if err != nil {
    panic("ListenAndServe: ", err.String())
  }
}

// serve static files as appropriate
func fileServe(c *http.Conn, req *http.Request) {
  cwd, err := os.Getwd();
  if err==nil {
    http.ServeFile(c, req, cwd + "/client/graph.js");
  } else {
    serveError(c, req, http.StatusInternalServerError); // 500
  }
}

// Send the given error code.
func serveError(c *http.Conn, req *http.Request, code int) {
    c.SetHeader("Content-Type", "text/plain; charset=utf-8");
    c.WriteHeader(code);
    io.WriteString(c, fmt.Sprintf("%d\n",code));
}



// processes data samples, sends back data to plot along with regression lines
func dataSampleServer(c *http.Conn, req *http.Request) {
  switch req.Method {
    case "GET":
      cwd, err := os.Getwd();
      if err==nil {
        http.ServeFile(c, req, cwd + "/client/viz.html");
      } else {
        serveError(c, req, http.StatusInternalServerError); // 500
      }
    case "POST":
      src := req.FormValue("dataseries");
      result := dataSampleProcess(src);
      // send the response
      _,_=io.WriteString(c, result);
    default :
      serveError(c, req, http.StatusMethodNotAllowed);
  }
}

// processes data samples, sends back data to plot along with regression lines
func dataSampleProcess(src string) (results string) {
  const MAXLINES = 1000000;
  
  // split the buffer into an array of strings, one per source line
  srcLines := strings.Split(src,"\n",MAXLINES);

  lineCount := len(srcLines);
  series := vector.New(0);

  for ix:=0; ix < lineCount; ix++ {
    stmp , err := parseLine(srcLines[ix]);
    if err == nil {
      series.Push(stmp);
    }
  }
  fmt.Println(series);
  jsonStr:="{series:[";
  for ix:=0; ix < series.Len(); ix++ {
    fmt.Println(series.At(ix));
    jsonStr += "{x:" + strconv.Ftoa(series.At(ix).(Point).x,'f',3) + ",y:" + strconv.Ftoa(series.At(ix).(Point).y,'f',3) + "},";
  }
  jsonStr += "],\n";
  
  slope, intercept, stdError, correlation := linearRegression(series);
  jsonStr += fmt.Sprintf("regressionLine:{slope:%f,intercept:%f,stdError:%f,correlation:%f},",slope, intercept, stdError, correlation);
  jsonStr += "}";
  
  return jsonStr;
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