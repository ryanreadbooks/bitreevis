# bitreevis

This repo is a tool which helps visualize binary tree structure in Golang. It is useful when debugging a program with binary tree structure.

# Prerequisites
* **[Go](https://golang.org/)**

# Installation
With Go module support (Go 1.11+), simply add the following import
```go
import "github.com/ryanreadbooks/bitreevis"
```
to your code. Or you can use the following the get this package.
```bash
$ go get -u github.com/ryanreadbooks/bitreevis
```

# Quick start

Details can be found in [examples/svg_demo.go](examples/svg_demo.go).

### bitreevis.BiNode

In order to visualize your own binary tree, you should implement the `bitreevis.BiNode` interface. Then you can use `bitreevis.VisAsSvg()` function to visualize the binary tree in svg graphic format.

### bitreevis.RenderOption

`bitreevis.RenderOption` is used to define the output style of the visualization. The size of nodes, color of nodes, the width of edges, etc. can be customized by setting option.

# Learn more

* [Tidier Drawings of Trees](https://ieeexplore.ieee.org/document/1702828) algorithm, which is the layout calculation used in bitreevis
* [Examples](examples)

