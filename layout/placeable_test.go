package layout_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/ryanreadbooks/bitreevis/bitree"
	"github.com/ryanreadbooks/bitreevis/layout"
)

type myNode struct {
	Left  *myNode
	Right *myNode
	Value int
}

func (m *myNode) GetLeftChild() bitree.BiNode {
	return m.Left
}

func (m *myNode) GetRightChild() bitree.BiNode {
	return m.Right
}

func (m *myNode) GetField() string {
	return strconv.Itoa(m.Value)
}

func TestBuildPlaceableTree_Dev(t *testing.T) {

	node6 := &myNode{Value: 6}
	node19 := &myNode{Value: 19}
	root := &myNode{Value: 5, Left: node6, Right: node19}

	node4 := &myNode{Value: 4}
	node3 := &myNode{Value: 3}
	node6.Left = node4

	node19.Right = node3

	pRoot := layout.NewPlaceableTreeFromBiNode(root)

	pNodes := bitree.CollectNodeByLevelOrder(pRoot)

	for _, floor := range pNodes {
		for _, node := range floor {
			fmt.Printf("%v, ", node.GetField())
		}
		fmt.Println("\n---------------")
	}
}
