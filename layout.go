package bitreevis

import (
	"math"
)

func layoutSetup(root *PlaceableNode, level int, lmost, rmost *extreme, siblingSeparation, nodeWidth, levelSeparation int) {
	var l, r *PlaceableNode
	ll := &extreme{}
	lr := &extreme{}
	rl := &extreme{}
	rr := &extreme{}

	var currSep, rootSep, lOffSum, rOffSum float32 = 0.0, 0.0, 0.0, 0.0

	if root == nil {
		lmost.level = -1
		rmost.level = -1
	} else {
		root.Y = float32(level) * float32(nodeWidth*2+levelSeparation)
		l = root.Left
		r = root.Right
		layoutSetup(l, level+1, lr, ll, siblingSeparation, nodeWidth, levelSeparation)
		layoutSetup(r, level+1, rr, rl, siblingSeparation, nodeWidth, levelSeparation)
		if r == nil && l == nil {
			rmost.addr = root
			lmost.addr = root
			rmost.level = level
			lmost.level = level
			rmost.offset = 0
			lmost.offset = 0
			root.Offset = 0
		} else {
			currSep = float32(siblingSeparation)
			rootSep = float32(siblingSeparation)
			lOffSum = 0
			rOffSum = 0

			for l != nil && r != nil {
				if currSep < float32(siblingSeparation) {
					rootSep += float32(siblingSeparation) - currSep
					currSep = float32(siblingSeparation)
				}
				if l.Right != nil {
					lOffSum += l.Offset
					currSep -= l.Offset
					l = l.Right
				} else {
					lOffSum -= l.Offset
					currSep += l.Offset
					l = l.Left
				}

				if r.Left != nil {
					rOffSum -= r.Offset
					currSep -= r.Offset
					r = r.Left
				} else {
					rOffSum += r.Offset
					currSep += r.Offset
					r = r.Right
				}
			}
			root.Offset = (rootSep + float32(nodeWidth)) / 2
			lOffSum -= root.Offset
			rOffSum += root.Offset

			if rl.level > ll.level || root.Left == nil {
				lmost.addr = rl.addr
				lmost.level = rl.level
				lmost.offset = rl.offset
				lmost.offset += root.Offset
			} else {
				lmost.addr = ll.addr
				lmost.level = ll.level
				lmost.offset = ll.offset
				lmost.offset -= root.Offset
			}

			if lr.level > rr.level || root.Right == nil {
				rmost.addr = lr.addr
				rmost.level = lr.level
				rmost.offset = lr.offset
				rmost.offset -= root.Offset
			} else {
				rmost.addr = rr.addr
				rmost.level = rr.level
				rmost.offset = rr.offset
				rmost.offset += root.Offset
			}

			if l != nil && l != root.Left {
				rr.addr.Thread = true
				rr.addr.Offset = float32(math.Abs(float64(rr.offset) + float64(root.Offset) - float64(lOffSum)))
				if (lOffSum - root.Offset) <= rr.offset {
					rr.addr.Left = l
				} else {
					rr.addr.Right = l
				}
			} else if r != nil && r != root.Right {
				ll.addr.Thread = true
				ll.addr.Offset = float32(math.Abs(float64(ll.offset) - float64(root.Offset) - float64(rOffSum)))
				if (rOffSum + root.Offset) >= ll.offset {
					ll.addr.Right = r
				} else {
					ll.addr.Left = r
				}
			}
		}
	}
}

func layoutPetrify(root *PlaceableNode, xPos float32) {
	if root != nil {
		root.X = float32(xPos)
		if root.Thread {
			root.Thread = false
			root.Left = nil
			root.Right = nil
		}
		layoutPetrify(root.Left, xPos-root.Offset)
		layoutPetrify(root.Right, xPos+root.Offset)
	}
}

func peformLayout(root *PlaceableNode, siblingSeparation, nodeWidth, levelSeparation int) *PlaceableNode {
	lm, rm := &extreme{}, &extreme{}
	layoutSetup(root, 0, lm, rm, siblingSeparation, nodeWidth, levelSeparation)
	layoutPetrify(root, root.X)

	return root
}

func PerformLayout(root *PlaceableNode, siblingSeparation, nodeWidth, levelSeparation int) *PlaceableNode {
	return peformLayout(root, siblingSeparation+nodeWidth*2, nodeWidth, levelSeparation)
}
