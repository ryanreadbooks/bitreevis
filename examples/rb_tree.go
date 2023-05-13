package main

import (
	"log"
	"strconv"

	"github.com/ryanreadbooks/bitreevis"
)

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

func main() {
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

	opt := bitreevis.RenderOption{
		SiblingSeparation: 20,
		LevelSeparation:   20,
		NodeRadius:        20,
		HorizontalPadding: 15,
		VerticalPadding:   20,
		BackgroundColor:   "#eeffff",
		NodeStrokeColor:   "black",
		NodeFieldTextColor: "white",
		NodeFieldTextSize: 14,
		EdgeWithArrow:     true,
		EdgeArrowSize:     8,
	}

	err := bitreevis.VisAsSvg(node1, "rbnode.svg", &opt)
	if err != nil {
		log.Fatalf("can not render tree into svg: %v\n", err)
	}

}