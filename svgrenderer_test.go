package bitreevis_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ryanreadbooks/bitreevis"
)

type myNode struct {
	Left  *myNode
	Right *myNode
	Value int
}

func (m *myNode) GetLeftChild() bitreevis.BiNode {
	return m.Left
}

func (m *myNode) GetRightChild() bitreevis.BiNode {
	return m.Right
}

func (m *myNode) GetField() string {
	return strconv.Itoa(m.Value)
}

func TestSvgRenderer_Dev(t *testing.T) {
	node6 := &myNode{Value: 6}
	node19 := &myNode{Value: 19}
	root := &myNode{Value: 5, Left: node6, Right: node19}

	node4 := &myNode{Value: 4}
	node3 := &myNode{Value: 3}
	node6.Left = node4
	node2 := &myNode{Value: 2}
	node6.Right = node2
	node19.Left = node3

	pRoot := bitreevis.NewPlaceableTreeFromBiNode(root)
	opt := bitreevis.RenderOption{
		SiblingSeparation: 30,
		LevelSeparation:   40,
		NodeRadius:        50,
		HorizontalPadding: 30,
		VerticalPadding:   10,
		BackgroundColor:   "#0daaf4",
		NodeStrokeColor:   "black",
		NodeFieldTextSize: 16,
		NodeLeafColor:     "#00eeac",
		EdgeWithArrow:     true,
		EdgeArrowSize:     5,
	}
	pRoot = bitreevis.PerformLayout(pRoot, opt.SiblingSeparation, opt.NodeRadius, opt.LevelSeparation)

	renderer := bitreevis.NewSvgRenderer()
	result := renderer.Render(pRoot, &opt)
	require.Nil(t, result.Error())
	require.Nil(t, result.Save("dev.svg"))
	os.Remove("dev.svg")
}

type rbNode struct {
	Left  *rbNode
	Right *rbNode
	Value int
	Color string
}

func (m *rbNode) GetLeftChild() bitreevis.BiNode {
	return m.Left
}

func (m *rbNode) GetRightChild() bitreevis.BiNode {
	return m.Right
}

func (m *rbNode) GetField() string {
	return strconv.Itoa(m.Value)
}

func (m *rbNode) GetColor() string {
	return m.Color
}

func TestSvgRenderer_LocalNodeColor(t *testing.T) {

	node1 := &rbNode{Value: 1, Color: "black"}
	node2 := &rbNode{Value: 2, Color: "black"}
	node3 := &rbNode{Value: 3, Color: "red"}
	node4 := &rbNode{Value: 4, Color: "black"}
	node5 := &rbNode{Value: 5, Color: "black"}
	node6 := &rbNode{Value: 6, Color: "red"}
	node7 := &rbNode{Value: 7, Color: "red"}
	node8 := &rbNode{Value: 8, Color: "red"}
	node9 := &rbNode{Value: 9, Color: "red"}

	node1.Left = node2
	node1.Right = node3
	node3.Left = node4
	node3.Right = node5
	node4.Left = node6
	node4.Right = node7
	node5.Left = node8
	node5.Right = node9

	pRoot := bitreevis.NewPlaceableTreeFromBiNode(node1)
	opt := bitreevis.RenderOption{
		SiblingSeparation:  20,
		LevelSeparation:    20,
		NodeRadius:         20,
		HorizontalPadding:  15,
		VerticalPadding:    20,
		BackgroundColor:    "#eeffff",
		NodeStrokeColor:    "black",
		NodeFieldTextColor: "white",
		NodeFieldTextSize:  14,
		EdgeWithArrow:      true,
		EdgeArrowSize:      8,
	}
	pRoot = bitreevis.PerformLayout(pRoot, opt.SiblingSeparation, opt.NodeRadius, opt.LevelSeparation)

	renderer := bitreevis.NewSvgRenderer()
	result := renderer.Render(pRoot, &opt)
	require.Nil(t, result.Error())
	require.Nil(t, result.Save("dev.svg"))
	os.Remove("dev.svg")
}
