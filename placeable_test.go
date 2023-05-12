package bitreevis_test

import (
	"fmt"
	"testing"

	"github.com/ryanreadbooks/bitreevis"
)

func TestBuildPlaceableTree_Dev(t *testing.T) {

	node6 := &myNode{Value: 6}
	node19 := &myNode{Value: 19}
	root := &myNode{Value: 5, Left: node6, Right: node19}

	node4 := &myNode{Value: 4}
	node3 := &myNode{Value: 3}
	node6.Left = node4

	node19.Right = node3

	pRoot := bitreevis.NewPlaceableTreeFromBiNode(root)

	pNodes := bitreevis.CollectNodeByLevelOrder(pRoot)

	for _, floor := range pNodes {
		for _, node := range floor {
			fmt.Printf("%v, ", node.GetField())
		}
		fmt.Println("\n---------------")
	}
}
