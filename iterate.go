package bitreevis

// CalHeight calculates the height of a binary tree in a recursive manner.
//
// If root does not have any children, CalHeight returns 0.
func CalHeight(root BiNode) int {
	if BiNodeIsNil(root) {
		return 0
	}
	return maxInt(CalHeight(root.GetLeftChild()), CalHeight(root.GetRightChild())) + 1
}

// CollectNodeByLevelOrder collects all nodes in a binary tree in level order.
//
// CollectNodeByLevelOrder returns nodes from top to down, from left to right, in the form of [][]BiNode
func CollectNodeByLevelOrder(root BiNode) [][]BiNode {
	nodes := make([]BiNode, 0)
	nodes = append(nodes, root)

	curLevel := 1
	height := CalHeight(root)
	levels := make([][]BiNode, 0, height)

	for len(nodes) > 0 && curLevel <= height {
		n := len(nodes)
		level := make([]BiNode, 0, n)
		for i := 0; i < n; i++ {
			cur := nodes[0]
			if !BiNodeIsNil(cur) {
				level = append(level, cur)
				nodes = append(nodes, cur.GetLeftChild())
				nodes = append(nodes, cur.GetRightChild())
			}
			nodes = nodes[1:]
		}
		levels = append(levels, level)
		curLevel += 1
	}

	return levels
}
