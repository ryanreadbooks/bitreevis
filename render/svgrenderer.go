package render

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	svg "github.com/ajstarks/svgo"

	"github.com/ryanreadbooks/bitreevis/layout"
)

type SvgRenderer struct {
	Canvas svg.SVG
	buf    *strings.Builder
}

const (
	selfDefinedArrowName = "self-defined-arrow-marker"
)

// NewSvgRenderer returns a new SvgRenderer.
func NewSvgRenderer() *SvgRenderer {
	r := &SvgRenderer{
		buf: &strings.Builder{},
	}
	r.Canvas = *svg.New(r.buf)

	return r
}

// Render performs rendering process for specified binary tree.
func (sr *SvgRenderer) Render(root *layout.PlaceableNode, option *RenderOption) *RenderResult {

	// init svg renderer
	nodes, stats := root.CollectNodesWithStat()
	sr.initRenderer(stats, option)

	// we should do global shift here to place the element in the absolute positions
	shiftX := float32(math.Abs(float64(stats.MinX))) + float32(option.NodeRadius) + float32(option.HorizontalPadding)
	shiftY := float32(option.NodeRadius) + float32(option.VerticalPadding)
	sr.Canvas.Group(fmt.Sprintf(`transform="translate(%.3f,%.3f)"`, shiftX, shiftY))

	// render nodes and edges
	for _, node := range nodes {
		sr.addNode(node, option.NodeRadius, option)
		if !node.IsLeaf() {
			sr.addEdge(node, option)
		}
	}

	sr.Canvas.Gend()
	sr.Canvas.End()

	// organize RenderResult instance
	rr := &RenderResult{
		Content: strings.NewReader(sr.buf.String()),
		Error:   nil,
	}

	return rr
}

func (sr *SvgRenderer) initRenderer(stats *layout.SizeLimitStat, opt *RenderOption) {
	// we use boxWidth to ensure root node is at the center of graphic
	// because root is always at (0,0) in relative coordinate
	boxWidth := math.Max(math.Abs(float64(stats.MinX)), math.Abs(float64(stats.MaxX))) * 2
	// width := math.Abs(float64(stats.MinX)-float64(stats.MaxX)) + float64(option.NodeRadius)*2 + float64(option.HorizontalPadding)*2
	width := boxWidth + float64(opt.NodeRadius)*2 + float64(opt.HorizontalPadding)*2
	height := math.Abs(float64(stats.MinY)-float64(stats.MaxY)) + float64(opt.NodeRadius)*2 + float64(opt.VerticalPadding)*2

	sr.Canvas.Start(int(width), int(height))

	if opt.EdgeWithArrow {
		sr.defineArrow(opt)
	}

	sr.setGlobalBackgroundColor(int(width), int(height), opt.BackgroundColor)
}

func (sr *SvgRenderer) defineArrow(opt *RenderOption) {
	var arrowSize float32 = float32(DefaultEdgeArrowSize)
	if opt.EdgeArrowSize != 0 {
		arrowSize = float32(opt.EdgeArrowSize)
	}
	arrowColor := DefaultEdgeColor
	if opt.EdgeLineColor != "" {
		arrowColor = opt.EdgeLineColor
	}

	sr.Canvas.Def()

	sr.beginMarker(selfDefinedArrowName, 0, float32(arrowSize)/2, arrowSize, arrowSize, arrowColor)
	// define the path for marker
	sr.Canvas.Path(fmt.Sprintf("M 0 0 L %.3f %.3f L 0 %.3f Z", arrowSize, float32(arrowSize)/2, arrowSize))

	sr.endMarker()

	sr.Canvas.DefEnd()
}

