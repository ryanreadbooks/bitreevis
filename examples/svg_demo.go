package main

import (
	"log"
	"strconv"

	"github.com/ryanreadbooks/bitreevis"
)

// Define our own binary tree node. You can put any members in it.
// Just make sure your own binary tree node implements bitreevis.BiNode interface.

type MyBiNode struct {
	Left  *MyBiNode
	Right *MyBiNode
	Value int
}

func (m *MyBiNode) GetLeftChild() bitreevis.BiNode {
	return m.Left
}

func (m *MyBiNode) GetRightChild() bitreevis.BiNode {
	return m.Right
}

func (m *MyBiNode) GetField() string {
	return strconv.Itoa(m.Value)
}

func NewNode(v int) *MyBiNode {
	return &MyBiNode{Value: v}
}

func main() {

	// we construct a binary tree with our self-defined MyBiNode node.
	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)
	node4 := NewNode(4)
	node5 := NewNode(5)
	node6 := NewNode(6)
	node7 := NewNode(7)
	node8 := NewNode(8)

	node8.Left = node4
	node8.Right = node7

	node4.Left = node1
	node4.Right = node3

	node3.Right = node2

	node7.Left = node5
	node7.Right = node6

	// render tree to svg file
	err := bitreevis.VisAsSvg(node8, "dev.svg", &bitreevis.RenderOption{
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
	})
	if err != nil {
		log.Fatalf("can not render svg file: %v", err)
	}
}
