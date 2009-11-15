var graph1;
var board;

function makeGraph(pack) {
  dataSeries = pack.series;
  plotRegression = false;
  if (pack.regressionLine) {
    plotRegression = true;
    regressionLine = pack.regressionLine;
  }
  
  var i, x1, y1;
  var p;
  var points = [];
  var x = [];
  var y = [];
  var start = 0;
  var end = dataSeries.length;


  var xmin=0, xmax=0, ymin=0, ymax=0; 
  for (i=start;i<end;i++) {
    // todo: there's a faster way to do this...
    if (dataSeries[i].x < xmin) {
      xmin = dataSeries[i].x;
    } else if (dataSeries[i].x > xmax) {
      xmax = dataSeries[i].x;
    }
    if (dataSeries[i].y < ymin) {
      ymin = dataSeries[i].y;
    } else if (dataSeries[i].y > ymax) {
      ymax = dataSeries[i].y;
    }
  }

  brd = JXG.JSXGraph.initBoard('jxgbox', {boundingbox: [xmin - 4, ymax + 4, xmax + 4, ymin - 4], axis: true, showNavigation: true});
  brd.suspendUpdate();

  points.push(brd.createElement('point', [xmin,0], {visible:false, name:'', fixed:true}));
  for (i=start;i<end;i++) {

    x1 = dataSeries[i].x;
    y1 = dataSeries[i].y;

    // Plot it
    p = brd.createElement('point', [x1,y1], 
                  {strokeWidth:2, strokeColor:'#ffffff', 
                   highlightStrokeColor:'#0077cc', fillColor:'#0077cc',  
                   highlightFillColor:'#0077cc', style:6, name:'', fixed:true}
                ); 
    points.push(p);
    x.push(x1);
    y.push(y1);
  }
  // Filled area. We need two additional points [start,0] and [end,0]
  points.push(brd.createElement('point', [xmax,0], {visible:false, name:'', fixed:true}));
  brd.createElement('polygon',points, {withLines:false,fillColor:'#e6f2fa'});
 
  // Curve:
  brd.createElement('curve', [x,y], 
                 {strokeWidth:3, strokeColor:'#0077cc', 
                  highlightStrokeColor:'#0077cc'}
               );
  
  if (plotRegression) {
    // Regression line
    var rx=[];
    var ry=[];
    // left side
    rx.push(0); // at the y-intercept
    ry.push(regressionLine.intercept);
    // right side
    rx.push(ymax);
    ry.push(regressionLine.slope * xmax + regressionLine.intercept);  // y = mx + b
    // plot it
    brd.createElement('curve', [rx,ry], 
                 {strokeWidth:3, strokeColor:'#eeaacc', 
                  highlightStrokeColor:'#eeaacc'}
               );
  }

  brd.unsuspendUpdate();
  
  return brd;
}

function updateChart(data, textStatus) {
  JXG.JSXGraph.freeBoard(board);
  makeGraph(data);
  return false;
}