func (sr *SvgRenderer) addNode(node *layout.PlaceableNode, radius int, opt *RenderOption) {
	// render node as a circle with radius centered at (node.x, node.y)
	nodeStyleAttr := make([]svgStyleAttribute, 0, 1)

	var nodeColor string
	isLeaf := node.IsLeaf()
	if isLeaf {
		nodeColor = opt.NodeLeafColor
	} else {
		nodeColor = opt.NodeColor
	}
	if nodeColor == "" {
		nodeColor = DefaultNodeColor
	}
	nodeStyleAttr = append(nodeStyleAttr, svgStyleAttribute{key: "fill", value: nodeColor})

	if opt.NodeStrokeColor != "" {
		nodeStyleAttr = append(nodeStyleAttr, svgStyleAttribute{key: "stroke", value: opt.NodeStrokeColor})
		var strokeWidth int = DefaultNodeStrokeWidth
		if opt.NodeStrokeWidth != 0 {
			strokeWidth = opt.NodeStrokeWidth
		}
		nodeStyleAttr = append(nodeStyleAttr, svgStyleAttribute{key: "stroke-width", value: strconv.Itoa(strokeWidth)})
	}

	sr.constructCircle(node.X, node.Y, float32(opt.NodeRadius), []svgAttribute{
		{key: "style", value: setSvgStyleAttributes(nodeStyleAttr)},
	})

	sr.addText(node.X, node.Y, node.GetField(), opt)
}

func (sr *SvgRenderer) addText(x, y float32, text string, opt *RenderOption) {
	var fontsize int = DefaultNodeFieldTextSize
	if opt.NodeFieldTextSize != 0 {
		fontsize = opt.NodeFieldTextSize
	}

	// text style
	textAttrs := make([]svgStyleAttribute, 0, 3)
	textAttrs = append(textAttrs, svgStyleAttribute{key: "text-anchor", value: "middle"})
	textAttrs = append(textAttrs, svgStyleAttribute{key: "font-size", value: fmt.Sprintf("%d", fontsize)})
	textcolor := DefaultNodeFieldTextColor
	if opt.NodeFieldTextColor != "" {
		textcolor = opt.NodeFieldTextColor
	}
	textAttrs = append(textAttrs, svgStyleAttribute{key: "fill", value: textcolor})

	sr.constructText(x, y, text, []svgAttribute{
		{key: "style", value: setSvgStyleAttributes(textAttrs)},
		{key: "dy", value: fmt.Sprintf("%.3f", float32(fontsize)/3)},
	})
}

func (sr *SvgRenderer) addEdge(node *layout.PlaceableNode, opt *RenderOption) {
	// set edge style attributes
	edgeStyleAttr := make([]svgStyleAttribute, 0, 2)
	var linewidth int = DefaultEdgeLineWidth
	if opt.EdgeLineWidth != 0 {
		linewidth = opt.EdgeLineWidth
	}
	var linecolor string = DefaultEdgeColor
	if opt.EdgeLineColor != "" {
		linecolor = opt.EdgeLineColor
	}
	edgeStyleAttr = append(edgeStyleAttr, svgStyleAttribute{key: "stroke-width", value: fmt.Sprintf("%d", linewidth)})
	edgeStyleAttr = append(edgeStyleAttr, svgStyleAttribute{key: "stroke", value: linecolor})
	edgeAttrStyleStr := setSvgStyleAttributes(edgeStyleAttr)

	edgeAttr := []svgAttribute{
		{key: "style", value: edgeAttrStyleStr},
	}

	var edgeOffsetEnd float64 = 0
	if opt.EdgeWithArrow {
		// arrow marker is specified
		var arrowSize = DefaultEdgeArrowSize
		if opt.EdgeArrowSize != 0 {
			arrowSize = opt.EdgeArrowSize
		}
		edgeOffsetEnd = float64(arrowSize)
		edgeAttr = append(edgeAttr, svgAttribute{key: "marker-end", value: fmt.Sprintf("url(#%s)", selfDefinedArrowName)})
	}

	if node.Left != nil {
		// left edge
		edgeStartX, edgeStartY, edgeEndX, edgeEndY := measureEdgeStartEnd(
			float64(node.X),
			float64(node.Y),
			float64(node.Left.X),
			float64(node.Left.Y),
			float64(opt.NodeRadius),
			0,
			edgeOffsetEnd,
		)

		sr.constructLine(edgeStartX, edgeStartY, edgeEndX, edgeEndY, edgeAttr)
	}
	if node.Right != nil {
		// right edge
		edgeStartX, edgeStartY, edgeEndX, edgeEndY := measureEdgeStartEnd(
			float64(node.X),
			float64(node.Y),
			float64(node.Right.X),
			float64(node.Right.Y),
			float64(opt.NodeRadius),
			0,
			edgeOffsetEnd,
		)
		sr.constructLine(edgeStartX, edgeStartY, edgeEndX, edgeEndY, edgeAttr)
	}
}

