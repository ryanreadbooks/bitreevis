package bitreevis

import (
	"io"
	"math"
)

// Default settings for the graphic
const (
	DefaultBackgroundColor    = "white"
	DefaultNodeColor          = "#868383"
	DefaultNodeStrokeWidth    = 1
	DefaultNodeFieldTextSize  = 16
	DefaultNodeFieldTextColor = "black"

	DefaultEdegLineWidth = 2
	DefaultEdgeColor     = "black"
	DefaultEdgeArrowSize = 2
	DefaultEdgeLineWidth = 2
)

// RenderResult contains rendered output from renderer.
//
// RenderResult should contain a reader which can be read (for example, bytes, string, etc.) from the outsize.
// Content() method returns a io.Reader which provides access to the rendered data.
//
// Save method can save the rendered result to the given filepath.
//
// Error method provides a way for the user to check the error occurred during the rendering process.
// If no error occurs during rendering, Error should return nil.
type RenderResult interface {
	// GetContent return an io.Reader which can be used to access the rendered data.
	GetContent() io.Reader

	// Save helps save the rendered data into a given file.
	Save(string) error

	// Error returns the error occurred during rendering. If no error occurs, it returns nil.
	Error() error
}

// Render is the interface which defines how to render the binary tree.
type Renderer interface {
	Render(*PlaceableNode, RenderOption) RenderResult
}

// RenderOption is the options of graphics when rendering.
//
// All units related to position and size are pixel(px).
type RenderOption struct {
	// BackgroundColor specifies the global background color of the whole graphic.
	BackgroundColor string
	// BackgroundColor specifies the horizontal padding of the graphic on one size.
	HorizontalPadding int
	// BackgroundColor specifies the vertical padding of the graphic on one size.
	VerticalPadding int

	// SiblingSeparation specifies the minimum gap between two sibling nodes.
	SiblingSeparation int
	// LevelSeparation specifies the gap between two different levels.
	LevelSeparation int

	// NodeRadius specifies the radius of node.
	NodeRadius int
	// NodeColor specifies the color of nodes.
	NodeColor string
	// NodeColor specifies the color of leaf nodes, if not specified, NodeColor is used.
	NodeLeafColor string
	// NodeStrokeColor specifies the stroke color of nodes.
	NodeStrokeColor string
	// NodeStrokeWidth specifies the stroke-width of nodes.
	NodeStrokeWidth int
	// NodeFieldTextSize specifies the font size inside of node.
	NodeFieldTextSize int
	// NodeFieldTextColor specifies the color of font inside of node.
	NodeFieldTextColor string

	// EdgeLineWidth specifies the width of edges which connects nodes.
	EdgeLineWidth int
	// EdgeLineWidth specifies the color of edges which connects nodes.
	EdgeLineColor string
	// EdgeWithArrow specifies whether to use arrow at the end of edge
	EdgeWithArrow bool
	// EdgeArrowSize specifies the arrow size of edge
	EdgeArrowSize int
}

// measureEdgeStartEnd is a helper function for calculating the start and end coordinate of an edge
// which connecting two nodes (parent node and child node).
//
// Specifically, (x1, y1) is the parent node, (x2, y2) is the child node.
// radius is the size of node; offsetStart specifies how far the edge start coordinate will go forward;
// offsetEnd specifies how far the edge end coordinate will go backward;
func measureEdgeStartEnd(x1, y1, x2, y2, radius float64, offsetStart, offsetEnd float64) (edgeStartX, edgeStartY, edgeEndX, edgeEndY float64) {
	distance := calDistanceBetweenPoints(x1, y1, x2, y2)
	edgeLength := distance - 2*radius
	// calculate the start coordinate and end coordinate of edge
	edgeDirectionXRaw := x2 - x1
	edgeDirectionYRaw := y2 - y1
	norm := math.Sqrt(math.Pow(edgeDirectionXRaw, 2) + math.Pow(edgeDirectionYRaw, 2))
	// normalization
	edgeDirectionX := edgeDirectionXRaw / norm
	edgeDirectionY := edgeDirectionYRaw / norm

	// edge starts at (edgeStartX, edgeStartY)
	edgeStartX = x1 + edgeDirectionX*(radius+offsetStart)
	edgeStartY = y1 + edgeDirectionY*(radius+offsetStart)
	// edge ends at (edgeEndX, edgeEndY)
	finalLength := edgeLength - offsetStart - offsetEnd
	edgeEndX = edgeStartX + edgeDirectionX*finalLength
	edgeEndY = edgeStartY + edgeDirectionY*finalLength

	return
}

// calDistanceBetweenPoints calculates the distance between point(x1,y1) and point(x2,y2)
func calDistanceBetweenPoints(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}
