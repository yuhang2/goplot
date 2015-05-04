GoPlot is a learning vehicle for experimenting in the Go programming language. The direction of the code is toward a graphing utility with some curve-fitting features. So far it's capable of doing a simple linear regression on the source data.

As of version 0.3.0, GoPlot runs as a web server. The one page it serves presents a graph and a textarea for entering data as x,y pairs, one per line. For example:
```
1,2
2,8.5
3,9
4,13
5,13.5
6,20
```
Input is simple value pairs (2D points) and results are provided as JSON. The JSON is interpreted on the client side and plotted in SVG using <a href='http://www.jsxgraph.org/'>jsxGraph</a>.

There's a server occasionally demonstrating this code at <a href='http://viridian.lnpc.net/goplot/viz'>this page</a>.