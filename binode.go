package bitreevis

import "reflect"

// A FieldHolder is an interface for storing content of a node in binary tree.
//
// GetField() should return a string .
type FieldHolder interface {
	GetField() string
}

// A BiNode represents a node in binary tree.
//
// BiNode specify the minumum elements of a node in binary tree, which are left child, right child and content of node.
type BiNode interface {
	FieldHolder
	GetLeftChild() BiNode
	GetRightChild() BiNode
}

// BiNodeIsNil reports whether the BiNode interface is nil.
//
// If BiNode interface itself is nil, BiNodeIsNil returns true.
// If BiNode interface itself is not nil, but the data of interface is nil, BiNodeIsNil returns true
func BiNodeIsNil(node BiNode) bool {
	if node == nil {
		return true
	}
	// we still need to check the data underneath
	v := reflect.ValueOf(node)
	if v.Kind() == reflect.Ptr {
		return v.IsNil()
	}
	return false
}

// PaintableBiNode represents a node whose color is private.
// You can implement this interface if you want each of your node to have different colors.
type PaintableBiNode interface {
	BiNode

	// GetColor returns a color string for this node.
	GetColor() string
}

// isPaintable helps check the input data of BiNode has a method called 'GetColor'
// If 'GetColor' method exists
func isPaintable(root BiNode) (string, bool) {
	v, ok := root.(PaintableBiNode)
	if ok {
		return v.GetColor(), true
	}

	return "", false
}
