package bitreevis

import ()

type extreme struct {
	addr   *PlaceableNode
	offset float32
	level  int
}

// PlaceableNode represents a node which can be placed and rendered.
//
// PlaceableNode's attributes follow the definition in algorithm of Tidier Drawings of Trees by Edward M. Reingold and John S. Tilford.
type PlaceableNode struct {
	Parent *PlaceableNode
	Left   *PlaceableNode
	Right  *PlaceableNode
	X      float32
	Y      float32
	Offset float32
	Thread bool
	Field  string
}

func (p *PlaceableNode) IsLeaf() bool {
	if p == nil {
		return false
	}
	return p.Left == nil && p.Right == nil
}

// Implement interface BiNode for placeableNode
func (p *PlaceableNode) GetLeftChild() BiNode {
	return p.Left
}

func (p *PlaceableNode) GetRightChild() BiNode {
	return p.Right
}

func (p *PlaceableNode) GetField() string {
	return p.Field
}

func (p *PlaceableNode) CollectNodes() []*PlaceableNode {
	nodes := make([]*PlaceableNode, 0, 16) // pre-allocation
	return inOrderTraverse(p, nodes)
}

func inOrderTraverse(root *PlaceableNode, nodes []*PlaceableNode) []*PlaceableNode {
	if root == nil {
		return nodes
	}
	nodes = inOrderTraverse(root.Left, nodes)
	nodes = append(nodes, root)
	nodes = inOrderTraverse(root.Right, nodes)

	return nodes
}

type SizeLimitStat struct {
	MinX float32
	MaxX float32
	MinY float32
	MaxY float32
}

func (p *PlaceableNode) CollectNodesWithStat() (nodes []*PlaceableNode, limit *SizeLimitStat) {
	nodes = make([]*PlaceableNode, 0, 16)
	limit = &SizeLimitStat{}
	nodes = inOrderTraverseWithStat(p, nodes, limit)

	return
}

func inOrderTraverseWithStat(root *PlaceableNode, nodes []*PlaceableNode, limit *SizeLimitStat) []*PlaceableNode {
	if root == nil {
		return nodes
	}
	limit.MinX = minFloat32(root.X, limit.MinX)
	limit.MaxX = maxFloat32(root.X, limit.MaxX)
	limit.MinY = minFloat32(root.Y, limit.MinY)
	limit.MaxY = maxFloat32(root.Y, limit.MaxY)

	nodes = inOrderTraverseWithStat(root.Left, nodes, limit)
	nodes = append(nodes, root)
	nodes = inOrderTraverseWithStat(root.Right, nodes, limit)

	return nodes
}

// NewPlaceableNode returns a new *PlaceableNode with spefified field value.
func NewPlaceableNode(field string) *PlaceableNode {
	return &PlaceableNode{Field: field}
}

// NewPlaceableTreeFromBiNode builds a tree made of placeableNode from a tree made of BiNode
func NewPlaceableTreeFromBiNode(root BiNode) *PlaceableNode {
	return buildPlaceableTreeRecursive(root)
}

// buildPlaceableTreeRecursive help build tree in a recursive manner
func buildPlaceableTreeRecursive(root BiNode) *PlaceableNode {
	if BiNodeIsNil(root) {
		return nil
	}
	pRoot := NewPlaceableNode(root.GetField())
	pRoot.Left = buildPlaceableTreeRecursive(root.GetLeftChild())
	if pRoot.Left != nil {
		pRoot.Left.Parent = pRoot
	}
	pRoot.Right = buildPlaceableTreeRecursive(root.GetRightChild())
	if pRoot.Right != nil {
		pRoot.Right.Parent = pRoot
	}

	return pRoot
}