func (sr *SvgRenderer) setGlobalBackgroundColor(w, h int, color string) {
	sr.addRect(0, 0, w, h, color)
}

func (sr *SvgRenderer) addRect(x, y, w, h int, color string) {
	sr.Canvas.Rect(x, y, w, h, "fill:"+color)
}

// svgStyleAttribute represents svg style attributes
type svgStyleAttribute struct {
	key   string
	value string
}

func setSvgStyleAttributes(attr []svgStyleAttribute) string {
	attrStr := strings.Builder{}
	for i, a := range attr {
		attrStr.WriteString(a.key + ":" + a.value)
		if i != len(attr)-1 {
			attrStr.WriteByte(';')
		}
	}

	return attrStr.String()
}

type svgAttribute struct {
	key   string
	value string
}

func (sr *SvgRenderer) svgCanvasBeginCustomShape(shape string, attrs []svgAttribute) {
	attrBuilder := strings.Builder{}
	for _, attr := range attrs {
		attrBuilder.WriteString(fmt.Sprintf(`%s="%s" `, attr.key, attr.value))
	}
	sr.buf.WriteString(fmt.Sprintf("<%s %s>\n", shape, attrBuilder.String()))
}

func (sr *SvgRenderer) svgCanvasEndCustomShape(shape string) {
	sr.buf.WriteString(fmt.Sprintf("</%s>\n", shape))
}

func (sr *SvgRenderer) beginMarker(id string, refX, refY, width, height float32, color string) {
	sr.svgCanvasBeginCustomShape("marker", []svgAttribute{
		{key: "id", value: id},
		{key: "markerUnits", value: "userSpaceOnUse"},
		{key: "refX", value: fmt.Sprintf("%.3f", refX)},
		{key: "refY", value: fmt.Sprintf("%.3f", refY)},
		{key: "markerWidth", value: fmt.Sprintf("%.3f", width)},
		{key: "markerHeight", value: fmt.Sprintf("%.3f", height)},
		{key: "fill", value: color},
		{key: "orient", value: "auto"},
	})
}

func (sr *SvgRenderer) endMarker() {
	sr.svgCanvasEndCustomShape("marker")
}

func (sr *SvgRenderer) svgCanvasAddCustomShape(shape string, attrs []svgAttribute) {
	attrBuilder := strings.Builder{}
	for _, attr := range attrs {
		attrBuilder.WriteString(fmt.Sprintf(`%s="%s" `, attr.key, attr.value))
	}
	sr.buf.WriteString(fmt.Sprintf("<%s %s/>\n", shape, attrBuilder.String()))
}

func (sr *SvgRenderer) constructLine(startX, startY, endX, endY float64, attrs []svgAttribute) {
	locAttrs := []svgAttribute{
		{key: "x1", value: fmt.Sprintf("%.3f", startX)},
		{key: "y1", value: fmt.Sprintf("%.3f", startY)},
		{key: "x2", value: fmt.Sprintf("%.3f", endX)},
		{key: "y2", value: fmt.Sprintf("%.3f", endY)},
	}
	attrs = append(attrs, locAttrs...)
	sr.svgCanvasAddCustomShape("line", attrs)
}

func (sr *SvgRenderer) constructCircle(cx, cy, radius float32, attrs []svgAttribute) {
	locAttrs := []svgAttribute{
		{key: "cx", value: fmt.Sprintf("%.3f", cx)},
		{key: "cy", value: fmt.Sprintf("%.3f", cy)},
		{key: "r", value: fmt.Sprintf("%.3f", radius)},
	}
	attrs = append(attrs, locAttrs...)
	sr.svgCanvasAddCustomShape("circle", attrs)
}

func (sr *SvgRenderer) constructText(x, y float32, text string, attrs []svgAttribute) {
	locAttrs := []svgAttribute{
		{key: "x", value: fmt.Sprintf("%.3f", x)},
		{key: "y", value: fmt.Sprintf("%.3f", y)},
	}
	attrs = append(attrs, locAttrs...)
	sr.svgCanvasBeginCustomShape("text", attrs)
	sr.Canvas.Writer.Write([]byte(text))
	sr.svgCanvasEndCustomShape("text")
}
