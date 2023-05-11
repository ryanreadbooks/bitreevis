package render_test

import (
	"io"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ryanreadbooks/bitreevis/bitree"
	"github.com/ryanreadbooks/bitreevis/layout"
	"github.com/ryanreadbooks/bitreevis/render"
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

	pRoot := layout.NewPlaceableTreeFromBiNode(root)
	opt := render.RenderOption{
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
	pRoot = layout.PerformLayout(pRoot, opt.SiblingSeparation, opt.NodeRadius, opt.LevelSeparation)

	renderer := render.NewSvgRenderer()
	result := renderer.Render(pRoot, &opt)
	require.Nil(t, result.Error)
	f, err := os.Create("dev.svg")
	require.Nil(t, err)
	defer f.Close()
	_, err = io.Copy(f, result.Content)
	require.Nil(t, err)
	os.Remove("dev.svg")
}
