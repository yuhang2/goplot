package main

import (
	"./file";
  "flag";
  "os";
  "io";
  "fmt";
  "bufio";
  "bytes";
)

// Plan:
// - read a data file
// - create path data
// - write out the path in an svg file

func main() {
  
  sourceFileName := flag.String("i","source.txt","Source data file name");
  destFileName := flag.String("o","viz.svg","Output file name");

  f, err := file.Open(*sourceFileName, 0, 0);
  if f == nil {
      fmt.Fprintf(os.Stderr, "can't open %s: error %s\n", sourceFileName, err);
      os.Exit(1);
  } defer f.Close();

  fmt.Println(*sourceFileName);
  fmt.Println(*destFileName);

   //read the data
   //f is a file.File*
  //var x float;
  //var y float;
  const NBUF = 512;
  var buf [NBUF]byte;


  cont := true;
  for cont {
    // reads into buf, gets number of bytes & error
    nr, er := f.Read(&buf);
    switch {
    case nr < 0:
      fmt.Fprintf(os.Stderr, "cat: error reading from %s: %s\n", f.String(), er.String());
      os.Exit(1);
    case nr == 0:  // EOF
      fmt.Fprintf(os.Stderr, "done\n");
      cont = false;
    case nr > 0:
      nw, ew := file.Stdout.Write(buf[0:nr]);
      if nw != nr {
        fmt.Fprintf(os.Stderr, "cat: error writing from %s: %s\n", f.String(), ew.String());
      }
    }
  }


   
  fmt.Println(*destFileName);

   
   var rd *bufio.Reader;
   rd = bufio.NewReader(f);
   line:="";
   if err != nil {
    line, err = rd.ReadString('\n');
    fmt.Printf("line %s.",line);
   }
   
  var sourceData []byte;
  sourceData, err = io.ReadFile(*sourceFileName);
  if err != nil {
    errStr := err.String();
    fmt.Fprintf(os.Stderr, "failed to read %s: %s\n", *sourceFileName, errStr);
  }
  sourceBuf := bytes.NewBuffer(sourceData);
  fmt.Println(sourceBuf.String());

}



